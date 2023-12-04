package main

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/DDP-Projekt/Kompilierer/src/ast"
	"github.com/DDP-Projekt/Kompilierer/src/ddperror"
	"github.com/DDP-Projekt/Kompilierer/src/parser"

	"github.com/google/go-github/v55/github"
)

var nameMap = map[string]map[string]string{
	"DE": {
		"type":        "Typ",
		"var":         "Variablen",
		"func":        "Funktionen",
		"moduleEmpty": "Dieses Modul ist Leer",
	},
	"EN": {
		"type":        "Type",
		"var":         "variables",
		"func":        "functions",
		"moduleEmpty": "This module is empty",
	},
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func clearDirectory(dir string) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if path == dir || d.Name() == "_index.md" {
			return nil
		}

		return os.Remove(path)
	})
}

func main() {
	outputDirDe := "content/DE/Programmierung/Standardbibliothek"
	outputDirEn := "content/EN/Programmierung/Standardbibliothek"

	// delete old articles
	panicIfErr(clearDirectory(outputDirDe))

	gh := github.NewClient(nil)
	_, dir, _, err := gh.Repositories.GetContents(context.Background(), "DDP-Projekt", "Kompilierer", "lib/stdlib/Duden", nil)
	panicIfErr(err)

	for _, entry := range dir {
		outputPathDe := filepath.Join(outputDirDe, strings.Replace(entry.GetName(), "ddp", "md", 1))
		outputPathEn := filepath.Join(outputDirEn, strings.Replace(entry.GetName(), "ddp", "md", 1))

		file, _, _, err := gh.Repositories.GetContents(context.Background(), "DDP-Projekt", "Kompilierer", entry.GetPath(), nil)
		panicIfErr(err)
		inputFile, err := file.GetContent()
		panicIfErr(err)

		MakeMdFiles(inputFile, outputPathDe, "DE")
		MakeMdFiles(inputFile, outputPathEn, "EN")
	}
}

func MakeMdFiles(inputFile, outputFilePath, lang string) {
	fmt.Printf("creating output file '%s'...\n", outputFilePath)
	os.MkdirAll(filepath.Dir(outputFilePath), os.ModeDir|os.ModePerm)
	outputFile, err := os.Create(outputFilePath)
	panicIfErr(err)
	defer outputFile.Close()

	fmt.Printf("parsing module...\n")
	module, err := parser.Parse(parser.Options{
		Source:       []byte(inputFile),
		ErrorHandler: ddperror.EmptyHandler,
	})
	panicIfErr(err)

	// turn the map into a slice
	publicDecls := make([]ast.Declaration, 0, len(module.PublicDecls))
	for _, decl := range module.PublicDecls {
		publicDecls = append(publicDecls, decl)
	}

	// sort the decls by order of appereance
	sort.Slice(publicDecls, func(i, j int) bool {
		return publicDecls[i].GetRange().Start.IsBefore(publicDecls[j].GetRange().Start)
	})

	fmt.Println("writing md file...")

	writeMD(inputFile, outputFile, publicDecls, lang)

	fmt.Println("done writing md file.")
	fmt.Println()
}

func writeMD(inputFile string, outputFile *os.File, publicDecls []ast.Declaration, lang string) {
	funcBldr := &bytes.Buffer{}
	varBldr := &bytes.Buffer{}

	hasVars, hasFuncs := false, false

	for _, decl := range publicDecls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			hasFuncs = true

			descr := ""
			if decl.Comment() != nil {
				descr = strings.Replace(strings.Trim(decl.Comment().String(), "[] \r\n"), "\t", "", -1)
				descr = strings.ReplaceAll(descr, "\r", "")
				descr = strings.ReplaceAll(descr, "\n", "<br>")
				descr = strings.ReplaceAll(descr, "\"", "\\\"")
			}

			params := ""
			paramTypes := ""
			if len(decl.ParamNames) > 0 {
				names := make([]string, 0, len(decl.ParamNames))
				types := make([]string, 0, len(decl.ParamTypes))
				for i := range decl.ParamNames {
					names = append(names, decl.ParamNames[i].String())
					types = append(types, decl.ParamTypes[i].String())
				}

				params = strings.Join(names, ",")
				paramTypes = strings.Join(types, ",")
			}

			aliases := ""
			for i, alias := range decl.Aliases {
				aliases += strings.Trim(alias.Original.Literal, "\"\n")
				if i+1 < len(decl.Aliases) {
					aliases += "\\\""
				}
			}

			impl := ""
			isExtern := ast.IsExternFunc(decl)
			if ast.IsExternFunc(decl) {
				impl = strings.Trim(decl.ExternFile.String(), "\"")
			} else {
				file := strings.Trim(inputFile, "\r\n")
				inputStr := strings.Split(file, "\n")
				lines := inputStr[decl.Body.Range.Start.Line:decl.Body.Range.End.Line]

				impl = strings.Join(lines, "<br>")
				impl = strings.ReplaceAll(impl, "\r", "")
				impl = strings.ReplaceAll(impl, "\n", "<br>")
				impl = strings.ReplaceAll(impl, "\"", "\\\"")
				impl = strings.ReplaceAll(impl, "\t", "&emsp;")
			}

			fmt.Fprintf(funcBldr, "{{< duden-function name=\"%s\" desc=\"%s\" params=\"%s\" paramTypes=\"%s\" ret=\"%s\" impl=\"%s\" extern=\"%t\" aliases=\"%s\" >}}\n\n", // }}"
				decl.Name(), descr, params, paramTypes, decl.Type.String(), impl, isExtern, aliases)
		case *ast.VarDecl:
			hasVars = true

			fmt.Fprintf(varBldr, "## %s\n", decl.Name())
			fmt.Fprintf(varBldr, "* %s: `%s`\n", nameMap[lang]["type"], decl.Type)
			fmt.Fprintln(varBldr, "")
		case *ast.StructDecl:
			panic("TODO: implement struct decl")
		}
	}

	fileName := strings.Replace(filepath.Base(outputFile.Name()), ".md", "", 1)
	fmt.Fprintf(outputFile, "+++\ntitle = \"%s\"\nweight = 1\ntype = \"article\"\n+++\n", fileName)
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
}
