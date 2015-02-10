package parser

import (
	pn "petriNet"
	"strconv"
	"strings"
)

func MakePetriNet(lines []string) pn.PetriNet {
	var places []pn.Place
	var transitions []pn.Transition
	var edges []pn.Edge
	var marking []int
	for _, line := range lines {
		if line[0:5] == "place" {
			numPlaces, _ := strconv.Atoi(line[6:])
			for i := 0; i < numPlaces; i++ {
				places = append(places, pn.Place{Id: "P" + strconv.Itoa(i)})
			}
		} else if len(line) > 10 && line[0:10] == "transition" {
			numTrans, _ := strconv.Atoi(line[11:])
			for i := 0; i < numTrans; i++ {
				transitions = append(transitions, pn.Transition{Id: "T" + strconv.Itoa(i)})
			}
		} else if string(line[0]) == "(" {
			var initMarking string
			initMarking = strings.Trim(line[1:], ")")
			markingArr := strings.Split(initMarking, ",")
			for _, num := range markingArr {
				value, _ := strconv.Atoi(num)
				marking = append(marking, value)
			}
		} else {
			edgeParams := strings.Split(line, " : ")
			weight, _ := strconv.Atoi(edgeParams[2])
			edges = append(edges, pn.Edge{From: edgeParams[0], To: edgeParams[1], Weight: weight})
		}
	}
	pn0 := pn.PetriNet{edges, transitions, places, marking}
	pn0.UpdatePlaces(marking)
	return pn0
}
