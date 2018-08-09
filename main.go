package main

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
)

const (
	commandGit = "git"
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
		mods = append(mods, &mod)
	}
	if len(mods) > 0 {
		log.Fatalf("found %d modifications", len(mods))
	}
}
