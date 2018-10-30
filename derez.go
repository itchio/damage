package damage

import (
	"github.com/itchio/damage/hdiutil"
	"github.com/pkg/errors"
)

type UDIFResources struct {
	LPic []UDIFResource `plist:"LPic"`
	Str  []UDIFResource `plist:"STR#"`
	Text []UDIFResource `plist:"TEXT"`
	Tmpl []UDIFResource `plist:"TMPL"`
	Plst []UDIFResource `plist:"plst"`
	Styl []UDIFResource `plist:"styl"`
}

type UDIFResource struct {
	Data       []byte `plist:"Data"`
	ID         string `plist:"ID"`
	Name       string `plist:"Name"`
	Attributes string `plist:"Attributes"`
}

func GetUDIFResources(host hdiutil.Host, dmgpath string) (*UDIFResources, error) {
	var rez UDIFResources
	err := host.RunAndDecode(&rez, "udifderez", "-xml", dmgpath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &rez, nil
}
