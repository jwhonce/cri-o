package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	psDescription = "Only Brent knows..."
	psCommand = cli.Command{
		Name: "ps",
		Usage: "...",
		Description: psDescription,
		Flags: []cli.Flag{},
		Action: psCmd,
		ArgsUsage: "",
	}
)


func psCmd(c *cli.Context) error {
	store, err := getStore(c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}

	containerStore, err := store.ContainerStore()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}

	containers, err := containerStore.Containers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}

	for _, container := range containers {
		fmt.Fprintf(os.Stdout, "%s", container.ID)
	}

	return nil
}
