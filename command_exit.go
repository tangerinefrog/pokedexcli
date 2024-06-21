package main

import "os"

func exitCallback(*pagingParam) error {
	os.Exit(0)
	return nil
}
