# go-signals
Bind func to signal

```go
package main

import "github.com/9glt/go-signals"
import "runtime"

func main() {
    signals.USR1(func() {
        println("reload config")
    })
    runtime.Goexit()
}
```

```bash
pkill -usr1 app
```