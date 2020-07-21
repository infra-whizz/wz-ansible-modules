package main

import (
	"fmt"

	"github.com/infra-whizz/wzmodlib"
)

/*
	name:
        description:
        - Package name C(name) or package specifier or a list of either.
        - Can include a version like C(name=1.0), C(name>3.4) or C(name<=2.7). If a version is given, C(oldpackage) is implied and zypper is allowed to
          update the package within the version range given.
        - You can also pass a url or a local path to a rpm file.
        - When using state=latest, this can be '*', which updates all installed packages.
        required: true
        aliases: [ 'pkg' ]
    state:
        description:
          - C(present) will make sure the package is installed.
            C(latest)  will make sure the latest version of the package is installed.
            C(absent)  will make sure the specified package is not installed.
            C(dist-upgrade) will make sure the latest version of all installed packages from all enabled repositories is installed.
          - When using C(dist-upgrade), I(name) should be C('*').
        required: false
        choices: [ present, latest, absent, dist-upgrade ]
        default: "present"
    type:
        description:
          - The type of package to be operated on.
        required: false
        choices: [ package, patch, pattern, product, srcpackage, application ]
        default: "package"
        version_added: "2.0"
    extra_args_precommand:
       version_added: "2.6"
       required: false
       description:
         - Add additional global target options to C(zypper).
         - Options should be supplied in a single line as if given in the command line.
    disable_gpg_check:
        description:
          - Whether to disable to GPG signature checking of the package
            signature being installed. Has an effect only if state is
            I(present) or I(latest).
        required: false
        default: "no"
        type: bool
    disable_recommends:
        version_added: "1.8"
        description:
          - Corresponds to the C(--no-recommends) option for I(zypper). Default behavior (C(yes)) modifies zypper's default behavior; C(no) does
            install recommended packages.
        required: false
        default: "yes"
        type: bool
    force:
        version_added: "2.2"
        description:
          - Adds C(--force) option to I(zypper). Allows to downgrade packages and change vendor or architecture.
        required: false
        default: "no"
        type: bool
    update_cache:
        version_added: "2.2"
        description:
          - Run the equivalent of C(zypper refresh) before the operation. Disabled in check mode.
        required: false
        default: "no"
        type: bool
        aliases: [ "refresh" ]
    oldpackage:
        version_added: "2.2"
        description:
          - Adds C(--oldpackage) option to I(zypper). Allows to downgrade packages with less side-effects than force. This is implied as soon as a
            version is specified as part of the package name.
        required: false
        default: "no"
        type: bool
*/
type ZypperArgs struct {
	Name              string
	Pkg               string // Stupid alias to "Name"
	UpdateCache       string `json:"update_cache"`
	Refresh           string // Stupid alias to "update_cache"
	DisableGPGCheck   string `json:"disable_gpg_check"`
	PkgType           string `json:"type"` // package type
	State             string
	DisableRecommends string `json:"disable_recommends"`
	OldPackage        string
	Force             string
	Root              string   // Installation root
	Packages          []string // List of packages (instead of one name)
	ExtraArgs         []string `json:"extra_args"`            // List of string options to "zypper"
	ExtraPreArgs      []string `json:"extra_args_precommand"` // List of string options to "zypper" those before the main command
	PipeFile          string   `json:"pipe_file"`             // Path to a pipefile. If specified, zypper's output XML will be streamed there.
}

func (zarg *ZypperArgs) initArrays() {
	if zarg.Packages == nil {
		zarg.Packages = make([]string, 0)
	}
	if zarg.ExtraArgs == nil {
		zarg.ExtraArgs = make([]string, 0)
	}
	if zarg.ExtraPreArgs == nil {
		zarg.ExtraPreArgs = make([]string, 0)
	}
}

func (zarg *ZypperArgs) Validate() error {
	var err error

	zarg.initArrays()

	if zarg.Root == "" {
		zarg.Root = "/"
	}

	// Handle pkg stupid alias to name
	if zarg.Pkg != "" {
		zarg.Name = zarg.Pkg
	}

	if zarg.Name == "" {
		if len(zarg.Packages) == 0 {
			return fmt.Errorf("Package name required in 'name' or at least one in a 'packages' list!")
		}
	} else {
		if !wzmodlib.SInList(zarg.Name, zarg.Packages) {
			zarg.Packages = append(zarg.Packages, zarg.Name)
		}
	}
	zarg.Pkg, zarg.Name = "", "" // It should be always just "packages"

	// Setup state
	if zarg.State == "" {
		zarg.State = "present"
	} else if err := wzmodlib.CheckAnsibleParameter("state", zarg.State,
		[]string{"present", "latest", "absent", "dist-upgrade"}); err != nil {
		return err
	}

	// Setup package type
	if zarg.PkgType == "" {
		zarg.PkgType = "package"
	} else if err := wzmodlib.CheckAnsibleParameter("type", zarg.PkgType,
		[]string{"package", "patch", "pattern", "product", "srcpackage", "application"}); err != nil {
		return err
	}

	if zarg.DisableRecommends, err = wzmodlib.CheckAnsibleBool("disable_recomends", zarg.DisableRecommends, true); err != nil {
		return err
	}

	// Setup GPG checks
	if zarg.DisableGPGCheck, err = wzmodlib.CheckAnsibleBool("disable_gpg_check", zarg.DisableGPGCheck, false); err != nil {
		return err
	}

	if zarg.Force, err = wzmodlib.CheckAnsibleBool("force", zarg.Force, false); err != nil {
		return err
	}

	// Handle "refresh" stupid alias to "update_cache"
	if zarg.Refresh != "" {
		zarg.UpdateCache = zarg.Refresh
	}
	if zarg.UpdateCache, err = wzmodlib.CheckAnsibleBool("update_cache", zarg.UpdateCache, false); err != nil {
		return err
	}

	if zarg.OldPackage, err = wzmodlib.CheckAnsibleBool("oldpackage", zarg.OldPackage, false); err != nil {
		return err
	}

	return nil
}
