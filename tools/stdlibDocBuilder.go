package main

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/DDP-Projekt/Kompilierer/src/ast"
	"github.com/DDP-Projekt/Kompilierer/src/ddperror"
	"github.com/DDP-Projekt/Kompilierer/src/parser"
)

func clearDirectory(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(dir, os.ModeDir)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Printf("Argument Count: %d\n", len(os.Args))

	if len(os.Args) == 1 {

		// Folder structure:
		// - Parentfolder
		//   - Kompilierer
		//   - Bedienungsanleitung
		inputDir := "../../Kompilierer/lib/stdlib/Duden/"
		outputDir := "../Artikel/DE/Programmierung/Standardbibliothek/"

		// delete old articles
		clearDirectory(outputDir)

		files, err := os.ReadDir(inputDir)
		if err != nil {
			panic(err)
		}
		fileNames := make([]string, 0, len(files))
		for _, file := range files {
			inputPath := filepath.Join(inputDir, file.Name())
			outputPath := filepath.Join(outputDir, strings.Replace(file.Name(), "ddp", "md", 1))

			MakeMdFiles(inputPath, outputPath)
			fileNames = append(fileNames, strings.TrimSuffix(file.Name(), ".ddp"))
		}
		// correct the index
		MakeIndexHtml(fileNames, "../index.html")

	} else if len(os.Args) == 3 {
		// Args[1]: Input, Args[2]: Output
		MakeMdFiles(os.Args[1], os.Args[2])
		MakeIndexHtml([]string{os.Args[1]}, "../index.html")
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
				fmt.Fprintf(funcBldr, "\t<li>Parameter: ")
				for i, paramName := range decl.ParamNames {
					fmt.Fprintf(funcBldr, "<code>%s</code>", paramName)
					if i < len(decl.ParamNames)-1 {
						fmt.Fprintf(funcBldr, ", ")
					}
				}

				// Parameter Types
				fmt.Fprintf(funcBldr, "</li>\n\t<li>Parameter Typ")
				if len(decl.ParamNames) > 1 {
					fmt.Fprintf(funcBldr, "en")
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
			fmt.Fprintf(funcBldr, "</li>\n\t<li>Rückgabe Typ: <code>%s</code></li>\n", decl.Type)
			fmt.Fprintln(funcBldr, "</ul>")

			// Aliases
			fmt.Fprintln(funcBldr, "\n<h3>Aliase</h3>\n<ol>")
			for _, alias := range decl.Aliases {
				fmt.Fprintf(funcBldr, "\t<li><code>%s</code></li>\n", html.EscapeString(alias.Original.Literal))
			}
			fmt.Fprintln(funcBldr, "</ol>")

			// Implemetation
			fmt.Fprintln(funcBldr, "\n<h3>Implementation</h3>")
			if ast.IsExternFunc(decl) {
				fmt.Fprintf(funcBldr, "Implementiert in <code>%s</code>\n", decl.ExternFile)
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
			fmt.Fprintf(varBldr, "* Typ: `%s`\n", decl.Type)
			fmt.Fprintln(varBldr, "")
		}
	}

	fileName := strings.Replace(filepath.Base(inputFilePath), ".ddp", "", 1)
	if hasVars {
		fmt.Fprintf(outputFile, "# Duden/%s Variablen\n", fileName)
		fmt.Fprintln(outputFile, varBldr)
	}
	if hasFuncs {
		fmt.Fprintf(outputFile, "# Duden/%s Funktionen\n", fileName)
		fmt.Fprintln(outputFile, funcBldr)
	}
	if !hasFuncs && !hasVars {
		fmt.Fprintf(outputFile, "Dieses Modul ist Leer")
	}

	fmt.Println("done writing md file.")
}

func MakeIndexHtml(dudenFiles []string, outPath string) {
	tmpl := template.Must(template.ParseFiles("index.gohtml"))
	file, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0700)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = tmpl.Execute(file, struct {
		StdlibModules []string
	}{
		StdlibModules: dudenFiles,
	})
	if err != nil {
		panic(err)
	}
}
