package main

import (
	"fmt"
	"strings"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	wzlib_subprocess "github.com/infra-whizz/wzlib/subprocess"
)

type PinchMgr struct {
	packageManagerType         string // rpm, dpkg (others?)
	packageManagerExtraOptions []string
	root                       string // root of the system
	debug                      bool

	wzlib_logger.WzLogger
}

func NewPinchMgr(mgr string, root string) *PinchMgr {
	pm := new(PinchMgr)
	pm.packageManagerType = mgr
	pm.root = root
	pm.packageManagerExtraOptions = make([]string, 0)
	return pm
}

// SetAdditionalOptions sets additional options (type string) to the package manager call
func (pm *PinchMgr) SetAdditionalOptions(options []string) *PinchMgr {
	pm.packageManagerExtraOptions = append(pm.packageManagerExtraOptions, options...)
	return pm
}

// SetDebug mode
func (pm *PinchMgr) SetDebug(debug bool) *PinchMgr {
	pm.debug = debug
	if !pm.debug {
		pm.MuteLogger()
	}
	return pm
}

// Pinch a package from the system without dependencies
func (pm *PinchMgr) Pinch(name string) error {
	var err error
	switch pm.packageManagerType {
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
	var cmd *wzlib_subprocess.BufferedCmd
	var cname string
	var err error
	args := []string{}
	if pm.root != "" {
		cname = "chroot"
		args = []string{pm.root, "rpm"}
		args = append(args, pm.packageManagerExtraOptions...)
		args = append(args, []string{"-e", name}...)
	} else {
		cname = "rpm"
		args = append(args, pm.packageManagerExtraOptions...)
		args = append(args, []string{"-e", name}...)
	}

	cmd, err = wzlib_subprocess.BufferedExec(cname, args...)
	if err != nil {
		return err
	}

	stdout := cmd.StdoutString()
	stderr := cmd.StderrString()

	err = cmd.Wait()

	if pm.debug {
		for _, stdoutLine := range strings.Split(stdout, "\n") {
			stdoutLine = strings.TrimSpace(stdoutLine)
			if stdoutLine != "" {
				pm.GetLogger().Debug(stdoutLine)
			}
		}
	}

	if stderr != "" {
		for _, stderrLine := range strings.Split(stderr, "\n") {
			stderrLine = strings.TrimSpace(stderrLine)
			if stderrLine != "" {
				pm.GetLogger().Debug(stderrLine)
			}
		}
		if err != nil {
			err = fmt.Errorf("Error: %s. %s", err.Error(), strings.ReplaceAll(stderr, "\n", " "))
		}
	}

	return err
}
