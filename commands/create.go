package commands

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/dihedron/go-windows-junction/junction"
)

type Create struct {
	Command
	Source string `short:"s" long:"source" description:"The source filesystem object." required:"yes"`
	Target string `short:"m" long:"mountpoint" description:"The target mountpoint." required:"yes"`
}

func (cmd *Create) Execute(args []string) error {

	if err := junction.Create(cmd.Source, cmd.Target); err != nil {
		slog.Error("error creating junction", "source", cmd.Source, "mountpoint", cmd.Target, "error", err)
		return err
	}
	result, err := os.Readlink(cmd.Target)
	if err != nil {
		slog.Error("error reading junction", "mountpoint", cmd.Target, "error", err)
		return err
	}
	fmt.Printf("%s is linked to %s", cmd.Target, result)
	return nil
}
