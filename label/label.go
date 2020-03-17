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
