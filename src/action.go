package src

import (
	"context"

	"github.com/sethvargo/go-githubactions"
)

func Run(ctx context.Context, cfg *DrafterConfig) error {
	githubactions.Infof("Processing Repo %s/%s", cfg.RepoOwner, cfg.RepoName)
	return nil
}
