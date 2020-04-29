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

type Month struct {
	North []string
	South []string
}

type print interface {
	PrintByName(name string)
	PrintAll()
}
