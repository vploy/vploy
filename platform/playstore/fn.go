package playstore

import (
	"log"

	"vploy.cli/core"
)

func Start() {
	var engine xEngine

	action := core.Args.TakeFirst()

	switch action {
	case "upload":
		engine.Upload.Run()

	case "edit":
		engine.Edit.Run()

	default:
		log.Fatalln("invalid playstore action")
	}
}
