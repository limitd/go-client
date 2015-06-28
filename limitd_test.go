package limitd

import (
	"fmt"
	"github.com/limitd/go-client/fixture"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	server := fixture.Start()

	exitCode := m.Run()
	fmt.Println("teardown")
	err := server.Process.Kill()
	server.Wait()
	if err != nil {
		panic(err)
	}

	os.Exit(exitCode)
}

func TestPut(t *testing.T) {

	// cases := []struct {
	// 	in, want string
	// }{
	// 	{"Hello, world", "dlrow ,olleH"},
	// 	{"Hello, 世界", "界世 ,olleH"},
	// 	{"", ""},
	// }
	// for _, c := range cases {
	// 	got := Reverse(c.in)
	// 	if got != c.want {
	// 		t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
	// 	}
	// }
}
