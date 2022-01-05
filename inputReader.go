package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readFile(filePath string, verbose bool) (string, molecule) {
	// open file
	file, err := os.Open(filePath)
	fmt.Println("Reading file at " + filePath)
	if err != nil {
		fmt.Println("Failed to open molecule file: " + filePath)
		log.Fatal(err)
	}
	structureName := strings.Split(filepath.Base(filePath),".")[0]

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	// create line counter
	i := 1

	var molecules molecule

	for scanner.Scan() {
		// get next line
		line := scanner.Text()
		// split by whitespace
		tokens := strings.Fields(line)
		if len(tokens) > 5 {

			// create new atom
			var newAtom atom

			newAtom.element = tokens[0]
			pos := make([]float64,3)
			for j := 1; j < 4; j++ {
				pos[j-1], err = strconv.ParseFloat(tokens[j],64)
				if err != nil {
					newErr := errors.New("Failed to convert \"" + tokens[j] + "\" in position 0 on line " + strconv.Itoa(i) + " to a float64")
					log.Fatal(newErr)
				}
			}
			newAtom.pos = pos

			// add atom to map
			if verbose{
				fmt.Println("Found new atom " + newAtom.element + " at line " + strconv.Itoa(i) + " - adding to pos " + strconv.Itoa(atomCounter))
			}

			molecules[molCounter].atoms[atomCounter] = newAtom
			atomCounter++

		}
		i++
	}

	return structureName, molecules
}

func printMolecule(molecule molecule) {
	fmt.Println("\t" + molecule.charge + " " + molecule.multiplicity)
	for j, atom := range molecule.atoms {
		fmt.Println("\t" + strconv.Itoa(j) + " " + atom.element + " " + fmt.Sprintf("%.6f", atom.pos[0]) + " " +
			fmt.Sprintf("%.6f", atom.pos[1]) + " " + fmt.Sprintf("%.6f", atom.pos[2]))
	}
}
