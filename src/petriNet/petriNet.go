package petriNet

import (
	"strconv"
)

type Transition struct {
	Id string
}

type Edge struct {
	To     string
	From   string
	Weight int
}

type Place struct {
	Tokens int
	Id     string
}

type PetriNet struct {
	Edges       []Edge
	Transitions []Transition
	Places      []Place
	InitMarking []int
}

func (pn PetriNet) Fire(t Transition) {
	edges := pn.edgesTo(t.Id)
	fired := false
	for _, edge := range edges {
		placeId := edge.From
		place := pn.findPlace(placeId)
		if edge.Weight > place.Tokens {
			return
		}
	}

	for _, edge := range edges {
		placeId := edge.From
		index, _ := strconv.Atoi(placeId[1:])
		pn.Places[index].Tokens -= edge.Weight
	}
	fired = true

	edges = pn.edgesFrom(t.Id)
	for _, edge := range edges {
		placeId := edge.To
		if fired {
			index, _ := strconv.Atoi(placeId[1:])
			pn.Places[index].Tokens += edge.Weight
		}
	}
}

func (pn PetriNet) findPlace(id string) Place {
	for _, place := range pn.Places {
		if place.Id == id {
			return place
		}
	}
	return Place{}
}

func (pn PetriNet) State() []int {
	var marking []int
	for _, place := range pn.Places {
		marking = append(marking, place.Tokens)
	}
	return marking
}

func (pn PetriNet) UpdatePlaces(m []int) {
	for i, value := range m {
		pn.Places[i].Tokens = value
	}
}

func (pn PetriNet) edgesFrom(id string) []Edge {
	edges := []Edge{}
	for _, edge := range pn.Edges {
		if edge.From == id {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (pn PetriNet) edgesTo(id string) []Edge {
	edges := []Edge{}
	for _, edge := range pn.Edges {
		if edge.To == id {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (pn PetriNet) enabled(t Transition, e Edge, p Place) bool {
	if e.To == t.Id && p.Id == e.From {
		if e.Weight <= p.Tokens {
			return true
		}
	}
	return false
}
