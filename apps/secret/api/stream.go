package api

import (
	"github.com/gorilla/websocket"
	"github.com/qiaogy91/cmdb/apps/resource"
	"google.golang.org/grpc"
)

type Stream struct {
	grpc.ServerStream
	conn *websocket.Conn
}

func (s *Stream) Send(rsp *resource.Resource) error {
	return s.conn.WriteJSON(rsp)
}

func NewStream(c *websocket.Conn) *Stream {
	return &Stream{conn: c}
}
