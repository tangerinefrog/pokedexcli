package main

import "os"

func exitCallback(*pagingParam, string) error {
	os.Exit(0)
	return nil
}
