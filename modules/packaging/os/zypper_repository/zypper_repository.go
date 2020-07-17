package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	wzmodlib "github.com/infra-whizz/wzmodlib"
)

/*
Documentation:
module: zypper_repository
author: "Matthias Vogelgesang (@matze)"
version_added: "1.4"
short_description: Add and remove Zypper repositories
description:
    - Add or remove Zypper repositories on SUSE and openSUSE
options:
    name:
        description:
            - A name for the repository.
        type: string
        required
    repo:
        description:
            - URI of the repository or .repo file. Required when state=present.
    state:
        description:
            - A source string state.
        choices: [ "absent", "present" ]
        default: "present"
    description:
        description:
            - A description of the repository
    disable_gpg_check:
        description:
            - Whether to disable GPG signature checking of
              all packages. Has an effect only if state is
              I(present).
            - Needs zypper version >= 1.6.2.
        type: bool
        default: 'no'
    autorefresh:
        description:
            - Enable autorefresh of the repository.
        type: bool
        default: 'yes'
        aliases: [ "refresh" ]
    priority:
        description:
            - Set priority of repository. Packages will always be installed
              from the repository with the smallest priority number.
            - Needs zypper version >= 1.12.25.
    overwrite_multiple:
        description:
            - Overwrite multiple repository entries, if repositories with both name and
              URL already exist.
        type: bool
        default: 'no'
    auto_import_keys:
        description:
            - Automatically import the gpg signing key of the new or changed repository.
            - Has an effect only if state is I(present). Has no effect on existing (unchanged) repositories or in combination with I(absent).
            - Implies runrefresh.
            - Only works with C(.repo) files if `name` is given explicitly.
        type: bool
        default: 'no'
    runrefresh:
        description:
            - Refresh the package list of the given repository.
            - Can be used with repo=* to refresh all repositories.
        type: bool
        default: 'no'
    enabled:
        description:
            - Set repository to enabled (or disabled).
        type: bool
        default: 'yes'

    keep_packages:
        description:
            - Keep packages in the cache
        type: bool
        default: 'no'

Examples:

# Add Ansible repository
- zypper_repository:
    repo: 'https://download.opensuse.org/repositories/systemsmanagement/openSUSE_Leap_15.2/'

# Refresh all repos
- zypper_repository:
    repo: '*'
    runrefresh: yes

# Add a repo and add it's gpg key
- zypper_repository:
    repo: 'https://download.opensuse.org/repositories/systemsmanagement/openSUSE_Leap_15.2/'
    auto_import_keys: yes

# Force refresh of a repository
- zypper_repository:
    repo: 'https://download.opensuse.org/repositories/systemsmanagement/openSUSE_Leap_15.2/'
    name:
    state: present
    runrefresh: yes
'''

*/

// Remove unlink quietly, unless removal is impossible
func quietUnlink(fpath string) error {
	_, err := os.Stat(fpath)
	if !os.IsNotExist(err) {
		if err := os.Remove(fpath); err != nil {
			return err
		}
	}
	return nil
}

// Update repository by creating if does not exist, refresh content or just remove it
func updateRepo(args *ZypperRepositoryArgs, response *wzmodlib.Response) {
	repo := NewZypperRepository().Configure(args)
	fname, fbody := repo.Export()
	fpath := path.Join(args.Root, "etc", "zypp", "repos.d", fname)
	if args.State == "present" {
		if err := ioutil.WriteFile(fpath, []byte(fbody), 0644); err != nil {
			response.Msg = err.Error()
			response.Failed = true
		} else {
			response.Msg = fmt.Sprintf("Repository %s has been updated", fpath)
			response.Changed = true
		}
	} else {
		if err := quietUnlink(fpath); err == nil {
			response.Msg = fmt.Sprintf("Repository %s has been deleted", fpath)
			response.Changed = true
		} else {
			response.Msg = err.Error()
			response.Failed = true
		}
	}
}

func main() {
	var args ZypperRepositoryArgs
	response := wzmodlib.CheckModuleCall(&args)

	if err := args.Validate(); err != nil {
		response.Failed = true
		response.Msg = err.Error()
	} else {
		updateRepo(&args, response)
	}

	wzmodlib.ExitWithJSON(*response)

}
