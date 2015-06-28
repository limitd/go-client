package fixture

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func createWorkingDir() (name string, err error) {
	name, err = ioutil.TempDir("/tmp", "go-limitd-test")
	if err != nil {
		return
	}

	err = os.Mkdir(name+"/db", 0777)
	if err != nil {
		return
	}

	config := `
log_level: debug

buckets:
  ip:
    per_minute: 10
	`
	err = ioutil.WriteFile(name+"/config.yml", []byte(config), 0644)
	if err != nil {
		panic(err)
	}
	return
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//Start the test server
func Start() (cmd *exec.Cmd) {
	wdir, err := createWorkingDir()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("limitd",
		"--config-file",
		wdir+"/config.yml",
		"--db",
		wdir+"/db")

	// Create stdout, stderr streams of type io.Reader
	stdout, err := cmd.StdoutPipe()
	checkError(err)
	stderr, err := cmd.StderrPipe()
	checkError(err)

	// Start command
	err = cmd.Start()
	checkError(err)

	// Non-blockingly echo command output to terminal
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	return
}
