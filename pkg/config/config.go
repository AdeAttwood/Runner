// Copyright 2022 Practically.io All rights reserved
//
// Use of this source is governed by a BSD-style
// licence that can be found in the LICENCE file or at
// https://www.practically.io/copyright
package config

import (
	"encoding/json"
	"errors"

	"github.com/google/go-jsonnet"
)

type Command struct {
	// The excitable for this task
	Command string `json:"command"`
	// A list of arguments that will be passed to the executable
	Arguments []string `json:"arguments"`
}

// A single task definition
type Task struct {
	ID string
	// The title of the task. This is what will be displayed in output
	Title string `json:"title"`
	// The log description of this task
	Description string `json:"description"`
	// The container that the command will be run in
	Image string `json:"image"`
	// The excitable for this task
	Commands []Command `json:"commands"`
	// A list of task ids that this task depends on
	Requires []string `json:"requires"`
}

// The main configuration for all of the tasks
type Config struct {
	// The version number for this configuration. Currently this can only be
	// "0.0.1"
	Version string `json:"version"`
	// A map of tasks that are keyed by the task identifier
	Tasks map[string]Task `json:"tasks"`
}

// Finds a task by the identifier
func (config *Config) FindTask(id string) (Task, error) {
	task, found := config.Tasks[id]
	task.ID = id

	if found {
		return task, nil
	}

	return Task{}, errors.New("Task not found")
}

func (config *Config) BuildTaskList(id string) ([]Task, error) {
	task_list := []Task{}
	return config.buildTaskListInternal(id, task_list)
}

func isTaskInTaskList(id string, task_list []Task) bool {
	for _, task := range task_list {
		if task.ID == id {
			return true
		}
	}

	return false
}

func (config *Config) buildTaskListInternal(id string, task_list []Task) ([]Task, error) {
	task, err := config.FindTask(id)
	if err != nil {
		return task_list, err
	}

	if !isTaskInTaskList(id, task_list) {
		task_list = append([]Task{task}, task_list...)
	}

	for _, value := range task.Requires {
		task_list, err = config.buildTaskListInternal(value, task_list)
		if err != nil {
			return task_list, err
		}
	}

	return task_list, err
}

// Get and creates a new configuration struct from a jsonnet file path
func Get(file string) (*Config, error) {
	vm := jsonnet.MakeVM()
	config := &Config{}
	json_config, err := vm.EvaluateFile(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal([]byte(json_config), config)
	return config, err
}
