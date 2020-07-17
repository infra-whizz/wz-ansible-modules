package main

import (
	"fmt"
)

type ZypperRepositoryArgs struct {
	Enabled           string // Set repository to enabled (or disabled).
	Runrefresh        string // Refresh the pkglist
	AutoImportKeys    string `json:"auto_import_keys"`   // Automatically import gpg signing keys
	OverwriteMultiple string `json:"overwrite_multiple"` // Overwrite multiple repository entries, if repositories with both name and URL exist
	Priority          int    // Repo prio
	Autorefresh       string // Enable autorefresh
	Refresh           string // Enable autorefresh (stupid alias)
	DisableGPGCheck   string `json:"disable_gpg_check"` // Disable or not GPG sigcheck
	State             string // State of the repo
	Repo              string // URI of the repo
	Name              string // Name or an alias of the repo (required)
	KeepPackages      string `json:"keep_packages"` // Keep packages in the cache
	Root              string // System root. Default "/"
}

// Default values are set if none found.
func (arg *ZypperRepositoryArgs) Validate() error {
	if arg.Root == "" {
		arg.Root = "/"
	}

	if arg.KeepPackages == "" {
		arg.KeepPackages = "no"
	}

	if arg.State == "" {
		arg.State = "present"
	}

	if arg.Enabled == "" {
		arg.Enabled = "yes"
	}

	if arg.Runrefresh == "" {
		arg.Runrefresh = "no"
	}

	if arg.Autorefresh == "" {
		arg.AutoImportKeys = "no"
	}

	if arg.OverwriteMultiple == "" {
		arg.OverwriteMultiple = "no"
	}

	// Alias of the autorefresh overrides it
	if arg.Refresh != "" {
		arg.Autorefresh = arg.Refresh
	}

	if arg.Autorefresh == "" {
		arg.Autorefresh = "yes"
		arg.Refresh = arg.Autorefresh
	}

	if arg.DisableGPGCheck == "" {
		arg.DisableGPGCheck = "no"
	}

	if arg.Repo == "" && arg.State == "present" {
		return fmt.Errorf("Repository should be defined in 'repo' argument!")
	}

	if arg.Name == "" {
		return fmt.Errorf("Repository alias/name should be defined in 'name' argument!")
	}

	return nil
}
