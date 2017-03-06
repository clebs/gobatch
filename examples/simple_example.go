package main

import (
	"os"
	"os/exec"

	"github.com/clebs/gobatch"
)

func main() {
	cmd1 := gobatch.CommandRunner{Command: exec.Command("ls"), Output: os.Stdout}
	cmd2 := gobatch.CommandRunner{Command: exec.Command("ps"), Output: os.Stdout}
	var batch gobatch.AsyncRunner
	batch.Add(cmd1, cmd2).Run()
}
