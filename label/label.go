// Copyright 2020 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Package label contains common definitions for workload labels used by Istio.
package label

const (
	// TLSMode is the name of label given to service instances to determine whether to use mTLS or
	// fallback to plaintext/tls
	TLSMode = "security.istio.io/tlsMode"

	// IstioCanonicalServiceName is the name of label for the Istio Canonical Service for a workload instance.
	IstioCanonicalServiceName = "service.istio.io/canonical-name"

	// IstioCanonicalServiceRevision is the name of label for the Istio Canonical Service revision for a workload instance.
	IstioCanonicalServiceRevision = "service.istio.io/canonical-revision"
)
