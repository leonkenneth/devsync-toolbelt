package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"
)

func runOriginal() {
	commandPath := path.Join(homedir(), "/.herodev/", "heroku-orig")
	run(commandPath, os.Args[1:])
}

func fileExists(file string) bool {
	_, err := os.Stat(file)

	return !os.IsNotExist(err)
}

func main() {
	if strings.HasPrefix(os.Args[1], "dev:") {
		subcommand := strings.Replace(os.Args[1], "dev:", "", 1)

		commandPath := path.Join(homedir(), "/.herodev/", "herodev-"+subcommand)
		if fileExists(commandPath) {
			run(commandPath, os.Args[2:])
		} else {
			runOriginal()
		}
	} else {
		runOriginal()
	}
}

func homedir() string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	return dir
}

func run(cmdPath string, args []string) {
	cmd := exec.Command(cmdPath, args...)
	cmd.Env = os.Environ()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	err = cmd.Wait()

	// FIXME: Error handling
	if err != nil {
		fmt.Printf("")
	}
}
