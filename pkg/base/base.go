/*
 * @Date: 2020-04-27 22:19:27
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @LastEditors: ferried
 * @LastEditTime: 2020-04-27 22:20:50
 * @Editor: Visual Studio Code
 * @Desc: nil
 * @License: nil
 */

package base

import (
	"strings"
)

func Get(typ string) interface{} {
	if strings.EqualFold(typ, "FISH") {
		return *fishes
	}
	if strings.EqualFold(typ, "insect") {
		return *insects
	}
	return nil
}
