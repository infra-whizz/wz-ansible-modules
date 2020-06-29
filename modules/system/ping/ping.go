// Ping module

package main

import (
	"fmt"

	wzmodlib "github.com/infra-whizz/wzmodlib"
)

type ModuleArgs struct {
	Text string
}

func main() {
	var args ModuleArgs
	response := wzmodlib.CheckModuleCall(&args)

	if args.Text != "" {
		response.Msg = fmt.Sprintf("You say: %s", args.Text)
	} else {
		response.Msg = "Pong!"
	}

	wzmodlib.ExitWithJSON(response)
}
