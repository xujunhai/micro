package json

import (
	"bytes"
	"github.com/json-iterator/go"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"gitlab.ziroom.com/rent-web/micro/util/buf"
)

var Json = jsoniter.ConfigFastest
var jsonpbMarshaler = &jsonpb.Marshaler{}

// create buffer pool with 16 instances each preallocated with 256 bytes
var bufferPool = buf.NewPool()

type Marshaler struct{}

func (j Marshaler) Marshal(v interface{}) ([]byte, error) {
	if pb, ok := v.(proto.Message); ok {
		buf := bufferPool.Get()
		defer bufferPool.Put(buf)
		if err := jsonpbMarshaler.Marshal(buf, pb); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
	return Json.Marshal(v)
}

func (j Marshaler) Unmarshal(d []byte, v interface{}) error {
	if pb, ok := v.(proto.Message); ok {
		return jsonpb.Unmarshal(bytes.NewReader(d), pb)
	}
	return Json.Unmarshal(d, v)
}

func (j Marshaler) String() string {
	return "json"
}
