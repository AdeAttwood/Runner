<div align="center">

# Runner

A configurable task runner

</div>

Runner is a task runner that allows you to configure your tasks with
[Jsonnet](https://jsonnet.org/). This empowers you to use proper templating of
task rather then writing Yaml.

## Configuration

Configuration is done by crating a `runner.jsonnet` file in your project. In this
file is where all of your task go.

```jsonnet
{
  version: '0.0.1',
  tasks: {
  list: {
    title: 'List all the files',
    commands: [
      { command: 'ls', arguments: ['-al'] },
    ],
  },
}
```

To run the list command you can run `run`

```shell
run list
```