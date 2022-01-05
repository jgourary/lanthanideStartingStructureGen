package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

// will place ion at specified distance from numbered (from zero) atom in ligand
func genStartingStructure(ligand molecule, ionElement string, interactingAtom int, dist float64) molecule {
	endAtom := ligand.atoms[interactingAtom]

	var startAtom atom
	foundHeavyAtom := false
	for _, bond := range endAtom.bondedAtoms {
		if ligand.atoms[bond].element == "C" {
			startAtom = ligand.atoms[bond]
			foundHeavyAtom = true
		}
	}
	if !foundHeavyAtom {
		err := errors.New("failed to find linked heavy atom to target atom in order to align ion with")
		log.Fatal(err)
		var ionMol molecule
		return ionMol
	} else {
		// get unit vector
		unitVector := getUnitAxis(startAtom.pos, endAtom.pos)

		// add ion
		var ion atom

		ion.element = ionElement
		ion.pos = getIonPos(endAtom.pos, unitVector, dist)
		ion.bondedAtoms = []int{}

		// check dist and report
		nearestAtom, nearestDist := getClosestLigandAtom(ligand, ion.pos)
		fmt.Println("Placed ion...")
		fmt.Println("Closest ligand atom was atom #" + strconv.Itoa(nearestAtom) + " (element = " + ligand.atoms[nearestAtom].element + ") at a distance of " + fmt.Sprintf("%.3f", nearestDist) + " A")

		// add ion to ligand
		var ionMol molecule
		ionAtoms := make(map[int]atom)
		ionAtoms[0] = ion
		ionMol.atoms = ionAtoms
		ionMol.charge = ionCharge
		ionMol.multiplicity = ionMultiplicity

		return ionMol
	}
}

func getClosestLigandAtom(ligand molecule, pos []float64) (int, float64) {
	nearestDist := 1000000.0
	nearestAtom := -1
	for i, atom := range ligand.atoms {
		dist := getDistance(atom.pos, pos)
		if dist < nearestDist {
			nearestAtom = i
			nearestDist = dist
		}
	}
	return nearestAtom, nearestDist
}

func getIonPos(endPos []float64, unitVector []float64, dist float64) []float64 {
	ionPos := make([]float64,3)
	ionPos[0] = endPos[0] + unitVector[0] * dist
	ionPos[1] = endPos[1] + unitVector[1] * dist
	ionPos[2] = endPos[2] + unitVector[2] * dist
	return ionPos
}

func getUnitAxis(startPos []float64, endPos []float64) []float64 {
	closestDist := getDistance(endPos, startPos)

	vector := []float64{(endPos[0]-startPos[0])/closestDist, (endPos[1]-startPos[1]) / closestDist, (endPos[2]-startPos[2])/closestDist}
	return vector
}

func getDistance(pos1 []float64, pos2 []float64) float64 {
	dx2 := math.Pow(pos1[0] - pos2[0],2)
	dy2 := math.Pow(pos1[1] - pos2[1],2)
	dz2 := math.Pow(pos1[2] - pos2[2],2)
	return math.Sqrt(dx2+dy2+dz2)
}
