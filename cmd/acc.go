/*
 * @Date: 2020-04-23 10:51:44
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @LastEditors: ferried
 * @LastEditTime: 2020-04-26 16:51:01
 * @Editor: Visual Studio Code
 * @Desc: nil
 * @License: nil
 */
package cmd

import (
	"animal-crossing-cli/cmd/list"

	"github.com/spf13/cobra"
)

var (
	animalCrossingCli = &cobra.Command{
		Use:     "acc",
		Short:   "Animal crossing's command line tool",
		Long:    "The villagers are right\nLicense:MIT\nAuthor:ferried<harlancui@outlook.com>",
		Example: "acc list insect",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Help()
			}
		},
	}
)

func Execute() error {
	return animalCrossingCli.Execute()
}

func init() {
	animalCrossingCli.AddCommand(list.Cmd())
}
