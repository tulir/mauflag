// mauflag - An extendable command-line argument parser for Golang
// Copyright (C) 2016 Maunium
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	flag "maunium.net/go/mauflag"
)

var bf = flag.Make().LongKey("bff").ShortKey("b").Bool()
var xb = flag.Make().ShortKey("x").Bool()
var yb = flag.Make().ShortKey("y").Bool()
var zb = flag.Make().ShortKey("z").Bool()
var str = flag.Make().LongKey("this-is-a-string").ShortKey("s").ShortKey("i").String()
var array = flag.Make().LongKey("array").ShortKey("a").LongKey("arr").StringArray()
var def = flag.Make().LongKey("asd").Default("lorem").String()

func main() {
	fmt.Println(flag.Parse())
	fmt.Println(*bf, *xb, *yb, *zb)
	fmt.Println(*str)
	fmt.Println(*def)
	for _, arg := range *array {
		fmt.Print(arg, ", ")
	}
	fmt.Print("\n")
	for _, arg := range flag.Args() {
		fmt.Print(arg, ", ")
	}
	fmt.Print("\n")
}
