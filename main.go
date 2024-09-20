package main

import (
	"log"

	"vploy.cli/core"
	"vploy.cli/platform/appstore"
	"vploy.cli/platform/playstore"
)

func main() {
	if core.Args.Len() < 2 {
		log.Fatalln("invalid arguments")
	}

	platform := core.Args.TakeFirst()

	switch platform {
	case "playstore":
		playstore.Start()

	case "appstore":
		appstore.Start()

	default:
		log.Fatalln("unknown platform")
	}
}
