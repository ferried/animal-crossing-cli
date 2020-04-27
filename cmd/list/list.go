/*
 * @Date: 2020-04-23 11:26:03
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @LastEditors: ferried
 * @LastEditTime: 2020-04-27 22:29:50
 * @Editor: Visual Studio Code
 * @Desc: nil
 * @License: nil
 */
package list

import (
	"animal-crossing-cli/pkg/base"
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/cobra"
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
			if len(ins) > 0 {
				HandleCommand(ins, reflect.TypeOf(base.Insect{}).Name())
			}
			if len(fi) > 0 {
				HandleCommand(fi, reflect.TypeOf(base.Fish{}).Name())
			}
			os.Exit(0)
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

func HandleCommand(name string, ty string) {
	datas := base.Get(ty)
	fmt.Println(datas)
	os.Exit(0)
}
