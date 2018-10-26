package damage

import (
	"log"

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

	log.Printf("plist: %#v", l)
	return nil
}
