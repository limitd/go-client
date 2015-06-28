// Package stringutil contains utility functions for working with strings.
package limitd

import (
	"github.com/dchest/uniuri"
	"github.com/golang/protobuf/proto"
	"github.com/limitd/go-client/messages"
	"net"
)

//Client defines the limitd client
type Client struct {
	Conn net.Conn
}

// Take n tokens from bucket t, key k
func (client *Client) Take(t string, k string, n int32) {
	request := &limitd.Request{
		Id:     proto.String(uniuri.New()),
		Method: limitd.Request_TAKE.Enum(),
		Type:   proto.String(t),
		Key:    proto.String(k),
		Count:  proto.Int32(12),
	}
	// goprotobuf.EncodeVarint followed by proto.Marshal
	data, _ := proto.Marshal(request)
	data = append(proto.EncodeVarint(uint64(len(data))), data...)
	client.Conn.Write(data)
	return
}

// Dial connect to a limitd server
func Dial(address string) (client *Client, err error) {

	conn, err := net.Dial("tcp", address)

	client = new(Client)
	client.Conn = conn

	return
}
