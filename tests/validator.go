// Copyright Istio Authors
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

package crd

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextval "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/validation"
	structuralschema "k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/schema/cel"
	structuraldefaulting "k8s.io/apiextensions-apiserver/pkg/apiserver/schema/defaulting"
	structurallisttype "k8s.io/apiextensions-apiserver/pkg/apiserver/schema/listtype"
	structuralpruning "k8s.io/apiextensions-apiserver/pkg/apiserver/schema/pruning"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	kubeyaml "k8s.io/apimachinery/pkg/util/yaml"
	celconfig "k8s.io/apiserver/pkg/apis/cel"
)

// Validator returns a new validator for custom resources
// Warning: this is meant for usage in tests only
type Validator struct {
	byGvk      map[schema.GroupVersionKind]validation.SchemaCreateValidator
	structural map[schema.GroupVersionKind]*structuralschema.Structural
	cel        map[schema.GroupVersionKind]*cel.Validator
	// If enabled, resources without a validator will be ignored. Otherwise, they will fail.
	SkipMissing bool
}

// Split where the '---' appears at the very beginning of a line. This will avoid
// accidentally splitting in cases where yaml resources contain nested yaml (which
// is indented).
var splitRegex = regexp.MustCompile(`(^|\n)---`)

// SplitString splits the given yaml doc if it's multipart document.
func SplitString(yamlText string) []string {
	out := make([]string, 0)
	parts := splitRegex.Split(yamlText, -1)
	for _, part := range parts {
		part := strings.TrimSpace(part)
		if len(part) > 0 {
			out = append(out, part)
		}
	}
	return out
}

func (v *Validator) ValidateCustomResource(o runtime.Object) error {
	content, err := runtime.DefaultUnstructuredConverter.ToUnstructured(o)
	if err != nil {
		return err
	}

	un := &unstructured.Unstructured{Object: content}
	vd, f := v.byGvk[un.GroupVersionKind()]
	if !f {
		if v.SkipMissing {
			return nil
		}
		return fmt.Errorf("failed to validate type %v: no validator found", un.GroupVersionKind())
	}
	// Fill in defaults
	structural := v.structural[un.GroupVersionKind()]
	structuraldefaulting.Default(un.Object, structural)
	if err := validation.ValidateCustomResource(nil, un.Object, vd).ToAggregate(); err != nil {
		return fmt.Errorf("%v/%v/%v: %v", un.GroupVersionKind().Kind, un.GetName(), un.GetNamespace(), err)
	}
	if err := structurallisttype.ValidateListSetsAndMaps(nil, structural, un.Object).ToAggregate(); err != nil {
		return fmt.Errorf("%v/%v/%v: %v", un.GroupVersionKind().Kind, un.GetName(), un.GetNamespace(), err)
	}
	pruneOpts := structuralschema.UnknownFieldPathOptions{TrackUnknownFieldPaths: true}
	if unknownFieldPaths := structuralpruning.PruneWithOptions(un.DeepCopy().Object, structural, false, pruneOpts); len(unknownFieldPaths) > 0 {
		return fmt.Errorf("%v/%v/%v: unknown fields %v", un.GroupVersionKind().Kind, un.GetName(), un.GetNamespace(), unknownFieldPaths)
	}

	errs, _ := v.cel[un.GroupVersionKind()].Validate(context.Background(), nil, structural, un.Object, nil, celconfig.RuntimeCELCostBudget)
	if errs.ToAggregate() != nil {
		return fmt.Errorf("%v/%v/%v: %v", un.GroupVersionKind().Kind, un.GetName(), un.GetNamespace(), errs.ToAggregate().Error())
	}
	return nil
}

func NewValidatorFromFiles(files ...string) (*Validator, error) {
	crds := []apiextensions.CustomResourceDefinition{}
	closers := make([]io.Closer, 0, len(files))
	defer func() {
		for _, closer := range closers {
			closer.Close()
		}
	}()
	for _, file := range files {
		data, err := os.Open(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read input yaml file: %v", err)
		}
		closers = append(closers, data)

		yamlDecoder := kubeyaml.NewYAMLOrJSONDecoder(data, 512*1024)
		for {
			un := &unstructured.Unstructured{}
			err = yamlDecoder.Decode(&un)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			crd := apiextensions.CustomResourceDefinition{}
			switch un.GroupVersionKind() {
			case schema.GroupVersionKind{
				Group:   "apiextensions.k8s.io",
				Version: "v1",
				Kind:    "CustomResourceDefinition",
			}:
				crdv1 := apiextensionsv1.CustomResourceDefinition{}
				if err := runtime.DefaultUnstructuredConverter.
					FromUnstructured(un.UnstructuredContent(), &crdv1); err != nil {
					return nil, err
				}
				if err := apiextensionsv1.Convert_v1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(&crdv1, &crd, nil); err != nil {
					return nil, err
				}
			case schema.GroupVersionKind{
				Group:   "apiextensions.k8s.io",
				Version: "v1beta1",
				Kind:    "CustomResourceDefinition",
			}:
				crdv1beta1 := apiextensionsv1beta1.CustomResourceDefinition{}
				if err := runtime.DefaultUnstructuredConverter.
					FromUnstructured(un.UnstructuredContent(), &crdv1beta1); err != nil {
					return nil, err
				}
				if err := apiextensionsv1beta1.Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(&crdv1beta1, &crd, nil); err != nil {
					return nil, err
				}
			default:
				return nil, fmt.Errorf("unknown CRD type: %v", un.GroupVersionKind())
			}
			crds = append(crds, crd)
		}
	}
	return NewValidatorFromCRDs(crds...)
}

func NewValidatorFromCRDs(crds ...apiextensions.CustomResourceDefinition) (*Validator, error) {
	v := &Validator{
		byGvk:      map[schema.GroupVersionKind]validation.SchemaCreateValidator{},
		structural: map[schema.GroupVersionKind]*structuralschema.Structural{},
		cel:        map[schema.GroupVersionKind]*cel.Validator{},
	}
	for _, crd := range crds {
		versions := crd.Spec.Versions
		if len(versions) == 0 {
			versions = []apiextensions.CustomResourceDefinitionVersion{{Name: crd.Spec.Version}} // nolint: staticcheck
		}
		for _, v := range versions {
			crd.Status.StoredVersions = append(crd.Status.StoredVersions, v.Name)
		}
		errs := apiextval.ValidateCustomResourceDefinition(context.Background(), &crd)
		if len(errs) > 0 {
			return nil, fmt.Errorf("CRD %v is not valid: %v", crd.Name, formatError(errs))
		}
		for _, ver := range versions {
			gvk := schema.GroupVersionKind{
				Group:   crd.Spec.Group,
				Version: ver.Name,
				Kind:    crd.Spec.Names.Kind,
			}
			crdSchema := ver.Schema
			if crdSchema == nil {
				crdSchema = crd.Spec.Validation
			}
			if crdSchema == nil {
				return nil, fmt.Errorf("crd did not have validation defined")
			}

			schemaValidator, _, err := validation.NewSchemaValidator(crdSchema.OpenAPIV3Schema)
			if err != nil {
				return nil, err
			}
			structural, err := structuralschema.NewStructural(crdSchema.OpenAPIV3Schema)
			if err != nil {
				return nil, err
			}

			v.byGvk[gvk] = schemaValidator
			v.structural[gvk] = structural
			// CEL programs are compiled and cached here
			if celv := cel.NewValidator(structural, true, celconfig.PerCallLimit); celv != nil {
				v.cel[gvk] = celv
			}

		}
	}

	return v, nil
}

func formatError(errs field.ErrorList) error {
	if len(errs) == 0 {
		return nil
	}
	s := strings.Builder{}
	s.WriteString("\n")
	for _, e := range errs.ToAggregate().Errors() {
		s.WriteString(fmt.Sprintf("- %v\n", strings.TrimPrefix(e.Error(), "spec.validation.openAPIV3Schema")))
	}
	return fmt.Errorf(s.String())
}

func NewIstioValidator(t *testing.T) *Validator {
	v, err := NewValidatorFromFiles(filepath.Join("../kubernetes/customresourcedefinitions.gen.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	return v
}
