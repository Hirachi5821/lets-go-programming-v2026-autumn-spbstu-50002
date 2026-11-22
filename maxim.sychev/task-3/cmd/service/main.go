package main

import (
	"flag"

	"github.com/Hirachi5821/task-3/internal/converter"
)

func main() {
	configPath := flag.String("config", "", "path to YAML configuration")
	flag.Parse()

	if err := converter.Run(*configPath); err != nil {
		panic(err)
	}
}
