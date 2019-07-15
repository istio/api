// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"

	"io/ioutil"
	"istio.io/api/annotation"
)

const (
	outputTemplate = `
// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// GENERATED FILE -- DO NOT EDIT
//
// nolint: lll
//go:generate go run $GOPATH/src/istio.io/api/annotation/cmd/gen/main.go --input=$GOPATH/src/istio.io/api/annotation/annotations.yaml --output=$GOPATH/src/istio.io/api/annotation/annotations.gen.go
//go:generate goimports -w annotations.gen.go

package {{ .Package }}

var (
	{{ range .Annotations }}
		{{.VariableName}} = Instance {
          Name: "{{ .Name }}",
          Description: {{ wordWrap .Description 24 }},
          Hidden: {{ .Hidden }},
          Deprecated: {{ .Deprecated }},
        }
	{{ end }}
)

`
)

var (
	input  string
	output string

	nameSeparator = regexp.MustCompile("[._\\-]")

	rootCmd = cobra.Command{
		Use:   "gen",
		Short: "Generates a Go source file containing annotations.",
		Long:  "Generates a Go source file containing annotation definitions based on an input YAML file.",
		Run: func(cmd *cobra.Command, args []string) {
			yamlContent, err := ioutil.ReadFile(input)
			if err != nil {
				log.Fatalf("unable to read input file: %v", err)
			}

			// Unmarshal the file.
			var cfg Configuration
			if err = yaml.Unmarshal(yamlContent, &cfg); err != nil {
				log.Fatalf("error parsing input file: %v", err)
			}

			// Generate variable names if not provided in the yaml.
			for i := range cfg.Annotations {
				if cfg.Annotations[i].Name == "" {
					log.Fatalf("annotation %d in input file missing name", i)
				}
				if cfg.Annotations[i].VariableName == "" {
					cfg.Annotations[i].VariableName = camelCase(cfg.Annotations[i].Name)
				}
			}

			// Create the output file template.
			t, err := template.New("annoTemplate").Funcs(template.FuncMap{
				"wordWrap": wordWrap,
			}).Parse(outputTemplate)
			if err != nil {
				log.Fatalf("failed parsing annotation template: %v", err)
			}

			// Generate the Go source.
			var goSource bytes.Buffer
			if err := t.Execute(&goSource, map[string]interface{}{
				"Package":     getPackage(),
				"Annotations": cfg.Annotations,
			}); err != nil {
				log.Fatalf("failed generating output Go source code %s: %v", output, err)
			}

			if err := ioutil.WriteFile(output, goSource.Bytes(), 0666); err != nil {
				log.Fatalf("Failed writing to output file %s: %v", output, err)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&input, "input", "",
		"Input YAML file to be parsed.")
	rootCmd.PersistentFlags().StringVar(&output, "output", "",
		"Output Go file to be generated.")

	flag.CommandLine.VisitAll(func(gf *flag.Flag) {
		rootCmd.PersistentFlags().AddGoFlag(gf)
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

type AnnotationVariable struct {
	annotation.Instance
	VariableName string `json:"variableName"`
}

type Configuration struct {
	Annotations []AnnotationVariable `json:"annotations"`
}

func getPackage() string {
	return filepath.Base(filepath.Dir(output))
}

func camelCase(name string) string {
	// Take the portion of the name after "/"
	name = strings.Split(name, "/")[1]

	// Replace any separator characters with spaces.
	name = nameSeparator.ReplaceAllString(name, " ")

	// Capitalize the first letter of each word
	name = strings.Title(name)

	// Remove the spaces to generate a camel case variable name.
	return strings.ReplaceAll(name, " ", "")
}

func wordWrap(in string, indent int) string {
	words := strings.Split(in, " ")

	maxLineLength := 80

	out := ""
	line := ""
	for len(words) > 0 {
		// Take the next word.
		word := words[0]
		words = words[1:]

		if indent+len(line)+len(word) > maxLineLength {
			// Need to word wrap - emit the current line.
			out += "\"" + line + " \""
			line = ""

			// Wrap to the start of the next line.
			out += "+\n"

			// Indent to the start position of the next line.
			for i := 0; i < indent; i++ {
				out += " "
			}
		}

		// Add the word to the current line.
		if len(line) > 0 {
			line += " "
		}
		line += word
	}

	// Emit the final line
	out += "\"" + line + "\""

	return out
}
