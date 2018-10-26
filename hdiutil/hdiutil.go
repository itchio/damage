package hdiutil

import (
	"os/exec"

	plist "github.com/DHowett/go-plist"
	"github.com/pkg/errors"
)

type Any map[string]interface{}

func GetPlist(name string, args ...string) (Any, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	result := make(Any)
	_, err = plist.Unmarshal(output, &result)
	if err != nil {
		return errors.WithStack(err)
	}

	return result, nil
}
