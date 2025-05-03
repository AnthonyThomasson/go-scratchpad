package agent

import (
	"context"

	"github.com/tmc/langchaingo/llms"
)

type Agent struct {
	llm llms.Model
}

func (a *Agent) Run(ctx context.Context, input string) (string, error) {
	return "", nil
}
