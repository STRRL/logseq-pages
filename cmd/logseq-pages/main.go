package main

import "github.com/strrl/logseq-pages/pkg/command"

func main() {
	err := command.NewRootCommand().Execute()
	if err != nil {
		panic(err)
	}
}
