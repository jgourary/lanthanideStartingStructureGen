package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func writeInput(ion molecule, ligand molecule, outDir string, structName string) {
	outPath := filepath.Join(outDir, structName + ".inp")
	_ = os.MkdirAll(outDir, 0755)
	thisFile, err := os.Create(outPath)
	if err != nil {
		fmt.Println("Failed to create new fragment file: " + outPath)
		log.Fatal(err)
	}
	_, _ = thisFile.WriteString("memory " + memory + "\n")
	_, _ = thisFile.WriteString("set basis " + basis + "\n")
	_, _ = thisFile.WriteString("molecule{\n")
	_, _ = thisFile.WriteString("\t" + ion.charge + " " + ion.multiplicity + "\n")
	for _, atom := range ion.atoms {
		_, _ = thisFile.WriteString("\t" + atom.element + " " + fmt.Sprintf("%.6f", atom.pos[0]) + " " +
			fmt.Sprintf("%.6f", atom.pos[1])  + " " + fmt.Sprintf("%.6f", atom.pos[2])  + "\n")
	}
	_, _ = thisFile.WriteString("\t--\n")
	_, _ = thisFile.WriteString("\t" + ligand.charge + " " + ligand.multiplicity + "\n")
	for _, atom := range ligand.atoms {
		_, _ = thisFile.WriteString("\t" + atom.element + " " + fmt.Sprintf("%.6f", atom.pos[0]) + " " +
			fmt.Sprintf("%.6f", atom.pos[1])  + " " + fmt.Sprintf("%.6f", atom.pos[2])  + "\n")
	}
	_, _ = thisFile.WriteString("}\n\n")
	_, _ = thisFile.WriteString("energy('" + energy + "')\n")
}
