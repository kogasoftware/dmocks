package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kogasoftware/dmocks"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error:%s\n", err)
			os.Exit(1)
		}
	}()

	os.Exit(_main())
}

func _main() int {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(1)
	}

	server := dmocks.NewDmocks()
	server.Run()

	return 0
}
