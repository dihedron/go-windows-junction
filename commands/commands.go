package commands

// Commands is the set of root command groups.
type All struct {
	Create Create `command:"create" alias:"c" description:"Create a Windows filesystem junction."`

	//lint:ignore SA5008 the github.com/jessevdk/go-flags library supports multiple alias tags on struct fields
	Version Version `command:"version" alias:"ver" alias:"v" description:"Print the application version information and exit."`

	//Settings Settings `command:"settings" alias:"s" description:"Apply the provided migration to a settings.xml file."`
}
