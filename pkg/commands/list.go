// Copyright 2022 Practically.io All rights reserved
//
// Use of this source is governed by a BSD-style
// licence that can be found in the LICENCE file or at
// https://www.practically.io/copyright
package commands

import (
	"github.com/AdeAttwood/Runner/pkg/config"
	"github.com/AdeAttwood/Runner/pkg/console"
)

const COMMAND_LIST_PADDING = 1

// Gets the string length of longest task ID
func maxTaskLength(task_list map[string]config.Task) int {
	id_length := 0
	for id, _ := range task_list {
		if len(id) > id_length {
			id_length = len(id)
		}
	}

	return id_length
}

func List(application Application, config config.Config) error {
	padding := maxTaskLength(config.Tasks) + COMMAND_LIST_PADDING
	console.WriteLine("")
	console.WriteLine("Run Commands:")
	console.WriteLine("")

	for id, task := range config.Tasks {
		console.WriteLine("  %-*s%s", padding, id, task.Title)
	}

	return nil
}
