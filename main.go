package main

import (
	"os"
	"os/exec"
	"regexp"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile("^golang\\.org/x/(.+)$")
}

func main() {
	cmdSlice := makeCmd(os.Args)
	cmd := exec.Command("go", cmdSlice...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	os.Exit(0)
}

func replace(origin string) string {
	return re.ReplaceAllString(origin, "github.com/golang/$1")
}

func makeCmd(args []string) []string {
	cmd := make([]string, len(args))
	cmd[0] = "get"

	for i := 1; i < len(args); i++ {
		cmd[i] = replace(args[i])
	}

	return cmd
}
