package main

import (
	"bytes"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/DDP-Projekt/Kompilierer/src/ast"
	"github.com/DDP-Projekt/Kompilierer/src/ddperror"
	"github.com/DDP-Projekt/Kompilierer/src/parser"
)

var nameMap = map[string]map[string]string{
	"DE": {
		"params":      "Parameter",
		"paramType":   "Parameter Typ",
		"returnType":  "Rückgabe Typ",
		"aliases":     "Aliase",
		"impl":        "Implementation",
		"externImpl":  "Implementiert in",
		"type":        "Typ",
		"var":         "Variablen",
		"func":        "Funktionen",
		"moduleEmpty": "Dieses Modul ist Leer",
	},
	"EN": {
		"params":      "Parameters",
		"paramType":   "Parameter type",
		"returnType":  "Return type",
		"aliases":     "Aliases",
		"impl":        "Implementation",
		"externImpl":  "Implemented in",
		"type":        "Type",
		"var":         "variables",
		"func":        "functions",
		"moduleEmpty": "This module is empty",
	},
}

func clearDirectory(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Printf("Argument Count: %d\n", len(os.Args))

	if len(os.Args) != 1 {
		return
	}

	inputDir := os.Getenv("DDPPATH") + "Duden/"
	outputDirDe := "../content/DE/Programmierung/Standardbibliothek/gen"
	outputDirEn := "../content/EN/Programmierung/Standardbibliothek/gen"

	files, err := os.ReadDir(inputDir)
	if err != nil {
		panic(err)
	}

	// delete old articles
	clearDirectory(outputDirDe)

	for _, file := range files {
		inputPath := filepath.Join(inputDir, file.Name())
		outputPathDe := filepath.Join(outputDirDe, strings.Replace(file.Name(), "ddp", "md", 1))
		outputPathEn := filepath.Join(outputDirEn, strings.Replace(file.Name(), "ddp", "md", 1))

		MakeMdFiles(inputPath, outputPathDe, "DE")
		MakeMdFiles(inputPath, outputPathEn, "EN")
	}
}

func MakeMdFiles(inputFilePath, outputFilePath, lang string) {
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

	// turn the map into a slice
	publicDecls := make([]ast.Declaration, 0, len(module.PublicDecls))
	for _, decl := range module.PublicDecls {
		publicDecls = append(publicDecls, decl)
	}

	// sort the decls by order of appereance
	sort.Slice(publicDecls, func(i, j int) bool {
		return publicDecls[i].GetRange().Start.IsBefore(publicDecls[j].GetRange().Start)
	})

	for _, decl := range publicDecls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			hasFuncs = true

			fmt.Fprintln(funcBldr, "<details>")
			// Funktionsname
			fmt.Fprintf(funcBldr, "<summary><h2>%s</h2></summary>\n", decl.Name())

			fmt.Fprintln(funcBldr, "<ul>")
			// Kommentar/Beschreibung
			if decl.Comment != nil {
				descr := strings.Replace(strings.Trim(decl.Comment.String(), "[] \r\n"), "\t", "", -1)
				fmt.Fprintf(funcBldr, "<pre>\n%s\n</pre>\n", descr)
			}

			if len(decl.ParamNames) > 0 {
				// Parameter Names
				fmt.Fprintf(funcBldr, "\t<li>%s: ", nameMap[lang]["params"])
				for i, paramName := range decl.ParamNames {
					fmt.Fprintf(funcBldr, "<code>%s</code>", paramName)
					if i < len(decl.ParamNames)-1 {
						fmt.Fprintf(funcBldr, ", ")
					}
				}

				// Parameter Types
				fmt.Fprintf(funcBldr, "</li>\n\t<li>%s", nameMap[lang]["paramType"])
				if len(decl.ParamNames) > 1 {
					if lang == "DE" {
						fmt.Fprintf(funcBldr, "en")
					} else if lang == "EN" {
						fmt.Fprintf(funcBldr, "s")
					}
				}
				fmt.Fprintf(funcBldr, ": ")

				for i, paramTypes := range decl.ParamTypes {
					fmt.Fprintf(funcBldr, "<code>%s</code>", paramTypes)
					if i < len(decl.ParamTypes)-1 {
						fmt.Fprintf(funcBldr, ", ")
					}
				}
			}

			// Rückgabe Typ
			fmt.Fprintf(funcBldr, "</li>\n\t<li>%s: <code>%s</code></li>\n", nameMap[lang]["returnType"], decl.Type)
			fmt.Fprintln(funcBldr, "</ul>")

			// Aliases
			fmt.Fprintf(funcBldr, "\n<h3>%s</h3>\n<ol>\n", nameMap[lang]["aliases"])
			for _, alias := range decl.Aliases {
				fmt.Fprintf(funcBldr, "\t<li><code>%s</code></li>\n", html.EscapeString(alias.Original.Literal))
			}
			fmt.Fprintln(funcBldr, "</ol>")

			// Implemetation
			fmt.Fprintf(funcBldr, "\n<h3>%s</h3>\n", nameMap[lang]["impl"])
			if ast.IsExternFunc(decl) {
				fmt.Fprintf(funcBldr, "%s <code>%s</code>\n", nameMap[lang]["externImpl"], decl.ExternFile)
			} else {
				fmt.Fprintln(funcBldr, "<pre class=\"language-ddp\" tabindex=\"0\">\n<code class=\"language-ddp\">")

				inputStr := strings.Split(string(inputFile), "\n")
				lines := inputStr[decl.Body.Range.Start.Line:decl.Body.Range.End.Line]
				for _, line := range lines {
					fmt.Fprintln(funcBldr, strings.Replace(line, "\t", "", 1))
				}

				fmt.Fprintln(funcBldr, "\n</code>\n</pre>")
			}

			fmt.Fprint(funcBldr, "</details>\n\n")
		case *ast.VarDecl:
			hasVars = true

			fmt.Fprintf(varBldr, "## %s\n", decl.Name())
			fmt.Fprintf(varBldr, "* %s: `%s`\n", nameMap[lang]["type"], decl.Type)
			fmt.Fprintln(varBldr, "")
		}
	}

	fileName := strings.Replace(filepath.Base(inputFilePath), ".ddp", "", 1)
	fmt.Fprintf(outputFile, "+++\ntitle = \"%s\"\nweight = 1\n+++\n", fileName)
	if hasVars {
		fmt.Fprintf(outputFile, "# Duden/%s %s\n", fileName, nameMap[lang]["var"])
		fmt.Fprintln(outputFile, varBldr)
	}
	if hasFuncs {
		fmt.Fprintf(outputFile, "# Duden/%s %s\n", fileName, nameMap[lang]["func"])
		fmt.Fprintln(outputFile, funcBldr)
	}
	if !hasFuncs && !hasVars {
		fmt.Fprintf(outputFile, nameMap[lang]["moduleEmpty"])
	}

	fmt.Println("done writing md file.")
}
