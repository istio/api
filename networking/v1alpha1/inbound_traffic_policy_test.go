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

package v1alpha1_test

import (
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"

	api "istio.io/api/networking/v1alpha1"
	outbound "istio.io/api/networking/v1alpha3"
)

func TestCompressionRoundTrip(t *testing.T) {
	inbound := &api.InboundTrafficPolicy{Http: &api.HttpSettings{Compression: &api.CompressionSettings{
		RemoveAcceptEncodingHeader: wrapperspb.Bool(false),
		Compressors: []*api.Compressor{
			{MinContentLength: wrapperspb.UInt32(0), CompressorConfig: &api.Compressor_Brotli{Brotli: &api.BrotliCompressor{Quality: wrapperspb.UInt32(0)}}},
			{CompressorConfig: &api.Compressor_Gzip{Gzip: &api.GzipCompressor{}}},
			{CompressorConfig: &api.Compressor_Zstd{Zstd: &api.ZstdCompressor{}}},
		},
	}}}
	for _, enabled := range []*wrapperspb.BoolValue{nil, wrapperspb.Bool(false), wrapperspb.Bool(true)} {
		dr := &outbound.DestinationRule{Host: "reviews", TrafficPolicy: &outbound.TrafficPolicy{
			Compression: &outbound.ClientCompressionSettings{Response: &outbound.ClientCompressionSettings_Response{
				Decompress: &outbound.ClientCompressionSettings_Decompression{Enabled: enabled},
			}},
		}}
		for name, original := range map[string]proto.Message{"inbound": inbound, "outbound": dr} {
			t.Run(name, func(t *testing.T) {
				codecs := []struct {
					name      string
					marshal   func(proto.Message) ([]byte, error)
					unmarshal func([]byte, proto.Message) error
				}{
					{"protobuf", proto.Marshal, proto.Unmarshal},
					{"protojson", protojson.Marshal, protojson.Unmarshal},
					{"jsonshim", func(m proto.Message) ([]byte, error) { return json.Marshal(m) }, func(b []byte, m proto.Message) error { return json.Unmarshal(b, m) }},
				}
				for _, codec := range codecs {
					t.Run(codec.name, func(t *testing.T) {
						data, err := codec.marshal(original)
						if err != nil {
							t.Fatal(err)
						}
						decoded := original.ProtoReflect().New().Interface()
						if err := codec.unmarshal(data, decoded); err != nil {
							t.Fatal(err)
						}
						if !proto.Equal(original, decoded) {
							t.Fatalf("round trip changed presence, order, or codec: %s", data)
						}
					})
				}
			})
		}
	}
	// Deep copies must not share codec tuning, wrapper values, or list entries.
	copied := inbound.DeepCopy()
	copied.Http.Compression.Compressors[0].GetBrotli().Quality.Value = 5
	copied.Http.Compression.Compressors[0].MinContentLength.Value = 1024
	copied.Http.Compression.RemoveAcceptEncodingHeader.Value = true
	if inbound.Http.Compression.Compressors[0].GetBrotli().Quality.Value != 0 ||
		inbound.Http.Compression.Compressors[0].MinContentLength.Value != 0 ||
		inbound.Http.Compression.RemoveAcceptEncodingHeader.Value {
		t.Fatal("deep copy aliases original compression settings")
	}
}
