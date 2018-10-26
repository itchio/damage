package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/itchio/damage"
	"github.com/itchio/damage/hdiutil"
	"github.com/itchio/wharf/state"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app     = kingpin.New("damage", "Your devilish little DMG helper")
	verbose = app.Flag("verbose", "Enable verbose output").Short('v').Bool()

	infoCmd  = app.Command("info", "Print information about a DMG file")
	infoFile = infoCmd.Arg("file", "The .dmg file to analyze").ExistingFile()

	consumer *state.Consumer
	host     hdiutil.Host
)

func main() {
	app.UsageTemplate(kingpin.CompactUsageTemplate)

	app.HelpFlag.Short('h')
	app.Version("head")
	app.VersionFlag.Short('V')
	app.Author("Amos Wenger <amos@itch.io>")

	consumer = &state.Consumer{
		OnMessage: func(lvl string, msg string) {
			log.Printf("[%s] %s", lvl, msg)
		},
	}
	host = hdiutil.NewHost(consumer)
	if *verbose {
		host.SetDump(spew.Dump)
	}

	args := os.Args[1:]
	cmd, err := app.Parse(args)
	if err != nil {
		ctx, _ := app.ParseContext(os.Args[1:])
		if ctx != nil {
			app.FatalUsageContext(ctx, "%s\n", err.Error())
		} else {
			app.FatalUsage("%s\n", err.Error())
		}
	}

	fullCmd := kingpin.MustParse(cmd, err)
	switch fullCmd {
	case infoCmd.FullCommand():
		info()
	}
}

func info() {
	file := *infoFile

	log.Printf("Analyzing (%s)", file)
	must(damage.GetInfo(host, file))
}

func must(err error) {
	if err != nil {
		panic(fmt.Sprintf("fatal error: %+v", err))
	}
}
