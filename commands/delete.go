package commands

import (
	"log/slog"
	"os"
	"path/filepath"
)

type Delete struct {
	Command
	Link string `short:"l" long:"link" description:"The link to be deleted." required:"yes"`
}

func (cmd *Delete) Execute(args []string) error {

	link, err := filepath.Abs(cmd.Link)
	if err != nil {
		slog.Error("error translating absolute path of link", "path", cmd.Link, "error", err)
		return err
	}
	slog.Debug("absolute path of link", "path", link)

	slog.Info("removing junction", "link", link)
	if err = os.Remove(link); err != nil {
		slog.Error("rerror removing junction", "link", link, "error", err)
		return err
	}
	slog.Info("junction removed", "link", link)
	return nil
}
