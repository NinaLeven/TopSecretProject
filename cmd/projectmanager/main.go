package main

import (
	"context"
	"flag"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	configPath := flag.String("c", "../..config/config.yaml", "path to config file")
	flag.Parse()

	if configPath == nil {
		fmt.Println("config path is not set")
		os.Exit(1)
	}

}
