package stream

import (
	mcli "github.com/trofkm/cli/cmd"
	"github.com/urfave/cli/v2"
)

func init() {
	mcli.Register(&cli.Command{
		Name:  "stream",
		Usage: "Create a service stream",
		Subcommands: []*cli.Command{
			{
				Name:    "bidi",
				Aliases: []string{"b"},
				Usage:   "Create a bidirectional service stream, e.g. " + mcli.App().Name + " stream bidirectional helloworld Helloworld.PingPong '{\"stroke\": 1}' '{\"stroke\": 2}'",
				Action:  Bidirectional,
			},
			{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "Create a server service stream, e.g. " + mcli.App().Name + " stream server helloworld Helloworld.ServerStream '{\"count\": 10}'",
				Action:  Server,
			},
		},
	})
}
