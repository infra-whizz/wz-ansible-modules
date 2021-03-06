package main

import (
	wzmodlib "github.com/infra-whizz/wzmodlib"
	"golang.org/x/sys/unix"
)

type ModuleArgs struct {
	Domainname bool
	Machine    bool
	Nodename   bool
	Release    bool
	Sysname    bool
	Version    bool
	All        bool
}

func main() {
	var args ModuleArgs
	response := wzmodlib.CheckModuleCall(&args)

	var uts unix.Utsname
	if err := unix.Uname(&uts); err != nil {
		response.Failed = true
		response.Msg = err.Error()
		wzmodlib.ExitWithFailedJSON(*response)
	}

	if args.Domainname || args.All {
		domainname := wzmodlib.Byte65toS(uts.Domainname)
		if domainname != "" {
			response.Return["domainname"] = domainname
		}
	}

	if args.Machine || args.All {
		machine := wzmodlib.Byte65toS(uts.Machine)
		if machine != "" {
			response.Return["machine"] = machine
		}
	}

	if args.Nodename || args.All {
		nodename := wzmodlib.Byte65toS(uts.Nodename)
		if nodename != "" {
			response.Return["nodename"] = nodename
		}
	}

	if args.Release || args.All {
		release := wzmodlib.Byte65toS(uts.Release)
		if release != "" {
			response.Return["release"] = release
		}
	}

	if args.Sysname || args.All {
		sysname := wzmodlib.Byte65toS(uts.Sysname)
		if sysname != "" {
			response.Return["sysname"] = sysname
		}
	}

	if args.Version || args.All {
		version := wzmodlib.Byte65toS(uts.Version)
		if version != "" {
			response.Return["version"] = version
		}
	}

	wzmodlib.ExitWithJSON(*response)
}
