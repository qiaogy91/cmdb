package impl_test

import (
	"fmt"
	"github.com/qiaogy91/cmdb/apps/resource"
	"google.golang.org/grpc"
)

type Stream struct {
	grpc.ServerStream
}

func (s *Stream) Send(rsp *resource.Resource) error {
	fmt.Printf("@@@stream response: %+v\n", rsp)
	return nil
}

func NewStream() *Stream {
	return &Stream{}
}
