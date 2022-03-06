// Copyright 2022 Practically.io All rights reserved
//
// Use of this source is governed by a BSD-style
// licence that can be found in the LICENCE file or at
// https://www.practically.io/copyright
package main

import (
	"flag"
	"log"
	"os"

	"github.com/AdeAttwood/Runner/pkg/commands"
	"github.com/AdeAttwood/Runner/pkg/config"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	application := commands.Application{}
	flag.StringVar(&application.Cwd, "cwd", cwd, "")
	flag.BoolVar(&application.Help, "help", false, "")
	flag.Usage = func() {
		commands.Help(application, config.Config{})
	}

	flag.Parse()
	c, err := config.Get(application.Cwd + "/runner.jsonnet")
	if err != nil {
		log.Fatal(err)
	}

	if flag.NArg() == 0 || flag.Arg(0) == "help" {
		commands.Help(application, *c)
		return
	}

	application.Arguments = flag.Args()[1:]
	switch flag.Arg(0) {
	case "run":
		err = commands.Run(application, *c)
	case "list":
		err = commands.List(application, *c)
	default:
		application.Arguments = flag.Args()
		err = commands.Run(application, *c)
	}

	exit_code := 0
	if err != nil {
		exit_code = 1
	}

	os.Exit(exit_code)
}
