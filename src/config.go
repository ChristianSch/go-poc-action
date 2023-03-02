package src

import (
	githubactions "github.com/sethvargo/go-githubactions"
)

type DrafterConfig struct {
	RepoOwner string
	RepoName  string
}

func NewFromInputs(action *githubactions.Action) (*DrafterConfig, error) {
	actionCtx, err := action.Context()
	if err != nil {
		return nil, err
	}

	owner, name := actionCtx.Repo()
	c := DrafterConfig{
		RepoOwner: owner,
		RepoName:  name,
	}

	return &c, nil
}
