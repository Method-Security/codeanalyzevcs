package main

import (
	"flag"
	"os"

	"github.com/Method-Security/codeanalyzevcs/cmd"
)

var version = "none"

func main() {
	flag.Parse()

	codeAnalyzeVCS := cmd.NewCodeAnalyzeVCS(version)
	codeAnalyzeVCS.InitRootCommand()
	codeAnalyzeVCS.InitGitlabCommand()

	if err := codeAnalyzeVCS.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
