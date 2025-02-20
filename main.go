//go:generate go run generate-asset.go

package main

import (
	"github.com/bmooreitul/docgen/cmd"
)

func main() {
	cmd.Execute()
}
