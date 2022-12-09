package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type dirEntry struct {
	name       string
	size       int
	dir        *dirEntry
	isDir      bool
	dirEntries []*dirEntry
}

func (e *dirEntry) chdir(d string) (*dirEntry, error) {
	if d == ".." {
		return e.dir, nil
	}
	for _, f := range e.dirEntries {
		if f.name == d && f.isDir {
			return f, nil
		}
	}
	// if the directory was not found, create a new one
	// we don't really know what actually exists, just
	// the shell history
	f := newDir(d)
	e.dirEntries = append(e.dirEntries, f)
	return f, nil
}

func (e *dirEntry) addEntry(f *dirEntry) error {
	if !e.isDir {
		return errors.New(fmt.Sprintf("cannot add entry; %v is not a directory", e.name))
	}
	e.dirEntries = append(e.dirEntries, f)
	return nil
}

func newDir(name string) *dirEntry {
	e := dirEntry{}
	e.name = name
	e.isDir = true
	e.dirEntries = []*dirEntry{}
	return &e
}

func newFile(name string, size int) *dirEntry {
	e := dirEntry{}
	e.name = name
	e.size = size
	e.isDir = false
	return &e
}

type state struct {
	cwd *dirEntry
}

func main() {
	root := newDir("/")

	i, _ := os.ReadFile("input")
	lines := bytes.Split(i, []byte("\n"))

	parseFS(lines, root)

	fmt.Println("part 1:", solve(root))

	// at least 30000000 is required to be made available
	spaceToFree := 30000000 - (70000000 - du(root))
	solve2(root, spaceToFree)
}

// du recurses through a directory tree, summarizing its size
func du(d *dirEntry) int {
	sum := 0
	for _, i := range d.dirEntries {
		if i.isDir {
			sum += du(i)
			continue
		}
		sum += i.size
	}
	return sum
}

type usage struct {
	name string
	size int
}

func dirDu(d *dirEntry) []usage {
	sums := []usage{}
	u := usage{
		name: d.name,
		size: du(d),
	}
	sums = append(sums, u)
	for _, i := range d.dirEntries {
		if !i.isDir {
			continue
		}
		u := dirDu(i)
		sums = append(sums, u...)
	}
	return sums
}

func solve(d *dirEntry) int {
	sums := dirDu(d)
	//fmt.Println(sums)

	solution := 0
	for _, s := range sums {
		if s.size <= 100000 {
			solution += s.size
		}
	}

	return solution
}

func solve2(d *dirEntry, c int) {
	sums := dirDu(d)
	filtered := []usage{}
	for _, s := range sums {
		if s.size >= c {
			filtered = append(filtered, s)
			fmt.Println(s.size)
		}
	}
	fmt.Println(filtered)
}

func parseFS(lines [][]byte, cwd *dirEntry) {
	numericRegexp := regexp.MustCompile(`^\d+$`)
	for i, l := range lines {
		_ = i
		//fmt.Printf("line number: %d; content: %q; cwd: %q\n", i, l, cwd.name)

		var err error
		splitln := strings.Split(string(l), " ")

		switch firstWord := splitln[0]; {

		// lines starting with $ are commands that were run
		case firstWord == "$":
			switch splitln[1] {
			case "cd":
				dir := splitln[2]
				cwd, err = cwd.chdir(dir)
				if err != nil {
					fmt.Println("failed to change directory to", dir, "with error", err)
				}
			case "ls":
				// don't do anything
			}

		// lines starting with dir are new directories
		case firstWord == "dir":
			name := splitln[1]
			d := newDir(name)
			d.dir = cwd
			cwd.addEntry(d)
		case numericRegexp.MatchString(firstWord):
			size, _ := strconv.Atoi(firstWord)
			name := splitln[1]
			f := newFile(name, size)
			f.dir = cwd
			cwd.addEntry(f)
		default:
			fmt.Println("not sure how we got here, line was", l)
		}
	}
}
