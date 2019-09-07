package config

import "fmt"

var (
	BuildTime = "<UNDEFINED>"
	GitCommit = "<UNDEFINED>"
	GitStatus = "<UNDEFINED>"
)

func ShowVersion() {
	fmt.Printf("BuildTime: %s\nGitCommit: %s\nGitStatus: %s\n", BuildTime, GitCommit, GitStatus)
}
