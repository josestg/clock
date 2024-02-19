# clock

A dependency injection-friendly `time.Now()`.

## Install

```shell
go get github.com/josestg/clock
```

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/josestg/clock"
)

func main() {
	{ // create a new clock from location name.
		jkt, err := clock.LoadLocation("Asia/Jakarta")
		if err != nil {
			panic(err)
		}

		now := jkt.Now()
		fmt.Println(now)
	}

	{ // create a new clock from location.
		local := clock.FromLocation(time.Local)
		fmt.Println(local.Now())

		// or using the singleton instance.
		fmt.Println(clock.Local.Now())
	}

	{
		// using the static clock for testing.
		now := time.Now()
		static := clock.Static(now)

		fmt.Println(static.Now() == now)
	}
}
```