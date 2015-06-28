package limitd

import (
	"fmt"
	"github.com/limitd/go-client/fixture"
	"github.com/limitd/go-client/messages"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//init test server
	server := fixture.Start()

	//run every test
	exitCode := m.Run()

	//stop test server
	fmt.Println("teardown")
	err := server.Process.Kill()
	server.Wait()
	if err != nil {
		panic(err)
	}

	//exit
	os.Exit(exitCode)
}

func TestTake(t *testing.T) {
	client, err := Dial(":9231")
	if err != nil {
		panic(err)
	}

	response, takeResponse, err := client.Take("ip", "127.0.0.1", 1)

	assert.Equal(t, limitd.Response_TAKE, response.GetType())
	assert.Equal(t, int32(9), takeResponse.GetRemaining())
	assert.Equal(t, true, takeResponse.GetConformant())
}
