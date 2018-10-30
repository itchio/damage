package damage

import (
	"github.com/itchio/damage/hdiutil"
	"github.com/pkg/errors"
)

// Mount a dmg file into a directory
func Mount(host hdiutil.Host, dmgpath string, dir string) (hdiutil.Any, error) {
	res := make(hdiutil.Any)
	err := host.Command("attach").WithArgs(
		"-plist",             // output format
		"-nobrowse",          // don't show in Finder
		"-noverify",          // we already verify image checksums when downloading
		"-noautofsck",        // nuh-huh
		"-noautoopen",        // please don't
		"-mount", "required", // if we can't mount why bother?
		"-mountpoint", dir,
		"-readonly", // we won't ever write to it
		"-noidme",   // some kind of scripting, disable
		dmgpath,
	).WithInput("y").RunAndDecode(&res)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, err
}
