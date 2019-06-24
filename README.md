# legs-message
Library for define legs protocol messages for legs-server and legs-client


## Basic Usage

### install
```bash
go get github.com/iij/legs-message
```
and
```go
import "github.com/iij/legs-message"
```

### create message
```go
msg := message.NewConsoleStartMessage("session-id", "shell cmd")
```

### Encode to msgpack
```go
b, err := message.Marshal(msg)
```

### Decode from msgpack
```go
msg := &message.ConsoleMessage{}
err = message.Unmarshal(b, msg)
```

## TODOs

- CI (testing, formatting)
