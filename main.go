package main

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
)

const (
	commandGit = "git"
	gitUnknown = "??"
	gitIgnored = "!!"
	gitDeleted = "D "
	gitAdded1  = "A "
	gitAdded2  = "AM"
	gitAdded3  = "AD"
)

var (
	version       = "dev"
	revision      = "dev"
	ignoredStatus = map[string]bool{
		gitUnknown: true,
		gitIgnored: true,
		gitDeleted: true,
		gitAdded1:  true,
		gitAdded2:  true,
		gitAdded3:  true,
	}
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
		log.Fatalf("found %d modifications", len(mods))
	}
}
