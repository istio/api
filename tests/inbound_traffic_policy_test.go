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
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestInboundTrafficPolicyBounds(t *testing.T) {
	v := NewIstioValidator(t)
	diagnostics := make([]any, 16)
	for i := range diagnostics {
		diagnostics[i] = map[string]any{"reason": "MoreSpecificPolicy", "message": strings.Repeat("x", 1024)}
	}
	for _, tt := range []struct {
		name    string
		status  map[string]any
		wantErr string
	}{
		{name: "maximum-diagnostics", status: map[string]any{"diagnostics": diagnostics}},
		{name: "too-many-diagnostics", status: map[string]any{"diagnostics": append(append([]any{}, diagnostics...), diagnostics[0])}, wantErr: "at most 16"},
		{name: "oversized-diagnostic", status: map[string]any{"diagnostics": []any{map[string]any{"reason": "MoreSpecificPolicy", "message": strings.Repeat("x", 1025)}}}, wantErr: "more than 1024 bytes"},
		{name: "inconsistent-counts", status: map[string]any{"workloads": map[string]any{"total": int64(1), "selected": int64(2)}}, wantErr: "must sum to total"},
		{name: "negative-count", status: map[string]any{"workloads": map[string]any{"total": int64(-1)}}, wantErr: "greater than or equal to 0"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			obj := &unstructured.Unstructured{Object: map[string]any{
				"apiVersion": "networking.istio.io/v1alpha1", "kind": "InboundTrafficPolicy",
				"metadata": map[string]any{"name": tt.name}, "spec": map[string]any{},
			}}
			if tt.status != nil {
				obj.Object["status"] = tt.status
			}
			err := v.ValidateCustomResource(obj)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatal(err)
				}
			} else if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("wanted error containing %q, got %v", tt.wantErr, err)
			}
		})
	}
}
