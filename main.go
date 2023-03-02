package main

import (
	"context"

	"github.com/ChristianSch/go-poc-action/src"
	githubactions "github.com/sethvargo/go-githubactions"
)

func run() error {
	ctx := context.Background()
	action := githubactions.New()

	cfg, err := src.NewFromInputs(action)
	if err != nil {
		return err
	}

	return src.Run(ctx, cfg)
}

func main() {
	err := run()
	if err != nil {
		githubactions.Fatalf("%v", err)
	}
}
