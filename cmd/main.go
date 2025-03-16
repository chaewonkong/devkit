package main

import (
	"github.com/chaewonkong/devkit/cmd/root"
	"github.com/chaewonkong/devkit/cmd/url"
	"github.com/chaewonkong/devkit/cmd/uuid"
)

func main() {
	rootCmd := root.NewRootCommand()
	rootCmd.AddCommand(uuid.NewUUIDCommand(), url.NewURLCommand())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
