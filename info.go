package damage

import (
	"encoding/json"
	"log"

	"github.com/itchio/damage/hdiutil"
	"github.com/pkg/errors"
)

type DiskInfo struct {
	Format            string              `plist:"Format"`
	FormatDescription string              `plist:"Format Description"`
	Partitions        Partitions          `plist:"partitions"`
	Properties        DiskProperties      `plist:"Properties"`
	SizeInformation   DiskSizeInformation `plist:"Size Information"`
}

type DiskSizeInformation struct {
	CompressedBytes int64 `plist:"Compressed Bytes"`
	TotalBytes      int64 `plist:"Total Bytes"`
}

type DiskProperties struct {
	// true if has EULA
	SoftwareLicenseAgreement bool `plist:"Software License Agreement" json:",omitempty"`
	// true if is compressed
	Checksummed bool `plist:"Checksummed" json:",omitempty"`
	// true if includes checksums
	Compressed bool `plist:"Compressed" json:",omitempty"`
	Encrypted  bool `plist:"Encrypted" json:",omitempty"`
}

type Partitions struct {
	Partition []Partition `plist:"partitions"`
}

type Partition struct {
	Hint        string                 `plist:"partition-hint" json:",omitempty"`
	Name        string                 `plist:"partition-name"`
	Length      int64                  `plist:"partition-length"`
	Synthesized bool                   `plist:"synthesized" json:",omitempty"`
	Filesystems map[string]interface{} `plist:"partition-filesystems" json:",omitempty"`
}

func GetInfo(host hdiutil.Host, dmgpath string) error {
	var info DiskInfo
	err := host.RunAndDecode(&info, "imageinfo", "-plist", dmgpath)
	if err != nil {
		return errors.WithStack(err)
	}
	// spew.Dump(info)

	out, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return errors.WithStack(err)
	}

	log.Printf("Info:\n%s", string(out))

	return nil
}
