package main

import (
	"testproject/confession/cmd"
	"testproject/confession/env"
)

var (
	// VERSION -
	VERSION string
	// COMMIT -
	COMMIT string
	// BUILDTIME -
	BUILDTIME string
)

func main() {
	env.Version = VERSION
	env.Commit = COMMIT
	env.BuildTime = BUILDTIME
	cmd.Execute()
}
