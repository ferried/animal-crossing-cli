package list

import (
	"animal-crossing-cli/pkg/insect"
	"github.com/spf13/cobra"
	"os"
)

var (
	ins     string
	fi      string
	ani     string
	vil     string
	listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List items",
		Long:    "List items",
		Example: "acc list insect",
		Run: func(cmd *cobra.Command, args []string) {
			if len(ins) != 0 {
				insect.HandleCommand(ins)
			} else {
				cmd.Help()
				os.Exit(0)
			}
		},
	}
)

func Cmd() *cobra.Command {
	return listCmd
}

func init() {
	listCmd.Flags().StringVarP(&ins, "insect", "i", "", "List all insects")
	listCmd.Flags().StringVarP(&fi, "fish", "f", "", "List all fishes")
	listCmd.Flags().StringVarP(&ani, "animal", "a", "", "List all animals")
	listCmd.Flags().StringVarP(&vil, "villager", "v", "", "List all villagers")
}
