package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/fs"
	"log"
	"maps"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"sync"

	"github.com/DDP-Projekt/Kompilierer/src/ast"
	"github.com/DDP-Projekt/Kompilierer/src/ddperror"
	"github.com/DDP-Projekt/Kompilierer/src/parser"
)

var nameMap = map[string]map[string]string{
	"DE": {
		"type":        "Typ",
		"value":       "Wert",
		"var":         "Variablen",
		"const":       "Konstante",
		"func":        "Funktionen",
		"comb":        "Kombinationen",
		"moduleEmpty": "Dieses Modul ist Leer",
	},
	"EN": {
		"type":        "Type",
		"value":       "Value",
		"var":         "variables",
		"const":       "constants",
		"func":        "functions",
		"comb":        "combinations",
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

type GitHubDir struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func main() {
	outputDirDe := "content/DE/Programmierung/Standardbibliothek"
	outputDirEn := "content/EN/Programmierung/Standardbibliothek"

	// delete old articles
	panicIfErr(clearDirectory(outputDirDe))

	// get all files in directory
	branch := "master"
	resp, err := http.Get("https://api.github.com/repos/DDP-Projekt/Kompilierer/contents/lib/stdlib/Duden?ref=" + branch)
	panicIfErr(err)
	defer resp.Body.Close()

	dir := make([]GitHubDir, 0)
	json.NewDecoder(resp.Body).Decode(&dir)

	wg := sync.WaitGroup{}
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// iterate through all files and generate MD files
	for _, entry := range dir {
		entry := entry // wtf???
		wg.Add(1)
		go func() {
			outputPathDe := filepath.Join(outputDirDe, strings.Replace(entry.Name, "ddp", "md", 1))
			outputPathEn := filepath.Join(outputDirEn, strings.Replace(entry.Name, "ddp", "md", 1))

			resp, err := http.Get("https://raw.githubusercontent.com/DDP-Projekt/Kompilierer/" + branch + "/" + entry.Path)
			panicIfErr(err)
			defer resp.Body.Close()
			inputFile, err := io.ReadAll(resp.Body)
			panicIfErr(err)

			log.Printf("parsing module %s...\n", entry.Name)
			module, err := parser.Parse(parser.Options{
				Source:       inputFile,
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

			makeMdFiles(publicDecls, inputFile, outputPathDe, "DE")
			makeMdFiles(publicDecls, inputFile, outputPathEn, "EN")

			wg.Done()
		}()
	}
	wg.Wait()
}

func makeMdFiles(publicDecls []ast.Declaration, inputFile []byte, outputFilePath, lang string) {
	log.Printf("creating output file '%s'...\n", outputFilePath)

	os.MkdirAll(filepath.Dir(outputFilePath), os.ModeDir|os.ModePerm)
	outputFile, err := os.Create(outputFilePath)
	panicIfErr(err)
	defer outputFile.Close()

	log.Printf("writing '%s' file...\n", outputFilePath)

	writeMD(string(inputFile), outputFile, publicDecls, lang)

	log.Printf("done writing '%s' file...\n", outputFilePath)
}

func writeMD(inputFile string, outputFile *os.File, publicDecls []ast.Declaration, lang string) {
	funcBldr, varBldr, constBldr, structBldr := &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}

	hasStructs, hasVars, hasConsts, hasFuncs := false, false, false, false

	for _, decl := range publicDecls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			hasFuncs = true

			descr := ""
			if decl.Comment() != nil {
				descr = strings.Replace(strings.Trim(decl.Comment().String(), "[] \r\n"), "\t", "", -1)
				descr = html.EscapeString(descr)
				descr = strings.ReplaceAll(descr, "\r", "")
				descr = strings.ReplaceAll(descr, "\n", "<br>")
				descr = strings.ReplaceAll(descr, "\"", "\\\"")
			}

			params := ""
			paramTypes := ""
			if len(decl.Parameters) > 0 {
				names := make([]string, 0, len(decl.Parameters))
				types := make([]string, 0, len(decl.Parameters))
				for _, param := range decl.Parameters {
					names = append(names, "<code>"+param.Name.String()+"</code>")
					types = append(types, "<code>"+param.Type.Type.String()+"</code>")
				}

				params = strings.Join(names, ", ")
				paramTypes = strings.Join(types, ", ")
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
			genericTypes := ""
			if isExtern {
				impl = strings.Trim(decl.ExternFile.String(), "\"")
			} else {
				file := strings.Trim(inputFile, "\r\n")
				inputStr := strings.Split(file, "\n")

				var start, end uint
				if !ast.IsGeneric(decl) {
					start, end = decl.Body.Range.Start.Line, decl.Body.Range.End.Line
				} else {
					genericTypes = "<code>" + strings.Join(slices.Collect(maps.Keys(decl.Generic.Types)), "</code>, <code>") + "</code>"
					start, end = decl.Generic.Tokens[0].Range.Start.Line, decl.Generic.Tokens[len(decl.Generic.Tokens)-1].Range.End.Line
				}
				lines := inputStr[start:end]

				impl = strings.Join(lines, "&#10;")
				impl = strings.ReplaceAll(impl, "\r", "")
				impl = strings.ReplaceAll(impl, "\n", "&#10;")
				impl = strings.ReplaceAll(impl, "\"", "\\\"")
				impl = strings.ReplaceAll(impl, "\t", "    ")
			}

			fmt.Fprintf(funcBldr, "{{< duden-function name=\"%s\" desc=\"%s\" params=\"%s\" paramTypes=\"%s\" genericTypes=\"%s\" ret=\"%s\" impl=\"%s\" extern=\"%t\" aliases=\"%s\" >}}\n\n", // }}"
				decl.Name(), descr, params, paramTypes, genericTypes, decl.ReturnType.String(), impl, isExtern, aliases)
		case *ast.VarDecl:
			hasVars = true

			fmt.Fprintf(varBldr, "## %s\n", decl.Name())
			fmt.Fprintf(varBldr, "* %s: <code>%s</code>\n", nameMap[lang]["type"], decl.Type)
			fmt.Fprintln(varBldr, "")
		case *ast.ConstDecl:
			hasConsts = true

			fmt.Fprintf(constBldr, "## %s\n", decl.Name())
			fmt.Fprintf(constBldr, "* %s: <code>%s</code>\n", nameMap[lang]["type"], decl.Type)
			fmt.Fprintf(constBldr, "* %s: <code>%s</code>\n", nameMap[lang]["value"], decl.Val.Token().Literal)
			fmt.Fprintln(constBldr, "")
		case *ast.StructDecl:
			hasStructs = true

			fields := ""
			for _, field := range decl.Fields {
				switch field := field.(type) {
				case *ast.VarDecl:
					if field.Public() {
						inputStr := strings.Split(inputFile, "\n")
						declRange := field.GetRange()
						declCode := inputStr[declRange.Start.Line-1][declRange.Start.Column-1 : declRange.End.Column]
						fields += html.EscapeString(declCode) + "|"
					}
				default:
				}
			}
			fields = strings.TrimRight(fields, "|")

			descr := ""
			if decl.Comment() != nil {
				descr = strings.Replace(strings.Trim(decl.Comment().String(), "[] \r\n"), "\t", "", -1)
				descr = html.EscapeString(descr)
				descr = strings.ReplaceAll(descr, "\r", "")
				descr = strings.ReplaceAll(descr, "\n", "<br>")
				descr = strings.ReplaceAll(descr, "\"", "\\\"")
			}

			aliases := ""
			for i, alias := range decl.Aliases {
				aliases += strings.Trim(alias.Original.Literal, "\"\n")
				if i+1 < len(decl.Aliases) {
					aliases += "\\\""
				}
			}

			fmt.Fprintf(structBldr, "{{< duden-combination name=\"%s\" desc=\"%s\" fields=\"%s\" aliases=\"%s\" >}}\n\n", // }}"
				decl.Name(), descr, fields, aliases)
		}
	}

	fileName := strings.Replace(filepath.Base(outputFile.Name()), ".md", "", 1)
	fmt.Fprintf(outputFile, "+++\ntitle = \"%s\"\nweight = 1\n+++\n", fileName)
	if hasStructs {
		fmt.Fprintf(outputFile, "# Duden/%s %s\n", fileName, nameMap[lang]["comb"])
		fmt.Fprintln(outputFile, structBldr)
	}
	if hasVars {
		fmt.Fprintf(outputFile, "# Duden/%s %s\n", fileName, nameMap[lang]["var"])
		fmt.Fprintln(outputFile, varBldr)
	}
	if hasConsts {
		fmt.Fprintf(outputFile, "# Duden/%s %s\n", fileName, nameMap[lang]["const"])
		fmt.Fprintln(outputFile, constBldr)
	}
	if hasFuncs {
		fmt.Fprintf(outputFile, "# Duden/%s %s\n", fileName, nameMap[lang]["func"])
		fmt.Fprintln(outputFile, funcBldr)
	}
	if !hasFuncs && !hasVars && hasConsts && !hasStructs {
		fmt.Fprint(outputFile, nameMap[lang]["moduleEmpty"])
	}
}
