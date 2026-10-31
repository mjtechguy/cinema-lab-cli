package main

import (
	"os"

	"github.com/cli/cli/v2/internal/ghcmd"
)

func main() {
	code := ghcmd.Main()
	os.Exit(int(code))
}

// lab: skip-in-place experiment 2026-10-23T07:15:20Z
