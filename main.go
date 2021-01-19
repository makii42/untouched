package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	commandGit = "git"
	rcError    = 23
	gitUnknown = "??"
	gitIgnored = "!!"
	gitDeleted = "D "
	gitAdded1  = "A "
	gitAdded2  = "AM"
	gitAdded3  = "AD"
)

var (
	version  = "dev"
	revision = "dev"

	ignoredStatus = map[string]bool{
		gitUnknown: true,
		gitIgnored: true,
		gitDeleted: true,
		gitAdded1:  true,
		gitAdded2:  true,
		gitAdded3:  true,
	}

	doDiff bool
)

type (
	modification struct {
		op, file string
	}
)

func main() {
	location, err := exec.LookPath(commandGit)
	if err != nil {
		log.Fatalf("could not find %s binary on path", commandGit)
	}
	cmd := exec.Command(location, "status", "--porcelain")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Could not obtain output: %s", err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(out))
	mods := []*modification{}
	for scanner.Scan() {
		line := scanner.Text()
		mod := modification{
			op:   line[:2],
			file: line[3:],
		}
		if _, ignore := ignoredStatus[mod.op]; !ignore {
			mods = append(mods, &mod)
		}
	}
	if len(mods) > 0 {
		fmt.Fprintf(os.Stderr, "found %d modifications\n", len(mods))
		for _, mod := range mods {
			fmt.Fprintf(os.Stderr, "%s %s\n", mod.op, mod.file)
		}
		if doDiff {
			cmd = exec.Command(location, "diff")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				log.Fatalf("Could not run '%s diff': %s", location, err)
			}
		}
		os.Exit(rcError)
	}
}

func init() {
	flag.BoolVar(&doDiff, "diff", false, "Execute 'git diff' if touched files are found.")
	flag.Parse()
}
