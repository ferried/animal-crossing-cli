/*
 * @File: handle.go
 * @Date: 2020/4/23 3:18 下午
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Software: GoLand
 * @Desc: null
 * @License: null
 */
package insect

import (
	"fmt"
	"os"
	"strings"

	"github.com/modood/table"
)

func HandleCommand(ins string) {
	if strings.EqualFold(ins, "insects") {
		table.Output(*Get())
		os.Exit(0)
	} else {
		var find []Insect
		for _, v := range *Get() {
			if strings.EqualFold(ins, v.Name) {
				find = append(find, v)
			}
		}
		if find == nil {
			fmt.Printf("没有找到名称为%s的项\n", ins)
		} else {
			table.Output(find)
		}
	}
}
