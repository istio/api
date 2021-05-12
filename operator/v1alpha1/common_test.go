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

package v1alpha1

import (
	"fmt"
	"testing"

	protobuf "github.com/gogo/protobuf/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func TestIntOrString_UnmarshalJSONPB(t *testing.T) {
	tests := []struct {
		in         []byte
		wantInt    bool
		wantString bool
		wantErr    bool
	}{
		{in: []byte(`"imastring"`), wantString: true},
		{in: []byte(`imalsoastring`), wantString: true},
		{in: []byte(`50imasneakystring`), wantString: true},
		{in: []byte(`50`), wantInt: true},
		// yaml spec defines this as a string so this is the behavior we want
		{in: []byte(`"50"`), wantString: true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s isString %v isInt %v", tt.in, tt.wantString, tt.wantInt), func(t *testing.T) {
			intorstring := &IntOrStringForPB{}
			if err := intorstring.UnmarshalJSONPB(nil, tt.in); (err != nil) != tt.wantErr {
				t.Errorf("IntOrString.UnmarshalJSONPB() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantInt && intorstring.Type == intstr.String {
				t.Error("wanted int got string")
			}
			if tt.wantString && intorstring.Type == intstr.Int {
				t.Error("wanted string got int")
			}
		})
	}
}

func TestBoolValueForPB_GetValue(t *testing.T) {
	var test *BoolValueForPB
	t.Run("nil value", func(t *testing.T) {
		if test.GetValue() != false {
			t.Fatalf("value of nil BoolValueForPB should be false")
		}
	})
	t.Run("false value", func(t *testing.T) {
		test = &BoolValueForPB{protobuf.BoolValue{Value: false}}
		if test.GetValue() != false {
			t.Fatalf("value of false BoolValueForPB should be false")
		}
	})
	t.Run("true value", func(t *testing.T) {
		test = &BoolValueForPB{protobuf.BoolValue{Value: false}}
		if test.GetValue() != false {
			t.Fatalf("value of true BoolValueForPB should be true")
		}
	})
}
