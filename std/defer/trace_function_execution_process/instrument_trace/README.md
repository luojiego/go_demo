# trace log injection tool
## cmd install
```
cd cmd/instrument/
go build
```
## usage

```
cd ../../
./cmd/instrument/instrument.exe example/demo/demo.go

```
output
```go
package main

import "trace"

func foo() {
        defer trace.Trace()()
        bar()
}

func bar() {
        defer trace.Trace()()

}

func main() {
        defer trace.Trace()()
        foo()
}

```
use -w option to rewrite file 
