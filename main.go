package main

import "os"

const memory = "4 GB"
const basis = "def2-TZVPD"
const energy = "MP2/def2-[TQ]ZVPD"

const ligandCharge = "-1"
const ligandMultiplicity = "1"
const ionCharge = "3"
const ionMultiplicity = "1"

func main() {
	inputFile := os.Args[1]
	outputDir := os.Args[2]

	structName, ligand := readFile(inputFile, false)
	ion := genStartingStructure(ligand, "La", 3, 3.0)
	writeInput(ion, ligand, outputDir, structName)
}