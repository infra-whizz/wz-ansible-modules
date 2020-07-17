package main

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

func (zr *ZypperRepository) bool2int(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
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
	zr.enabled = zr.bool2int(enabled)
	return zr
}

// Autorefresh the repository
func (zr *ZypperRepository) Autorefresh(refresh bool) *ZypperRepository {
	zr.autorefresh = zr.bool2int(refresh)
	return zr
}

// CheckGpg key of the repository
func (zr *ZypperRepository) CheckGpg(check bool) *ZypperRepository {
	zr.gpgcheck = zr.bool2int(check)
	return zr
}

// KeepPackages in the cache
func (zr *ZypperRepository) KeepPackages(keep bool) *ZypperRepository {
	zr.keeppackages = zr.bool2int(keep)
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
