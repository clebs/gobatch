# Go Batch
Go-batch  allows to prepare batches of tasks that can be run concurrently and/or in sequence, creating any level of nesting and combination between the two modes. Creating a custom task is as easy as implementing the `Runner` interface.

## How it works:
Create a `SyncRunner` to run tasks in sequence or an `AsyncRunner` to run them concurrently. A task is any type implementing the `Runner` interface in this package.
Both `AsyncRunner` and `SyncRunner` implement the `Runner` interface, so they can be combined and nested.

## Example:
```go
package main

import (
	"os"
	"os/exec"

	"github.com/clebs/gobatch"
)

func main() {
	cmd1 := gobatch.CommandRunner{Command: exec.Command("ls"), Output: os.Stdout}
	cmd2 := gobatch.CommandRunner{Command: exec.Command("ps"), Output: os.Stdout}

	var asyncBatch gobatch.AsyncRunner
	asyncBatch.Add(cmd1, cmd2).Run() // Both commands will be run in separate goroutines concurrently

	var syncBatch gobatch.SyncRunner
	syncBatch.Add(cmd1, cmd2).Run() // Both commands will be run sequentially

	cmd3 := gobatch.CommandRunner{Command: exec.Command("whoami"), Output: os.Stdout}
	var otherAsyncBatch gobatch.AsyncRunner
	otherAsyncBatch.Add(cmd3, syncBatch).Run() // cmd3 will run concurrently to cmd1 and cmd2 which will run sequentially
}
```