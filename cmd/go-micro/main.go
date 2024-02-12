package main

import (
	"github.com/trofkm/cli/cmd"

	// register commands
	_ "github.com/trofkm/cli/cmd/call"
	_ "github.com/trofkm/cli/cmd/completion"
	_ "github.com/trofkm/cli/cmd/describe"
	_ "github.com/trofkm/cli/cmd/generate"
	_ "github.com/trofkm/cli/cmd/new"
	_ "github.com/trofkm/cli/cmd/run"
	_ "github.com/trofkm/cli/cmd/services"
	_ "github.com/trofkm/cli/cmd/stream"

	// plugins
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
)

func main() {
	cmd.Run()
}
