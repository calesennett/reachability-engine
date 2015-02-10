package reachability

import (
	"fmt"
	pn "petriNet"
	"reflect"
)

type Edge struct {
	From []int
	To   []int
}

func Reachable(pn pn.PetriNet) {
	unexplored := [][]int{}
	reachable := [][]int{}
	edges := []Edge{}
	unexplored = append(unexplored, pn.State())
	reachable = append(reachable, pn.State())
	for len(unexplored) != 0 {
		state := unexplored[0]
		unexplored = unexplored[1:]
		pn.UpdatePlaces(state)
		for _, t := range pn.Transitions {
			pn.Fire(t)
			newState := pn.State()
			if !inSet(newState, reachable) {
				unexplored = append(unexplored, newState)
				reachable = append(reachable, newState)
			}
			if !reflect.DeepEqual(state, newState) {
				edges = append(edges, Edge{From: state, To: newState})
			}
		}
	}
	outputGraph(reachable, edges)
}

func outputGraph(reach [][]int, edges []Edge) {
	fmt.Printf("States: %v\n", len(reach))
	fmt.Printf("Edges: %v\n\n", len(edges))
	for _, edge := range edges {
		fmt.Printf("%v : %v\n", edge.From, edge.To)
	}
}

func inSet(state []int, states [][]int) bool {
	for _, s := range states {
		if reflect.DeepEqual(state, s) {
			return true
		}
	}
	return false
}
