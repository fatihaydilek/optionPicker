package main

import (
	"fmt"
	"github.com/fatihaydilek/optionPicker/opts"
)

func main() {
	s := opts.NewRuleSet()
	s.AddDep("a", "b")
	s.AddDep("a", "c")
	s.AddConflict("b", "d")
	s.AddConflict("b", "e")
	if !s.IsCoherent() {
		fmt.Println("s.IsCoherent failed")
	}

	selected := opts.New(s)
	selected.Toggle("d")
	selected.Toggle("e")
	selected.Toggle("a")
	fmt.Println(selected.StringSlice())
}
