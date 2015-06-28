// Package stringutil contains utility functions for working with strings.
package limitd

import (
	"bytes"
	"github.com/dchest/uniuri"
	"github.com/golang/protobuf/proto"
	"github.com/limitd/go-client/messages"
	"io"
	"log"
	"net"
)

//Client defines the limitd client
type Client struct {
	Conn            net.Conn
	PendingRequests map[string]chan<- *limitd.Response
}

// Take n tokens from bucket t, key k
func (client *Client) Take(t string, k string, n int32) (response *limitd.Response, takeResponse *limitd.TakeResponse, err error) {
	requestID := uniuri.New()

	request := &limitd.Request{
		Id:     proto.String(requestID),
		Method: limitd.Request_TAKE.Enum(),
		Type:   proto.String(t),
		Key:    proto.String(k),
		Count:  proto.Int32(n),
	}

	// goprotobuf.EncodeVarint followed by proto.Marshal
	responseChan := make(chan *limitd.Response)
	client.PendingRequests[requestID] = responseChan

	data, _ := proto.Marshal(request)
	data = append(proto.EncodeVarint(uint64(len(data))), data...)
	client.Conn.Write(data)

	response = <-responseChan
	takeR, err := proto.GetExtension(response, limitd.E_TakeResponse_Response)
	if err != nil {
		return
	}

	if takeResponseCasted, ok := takeR.(*limitd.TakeResponse); ok {
		takeResponse = takeResponseCasted
	}

	return
}

func (client *Client) listen() {
	buffer := new(bytes.Buffer)
	data := make([]byte, 8192)

	for {
		n, err := client.Conn.Read(data)

		if err != nil {
			if err == io.EOF {
				log.Println("Client exited")
				return
			}
			log.Println(err.Error())
		}

		buffer.Write(data[0:n])

		messageLength, bytesRead := proto.DecodeVarint(buffer.Bytes())

		if messageLength > 0 && messageLength < uint64(buffer.Len()) {
			response := &limitd.Response{}

			err = proto.Unmarshal(buffer.Bytes()[bytesRead:messageLength+uint64(bytesRead)], response)
			if err != nil {
				log.Printf("Failed to read proto response: %s\n", err.Error())
			} else {
				responseChannel := client.PendingRequests[*response.RequestId]
				responseChannel <- response
				delete(client.PendingRequests, *response.RequestId)
				buffer.Reset()
			}
		}
	}
}

// Dial connect to a limitd server
func Dial(address string) (client *Client, err error) {

	conn, err := net.Dial("tcp", address)

	client = new(Client)
	client.Conn = conn
	client.PendingRequests = make(map[string]chan<- *limitd.Response)

	go client.listen()
	return
}
