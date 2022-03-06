// Copyright 2022 Practically.io All rights reserved
//
// Use of this source is governed by a BSD-style
// licence that can be found in the LICENCE file or at
// https://www.practically.io/copyright

{
  version: '0.0.1',
  tasks: {
    // Build tasks
    'build:local': {
      title: 'Build the Runner executable',
      commands: [
        { command: 'go', arguments: ['build', '-o', 'bin/run', '-i', 'github.com/AdeAttwood/Runner/cmd/run'] },
      ],
    },

    // Install tasks
    install: {
      title: 'Install the Runner executable',
      commands: [
        { command: 'go', arguments: ['install', '-i', 'github.com/AdeAttwood/Runner/cmd/run'] },
      ],
    },

    // Testing tasks
    'test:go': {
      title: 'Test all of the go code',
      commands: [
        { command: 'go', arguments: ['test', './...'] },
      ],
    },

    // Code formatting tasks
    format: {
      title: 'Formats all of the code',
      commands: [],
      requires: ['format:go', 'format:jsonnet'],
    },
    'format:go': {
      title: 'Formats all of the go code',
      commands: [
        { command: 'gofmt', arguments: ['-w', '.'] },
      ],
    },
    'format:jsonnet': {
      title: 'Formats all of the jsonnet code',
      commands: [
        { command: 'jsonnetfmt', arguments: ['-i', 'runner.jsonnet'] },
      ],
    },

    // Code linting tasks
    lint: {
      title: 'Lints all of the code',
      commands: [],
      requires: ['lint:go', 'lint:jsonnet'],
    },
    'lint:go': {
      title: 'Lints all of the go code',
      commands: [
        { command: 'test', arguments: ['-z', '"$(gofmt -d .)"'] },
      ],
    },
    'lint:jsonnet': {
      title: 'Lints all of the jsonnet code',
      commands: [
        { command: 'jsonnetfmt', arguments: ['--test', 'runner.jsonnet'] },
      ],
    },
  },
}
