package hdiutil

import (
	"os/exec"
	"strings"

	plist "github.com/DHowett/go-plist"
	"github.com/itchio/wharf/state"
	"github.com/pkg/errors"
)

// Any represent any data from a plist file
type Any map[string]interface{}

type DumpFunc func(p ...interface{})

// Host allows communicating with hdiutil, and
// handles logging, parsing, etc.
type Host interface {
	SetDump(dump DumpFunc)
	RunAndDecode(dst interface{}, name string, args ...string) error
}

type host struct {
	consumer *state.Consumer
	dump     DumpFunc
}

// NewHost configures and returns a new hdiutil host
func NewHost(consumer *state.Consumer) Host {
	return &host{
		consumer: consumer,
	}
}

func (h *host) SetDump(dump DumpFunc) {
	h.dump = dump
}

func (h *host) RunAndDecode(dst interface{}, name string, args ...string) error {
	output, err := h.run(name, args...)
	if err != nil {
		return errors.WithStack(err)
	}

	if h.dump != nil {
		result := make(Any)
		_, err = plist.Unmarshal(output, &result)
		h.dump(result)
	}

	_, err = plist.Unmarshal(output, dst)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (h *host) run(subcmd string, args ...string) ([]byte, error) {
	h.consumer.Debugf("hdiutil ::: %s ::: %s", subcmd, strings.Join(args, " ::: "))

	hdiArgs := []string{subcmd}
	hdiArgs = append(hdiArgs, args...)
	cmd := exec.Command("hdiutil", hdiArgs...)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	result := make(Any)
	_, err = plist.Unmarshal(output, &result)
	if err != nil {
		return nil, errors.WithMessagef(err, "While running hdiutil.%s", subcmd)
	}

	return output, nil
}
