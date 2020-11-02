# ocher

Ocher is a PostgreSQL-based task-queue.

# Usage

## Server

```go
package main

import (
	"net"
	"mkuznets.com/go/ocher"
)

func main() {
    server := ocher.NewServer("postgres://postgres@127.0.0.1:5432/postgres")
    
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        panic(err)
    }
    if err := server.Serve(lis); err != nil {
        panic(err)
    }
}
```

## Worker

```go
package main

import (
    "context"
    "mkuznets.com/go/ocher"
)

func getTheAnswer(ctx context.Context, task ocher.Task) ([]byte, error) {
    // some long processing...
    return []byte(`{"result": 42}`), nil
}

func main() {
    worker := ocher.NewWorker("127.0.0.1:50051")
    worker.RegisterTask("ultimate_question", getTheAnswer)
    if err := worker.Serve(); err != nil {
        panic(err)
    }
}
```
