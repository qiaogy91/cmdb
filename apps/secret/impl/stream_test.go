package impl_test

import (
	"fmt"
	"github.com/qiaogy91/cmdb/apps/secret"
	"google.golang.org/grpc"
)

type Stream struct {
	grpc.ServerStream
}

func (s *Stream) Send(rsp *secret.SyncResourceResponse) error {
	fmt.Printf("@@@stream response: %+v\n", rsp)
	return nil
}

func NewStream() *Stream {
	return &Stream{}
}
