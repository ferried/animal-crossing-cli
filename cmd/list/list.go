package list

import (
	"animal-crossing-cli/pkg/insect"
	"github.com/modood/table"
	"github.com/spf13/cobra"
	"os"
	"strings"
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
				if strings.EqualFold(ins, "insects") {
					table.Output(*insect.Get())
					os.Exit(0)
				} else {
					for _, v := range *insect.Get() {
						if strings.EqualFold(ins, v.Name) {
							table.Output([]insect.Insect{v})
						}
					}
				}
			}
		},
	}
)

func Cmd() *cobra.Command {
	return listCmd
}

func init() {
	listCmd.Flags().StringVarP(&ins, "insect", "i", "insects", "List all insects")
	listCmd.Flags().StringVarP(&fi, "fish", "f", "fishes", "List all fishes")
	listCmd.Flags().StringVarP(&ani, "animal", "a", "animals", "List all animals")
	listCmd.Flags().StringVarP(&vil, "villager", "v", "villagers", "List all villagers")
}
