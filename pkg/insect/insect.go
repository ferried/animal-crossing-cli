/*
 * @Date: 2020-04-23 09:31:45
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @LastEditors: ferried
 * @LastEditTime: 2020-04-23 09:43:43
 * @Editor: Visual Studio Code
 * @Desc: nil
 * @License: nil
 */

package insect

import (
	"animal-crossing-cli/pkg/common"
	"animal-crossing-cli/pkg/utils"
	"encoding/json"
)

var (
	Insects *[]Insect
)

type Insect struct {
	Name      string       `json:"name"`
	Price     int          `json:"price"`
	Month     common.Month `json:"month"`
	Time      []string     `json:"time"`
	Place     []string     `json:"place"`
	Condition string       `json:"condition"`
}

func (insect *Insect) Get() *[]Insect {
	return Insects
}

func Store() {
	d := utils.Store("insect.json")
	json.Unmarshal(d, &Insects)
}

func init() {
	Store()
}
