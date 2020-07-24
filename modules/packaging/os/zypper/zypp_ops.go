package main

import (
	"encoding/xml"
	"fmt"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	"github.com/infra-whizz/wzmodlib"
	"github.com/sirupsen/logrus"
)

type ZypperOperations struct {
	args *ZypperArgs
	zypp *Zypper
	wzlib_logger.WzLogger
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

// FilterNew removes all the packages that are not installed.
func (zr *ZypperOperations) FilterNew(packages []string) ([]string, error) {
	installedPackages := make([]string, 0)
	stout, sterr, err := zr.zypp.New().Search().InstalledOnly().Packages(packages...).Call(zr.args.PipeFile)

	if err != nil {
		panic(err)
	}

	if sterr != "" {
		return installedPackages, fmt.Errorf("Error: %s", sterr)
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

		if installed {
			installedPackages = append(installedPackages, pkgname)
		}
	}

	return installedPackages, nil
}

// FilterInstalled cleans up already installed packages.
func (zr *ZypperOperations) FilterInstalled(packages []string) ([]string, error) {
	newPackages := make([]string, 0)
	stout, sterr, err := zr.zypp.New().Search().InstalledOnly().Packages(packages...).Call(zr.args.PipeFile)
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

func (zr *ZypperOperations) Install() error {
	packages, err := zr.FilterInstalled(zr.args.Packages)
	if err != nil {
		return err
	}
	if len(packages) == 0 {
		return fmt.Errorf("No packages has been selected for installation")
	}

	_, sterr, err := zr.zypp.New().Install().Packages(packages...).Call(zr.args.PipeFile)
	if err != nil {
		return err
	}
	if sterr != "" {
		return fmt.Errorf("Error: %s", sterr)
	}

	return nil
}

// Remove installedd packages
func (zr *ZypperOperations) Remove() error {
	packages, err := zr.FilterNew(zr.args.Packages)
	if err != nil {
		return err
	}

	if len(packages) == 0 {
		return fmt.Errorf("No packages has been selected for removal")
	}

	_, sterr, err := zr.zypp.New().Remove().Packages(packages...).Call(zr.args.PipeFile)
	if err != nil {
		return err
	}
	if sterr != "" {
		return fmt.Errorf("Error: %s", sterr)
	}

	return nil
}

// Configure from arguments
func (zr *ZypperOperations) Configure(args *ZypperArgs) *ZypperOperations {
	zr.args = args
	zr.zypp.SetRoot(zr.args.Root)

	if zr.args.Debug == "yes" {
		zr.GetLogger().SetLevel(logrus.TraceLevel)
	} else {
		zr.GetLogger().SetLevel(logrus.PanicLevel)
	}

	return zr
}

// Run zypper, configured by the state
func (zr *ZypperOperations) Run() (bool, error) {
	switch zr.args.State {
	case "present":
		if err := zr.Install(); err != nil {
			return false, err
		}
	case "latest":
		return false, fmt.Errorf("State %s not yet implemented", zr.args.State)
	case "absent":
		if err := zr.Remove(); err != nil {
			return false, err
		}
	case "dist-upgrade":
		return false, fmt.Errorf("State %s not yet implemented", zr.args.State)
	default:
		return false, fmt.Errorf("Unknown state: %s", zr.args.State) // This should not happen though
	}
	return true, nil
}
