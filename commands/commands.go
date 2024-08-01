package commands

// Commands is the set of root command groups.
type All struct {
	Create Create `command:"create" alias:"c" description:"Create a Windows filesystem junction."`

	//lint:ignore SA5008 the github.com/jessevdk/go-flags library supports multiple alias tags on struct fields
	Delete Delete `command:"delete" alias:"d" alias:"remove" alias:"rm" alias:"r" description:"Remove a Windows filesystem junction or UNIX symbolic link."`

	//lint:ignore SA5008 the github.com/jessevdk/go-flags library supports multiple alias tags on struct fields
	Version Version `command:"version" alias:"ver" alias:"v" description:"Print the application version information and exit."`
}
