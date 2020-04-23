/*
 * @Date: 2020-04-23 09:37:41
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @LastEditors: ferried
 * @LastEditTime: 2020-04-23 09:38:34
 * @Editor: Visual Studio Code
 * @Desc: nil
 * @License: nil
 */

package fish

import "animal-crossing-cli/pkg/base"

type Fish struct {
	Name      string     `json:"name"`
	Price     int        `json:"price"`
	Month     base.Month `json:"month"`
	Time      []string   `json:"time"`
	Place     []string   `json:"place"`
	Condition string     `json:"condition"`
	Shadow    string     `json:"shadow"`
}
