package commands

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/dihedron/go-windows-junction/junction"
)

type Create struct {
	Command
	Target     string `short:"t" long:"target" description:"The target filesystem object." required:"yes"`
	Mountpoint string `short:"m" long:"mountpoint" description:"The mountpoint linking to the target." required:"yes"`
}

func (cmd *Create) Execute(args []string) error {

	if err := junction.Create(cmd.Target, cmd.Mountpoint); err != nil {
		slog.Error("error creating junction", "target", cmd.Target, "mountpoint", cmd.Mountpoint, "error", err)
		return err
	}
	result, err := os.Readlink(cmd.Mountpoint)
	if err != nil {
		slog.Error("error reading junction", "mountpoint", cmd.Mountpoint, "error", err)
		return err
	}
	fmt.Printf("%s is linked to %s", cmd.Mountpoint, result)
	return nil
}
