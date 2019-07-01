// genoapi generates OpenAPI files from .proto definitions and other sources
// of CUE constraints.
//
// Generation adopts the Proto <-> JSON mappings conventions. Most notably,
// field names are converted to JSON names.
//
// Generation involves the following steps:
//
//   1. Convert .proto files to CUE files
//   2. Validate the consistency of the CUE defintions
//   3. Convert CUE files to self-contained OpenAPI files.
//
// Each of which is documented in more detail below.
//
//
// 1. Converting Proto to CUE
//
// genoapi generates all .proto files using a single builder. As the Istio
// OpenAPI files are self-contained, this is not strictly necessary, but it
// allows for better checking and works better if one wants to generate the
// intermediate CUE results for evaluation.
//
// Field names are mapped using JSON naming.
//
// Protobuf definitions may contain (cue.val) and (cue.opt) options to annotate
// fields with constraints.
//
// Caveats:
// - It is assumed that the input .proto files are valid and compile with
//   protoc. The conversion may ignore errors if files are invalid.
// - CUE package names share the same naming conventions as Go packages. CUE
//   requires the go_package option to exist and be well-defined. Note that
//   some of the gogoproto go_package definition are illformed. Be sure to use
//   the original .proto files for the google protobuf types.
//
// 2. Validating generated CUE
//
// The generated CUE from the previous step may be combined with other sources
// of CUE constraints. This step validates the combined sources for consistency.
//
//
// 3. Converting CUE to OpenAPI
//
// In this step a self-contained OpenAPI definition is generated for each
// directory containing proto definitions. Files are made self-contained by
// including a schema definition for each imported type within the OpenAPI
// spec itself. To avoid name collissions, types are, by convention, prefixed
// with their proto package name.
//
//
// Possible extensions to the generation pipeline
//
// The generation pipeline can be augmented by injecting CUE from other sources
// before step 2. As combining CUE sources is a commutative operation, order
// of injection does not matter and there is no need for the user to be
// explicit about any order or injection points.
//
// Examples of other possible CUE sources are:
// - hand-written .cue files in each of the cue directories
// - constraints extracted from Go code
//
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/openapi"
	"cuelang.org/go/encoding/protobuf"
	"github.com/getkin/kin-openapi/openapi3"
)

var (
	inplace = flag.Bool("inplace", false, "generate configurations in place")
	paths   = flag.String("paths", "/protobuf", "comma-separated path to search for .proto imports")
)

func main() {
	flag.Parse()

	// Include paths for proto evaluation.
	cwd, _ := os.Getwd()
	importPaths := []string{cwd}
	for _, p := range strings.Split(*paths, ",") {
		if p != "" {
			abs, _ := filepath.Abs(p)
			importPaths = append(importPaths, abs)
		}
	}

	protoConfig := &protobuf.Config{
		Module: "istio.io/api",
		Paths:  importPaths,
	}

	b := protobuf.NewBuilder(protoConfig)

	err := filepath.Walk(cwd, func(path string, f os.FileInfo, _ error) (err error) {
		if !strings.HasSuffix(path, ".proto") {
			return nil
		}
		return b.AddFile(path, nil)
	})

	files, err := b.Files()
	if err != nil {
		fatal(err, "Error generating CUE from proto")
	}

	modRoot := cwd
	if !*inplace {
		modRoot, err = ioutil.TempDir("", "istio-openapi-gen")
		if err != nil {
			log.Fatalf("Error creating temp dir: %v", err)
		}
		defer os.RemoveAll(modRoot)
	}

	for _, f := range files {
		b, err := format.Node(f)
		if err != nil {
			fatal(err, "Error formatting file: ")
		}
		filename := f.Filename
		relPath, _ := filepath.Rel(cwd, filename)
		filename = filepath.Join(modRoot, relPath)

		_ = os.MkdirAll(filepath.Dir(filename), 0755)
		if err := ioutil.WriteFile(filename, b, 0644); err != nil {
			log.Fatalf("Error writing file: %v", err)
		}
	}

	// Gernate the openAPI
	protoNames := map[string]string{}

	// build map of CUE package import paths (Go paths) to proto package paths.
	builds, _ := b.Instances()
	for _, i := range builds {
		protoNames[i.ImportPath] = i.DisplayPath
	}

	// Set up the configuration for the OpenAPI generator.
	oapiConfig := &openapi.Config{
		// Resolve all references to objects within the file.
		SelfContained: true,

		// Computes the reference format for Istio types.
		ReferenceFunc: func(p *cue.Instance, path []string) string {
			name := strings.Join(path, ".")
			// Map CUE names to proto names.
			name = strings.Replace(name, "_", ".", -1)

			pkg := protoNames[p.ImportPath]
			if pkg == "" {
				log.Fatalf("No protoname for pkg with import path %q", p.ImportPath)
			}
			return pkg + "." + name
		},
	}

	// Build the OpenAPI files.
	instances := load.Instances([]string{"./..."}, &load.Config{
		Dir:    modRoot,
		Module: "istio.io/api",
	})
	for _, inst := range cue.Build(instances) {
		if inst.Err != nil {
			fatal(err, "Instance failed.")
		}
		basename := protoNames[inst.ImportPath] + ".json"
		fmt.Println("Building", basename+"...")

		if err := inst.Value().Validate(); err != nil {
			fatal(err, "Validation failed.")
		}

		b, err := openapi.Gen(inst, oapiConfig)
		if err != nil {
			fatal(err, "Error generating OpenAPI file")
		}

		// Note: this just tests basic OpenAPI 3 validity. It cannot, of course,
		// know if the the proto files were correctly mapped.
		_, err = openapi3.NewSwaggerLoader().LoadSwaggerFromData(b)
		if err != nil {
			log.Fatalf("Invalid OpenAPI generated: %v", err)
		}

		var buf bytes.Buffer
		_ = json.Indent(&buf, b, "", "  ")

		relPath, _ := filepath.Rel(modRoot, inst.Dir)
		filename := filepath.Join(cwd, relPath, basename)

		err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
		if err != nil {
			log.Fatalf("Error writing OpenAPI file: %v", err)
		}
	}
}

func fatal(err error, msg string) {
	errors.Print(os.Stderr, err, nil)
	log.Fatal(msg)
}
