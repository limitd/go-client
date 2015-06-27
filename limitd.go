// Package stringutil contains utility functions for working with strings.
package limitd

import (
	"github.com/golang/protobuf/proto"
	"github.com/limitd/go-client/messages/Request"
	"log"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	test := &limitd.Request{
		Id:     proto.String("hello"),
		Method: limitd.Request_TAKE.Enum(),
		Type:   proto.String("ip"),
		Key:    proto.String("127.0.0.1"),
		Count:  proto.Int32(12),
	}

	data, err := proto.Marshal(test)

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(data)

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
