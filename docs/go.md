# Instructions for the Go Cook

## benchcmp
Displays performance changes between benchmarks

## Channels
* Closing a channel indicates that no more values will be sent on it.

### Links
* https://gobyexample.com/closing-channels

## Defer
If you have a chain of method calls in a defer statement all of them will be
executed but the last prior to the execution of the defer statement:

```
type P struct{}
func (p P) Second() P { fmt.Println("Second"); return p }
func (p P) Third() P { fmt.Println("Third"); return p }
func (p P) Last() { fmt.Println("Last") }

func main() {
    p := P{}
    fmt.Println("First")
    defer p.Second().Third().Last()
    fmt.Println("Fourth")
}
```

## Escape Analysis
* Unlike in C, it's perfectly OK to return the address of a local variable; the
storage associated with the variable survives after the function returns.

## gcflags
To get escape analysis and inlining information:

```
go build -gcflags -m
```

## GODEBUG
To get goroutine tracing, set `GODEBUG` like this:

```
GODEBUG=schedtrace=1000,scheddetail=1 ./main
```

## Interfaces
* Getters as interface methods are a thing in Go
    * pkg: database/sql, type: [Result](https://golang.org/pkg/database/sql/#Result)
    * pkg: net, type: [Conn](https://golang.org/pkg/net/#Conn)

## Links
* http://peter.bourgon.org/go-in-production/

## Projects
* https://github.com/kubernetes/kubernetes/tree/master/cmd
* https://github.com/camlistore/camlistore/tree/master/cmd
* https://github.com/hashicorp/consul
* https://github.com/coreos/etcd
