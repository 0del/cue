package main

import (
	"fmt"

	"cuelang.org/go/cue"
)

type ab struct{ A, B int }

var r cue.Runtime

var x ab

func main() {
	i, _ := r.Compile("test", `{A: 2, B: 4}`)
	_ = i.Value().Decode(&x)
	fmt.Println(x)

	i, _ = r.Compile("test", `{B: "foo"}`)
	_ = i.Value().Decode(&x)
	fmt.Println(x)

	// Output:
	// {2 4}
	// json: cannot unmarshal string into Go struct field ab.B of type int
}
