# mauflag
An extendable command-line argument parser for Golang.

Use `go get maunium.net/go/mauflag` to get this package

## Basic usage
A flag can have as many long or short keys as you like. Short keys are prefixed with one dash (`-`) whilst long keys are prefixed with two (`--`)

To create a flag, you must first call the `Make()` method in a flagset. Calling flagset methods as package functions uses the default flagset which takes input arguments from `os.Args`

After creating a flag, you can call the functions it has as you like. After configuring the flag, you can call one of the type functions to set the flag type and get a pointer to the value.

```go
var myStringFlag = flag.Make().ShortKey("s").LongKey("my-string-flag").Default("a string").String()
```
This example creates a flag that takes string values. The flag can be set using `-s` or `--my-string-flag`. If the input doesn't contain a value for the flag the value will be `a string`

More docs, including all values supported by default, can be found from [godoc.org/maunium.net/go/mauflag](https://godoc.org/maunium.net/go/mauflag)

## Custom values
All value containers must implement the `Value` interface. It contains two functions:
* `Set(string)` which is called whenever the parser finds a value associated with the flag.
* `Name()` which should return a human-readable name for the type of the value. If the value container contains multiple actual values, `Name()` should return the name of a single object (e.g. `integer` if the value container contains multiple integers)

After creating a value container you can use it by calling the `Custom` method of a `Flag` and giving an instance of the value container. The parser will then call the `Set` method each time it encounters a value associated with the flag.
