Nomad Dotnet Driver Plugin
==========

Nomad Task Driver for Dotnet Core
[Nomad task driver plugins](https://www.nomadproject.io/docs/drivers/index.html).

This project is the counterpart of [Nomad Java Driver](https://developer.hashicorp.com/nomad/docs/drivers/java) 
for Dotnet


- Website: [https://www.nomadproject.io](https://www.nomadproject.io)
- Mailing list: [Google Groups](http://groups.google.com/group/nomad-tool)

Requirements
-------------------

- [Go](https://golang.org/doc/install) v1.18 or later (to compile the plugin)
- [Nomad](https://www.nomadproject.io/downloads.html) v0.9+ (to run the plugin)


## Deploying Driver Plugins in Nomad

The initial version of the skeleton is a simple task that outputs a greeting.
You can try it out by starting a Nomad agent and running the job provided in
the `example` folder:

```sh
$ make build
$ nomad agent -dev -config=./example/agent.hcl -plugin-dir=$(pwd)

# in another shell
$ nomad run ./example/example.nomad
$ nomad logs <ALLOCATION ID>
```
## Dotnet Job Configuration

As you may know, we can use nomad to run task without traditional containers (Docker, OCI...) and instead
use OS-level or Runtime-level isolations. CLR as dotnet runtime, facilitates some isolations, 
for example thread limitations and GC memory restrictions. You can use both isolations within
Nomad job configurations.
The dotnet command line configuration for running the process is ported here in the 3 category of
Gc config, Threading config and globalization config.