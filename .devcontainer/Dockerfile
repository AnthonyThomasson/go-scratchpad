FROM mcr.microsoft.com/devcontainers/go:1-bullseye

# Install Docker CLI
RUN apt-get update && \
	apt-get install -y docker.io

# Install go tools
RUN go install golang.org/x/tools/gopls@latest
RUN go install golang.org/x/lint/golint@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Create a workspace directory
RUN mkdir -p /workspaces/go-scratchpad
WORKDIR /workspaces/go-scratchpad