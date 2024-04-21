package main

import "os"

func exit(cfg *config) error {
	os.Exit(0)
	return nil
}
