package main

import (
	"encoding/xml"
	"fmt"

	"github.com/infra-whizz/wzmodlib"
)

type ZypperOperations struct {
	args *ZypperArgs
	zypp *Zypper
}

func NewZypperOperations() *ZypperOperations {
	zr := new(ZypperOperations)
	zr.zypp = NewZypper().XML(true)
	return zr
}

// An alias to YesNo2Bool
func (zr *ZypperOperations) toBool(val string) bool {
	return wzmodlib.YesNo2Bool(val)
}

// FilterInstalled cleans up already installed packages
func (zr *ZypperOperations) FilterInstalled(packages []string) ([]string, error) {
	newPackages := make([]string, 0)
	stout, sterr, err := zr.zypp.New().Search().Installed().Packages(packages...).Call()
	if err != nil {
		panic(err)
	}
	if sterr != "" {
		return newPackages, fmt.Errorf("Error: %s", sterr)
	}

	var search ZypperSearch
	if err := xml.Unmarshal([]byte(stout), &search); err != nil {
		panic(err)
	}

	for _, pkgname := range packages {
		installed := false
		for _, solvable := range search.Result.SolvableList.Packages {
			if pkgname == solvable.Name {
				installed = true
				break
			}
		}
		if !installed {
			newPackages = append(newPackages, pkgname)
		}
	}

	return newPackages, nil
}

func (zr *ZypperOperations) Install(packages ...string) {
	stout, sterr, err := zr.zypp.New().Install().Packages(packages...).Call()
	if err != nil {
		panic(err)
	}
	if sterr != "" {
		panic(fmt.Errorf("Error: %s", sterr))
	}
	fmt.Println(stout)
}

// Configure from arguments
func (zr *ZypperOperations) Configure(args *ZypperArgs) *ZypperOperations {
	zr.args = args
	zr.zypp.SetRoot(zr.args.Root)
	return zr
}

// Run zypper, configured by the state
func (zr *ZypperOperations) Run() error {
	switch zr.args.State {
	case "present":
		packages, err := zr.FilterInstalled(zr.args.Packages)
		if err != nil {
			panic(err)
		}
		zr.Install(packages...)
	case "latest":
	case "absent":
	case "dist-upgrade":
	default:
		panic("Unknown state: " + zr.args.State) // This should not happen
	}

	return nil
}
