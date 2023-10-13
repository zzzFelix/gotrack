package main

import (
	"os"

	"github.com/zzzFelix/gotrack/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
