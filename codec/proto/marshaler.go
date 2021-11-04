package proto

import (
	"bytes"

	"github.com/golang/protobuf/proto"
	"gitlab.ziroom.com/rent-web/micro/codec"
	"gitlab.ziroom.com/rent-web/micro/util/buf"
)

// create buffer pool with 16 instances each preallocated with 256 bytes
var bufferPool = buf.NewPool()

type Marshaler struct{}

func (Marshaler) Marshal(v interface{}) ([]byte, error) {
	pb, ok := v.(proto.Message)
	if !ok {
		return nil, codec.ErrInvalidMessage
	}

	// looks not good, but allows to reuse underlining bytes
	buf := bufferPool.Get()
	pbuf := proto.NewBuffer(buf.Bytes())
	defer func() {
		bufferPool.Put(bytes.NewBuffer(pbuf.Bytes()))
	}()

	if err := pbuf.Marshal(pb); err != nil {
		return nil, err
	}

	return pbuf.Bytes(), nil
}

func (Marshaler) Unmarshal(data []byte, v interface{}) error {
	pb, ok := v.(proto.Message)
	if !ok {
		return codec.ErrInvalidMessage
	}

	return proto.Unmarshal(data, pb)
}

func (Marshaler) String() string {
	return "proto"
}
