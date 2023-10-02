package crd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

type TestExpectation struct {
	WantErr string `json:"_err,omitempty"`
}

func TestCRDs(t *testing.T) {
	v := NewIstioValidator(t)
	d, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range d {
		t.Run(f.Name(), func(t *testing.T) {
			f, err := os.ReadFile(filepath.Join("testdata", f.Name()))
			if err != nil {
				t.Fatal(err)
			}
			for _, item := range SplitString(string(f)) {
				obj := &unstructured.Unstructured{}
				if err := yaml.Unmarshal([]byte(item), obj); err != nil {
					t.Fatal(err)
				}
				delete(obj.Object, "_err")
				t.Run(obj.GetName(), func(t *testing.T) {
					want := TestExpectation{}
					if err := yaml.Unmarshal([]byte(item), &want); err != nil {
						t.Fatal(err)
					}
					res := v.ValidateCustomResource(obj)
					t.Log(res)
					if want.WantErr == "" {
						// Want no error
						if res != nil {
							t.Fatalf("configuration was invalid: %v", res)
						}
					} else {
						if res == nil {
							t.Fatalf("wanted error like %q, got none", want.WantErr)
						}
						if !strings.Contains(res.Error(), want.WantErr) {
							t.Fatalf("wanted error like %q, got %q", want.WantErr, res)
						}
					}
				})
			}
		})
	}
}
