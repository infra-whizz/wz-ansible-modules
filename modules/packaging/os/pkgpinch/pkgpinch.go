package main

import (
	"strings"

	wzmodlib "github.com/infra-whizz/wzmodlib"
)

/*
Documentation:
module: pkgpinch
author: "Bo Maryniuk (isbm)"
version_added: "2.9"
short_description: System pincher
description:
    - Direct package removal (dpkg, RPM) that supports chrooted environment
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
    options:
        description:
            - Additional package manager options
        required: false
        type: list

Examples:

# Remove bunch of packages
- pkgpinch:
    pkgs:
      - vim
      - suse_branding
      - dracut
    ignore_names: yes

- pkgpinch:
    pkgs:
      - vim
      - suse_branding
      - dracut
    ignore_names: yes
    options:
      - --noscripts
      - --nodeps
      - --dbpath
      - /var/lib/rpm

*/

// Pinch packages
func pinchPackages(args *PkgPinchArgs, response *wzmodlib.Response) {
	var out strings.Builder
	repo := NewPkgPinch().Configure(args)
	mgr := NewPinchMgr(repo.packageManager, repo.root).SetAdditionalOptions(repo.packageManagerOpts).SetDebug(repo.debug)
	changes := 0
	for _, pkgname := range repo.pkgNames {
		if err := mgr.Pinch(pkgname); err != nil {
			if !repo.ignoreErrors {
				// TODO: Check for wrong name error
				out.WriteString(err.Error() + ". ")
			} else {
				changes++
			}
		}
	}
	errMsg := strings.TrimSpace(out.String())
	if errMsg != "" {
		response.Msg = errMsg
		response.Failed = true
	}
	response.Changed = changes > 0
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
