package annotation

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

const (
	// OverrideAnnotation is used to store the overrides for injected containers
	OverrideAnnotation = "proxy.istio.io/overrides"

	// TemplatesAnnotation declares the set of templates to use for injection. If not specified, DefaultTemplates
	// will take precedence, which will inject a standard sidecar.
	// The format is a comma separated list. For example, `inject.istio.io/templates: sidecar,debug`.
	TemplatesAnnotation = "inject.istio.io/templates"
)
