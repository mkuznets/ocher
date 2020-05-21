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
	server := ocher.New("postgres://postgres@127.0.0.1:5432/postgres")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
```

## Client (Python)

```python
# encoding: utf-8 

import json
import ocher  # pip install ocher

client = ocher.Client('127.0.0.1:50051', 'client_id')

for tx, task in client.tasks(queue='ydl'):
    print('processing', task.id, task.queue, task.json())
    tx.finish(json.dumps({'result': 'ok'}).encode())
```
