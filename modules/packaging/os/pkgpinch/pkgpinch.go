package main

import (
	"fmt"

	wzmodlib "github.com/infra-whizz/wzmodlib"
)

/*
Documentation:
module: pkgpinch
author: "Bo Maryniuk (isbm)"
version_added: "2.9"
short_description: System pincher
description:
    - Direct package removal (dpkg, RPM)
options:
    packages:
        description:
            - List of package names
        type: list
        required: true
    ignore_names:
        description:
            - Ignore naming errors (e.g. specific package is not installed)
        type: bool
        required: false
        default: 'no'
    ignore_errors:
        description:
            - Ignore all errors. This will also set "ignore_names" to True
        type: bool
        required: false
        default: 'no'
    manager:
        description:
            - Package manager of the choice
        required: false
        choices: ['rpm', 'dpkg']
        default: "rpm"

Examples:

# Remove bunch of packages
- pkgpinch:
    pkgs:
      - vim
      - suse_branding
      - dracut
    ignore_names: yes
'''

*/

// Pinch packages
func pinchPackages(args *PkgPinchArgs, response *wzmodlib.Response) {
	repo := NewPkgPinch().Configure(args)
	response.Msg = fmt.Sprintf("Errors: %v, Names: %v, Packages: %v, Manager: %v, Root: %v",
		repo.ignoreErrors, repo.ignoreWrongPkgNames, repo.pkgNames, repo.packageManager, repo.root)
}

func main() {
	var args PkgPinchArgs
	response := wzmodlib.CheckModuleCall(&args)

	if err := args.Validate(); err != nil {
		response.Failed = true
		response.Msg = err.Error()
	} else {
		pinchPackages(&args, response)
	}

	wzmodlib.ExitWithJSON(*response)
}
