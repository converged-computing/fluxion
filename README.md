# Fluxion

> Hierarchical Memory Graph Database

Fluxion is the [flux-sched](https://github.com/flux-framework/flux-sched) project packaged and provided as a containerized graph database. You can interact with it using any language that can interact with the gRPC endpoints. Example and client libraries are provided here.

🚧️ **under development** 🚧️

## Usage

### Build

If you have flux-sched locally you can build on your local machine, however it's recommended to use a VS Code developer environment with the included [.devcontainer](.devcontainer) directory. 

To build the container:

```bash
make build
```

If you are developing from inside the container:

```bash
make proto
go mod tidy
go mod vendor
```

And then build:

```bash
make server
```

### Run the Server

```bash
./bin/server
```
```console
This is the fluxion graph server
[GRPCServer] gRPC Listening on [::]:4242
```

In another terminal you can try one of the client examples in [examples](examples). For example:

```bash
go run examples/go/example.go
```

And that's it! 

## TODO

- create clients for Go/Python
- automated container build for server
- example that takes JGF v2 and calls init to the graph database endpoint
- example that does a satisfy / match allocate
- remove custom parsing of jobspec and use jobspec-go
- get grow "unpack" bindings into fluxion-go and add update here
- cute gopher logo!
- replace jobspec here with jobspec-go
- assess use cases for service grpc
- when flux-sched supports JGF v2, upgrade here.

## Thank you

This code is based off of the work done for [fluence](https://github.com/flux-framework/flux-k8s) by the same authors.
The code base will change significantly with development.

## License

Release

SPDX-License-Identifier: Apache-2.0

LLNL-CODE-764420
