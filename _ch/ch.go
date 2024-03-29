package main

import (
	"bufio"
	"fmt"
	"os"
	fp "path/filepath"
	"sort"
	str "strings"
)

type itemT struct {
	section string
	desc    string
	info    string
}

type argsT struct {
	checkOnly     bool
	getAllFiles   bool
	caseSensitive bool
	query         []string
	queryNeg      []string
	rootDir       string
	cheatDirs     []string
	files         []string
	filesFixed    []string
}

type posT struct {
	start int
	end   int
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		printHelp()
		os.Exit(0)
	}
	args := parseArgs()

	var items []itemT
	for _, file := range args.files {
		if args.checkOnly {
			fmt.Printf("checking %s...\n",
				str.TrimPrefix(file, args.rootDir+"/"))
			checkFile(file)
			continue
		}
		items = append(items, parseFile(args, file)...)
	}

	if args.checkOnly {
		fmt.Println("all ok.")
	}

	if len(items) > 0 {
		printItems(items, args)
	}
}

func checkFile(file string) {
	var section string
	var i, emptyCount int
	var isSection, parStart bool

	fd, err := os.Open(file)
	errExit(err)
	defer fd.Close()

	input := bufio.NewScanner(fd)
	for input.Scan() {
		i++
		line := input.Text()

		// every file must start with section title
		if i == 1 {
			if len(line) == 0 {
				msg := "line 1: can't be empty\n"
				errExit(fmt.Errorf(msg))
			}
			continue
		}
		if i == 2 && !str.HasPrefix(line, "==") {
			msg := "line 1-2: must be a section title\n"
			errExit(fmt.Errorf(msg))
		}

		// check for correct spacing between paragraphs
		if line != "" {
			if emptyCount == 2 || emptyCount > 3 {
				msg := "line %d: wrong paragraph separation\n"
				errExit(fmt.Errorf(msg, i))
			}
			if emptyCount == 1 {
				parStart = true
			} else if emptyCount == 3 {
				section = line
				isSection = true
			}
			emptyCount = 0
		} else {
			emptyCount++
			continue
		}

		// every paragraph must start with a description
		if parStart && !str.HasPrefix(line, "# ") {
			msg := "line %d: paragraph must have a description\n"
			errExit(fmt.Errorf(msg, i))
		}

		// section title must be underlined
		if isSection && line != section && !str.HasPrefix(line, "==") {
			msg := "line %d: section title must be followed by "
			msg += "a line with ===\n"
			errExit(fmt.Errorf(msg, i))
		}
		if isSection && line != section && len(line) != len(section) {
			msg := "line %d: section underscore must be the same "
			msg += "length as the section title\n"
			errExit(fmt.Errorf(msg, i))
		}

		// clear section title after underline
		if isSection && line != section {
			isSection = false
		}

		// no trailing spaces or tabs
		if str.HasSuffix(line, " ") || str.HasSuffix(line, "\t") {
			msg := "line %d: trailing tab or space\n"
			errExit(fmt.Errorf(msg, i))
		}

		parStart = false
	}
}

func parseFile(args argsT, file string) []itemT {
	var section, prevLine string
	var item itemT
	var items []itemT
	var i, emptyCount int

	fd, err := os.Open(file)
	errExit(err)
	defer fd.Close()

	input := bufio.NewScanner(fd)
	for input.Scan() {
		i++
		line := input.Text()

		if line == "" {
			emptyCount++
			continue
		}

		cheatFile := fp.Base(file)
		var cheatDir string
		if fp.Dir(file) != args.rootDir {
			cheatDir = fp.Base(fp.Dir(file)) + "/"
		}

		lineType := getLineType(i, emptyCount, line, prevLine, item)
		switch lineType {
		case "section":
			item.section = section
			if itemMatch(item, args) {
				items = append(items, item)
			}
			item = itemT{}
			section = cheatDir + cheatFile + " :: " + line
		case "desc":
			item.desc += line + "\n"
		case "info":
			item.info += line + "\n"
		case "new_item":
			item.section = section
			if itemMatch(item, args) {
				items = append(items, item)
			}
			item = itemT{}
			item.desc = line + "\n"
		case "ignore":
			continue
		}
		emptyCount = 0
		prevLine = line
	}

	if itemMatch(item, args) {
		item.section = section
		items = append(items, item)
	}

	return items
}

// returns line types: "section", "desc", "info", "new_item" or "ignore"
func getLineType(i, emptyCount int, line, prevLine string, item itemT) string {
	switch {
	case emptyCount == 3 || i == 1:
		return "section"

	// ignore lines with underscores
	case str.Replace(line, "=", "", -1) == "":
		return "ignore"

	case emptyCount == 1:
		return "new_item"

	case str.HasPrefix(line, "# ") && item.info == "":
		return "desc"

	default:
		return "info"
	}

	return "empty"
}

func itemMatch(item itemT, args argsT) bool {
	section := item.section
	desc := item.desc
	info := item.info

	if desc == "" && info == "" {
		return false
	}

	if args.query[0] == "" {
		return true
	}

	if !args.caseSensitive {
		section = str.ToLower(section)
		desc = str.ToLower(desc)
		info = str.ToLower(info)
	}

	for _, q := range args.queryNeg {
		if !args.caseSensitive {
			q = str.ToLower(q)
		}

		sectionHasQ := str.Contains(section, q)
		descHasQ := str.Contains(desc, q)
		infoHasQ := str.Contains(info, q)

		if sectionHasQ || descHasQ || infoHasQ {
			return false
		}
	}

	for _, q := range args.query {
		if !args.caseSensitive {
			q = str.ToLower(q)
		}

		sectionHasQ := str.Contains(section, q)
		descHasQ := str.Contains(desc, q)
		infoHasQ := str.Contains(info, q)

		if !sectionHasQ && !descHasQ && !infoHasQ {
			return false
		}
	}

	return true
}

func printItems(items []itemT, args argsT) {
	var section string
	for _, item := range items {
		if section != item.section {
			section = item.section
			fmt.Printf("\033[1m%s\033[0m\n\n", section)
		}

		// highlight keywords from the query
		desc := highlightKeywords(item.desc, args)
		info := highlightKeywords(item.info, args)

		fmt.Printf("%s", desc)
		fmt.Printf("%s\n", info)
	}
}

func highlightKeywords(in string, args argsT) string {
	red := "\033[1;31m"
	reset := "\033[0m"

	inLower := in
	if !args.caseSensitive {
		inLower = str.ToLower(in)
	}

	var positions []posT

	for _, q := range args.query {
		if !args.caseSensitive {
			q = str.ToLower(q)
		}
		index := str.Index(inLower, q)
		if index == -1 {
			continue
		}
		var pos posT
		pos.start = index
		pos.end = index + len(q)
		positions = append(positions, pos)
	}

	sort.Slice(positions, func(i, j int) bool {
		return positions[i].start <= positions[j].start
	})

	for i, pos := range positions {
		// handle overlapping keywords by expanding next pos range
		if i < len(positions)-1 && pos.end > positions[i+1].start {
			(&positions[i+1]).start = pos.start
			continue
		}

		in = in[:pos.start] + red + in[pos.start:pos.end] + reset +
			in[pos.end:]
		movePos(positions, i, len(red)+len(reset))
	}

	return in
}

func movePos(positions []posT, index, offset int) {
	for i := index; i < len(positions); i++ {
		(&positions[i]).start += offset
		(&positions[i]).end += offset
	}
}

func parseArgs() argsT {
	var args argsT

	homeDir, err := os.UserHomeDir()
	errExit(err)
	args.rootDir = fp.Join(homeDir, ".local/share/cheatsheets")

	var skip bool
	for i := 1; i < len(os.Args); i++ {
		if skip {
			skip = false
			continue
		}

		arg := os.Args[i]
		switch {
		case arg == "-d":
			if i+1 < len(os.Args) && pathIsDir(os.Args[i+1]) {
				args.rootDir = os.Args[i+1]
				skip = true
			} else {
				msg := "-d requires a valid dir arg"
				errExit(fmt.Errorf(msg))
			}

		case arg == "-c":
			args.checkOnly = true
			args.getAllFiles = true

		case arg == "-s":
			args.caseSensitive = true

		case arg == ".":
			args.getAllFiles = true

		case str.HasPrefix(arg, "!"):
			args.queryNeg = append(args.queryNeg,
				str.TrimPrefix(arg, "!"))

		case str.HasPrefix(arg, ".."):
			args.filesFixed = append(args.filesFixed,
				str.TrimPrefix(arg, ".."))

		case str.HasPrefix(arg, "."):
			args.cheatDirs = append(args.cheatDirs,
				getCheatDirs(args.rootDir, os.Args[i])...)

		default:
			args.query = append(args.query, os.Args[i])
		}
	}

	if len(args.cheatDirs) == 0 {
		args.cheatDirs = append(args.cheatDirs, args.rootDir)
	}

	if args.getAllFiles {
		args.files = getAllFiles(args.rootDir)
	} else if len(args.filesFixed) != 0 {
		args.files = getFixedFiles(args)
	} else {
		for _, d := range args.cheatDirs {
			args.files = append(args.files, getFiles(d)...)
		}
	}

	if len(args.query) == 0 {
		args.query = append(args.query, "")
	}

	return args
}

func getAllFiles(rootDir string) []string {
	res := getFiles(rootDir)
	dirs := getDirs(rootDir)
	for _, d := range dirs {
		dir := fp.Join(rootDir, d)
		subFiles := getFiles(dir)
		res = append(res, subFiles...)
	}
	return res
}

func getFixedFiles(args argsT) []string {
	var res []string
	for _, d := range args.cheatDirs {
		files := getFiles(d)
		for _, ff := range args.filesFixed {
			for _, f := range files {
				if str.Contains(f, ff) {
					res = append(res, f)
				}
			}
		}
	}
	return res
}

func getDirs(path string) []string {
	var dirs []string
	files, err := os.ReadDir(path)
	errExit(err)
	for _, d := range files {
		dir := fp.Join(path, d.Name())
		if !pathIsDir(dir) || skipFile(d.Name()) {
			continue
		}
		dirs = append(dirs, d.Name())
	}
	return dirs
}

func getCheatDirs(rootDir, cheatDir string) []string {
	var cheatDirs []string
	files, err := os.ReadDir(rootDir)
	errExit(err)
	for _, f := range files {
		d := str.TrimPrefix(cheatDir, ".")
		name := f.Name()
		if f.IsDir() && str.Contains(name, d) && !skipFile(name) {
			filePath := fp.Join(rootDir, name)
			cheatDirs = append(cheatDirs, filePath)
		}
	}

	if len(cheatDirs) == 0 {
		errExit(fmt.Errorf("no dir found containing %s", cheatDir))
	}

	return cheatDirs
}

func getFiles(dir string) []string {
	var res []string
	files, err := os.ReadDir(dir)
	errExit(err)
	for _, f := range files {
		if !f.IsDir() && !skipFile(f.Name()) {
			res = append(res, fp.Join(dir, f.Name()))
		}
	}
	return res
}

func skipFile(file string) bool {
	if str.HasPrefix(file, ".") || str.HasPrefix(file, "_") {
		return true
	}
	return false
}

func pathIsDir(path string) bool {
	fh, err := os.Open(path)
	if err != nil {
		return false
	}
	defer fh.Close()
	fileInfo, err := fh.Stat()
	if fileInfo.IsDir() {
		return true
	}
	return false
}

func errExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printHelp() {
	help := `syntax: ch [-s] [-d <dir>] [-c] [.<dir>] [..<file>] [query]

parameters:
-s	case sensitive search
-d	dir containing cheatsheets (default: ~/.local/share/cheatsheets/)
-c	only check syntax of all cheatsheets
-h	this help

query:
.		search all dirs and files
.dir_str	search only in subdirectories containing string 'dir_str'
..file_str	search only in files containing string 'file_str'
keyword		find a string 'keyword' in cheatsheet items
!keyword	don't print items containing 'keyword'

examples:
ch remove user !wget
ch .go ..time string
ch -s ESS
`
	fmt.Printf(help)
}
