package main

import (
	"bufio"
	"os"
	"parser"
	"reachability"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	pn := parser.MakePetriNet(lines)
	reachability.Reachable(pn)
}
