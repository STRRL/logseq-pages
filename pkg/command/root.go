package command

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	rootOption := RootOption{
		WorkDirectory: "",
	}

	rootCommand := &cobra.Command{
		Use:   "logseq-pages",
		Short: "logseq-pages is a tool for mark pages as public or private",
	}
	rootCommand.PersistentFlags().StringVarP(&rootOption.WorkDirectory, "work-directory", "d", "", "path of logseq directory, default is current directory")
	rootCommand.AddCommand(
		NewListCommand(&rootOption),
		NewMarkCommand(&rootOption),
	)
	return rootCommand
}

type RootOption struct {
	WorkDirectory string
}
