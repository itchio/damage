package damage

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/itchio/damage/hdiutil"
	"github.com/pkg/errors"
)

type DiskInfo struct {
	Format            string     `plist:"Format"`
	FormatDescription string     `plist:"Format Description"`
	Partitions        Partitions `plist:"partitions"`
}

type Partitions struct {
	Partition []Partition `plist:"partitions"`
}

type Partition struct {
	Hint        string                 `plist:"partition-hint"`
	Name        string                 `plist:"partition-name"`
	Length      int64                  `plist:"partition-length"`
	Filesystems map[string]interface{} `plist:"partition-filesystems"`
}

func GetInfo(host hdiutil.Host, dmgpath string) error {
	var info DiskInfo
	err := host.RunAndDecode(&info, "imageinfo", "-plist", dmgpath)
	if err != nil {
		return errors.WithStack(err)
	}
	spew.Dump(info)

	return nil
}
