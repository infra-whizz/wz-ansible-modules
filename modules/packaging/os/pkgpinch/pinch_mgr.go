package main

import (
	"fmt"
	"strings"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	wzlib_subprocess "github.com/infra-whizz/wzlib/subprocess"
)

type PinchMgr struct {
	packageManagerType string // rpm, dpkg (others?)
	root               string // root of the system
	debug              bool

	wzlib_logger.WzLogger
}

func NewPinchMgr(mgr string, root string) *PinchMgr {
	pm := new(PinchMgr)
	pm.packageManagerType = mgr
	pm.root = root
	return pm
}

// SetDebug mode
func (pm *PinchMgr) SetDebug(debug bool) *PinchMgr {
	pm.debug = debug
	return pm
}

// Pinch a package from the system without dependencies
func (pm *PinchMgr) Pinch(name string) error {
	var err error
	switch name {
	case "rpm":
		err = pm.pinchRpm(name)
	case "dpkg":
		err = pm.pinchDpkg(name)
	default:
		err = fmt.Errorf("Unknown package manager: %s", name)
	}
	return err
}

// Use RPM to uninstall a package
func (pm *PinchMgr) pinchRpm(name string) error {
	//wzlib_subprocess.BufferedExec("rpm", "--root", name)
	return nil
}

// Use Dpkg to uninstall a package
func (pm *PinchMgr) pinchDpkg(name string) error {
	cmd, err := wzlib_subprocess.BufferedExec("dpkg", "--root", pm.root, "--remove", name)
	if err != nil {
		return err
	}

	stdout := cmd.StdoutString()
	stderr := cmd.StderrString()

	err = cmd.Wait()
	if err != nil {
		return err
	}

	if pm.debug {
		for _, stdoutLine := range strings.Split(stdout, "\n") {
			pm.GetLogger().Debug(stdoutLine)
		}
	}

	if stderr != "" {
		for _, stderrLine := range strings.Split(stderr, "\n") {
			pm.GetLogger().Error(stderrLine)
		}
	}

	return err
}
