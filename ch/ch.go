package main

// add help with syntax: ch [-d <dir>] [-c] [.dir] [..section] key1 !key2 key
// add negation
// add section
// add . for all
// add search in item title
// add case sensitive flag and make queries insensitive

import (
	"bufio"
	"fmt"
	"os"
	fp "path/filepath"
	str "strings"
)

type itemT struct {
	section string
	desc    string
	info    string
}

type argsT struct {
	checkOnly   bool
	getAllFiles bool
	query       []string
	rootDir     string
	cheatDirs   []string
	files       []string
}

func main() {
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
		printItems(items, args.query)
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

		lineType := getLineType(i, emptyCount, line, prevLine)
		switch lineType {
		case "section":
			item.section = section
			if itemMatch(item, args.query) {
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
			if itemMatch(item, args.query) {
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

	if itemMatch(item, args.query) {
		items = append(items, item)
	}

	return items
}

// returns line types: "section", "desc", "info", "new_item" or "ignore"
func getLineType(i, emptyCount int, line, prevLine string) string {
	switch {
	case emptyCount == 3 || i == 1:
		return "section"

	// ignore lines with underscores
	case str.Replace(line, "=", "", -1) == "":
		return "ignore"

	case emptyCount == 1:
		return "new_item"

	case str.HasPrefix(line, "# ") && str.HasPrefix(prevLine, "# "):
		return "desc"

	default:
		return "info"
	}

	return "empty"
}

func itemMatch(item itemT, query []string) bool {
	if item.desc == "" && item.info == "" {
		return false
	}

	if query[0] == "" {
		return true
	}

	for _, q := range query {
		if !str.Contains(item.desc, q) && !str.Contains(item.info, q) {
			return false
		}
	}

	return true
}

func printItems(items []itemT, query []string) {
	var section string
	for _, item := range items {
		if section != item.section {
			section = item.section
			fmt.Printf("\033[1m%s\033[0m\n\n", section)
		}

		// highlight keywords from the query
		desc := highlightKeywords(item.desc, query)
		info := highlightKeywords(item.info, query)

		fmt.Printf("%s", desc)
		fmt.Printf("%s\n", info)
	}
}

func highlightKeywords(s string, query []string) string {
	for i, q := range query {
		var qJoin, joinRepl string
		if i+1 < len(query) {
			qJoin = q + query[i+1][1:]
			joinRepl = "\033[1;31m" + qJoin + "\033[0m"
		}

		if qJoin != "" && str.Contains(s, qJoin) {
			s = str.Replace(s, qJoin, joinRepl, -1)
		} else {
			s = str.Replace(s, q, "\033[1;31m"+q+"\033[0m", -1)
		}
	}

	return s
}

func parseArgs() argsT {
	var args argsT
	args.rootDir = "~/cheatsheets"

	var skip bool
	for i := 0; i < len(os.Args); i++ {
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

		case arg == ".":
			args.getAllFiles = true

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
	if str.HasPrefix(file, ".") {
		return true
	}

	skipNames := []string{
		"ch",
		"dotfiles",
		"config",
		"benchmark",
		"template",
		"install",
		"todo",
	}

	for _, name := range skipNames {
		if file == name {
			return true
		}
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
