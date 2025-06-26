package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mr-nuno/pulumi-envoy/provider"
)

// main is the entry point for the Pulumi Loopia DNS provider binary.
func main() {
	prov, err := provider.NewProvider()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	err = prov.Run(context.Background(), "envoy", "0.1.0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
