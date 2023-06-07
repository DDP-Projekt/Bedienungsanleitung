import re
import sys
import codecs
import os

def makeMdFile(inputFilePath, outputFilePath):
	print("opening input file...")
	inputFile = codecs.open(inputFilePath, "r", "utf-8")
	print("opening output file...")
	outputFile = codecs.open(outputFilePath, "w", "utf-8")

	print("reading input file...")
	src = inputFile.readlines()

	print("compiling regexes...")
	funcRegexes = [
						re.compile("Die\\s+((ö|(oe))ffentliche\\s+)?Funktion\\s+([\\wäöüÄÖÜ]+)\\s+gibt\\s+(nichts|einen?\\s+[\\wäöüÄÖÜ]+)\\s+zur(ü|(ue))ck,(\\s*macht:)?"),
						re.compile("Die\\s+((ö|(oe))ffentliche\\s+)?Funktion\\s+([\\wäöüÄÖÜ]+)\\s+mit\\s+dem\\s+Parameter\\s+([\\wäöüÄÖÜ]+)\\s+vom\\s+Typ\\s+([\\w\\säöüÄÖÜ]+),\\s*gibt\\s+(nichts|einen?\\s+[\\w\\säöüÄÖÜ]+)\\s+zur(ü|(ue))ck,(\\s*macht:)?"),
						re.compile("Die\\s+((ö|(oe))ffentliche\\s+)?Funktion\\s+([\\wäöüÄÖÜ]+)\\s+mit\\s+den\\s+Parametern\\s+(\\w+(?:,\\s*[\\wäöüÄÖÜ]+)*)\\s+und\\s+([\\wäöüÄÖÜ]+)\\s+vom\\s+Typ\\s+([\\w\\säöüÄÖÜ]+(?:,\\s*[\\w\\säöüÄÖÜ]+)*)\\s+und\\s+([\\w\\säöüÄÖÜ]+),\\s*gibt\\s+(nichts|einen? [\\w\\säöüÄÖÜ]+)\\s+zur(ü|(ue))ck,(\\s*macht:)?")
					]

	print("writing MD file...")
	outputFile.write("# Duden/Ausgabe Funktionen\n")

	for line in src:
		line = line.strip()
		m1 = funcRegexes[0].match(line)
		m2 = funcRegexes[1].match(line)
		m3 = funcRegexes[2].match(line)

		funcName = ""
		funcParam = "keine"
		funcParamType = "keine"
		funcRet = ""
		if m1:
			funcName = m1.group(4)
			funcRet = m1.group(5)
		elif m2:
			funcName = m2.group(4)
			funcParam = m2.group(5)
			funcParamType = m2.group(6)
			funcRet = m2.group(7)
		elif m3:
			funcName = m3.group(4)
			funcParam = m3.group(5) + ", " + m3.group(6)
			funcParamType = m3.group(7) + ", " + m3.group(8)
			funcRet = m3.group(9)
		else:
			continue

		outputFile.write(f"## {funcName}\n* Parameter: {funcParam}\n* Parameter Typ(en): {funcParamType}\n* Rückgabe Typ: {funcRet}\n\n")

	print("finished writing MD file.")

print(f"Argument count: {len(sys.argv)}")

if len(sys.argv) == 1:
	inputDir = "C:\\Users\\mrshe\\source\\repos\\DDP-Projekt\\Kompilierer\\lib\\stdlib\\Duden\\"
	outputDir = "C:\\Users\\mrshe\\source\\repos\\DDP-Projekt\\Bedienungsanleitung\\Artikel\\DE\\Programmierung\\Standardbibliothek\\"

	inputFiles = os.listdir(inputDir)
	for file in inputFiles:
		inputPath = os.path.join(inputDir, file)
		outputPath = os.path.join(outputDir, file.replace("ddp", "md"))
		makeMdFile(inputPath, outputPath)

if len(sys.argv) == 3:
	makeMdFile(sys.argv[1], sys.argv[2]) 