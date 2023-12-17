package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/jroimartin/clilog"
)

func main() {
	h := clilog.NewCLIHandler(os.Stderr, &clilog.HandlerOptions{Level: slog.LevelInfo})
	slog.SetDefault(slog.New(h))

	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Fprintln(os.Stderr, "not ok")
		os.Exit(1)
	}

	fmt.Printf("%v %v\n", info.Main.Path, info.Main.Version)
	for _, dep := range info.Deps {
		fmt.Printf("%v %v\n", dep.Path, dep.Version)
	}

	fmt.Println("test")
}
