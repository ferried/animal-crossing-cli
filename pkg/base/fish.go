/*
 * @Date: 2020-04-23 09:37:41
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @LastEditors: ferried
 * @LastEditTime: 2020-04-27 22:15:40
 * @Editor: Visual Studio Code
 * @Desc: nil
 * @License: nil
 */

package base

import (
	"animal-crossing-cli/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/modood/table"
	"strings"
)

var (
	fishes *[]Fish
)

func init() {
	d := utils.Store("fish.json")
	json.Unmarshal(d, &fishes)
}

type Fish struct {
	Name   string   `json:"name"`
	Price  int      `json:"price"`
	Month  Month    `json:"month"`
	Time   []string `json:"time"`
	Place  []string `json:"place"`
	Shadow string   `json:"shadow"`
}

func (fish Fish) PrintByName(name string) {
	for _, i := range *fishes {
		if strings.EqualFold(i.Name, name) {
			table.Output([]Fish{i})
		}
	}
}

func (fish Fish) PrintAll() {
	fmt.Println("log")
	table.Output(*fishes)
}
