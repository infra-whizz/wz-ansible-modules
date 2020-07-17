package main

import (
	"fmt"
	"strings"

	"github.com/infra-whizz/wzmodlib"
)

type ZypperRepository struct {
	name         string
	baseurl      string
	repotype     string // Obsolete. Serialised to "type", always NONE for autodetect.
	enabled      int
	autorefresh  int
	priority     int
	gpgcheck     int
	keeppackages int
}

func NewZypperRepository() *ZypperRepository {
	zr := new(ZypperRepository)
	zr.repotype = "NONE"
	zr.enabled = 1
	zr.autorefresh = 0
	return zr
}

// Configure from arguments
func (zr *ZypperRepository) Configure(args *ZypperRepositoryArgs) *ZypperRepository {
	zr.SetName(args.Name)
	zr.SetBaseUrl(args.Repo)
	zr.Enable(wzmodlib.YesNo2Bool(args.Enabled))
	zr.Autorefresh(wzmodlib.YesNo2Bool(args.Autorefresh))
	zr.SetPriority(args.Priority)
	zr.CheckGpg(!wzmodlib.YesNo2Bool(args.DisableGPGCheck)) // Note: inverted parameter
	zr.KeepPackages(wzmodlib.YesNo2Bool(args.KeepPackages))
	return zr
}

// SetName of the repository
func (zr *ZypperRepository) SetName(name string) *ZypperRepository {
	zr.name = name
	return zr
}

// SetBaseUrl of the repository
func (zr *ZypperRepository) SetBaseUrl(url string) *ZypperRepository {
	zr.baseurl = url
	return zr
}

// Enable repository
func (zr *ZypperRepository) Enable(enabled bool) *ZypperRepository {
	zr.enabled = wzmodlib.Bool2Int(enabled)
	return zr
}

// Autorefresh the repository
func (zr *ZypperRepository) Autorefresh(refresh bool) *ZypperRepository {
	zr.autorefresh = wzmodlib.Bool2Int(refresh)
	return zr
}

// CheckGpg key of the repository
func (zr *ZypperRepository) CheckGpg(check bool) *ZypperRepository {
	zr.gpgcheck = wzmodlib.Bool2Int(check)
	return zr
}

// KeepPackages in the cache
func (zr *ZypperRepository) KeepPackages(keep bool) *ZypperRepository {
	zr.keeppackages = wzmodlib.Bool2Int(keep)
	return zr
}

// SetPriority to process the repository
func (zr *ZypperRepository) SetPriority(prio int) *ZypperRepository {
	if prio < 0 {
		prio = 0
	}
	zr.priority = prio
	return zr
}

// Export to config string, suitable to write it to the file
func (zr *ZypperRepository) Export() (string, string) {
	var out strings.Builder

	out.WriteString(fmt.Sprintf("[%s]\n", zr.name))
	out.WriteString(fmt.Sprintf("enabled=%d\n", zr.enabled))
	out.WriteString(fmt.Sprintf("autorefresh=%d\n", zr.autorefresh))
	out.WriteString(fmt.Sprintf("baseurl=%s\n", zr.baseurl))
	out.WriteString("type=NONE\n")
	out.WriteString(fmt.Sprintf("priority=%d\n", zr.priority))
	out.WriteString(fmt.Sprintf("gpgcheck=%d\n", zr.gpgcheck))
	out.WriteString(fmt.Sprintf("keeppackages=%d\n", zr.keeppackages))

	return fmt.Sprintf("%s.repo", zr.name), out.String()
}
