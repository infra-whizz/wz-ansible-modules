package main

import (
	"fmt"

	wzlib_subprocess "github.com/infra-whizz/wzlib/subprocess"
)

type Zypper struct {
	root         string
	xmlInterface bool
	opts         []string // internal options
	packages     []string
	exec         string
}

// NewZypper creates a whole Zypper object class
func NewZypper() *Zypper {
	zypp := new(Zypper)
	zypp.exec = "zypper"
	zypp.packages = make([]string, 0)
	zypp.opts = make([]string, 0)
	return zypp
}

// New instance of the volatile Zypper call
func (zypp *Zypper) New() *Zypper {
	instance := NewZypper()
	instance.root = zypp.root
	instance.xmlInterface = zypp.xmlInterface

	instance.opts = []string{"--root", zypp.root, "--non-interactive", "--gpg-auto-import-keys"}
	if zypp.xmlInterface {
		instance.opts = append(instance.opts, "-x")
	}
	return instance
}

func (zypp *Zypper) addOpts(opts ...string) *Zypper {
	zypp.opts = append(zypp.opts, opts...)
	return zypp
}

// XML sets zypper xml output interface streaming to ON or OFF
func (zypp *Zypper) XML(state bool) *Zypper {
	zypp.xmlInterface = state
	return zypp
}

// SetRoot of zypper
func (zypp *Zypper) SetRoot(root string) *Zypper {
	zypp.root = root
	return zypp
}

// Packages set
func (zypp *Zypper) Packages(names ...string) *Zypper {
	return zypp.addOpts(names...)
}

// Search mode via Zypper. This turns just search mode, the filtering is not included.
func (zypp *Zypper) Search() *Zypper {
	return zypp.addOpts("search")
}

// Install packages using Zypper
func (zypp *Zypper) Install() *Zypper {
	return zypp.addOpts("in")
}

// InstalledOnly is to notify Zypper to set --installed-only flag
func (zypp *Zypper) InstalledOnly() *Zypper {
	return zypp.addOpts("--installed-only")
}

// Call Zypper command
func (zypp *Zypper) Call(pipe string) (stout string, sterr string, err error) {
	if pipe != "" {
		p := NewTextProcessStream(pipe)
		stout, sterr = wzlib_subprocess.StreamedExec(p, "zypper", zypp.opts...)
		return stout, sterr, nil
	}

	cmd, err := wzlib_subprocess.BufferedExec("zypper", zypp.opts...)
	if err != nil {
		return "", "", err
	}
	stout = cmd.StdoutString()
	sterr = cmd.StderrString()

	fmt.Println(stout)
	cmd.Wait()

	return stout, sterr, nil
}
