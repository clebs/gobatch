//Package gobatch provides a set of types that allow to run tasks concurrently and sequentially. Being able to have any level of nesting and combination
package gobatch

import (
	"fmt"
	"io"
	"os/exec"
	"sync"
)

//Runner is any type that can execute logic on demand via the Run method
// Specific implementations may document their own behavior.
type Runner interface {
	Run()
}

// AsyncRunner runs a set of tasks concurrently
type AsyncRunner []Runner

// Add adds a new set of Runners at the end od the Runner slice
func (ab *AsyncRunner) Add(rs ...Runner) *AsyncRunner {
	*ab = append(*ab, rs...)
	return ab
}

// Run executes all tasks concurrently
func (ab AsyncRunner) Run() {
	var wg sync.WaitGroup
	wg.Add(len(ab))

	for _, task := range ab {
		go WaitRunner{Runner: task, Waiter: &wg}.Run()
	}
	wg.Wait()
}

// WaitRunner is a Runner that has a WaitGroup which is notified when the Run method finishes.
type WaitRunner struct {
	Runner Runner
	Waiter *sync.WaitGroup
}

// Run executes the Runner and notifies the WaitGroup when finished.
func (wr WaitRunner) Run() {
	wr.Runner.Run()
	defer wr.Waiter.Done()
}

// SyncRunner runs a set of tasks in sequence
type SyncRunner []Runner

// Add adds a new set of Runners at the end od the Runner slice
func (sb *SyncRunner) Add(rs ...Runner) *SyncRunner {
	*sb = append(*sb, rs...)
	return sb
}

// Run executes all tasks sequentially
func (sb SyncRunner) Run() {
	for _, task := range sb {
		task.Run()
	}
}

//CommandRunner executes a command and writes its output to the Writer.
type CommandRunner struct {
	Command *exec.Cmd
	Output  io.Writer
}

// Run executes the command and sends output (stdout and stderr) to the channel
func (cr CommandRunner) Run() {
	out, err := cr.Command.CombinedOutput()
	if err != nil {
		out = []byte(err.Error())
	}
	_, err = cr.Output.Write(out)
	if err != nil {
		panic(fmt.Sprintf("Error writing command output: %s", err.Error()))
	}
}
