package util

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

// DeepCopyInto is a helper function for implementing DeepCopy() functions for generated protobuf types
func DeepCopyInto(in proto.Marshaler, out proto.Unmarshaler) {
	if reflect.TypeOf(in) != reflect.TypeOf(out) {
		panic(fmt.Errorf("Types are not the same, cannot DeepCopyInto. %s != %s", reflect.TypeOf(in), reflect.TypeOf(out)))
	}
	dAtA, err := in.Marshal()
	if err != nil {
		panic(err)
	}
	err = out.Unmarshal(dAtA)
	if err != nil {
		panic(err)
	}
}
