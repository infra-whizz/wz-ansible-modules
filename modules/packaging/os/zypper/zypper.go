package main

import (
	"github.com/infra-whizz/wzmodlib"
)

func runZypper(args *ZypperArgs, response *wzmodlib.Response) {
	zypp := NewZypperOperations()
	zypp.Configure(args)
	zypp.Run()
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
	wzmodlib.ExitWithJSON(*response)
}
