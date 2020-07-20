package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/infra-whizz/wzmodlib"
)

func runZypper(args *ZypperArgs, response *wzmodlib.Response) {

}

func main() {
	var args ZypperArgs
	response := wzmodlib.CheckModuleCall(&args)

	if err := args.Validate(); err != nil {
		response.Failed = true
		response.Msg = err.Error()
	} else {
		runZypper(&args, response)
	}

	spew.Dump(args)
	wzmodlib.ExitWithJSON(*response)

}
