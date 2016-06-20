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
	"fmt"
	"os"
	"strings"
)

// Set is a set of flags with certain input arguments
type Set struct {
	// The list of strings used as input
	InputArgs []string
	// Whether or not to ignore all flags after the user has entered two dashes with no flag key ("--")
	// If enabled, all arguments after two dashes with no flag key will go into the args array (@see Args())
	DoubleLineEscape bool
	// Whether or not to exit the program when there's an error
	// If enabled, the error message will be printed to `stderr` after which `os.Exit(1)` will be called.
	ExitOnError bool

	args  []string
	flags []*Flag
}

// New creates a new flagset
func New(args []string) *Set {
	return &Set{InputArgs: args}
}

// Args returns the arguments that weren't associated with any flag
func (fs *Set) Args() []string {
	return fs.args
}

// Arg returns the string at the given index from the list Args() returns
// If the index does not exist, Arg will return an empty string.
func (fs *Set) Arg(i int) string {
	if len(fs.args) >= i {
		return ""
	}
	return fs.args[i]
}

func (fs *Set) err(format string, args ...interface{}) error {
	if fs.ExitOnError {
		fmt.Fprintf(os.Stderr, format, args...)
		fmt.Fprint(os.Stderr, "\n")
		os.Exit(1)
		return nil
	}
	return fmt.Errorf(format, args...)
}

// Parse the input arguments in this flagset into mauflag form
// Before this function is called all the flags will have either the type default or the given default value
func (fs *Set) Parse() error {
	var flag *Flag
	var key string
	var noMoreFlags = false

	for _, arg := range fs.InputArgs {
		arg = strings.ToLower(arg)
		if noMoreFlags {
			fs.args = append(fs.args, arg)
		} else if arg == "--" && fs.DoubleLineEscape {
			noMoreFlags = true
		} else if flag != nil {
			err := flag.setValue(arg)
			if err != nil {
				return fs.err("Flag %s was not a %s", key, flag.Value.Name())
			}
			flag = nil
		} else if arg[0] == '-' {
			var err error
			key, flag, err = fs.flagStart(arg)
			if err != nil {
				return err
			}
		} else {
			fs.args = append(fs.args, arg)
		}
	}
	return nil
}

func (fs *Set) flagStart(arg string) (string, *Flag, error) {
	key := arg[1:]

	var val string
	if strings.ContainsRune(key, '=') {
		val = key[strings.Index(arg, "=")+1:]
		key = key[:strings.Index(arg, "=")]
	}

	flag, key := fs.getFlag(key)
	if flag == nil {
		return "", nil, fs.err("Unknown flag: %s", key)
	} else if len(val) > 0 {
		flag.setValue(val)
		return "", nil, nil
	} else {
		_, ok := flag.Value.(*boolValue)
		if ok {
			flag.setValue("true")
			return "", nil, nil
		}
	}
	return key, flag, nil
}

func (fs *Set) getFlag(key string) (*Flag, string) {
	if key[0] == '-' {
		key = key[1:]
		for _, lflag := range flags {
			for _, lkey := range lflag.longKeys {
				if lkey == key {
					return lflag, lkey
				}
			}
		}
	} else {
		for _, lflag := range flags {
			for _, lkey := range lflag.shortKeys {
				if lkey == key {
					return lflag, lkey
				}
			}
		}
	}
	return nil, key
}
