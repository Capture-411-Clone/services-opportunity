package main

import (
	"github.com/Capture-411/services-opportunity/cmd/chef/cmd"
	"github.com/Capture-411/services-opportunity/config"
)

func main() {
	config.Load()

	cmd.Execute()
}
