package serialization

import (
	"io/ioutil"
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	"istio.io/api/pkg/kube/apis/networking/v1alpha3"
)

func TestDeserializeVirtualService(t *testing.T) {
	scheme := runtime.NewScheme()
	codec := serializer.NewCodecFactory(scheme)
	v1alpha3.AddToScheme(scheme)

	data, err := ioutil.ReadFile("../testdata/virtual_service_1.yaml")
	if err != nil {
		t.Fatalf("failed to read test data: %v", err)
	}

	obj, err := runtime.Decode(codec.UniversalDecoder(v1alpha3.SchemeGroupVersion), data)
	if err != nil {
		t.Fatalf("failed to decode data into VirtualService: %v", err)
	}

	vs := obj.(*v1alpha3.VirtualService)
	if vs.GetNamespace() != "somenamespace" {
		t.Errorf("namespace doesn't match somenamespace: %s", vs.GetNamespace())
	}
	if len(vs.Spec.Hosts) != 1 || vs.Spec.Hosts[0] != "istio-pilot.somenamespace.svc.cluster.local" {
		t.Errorf("failed to read Hosts: %v", vs.Spec.Hosts)
	}
	if len(vs.Spec.Gateways) != 1 || vs.Spec.Gateways[0] != "meshexpansion-ilb-gateway" {
		t.Errorf("failed to read Gateways: %v", vs.Spec.Gateways)
	}
	if len(vs.Spec.Tcp) != 3 {
		t.Fatalf("Failed to read Tcp: %v", vs.Spec.Tcp)
	}
	if vs.Spec.Tcp[0].Route[0].Destination.Host != "istio-pilot.somenamespace.svc.cluster.local" {
		t.Errorf("failed to read Destination.Host: %v", vs.Spec.Tcp[0].Route[0].Destination.Host)
	}
	if vs.Spec.Tcp[0].Route[0].Destination.Port.GetNumber() != 15011 {
		t.Errorf("failed to read Tcp.Rout.Destination.Port.Number: %v", vs.Spec.Tcp[0].Route[0].Destination.Port.GetNumber())
	}
}
