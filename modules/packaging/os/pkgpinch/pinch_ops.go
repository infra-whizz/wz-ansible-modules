package main

type PkgPinch struct {
	pkgNames            []string
	packageManager      string
	root                string
	ignoreErrors        bool
	ignoreWrongPkgNames bool
}

func NewPkgPinch() *PkgPinch {
	pp := new(PkgPinch)
	pp.pkgNames = make([]string, 0)
	pp.ignoreWrongPkgNames, pp.ignoreErrors = false, false // Default anyway, but explicitly mentioned

	return pp
}

// Configure from arguments
func (pp *PkgPinch) Configure(args *PkgPinchArgs) *PkgPinch {
	pp.SetPkgNames(args.Packages)
	pp.SetIgnoreErrors(args.IgnoreErrors)
	pp.SetIgnoreWrongPkgNames(args.IgnoreNames)
	pp.SetPackageManager(args.PackageManager)
	pp.SetRoot(args.Root)

	return pp
}

// SetPackageManager to work with (rpm, dpkg etc)
func (pp *PkgPinch) SetPackageManager(manager string) *PkgPinch {
	pp.packageManager = manager
	return pp
}

// SetPkgNames of the packages to be pinched from the system
func (pp *PkgPinch) SetPkgNames(names []string) *PkgPinch {
	pp.pkgNames = append(pp.pkgNames, names...)
	return pp
}

// SetIgnoreErrors
func (pp *PkgPinch) SetIgnoreErrors(ignoreErrors string) *PkgPinch {
	pp.ignoreErrors = ignoreErrors == "yes"
	return pp
}

// SetIgnoreWrongPkgNames
func (pp *PkgPinch) SetIgnoreWrongPkgNames(ignoreNames string) *PkgPinch {
	pp.ignoreWrongPkgNames = ignoreNames == "yes"
	return pp
}

// SetRoot of the packages
func (pp *PkgPinch) SetRoot(root string) *PkgPinch {
	pp.root = root
	return pp
}
