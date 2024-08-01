package commands

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/dihedron/junctions/junction"
)

type Create struct {
	Command
	Target string `short:"t" long:"target" description:"The target filesystem object." required:"yes"`
	Link   string `short:"l" long:"link" description:"The name of the link." required:"yes"`
}

func (cmd *Create) Execute(args []string) error {
	slog.Info("creating junction", "target", cmd.Target, "link", cmd.Link)

	target, err := filepath.Abs(cmd.Target)
	if err != nil {
		slog.Error("error translating absolute path of target", "path", cmd.Target, "error", err)
		return err
	}
	slog.Debug("absolute path of target", "path", target)

	link, err := filepath.Abs(cmd.Link)
	if err != nil {
		slog.Error("error translating absolute path of link", "path", cmd.Link, "error", err)
		return err
	}
	slog.Debug("absolute path of link", "path", link)

	if err := junction.Create(target, link); err != nil {
		slog.Error("error creating junction", "target", target, "link", link, "error", err)
		return err
	}
	slog.Info("junction created", "target", target, "link", link)
	result, err := os.Readlink(link)
	if err != nil {
		slog.Error("error reading junction", "link", link, "error", err)
		return err
	}
	slog.Info("junction checked", "target", target, "link", link)
	fmt.Printf("%s is linked to %s\n", link, result)
	return nil
}
