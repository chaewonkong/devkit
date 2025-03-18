package main

import (
	"github.com/chaewonkong/devkit/internal/base64"
	"github.com/chaewonkong/devkit/internal/root"
	"github.com/chaewonkong/devkit/internal/unix"
	"github.com/chaewonkong/devkit/internal/url"
	"github.com/chaewonkong/devkit/internal/uuid"
)

func main() {
	rootCmd := root.NewRootCommand()
	rootCmd.AddCommand(
		uuid.NewUUIDCommand(),
		url.NewURLCommand(),
		base64.NewBase64Command(),
		unix.NewUnixTimeCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErr(err)
	}
}
