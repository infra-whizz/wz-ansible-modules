package main

import (
	"fmt"

	wzmodlib "github.com/infra-whizz/wzmodlib"
)

type PkgPinchArgs struct {
	Packages       []string
	IgnoreNames    string `json:"ignore_names"`  // Ignore naming errors
	IgnoreErrors   string `json:"ignore_errors"` // Ignore all errors
	PackageManager string `json:"manager"`       // Package manager of a choice
	Root           string // System root. Default "/"
}

// Default values are set if none found.
func (arg *PkgPinchArgs) Validate() error {
	if arg.Root == "" {
		arg.Root = "/"
	}

	if len(arg.Packages) < 1 {
		return fmt.Errorf("No packages has been defined")
	}

	if arg.IgnoreNames != "yes" {
		arg.IgnoreNames = "no"
	}

	if arg.IgnoreErrors == "yes" {
		arg.IgnoreErrors = "no"
	}

	// Setup package manager
	if arg.PackageManager == "" {
		arg.PackageManager = "rpm"
	} else if err := wzmodlib.CheckAnsibleParameter("manager", arg.PackageManager,
		[]string{"rpm", "dpkg"}); err != nil {
		return err
	}

	return nil
}
