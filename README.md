# speedtest

a library that tests download and upload speeds

usage:

```
package main

import (
    "github.com/tvolkov/speedtest"
)

func main() {
    speedtest.Speedtest("speedtest")
    speedtest.Speedtest("fastcom")
}
```

alternatively you can run
```
go test
```
to see both implementations working

# N.B
I haven't had enough time to properly test it but I realize there should be varioud unit tests added for each module, so for now there are just some basic high level tests

