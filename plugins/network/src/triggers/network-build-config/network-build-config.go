package main

import (
	"flag"

	network "github.com/dokku/dokku/plugins/network"
)

// rebuilds network settings for an app
func main() {
	flag.Parse()
	appName := flag.Arg(0)

	network.BuildConfig(appName)
}