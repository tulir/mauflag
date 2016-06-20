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

package mauflag

import (
	"os"
)

var defaultSet = New(os.Args)

// DefaultSet returns the default flagset which takes its arguments from os.Args
func DefaultSet() *Set {
	return defaultSet
}

// Make creates and registers a flag
func Make() *Flag {
	return DefaultSet().Make()
}

// Parse the command line arguments into mauflag form
func Parse() error {
	return DefaultSet().Parse()
}

// Args returns the arguments that weren't associated with any flag
// Note: The first arg is the command
func Args() []string {
	return DefaultSet().Args()
}

// Arg returns the string at the given index from the list Args() returns
// If the index does not exist, Arg will return an empty string.
// Note: The first arg (`Arg(0)`) is the command
func Arg(i int) string {
	return DefaultSet().Arg(i)
}

// NArg returns the number of arguments not associated with any flags
func NArg() int {
	return len(DefaultSet().args)
}
