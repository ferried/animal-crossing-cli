package list

import (
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List items",
		Long:    "List items",
		Example: "acc list insect",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Help()
			}
		},
	}
)

func Cmd() *cobra.Command {
	return listCmd
}

func init() {
	listCmd.Flags().StringP("insect", "i", "insect", "List all insect")
	listCmd.Flags().StringP("fish", "f", "fish", "List all fish")
	listCmd.Flags().StringP("animal", "a", "animal", "List all animal")
	listCmd.Flags().StringP("villager", "v", "villager", "List all villager")
}
