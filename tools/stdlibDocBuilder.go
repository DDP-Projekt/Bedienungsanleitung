package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DDP-Projekt/Kompilierer/src/ast"
	"github.com/DDP-Projekt/Kompilierer/src/ddperror"
	"github.com/DDP-Projekt/Kompilierer/src/parser"
)

func main() {
	fmt.Printf("Argument Count: %d", len(os.Args))

	if len(os.Args) == 1 {

		inputDir := "C:\\Users\\mrshe\\source\\repos\\DDP-Projekt\\Kompilierer\\lib\\stdlib\\Duden\\"
		outputDir := "C:\\Users\\mrshe\\source\\repos\\DDP-Projekt\\Bedienungsanleitung\\Artikel\\DE\\Programmierung\\Standardbibliothek\\"

		files, err := os.ReadDir(inputDir)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			inputPath := filepath.Join(inputDir, file.Name())
			outputPath := filepath.Join(outputDir, strings.Replace(file.Name(), "ddp", "md", 1))

			MakeMdFiles(inputPath, outputPath)
		}

	} else if len(os.Args) == 3 {
		MakeMdFiles(os.Args[1], os.Args[2])
	}
}

func MakeMdFiles(inputFilePath, outputFilePath string) {
	fmt.Println("reading input file...")
	inputFile, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("creating output file...")
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	fmt.Printf("parsing module '%s'...\n", inputFilePath)
	module, err := parser.Parse(parser.Options{
		FileName:     inputFilePath,
		Source:       inputFile,
		ErrorHandler: ddperror.EmptyHandler,
	})
	if err != nil {
		panic(err)
	}

	funcBldr := &bytes.Buffer{}
	varBldr := &bytes.Buffer{}

	hasVars, hasFuncs := false, false

	fmt.Println("writing md file...")

	for name, decl := range module.PublicDecls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			hasFuncs = true

			// Funktionsname
			fmt.Fprintf(funcBldr, "## %s\n", name)

			if len(decl.ParamNames) > 0 {
				// Parameter Names
				fmt.Fprintf(funcBldr, "* Parameter: ")
				for i, paramName := range decl.ParamNames {
					fmt.Fprintf(funcBldr, "%s", paramName)
					if i < len(decl.ParamNames)-1 {
						fmt.Fprintf(funcBldr, ", ")
					}
				}

				// Parameter Types
				fmt.Fprintf(funcBldr, "\n* Parameter Typ/en: ")
				for i, paramTypes := range decl.ParamTypes {
					fmt.Fprintf(funcBldr, "%s", paramTypes)
					if i < len(decl.ParamTypes)-1 {
						fmt.Fprintf(funcBldr, ", ")
					}
				}
			}

			// Rückgabe Typ
			fmt.Fprintf(funcBldr, "\n* Rückgabe Typ: %s\n\n", decl.Type)
		case *ast.VarDecl:
			hasVars = true

			fmt.Fprintf(varBldr, "## %s\n", name)
			fmt.Fprintf(varBldr, "* Typ: %s\n", decl.Type)
		}
	}

	if hasVars {
		fmt.Fprintf(outputFile, "# Variablen\n")
		fmt.Fprintln(outputFile, varBldr)
	}
	if hasFuncs {
		fmt.Fprintf(outputFile, "# Funktionen\n")
		fmt.Fprintln(outputFile, funcBldr)
	}
	if !hasFuncs && !hasVars {
		fmt.Fprintf(outputFile, "Dieses Modul ist Leer")
	}

	fmt.Println("done writing md file.")
}
