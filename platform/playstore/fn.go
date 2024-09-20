package playstore

import (
	"log"

	"vploy.cli/core"
)

func Start() {
	action := core.Args.TakeFirst()

	switch action {
	case "upload":

	case "edit":

	default:
		log.Fatalln("invalid playstore action")
	}
}
