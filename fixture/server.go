package fixture

import (
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

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()

	time.Sleep(1000 * time.Millisecond)

	return
}
