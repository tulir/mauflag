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

package flag

import (
	"os"
)

var defaultSet = New(os.Args[1:])

// DefaultSet returns the default flagset
func DefaultSet() *Set {
	return defaultSet
}

// Make calls the Make() function of the default flagset
func Make() *Flag {
	return DefaultSet().Make()
}

// Parse calls the Parse() function of the default flagset
func Parse() error {
	return DefaultSet().Parse()
}

// Args calls the Args() function of the default flagset
func Args() []string {
	return DefaultSet().Args()
}

// Arg calls the Arg(i) function of the default flagset
func Arg(i int) string {
	return DefaultSet().Arg(i)
}
