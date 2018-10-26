package damage

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/itchio/damage/hdiutil"
	"github.com/pkg/errors"
)

func Info(host hdiutil.Host, dmgpath string) error {
	l, err := host.AsPlist(
		"imageinfo",
		"-plist",
		dmgpath,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	spew.Dump(l)
	return nil
}
