// Copyright 2022 Practically.io All rights reserved
//
// Use of this source is governed by a BSD-style
// licence that can be found in the LICENCE file or at
// https://www.practically.io/copyright
package config

import (
	"testing"
)

func makeTasks() map[string]Task {
	return map[string]Task{
		"task": {},
		"task:one": {
			Requires: []string{"task"},
		},
		"task:two": {
			Requires: []string{"task", "task:one"},
		},
		"task:three": {
			Requires: []string{"task:four"},
		},
		"task:four": {
			Requires: []string{"task:three"},
		},
	}
}

// Test you can get tasks from the configuration
func TestGettingTasks(t *testing.T) {
	var err error
	config := Config{
		Version: "0.0.1",
		Tasks:   makeTasks(),
	}

	_, err = config.FindTask("task")
	if err != nil {
		t.Fatal(err)
	}

	_, err = config.FindTask("not a task")
	if err == nil {
		t.Fatal(err)
	}
}

// Build a task list with a single task to run
func TestBuildTaskSingle(t *testing.T) {
	var err error
	config := Config{
		Version: "0.0.1",
		Tasks:   makeTasks(),
	}

	task_list, err := config.BuildTaskList("task")
	if err != nil {
		t.Fatal(err)
	}

	if len(task_list) != 1 {
		t.Fatal("Task list should only have one task in it")
	}
}

// Build a task list with multiple tasks
func TestBuildTaskMultiple(t *testing.T) {
	var err error
	config := Config{
		Version: "0.0.1",
		Tasks:   makeTasks(),
	}

	task_list, err := config.BuildTaskList("task:two")
	if err != nil {
		t.Fatal(err)
	}

	if len(task_list) != 3 {
		t.Fatal("Task list should only have one task in it")
	}
}
