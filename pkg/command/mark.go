package command

import "github.com/spf13/cobra"

func NewMarkCommand(rootOption *RootOption) *cobra.Command {
	result := &cobra.Command{
		Use: "mark",
	}
	return result
}
