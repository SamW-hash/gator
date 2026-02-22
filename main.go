package main

import (
	"fmt"
	"os"

	"github.com/SamW-hash/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := cfg.SetUser("Sam"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cfg2, _ := config.Read()
	fmt.Println(cfg2.DBURL)
}
