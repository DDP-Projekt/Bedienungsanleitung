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
	fmt.Printf("Argument Count: %d\n", len(os.Args))

	if len(os.Args) == 1 {

		// Folder structure:
		// - Parentfolder
		//   - Kompilierer
		//   - Bedienungsanleitung
		inputDir := "../../Kompilierer/lib/stdlib/Duden/"
		outputDir := "../Artikel/DE/Programmierung/Standardbibliothek/"

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
		// Args[1]: Input, Args[2]: Output
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

			// Kommentar/Beschreibung
			if decl.Comment != nil {
				descr := strings.Replace(strings.Trim(decl.Comment.String(), "[] \r\n"), "\t", "", -1)
				fmt.Fprintf(funcBldr, "```\n%s\n```\n", descr)
			}

			if len(decl.ParamNames) > 0 {
				// Parameter Names
				fmt.Fprintf(funcBldr, "* Parameter: ")
				for i, paramName := range decl.ParamNames {
					fmt.Fprintf(funcBldr, "`%s`", paramName)
					if i < len(decl.ParamNames)-1 {
						fmt.Fprintf(funcBldr, ", ")
					}
				}

				// Parameter Types
				fmt.Fprintf(funcBldr, "\n* Parameter Typ")
				if len(decl.ParamNames) > 1 {
					fmt.Fprintf(funcBldr, "en")
				}
				fmt.Fprintf(funcBldr, ": ")

				for i, paramTypes := range decl.ParamTypes {
					fmt.Fprintf(funcBldr, "`%s`", paramTypes)
					if i < len(decl.ParamTypes)-1 {
						fmt.Fprintf(funcBldr, ", ")
					}
				}
			}

			// Rückgabe Typ
			fmt.Fprintf(funcBldr, "\n* Rückgabe Typ: `%s`\n\n", decl.Type)

			// Aliases
			fmt.Fprintln(funcBldr, "### Aliase")
			for i, alias := range decl.Aliases {
				fmt.Fprintf(funcBldr, "%d. `%s`\n", i+1, alias.Original.Literal)
			}
			fmt.Fprintln(funcBldr, "")

			// Implemetation
			fmt.Fprintln(funcBldr, "### Implementation")
			if ast.IsExternFunc(decl) {
				fmt.Fprintf(funcBldr, "Implementiert in %s\n", decl.ExternFile)
			} else {
				fmt.Fprintln(funcBldr, "```ddp")

				inputStr := strings.Split(string(inputFile), "\n")
				lines := inputStr[decl.Body.Range.Start.Line:decl.Body.Range.End.Line]
				fmt.Fprintf(funcBldr, "%s", strings.Join(lines, "\n"))

				fmt.Fprintln(funcBldr, "\n```")
			}

		case *ast.VarDecl:
			hasVars = true

			fmt.Fprintf(varBldr, "## %s\n", name)
			fmt.Fprintf(varBldr, "* Typ: `%s`\n", decl.Type)
			fmt.Fprintln(varBldr, "")
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
