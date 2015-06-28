package limitd

import (
	"fmt"
	"github.com/limitd/go-client/fixture"
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

func TestPut(t *testing.T) {
	client, err := Dial(":9231")
	if err != nil {
		panic(err)
	}
	client.Take("ip", "127.0.0.1", 50)
}
