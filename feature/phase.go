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

package feature

// Phase defines the relative maturity and support level of an Istio feature.
// Please note that the phases are applied to individual features, not to the
// project as a whole.
type Phase string

const (
	// Development indicates that the feature is still under active development
	// and is not publicly documented (pre-Alpha).
	//
	// API: No guarantees on backward compatibility
	//
	// Performance: Not quantified or guaranteed
	//
	// Deprecation Policy: None
	//
	// Security: NA
	Development Phase = "development"

	// Alpha indicates that the feature is demo-able, works end-to-end but
	// has limitations. If you use it in production and encounter a serious
	// issue we may not be able to fix it for you, so be sure that you can
	// continue to function if you have to disable it
	//
	// API: No guarantees on backward compatibility
	//
	// Performance: Not quantified or guaranteed
	//
	// Deprecation Policy: None
	//
	// Security: Security vulnerabilities will be handled publicly as simple
	// bug fixes
	Alpha Phase = "alpha"

	// Beta indicates that the feature is usable in production, not a toy anymore
	//
	// API: APIs are versioned
	//
	// Performance: Not quantified or guaranteed
	//
	// Deprecation Policy: Weak - 3 months
	//
	// Security: Security vulnerabilities will be handled according to our
	// security vulnerability policy
	Beta Phase = "beta"

	// Stable indicates that the feature is dependable, production hardened
	//
	// API: Dependable, production-worthy. APIs are versioned, with automated
	// version conversion for backward compatibility
	//
	// Performance: Performance (latency/scale) is quantified, documented,
	// with guarantees against regression
	//
	// Deprecation Policy: Dependable, Firm. 1 year notice will be provided
	// before changes
	//
	// Security: Security vulnerabilities will be handled according to our
	// security vulnerability policy
	Stable Phase = "stable"
)
