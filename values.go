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
	"strconv"
)

// Value represents some value of a flag
type Value interface {
	Set(string) error
	Name() string
}

type boolValue bool

func (val *boolValue) Name() string {
	return "boolean"
}

func (val *boolValue) Set(newVal string) error {
	b, err := strconv.ParseBool(newVal)
	if err == nil {
		*val = boolValue(b)
	}
	return err
}

type stringValue string

func (val *stringValue) Name() string {
	return "string"
}

func (val *stringValue) Set(newVal string) error {
	*val = stringValue(newVal)
	return nil
}

type stringArrayValue []string

func (val *stringArrayValue) Name() string {
	return "string array"
}

func (val *stringArrayValue) Set(newVal string) error {
	*val = append(*val, newVal)
	return nil
}

type intValue int

func (val *intValue) Name() string {
	return "integer"
}

func (val *intValue) Set(newVal string) error {
	i, err := strconv.ParseInt(newVal, 10, 32)
	if err == nil {
		*val = intValue(i)
	}
	return err
}

type uintValue uint

func (val *uintValue) Name() string {
	return "unsigned integer"
}

func (val *uintValue) Set(newVal string) error {
	i, err := strconv.ParseUint(newVal, 10, 32)
	if err == nil {
		*val = uintValue(i)
	}
	return err
}

type int8Value int8

func (val *int8Value) Name() string {
	return "8-bit integer"
}

func (val *int8Value) Set(newVal string) error {
	i, err := strconv.ParseInt(newVal, 10, 8)
	if err == nil {
		*val = int8Value(i)
	}
	return err
}

type uint8Value uint8

func (val *uint8Value) Name() string {
	return "unsigned 8-bit integer"
}

func (val *uint8Value) Set(newVal string) error {
	i, err := strconv.ParseUint(newVal, 10, 8)
	if err == nil {
		*val = uint8Value(i)
	}
	return err
}

type int16Value int16

func (val *int16Value) Name() string {
	return "16-bit integer"
}

func (val *int16Value) Set(newVal string) error {
	i, err := strconv.ParseInt(newVal, 10, 16)
	if err == nil {
		*val = int16Value(i)
	}
	return err
}

type uint16Value uint16

func (val *uint16Value) Name() string {
	return "unsigned 16-bit integer"
}

func (val *uint16Value) Set(newVal string) error {
	i, err := strconv.ParseUint(newVal, 10, 16)
	if err == nil {
		*val = uint16Value(i)
	}
	return err
}

type int32Value int32

func (val *int32Value) Name() string {
	return "32-bit integer"
}

func (val *int32Value) Set(newVal string) error {
	i, err := strconv.ParseInt(newVal, 10, 32)
	if err == nil {
		*val = int32Value(i)
	}
	return err
}

type uint32Value uint32

func (val *uint32Value) Name() string {
	return "unsigned 32-bit integer"
}

func (val *uint32Value) Set(newVal string) error {
	i, err := strconv.ParseUint(newVal, 10, 32)
	if err == nil {
		*val = uint32Value(i)
	}
	return err
}

type int64Value int64

func (val *int64Value) Name() string {
	return "64-bit integer"
}

func (val *int64Value) Set(newVal string) error {
	i, err := strconv.ParseInt(newVal, 10, 64)
	if err == nil {
		*val = int64Value(i)
	}
	return err
}

type uint64Value uint64

func (val *uint64Value) Name() string {
	return "unsigned 64-bit integer"
}

func (val *uint64Value) Set(newVal string) error {
	i, err := strconv.ParseUint(newVal, 10, 64)
	if err == nil {
		*val = uint64Value(i)
	}
	return err
}

type float32Value float32

func (val *float32Value) Name() string {
	return "32-bit float"
}

func (val *float32Value) Set(newVal string) error {
	i, err := strconv.ParseFloat(newVal, 32)
	if err == nil {
		*val = float32Value(i)
	}
	return err
}

type float64Value float64

func (val *float64Value) Name() string {
	return "64-bit float"
}

func (val *float64Value) Set(newVal string) error {
	i, err := strconv.ParseFloat(newVal, 64)
	if err == nil {
		*val = float64Value(i)
	}
	return err
}
