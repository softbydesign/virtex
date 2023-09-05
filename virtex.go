package main

import (
	"fmt"
	"virtex/config"
)

func main() {
	virtexConfig, err := config.LoadVirtexConfig(config.DEFAULT_VIRTEX_CONFIG_PATH)

	if err != nil {
		panic("Unable to start Virtex server")
	}

	fmt.Println(virtexConfig)
}
