# Fluxion

> Hierarchical Memory Graph Database 🕶️

![img/fluxion-banner.jpg](img/fluxion-banner.jpg)

Fluxion is the [flux-sched](https://github.com/flux-framework/flux-sched) project packaged and provided as a containerized graph database. You can interact with it using any language that can interact with the gRPC endpoints. Example and client libraries are provided here.


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

### Go Example

In another terminal you can try one of the client examples in [examples](examples). For example:

```bash
go run examples/go/example.go --jobspec ./examples/go/jobspec.yaml
```
```console
$ go run examples/go/example.go --jobspec ./examples/go/jobspec.yaml
🦩️ This is the fluxion graph client
2024/05/03 13:27:13 🦩️ starting client (localhost:4242)...
2024/05/03 13:27:13 ⭐️ Init cluster status: INIT_SUCCESS
2024/05/03 13:27:13 ⭐️ Match status: MATCH_SUCCESS
2024/05/03 13:27:13      Allocation: {"graph": {"nodes": [{"id": "8", "metadata": {"type": "core", "basename": "core", "name": "core0", "id": 0, "uniq_id": 8, "rank": -1, "exclusive": true, "unit": "", "size": 1, "paths": {"containment": "/tiny0/rack0/node0/socket0/core0"}}}, {"id": "4", "metadata": {"type": "socket", "basename": "socket", "name": "socket0", "id": 0, "uniq_id": 4, "rank": -1, "exclusive": true, "unit": "", "size": 1, "paths": {"containment": "/tiny0/rack0/node0/socket0"}}}, {"id": "2", "metadata": {"type": "node", "basename": "node", "name": "node0", "id": 0, "uniq_id": 2, "rank": -1, "exclusive": false, "unit": "", "size": 1, "paths": {"containment": "/tiny0/rack0/node0"}}}, {"id": "1", "metadata": {"type": "rack", "basename": "rack", "name": "rack0", "id": 0, "uniq_id": 1, "rank": -1, "exclusive": false, "unit": "", "size": 1, "paths": {"containment": "/tiny0/rack0"}}}, {"id": "0", "metadata": {"type": "cluster", "basename": "tiny", "name": "tiny0", "id": 0, "uniq_id": 0, "rank": -1, "exclusive": false, "unit": "", "size": 1, "paths": {"containment": "/tiny0"}}}], "edges": [{"source": "4", "target": "8", "metadata": {"name": {"containment": "contains"}}}, {"source": "2", "target": "4", "metadata": {"name": {"containment": "contains"}}}, {"source": "1", "target": "2", "metadata": {"name": {"containment": "contains"}}}, {"source": "0", "target": "1", "metadata": {"name": {"containment": "contains"}}}]}}
2024/05/03 13:27:13        Overhead: 0.000818
2024/05/03 13:27:13        Reserved: false
2024/05/03 13:27:13           Jobid: 1
2024/05/03 13:27:13              At: 0
😴️ Sleeping for 3 seconds before cancel...
2024/05/03 13:27:18 ⭐️ Cancel status: CANCEL_SUCCESS
```

In the server terminal you will see:

```console
GOOS=linux CGO_CFLAGS="-I/opt/flux-sched -I/opt/flux-sched/resource/reapi/bindings/c" CGO_LDFLAGS="-L/usr/lib -L/usr/lib/flux -L/opt/flux-sched/resource/reapi/bindings -lreapi_cli -lflux-idset -lstdc++ -lczmq -ljansson -lhwloc -lboost_system -lflux-hostlist -lboost_graph -lyaml-cpp" go build -ldflags '-w' -o bin/server cmd/main.go
🦩️ This is the fluxion graph server
[GRPCServer] gRPC Listening on [::]:4242
[Fluxion] Created flux memory graph[Fluxion] match policy: {"matcher_policy": "lonode"}[Fluxion] There are no errors

💼️ Allocation Result
       Overhead: 0.000602
       Reserved: false
         Status: MATCH_SUCCESS
          Jobid: 1
             At: 0
     Allocation: {"graph": {"nodes": [{"id": "8", "metadata": {"type": "core", "basename": "core", "name": "core0", "id": 0, "uniq_id": 8, "rank": -1, "exclusive": true, "unit": "", "size": 1, "paths": {"containment": "/tiny0/rack0/node0/socket0/core0"}}}, {"id": "4", "metadata": {"type": "socket", "basename": "socket", "name": "socket0", "id": 0, "uniq_id": 4, "rank": -1, "exclusive": true, "unit": "", "size": 1, "paths": {"containment": "/tiny0/rack0/node0/socket0"}}}, {"id": "2", "metadata": {"type": "node", "basename": "node", "name": "node0", "id": 0, "uniq_id": 2, "rank": -1, "exclusive": false, "unit": "", "size": 1, "paths": {"containment": "/tiny0/rack0/node0"}}}, {"id": "1", "metadata": {"type": "rack", "basename": "rack", "name": "rack0", "id": 0, "uniq_id": 1, "rank": -1, "exclusive": false, "unit": "", "size": 1, "paths": {"containment": "/tiny0/rack0"}}}, {"id": "0", "metadata": {"type": "cluster", "basename": "tiny", "name": "tiny0", "id": 0, "uniq_id": 0, "rank": -1, "exclusive": false, "unit": "", "size": 1, "paths": {"containment": "/tiny0"}}}], "edges": [{"source": "4", "target": "8", "metadata": {"name": {"containment": "contains"}}}, {"source": "2", "target": "4", "metadata": {"name": {"containment": "contains"}}}, {"source": "1", "target": "2", "metadata": {"name": {"containment": "contains"}}}, {"source": "0", "target": "1", "metadata": {"name": {"containment": "contains"}}}]}}

[Fluxion] received cancel request jobID:1
```

And that's it!

### Container Example

As an alternative, you can run the container service instead.

```bash
docker run -p 4242:4242 ghcr.io/converged-computing/fluxion  --host 0.0.0.0
```

And from a different terminal (make sure you've closed VSCode which sometimes can claim ports).

```bash
go run examples/go/example.go  --jobspec examples/go/jobspec.yaml
```

And you will get the same response as above.

### Python Example

Finally, we have the same gRPC endpoints exposed via Python. With the server running:

```bash
make server && ./bin/server
```

Run the Python example:

```bash
python ./examples/python/example.py
```

Note that the protocol buffers are built (or updated) for Python like:

```bash
make python
```

You shouldn't need to run this command unless you change something.

## TODO

- get grow "unpack" bindings into fluxion-go and add update here
- when flux-sched supports JGF v2, upgrade here.

## Thank you

This code is based off of the work done for [fluence](https://github.com/flux-framework/flux-k8s) by the same authors.
The code base has changed significantly with development.

## License

Release

SPDX-License-Identifier: Apache-2.0

LLNL-CODE-764420
