package main

import (
	"github.com/infra-whizz/wzmodlib"
)

func runZypper(args *ZypperArgs, response *wzmodlib.Response) {
	zypp := NewZypperOperations()
	zypp.Configure(args)
	var err error
	response.Changed, err = zypp.Run()
	if err != nil {
		response.Failed = true
		response.Msg = err.Error()
	}
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
