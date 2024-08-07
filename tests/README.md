# Validation testing

This set of packages tests validation of our CRDs across a range of Kubernetes versions.
One common issue with CEL validation is exceeding the cost window.
This logic changes with each Kubernetes version.
Additionally, new functions are added in various Kubernetes versions.

In these tests, we run the same logic with different Kubernetes versions.
This is done by a go module per version we care about.

The "base" package has the core logic, and the minimum version we test (it must be the oldest one to appease Go).
The other tests packages test versions we care to tests.
Its a good idea to always test the latest version.
