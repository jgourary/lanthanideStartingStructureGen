package main

type molecule struct {
	charge string
	multiplicity string
	atoms map[int]atom

	shift float64
}

type atom struct {
	element string
	pos []float64
	bondedAtoms []int
}

func copyAtom(oldAtom atom) atom {
	var newAtom atom
	newAtom.element = oldAtom.element
	newAtom.pos = make([]float64, 3)
	newAtom.pos[0] = oldAtom.pos[0]
	newAtom.pos[1] = oldAtom.pos[1]
	newAtom.pos[2] = oldAtom.pos[2]
	copy(newAtom.bondedAtoms, oldAtom.bondedAtoms)
	return newAtom
}

func copyMolecule(oldMol molecule) molecule {
	var newMol molecule

	newMol.charge = oldMol.charge
	newMol.multiplicity = oldMol.multiplicity
	newMol.atoms = make(map[int]atom)

	for i, _ := range oldMol.atoms {
		newMol.atoms[i] = copyAtom(oldMol.atoms[i])
	}
	return newMol
}