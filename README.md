# WatchDock

WatchDock is an open-source project built around the idea of managing **Docker containers** running on **individual** hosts.


## Features 
Below is a list of features in WatchDock upcoming releases:
### v0.1.0

- [ ] Get all Docker containers from individual hosts
- [ ] Get stats from a Docker host (`docker stats`)
- [ ] Add a Docker Registry 
- [ ] Perform containers operations
  - [ ] Start
  - [ ] Stop
  - [ ] Remove
- [ ] Perform Docker Image operations
  - [ ] List images
  - [ ] List image tags

### v0.2.0
  - [ ] Remove a tag
  - [ ] Add new tag
  - [ ] Delete an image
<br />

## Generating Source Code

### Get `protoc` compiler
Download the `protoc` compiler from Protocol Buffers [release page](https://github.com/protocolbuffers/protobuf/releases).

### Install `gRPC` tooling

```
go get -u google.golang.org/grpc
```

### Install Protocol Buffers tooling 

```
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go

```

### Compile PB Messages -> Go Source Code
```
protoc 
    -I ./pb-messages
    ./pb-messages/*.proto
    --go_out=plugins=grpc:./pb
```
