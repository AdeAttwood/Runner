// Copyright 2022 Practically.io All rights reserved
//
// Use of this source is governed by a BSD-style
// licence that can be found in the LICENCE file or at
// https://www.practically.io/copyright
package runner

import (
	"errors"
	"os/exec"
	"strconv"

	"github.com/AdeAttwood/Runner/pkg/config"
)

type EventType int

const (
	Start EventType = 1
	End             = 2
)

type RunResult struct {
	Error  error
	Output string
}

type RunnerEvent struct {
	Type   EventType
	Task   *config.Task
	Result *RunResult
}

func buildScript(task config.Task) string {
	script := ""

	for _, task_command := range task.Commands {
		script += task_command.Command

		for _, argument := range task_command.Arguments {
			script += " " + argument
		}

		script += "; "
	}

	return script
}

func runTask(task config.Task) RunResult {
	var command *exec.Cmd
	script := buildScript(task)

	if len(task.Image) > 0 {
		command = exec.Command("docker", "run", "--rm", "/bin/bash", "-c", script)
	} else {
		command = exec.Command("/bin/bash", "-c", script)
	}

	output, err := command.CombinedOutput()
	if err == nil && command.ProcessState.ExitCode() != 0 {
		exit_code := strconv.Itoa(command.ProcessState.ExitCode())
		err = errors.New("Process exited with exit code '" + exit_code + "'")
	}

	return RunResult{Error: err, Output: string(output)}
}

// Runs a task
func Run(tasks []config.Task, handler func(RunnerEvent)) RunResult {
	var result RunResult
	for _, task := range tasks {
		handler(RunnerEvent{Type: Start, Task: &task, Result: &result})
		result = runTask(task)
		handler(RunnerEvent{Type: End, Task: &task, Result: &result})

		if result.Error != nil {
			return result
		}
	}

	return result
}
