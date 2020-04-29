/*
 * @Date: 2020-04-23 09:31:45
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @LastEditors: ferried
 * @LastEditTime: 2020-04-27 22:10:23
 * @Editor: Visual Studio Code
 * @Desc: nil
 * @License: nil
 */

package base

import (
	"animal-crossing-cli/pkg/utils"
	"encoding/json"
	"github.com/modood/table"
	"strings"
)

var (
	insects *[]Insect
)

func init() {
	d := utils.Store("insect.json")
	json.Unmarshal(d, &insects)

}

type Insect struct {
	Name      string   `json:"name"`
	Price     int      `json:"price"`
	Month     Month    `json:"month"`
	Time      []string `json:"time"`
	Place     []string `json:"place"`
	Condition string   `json:"condition"`
}

func (insect Insect) PrintByName(name string) {
	for _, i := range *insects {
		if strings.EqualFold(i.Name, name) {
			table.Output(i)
		}
	}
}

func (insect Insect) PrintAll() {
	table.Output(*insects)
}
