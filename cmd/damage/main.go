package main

import (
	"encoding/json"
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
	infoLong = infoCmd.Flag("long", "Show all info").Bool()

	consumer *state.Consumer
	host     hdiutil.Host
)

func main() {
	log.SetFlags(0)

	app.UsageTemplate(kingpin.CompactUsageTemplate)

	app.HelpFlag.Short('h')
	app.Version("head")
	app.VersionFlag.Short('V')
	app.Author("Amos Wenger <amos@itch.io>")

	consumer = &state.Consumer{
		OnMessage: func(lvl string, msg string) {
			if lvl == "debug" && !*verbose {
				return
			}
			log.Printf("[%s] %s", lvl, msg)
		},
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

	host = hdiutil.NewHost(consumer)
	if *verbose {
		log.Printf("Running in verbose mode")
		host.SetDump(spew.Dump)
	}

	fullCmd := kingpin.MustParse(cmd, err)
	switch fullCmd {
	case infoCmd.FullCommand():
		info()
	}
}

func info() {
	file := *infoFile

	info, err := damage.GetDiskInfo(host, file)
	must(err)

	if *infoLong {
		jsonDump(info)
	} else {
		log.Printf("============================")
		log.Printf("%s", file)
		log.Printf("----------------------------")
		log.Printf("%s", info)
		log.Printf("============================")
	}
}

func jsonDump(v interface{}) {
	out, err := json.MarshalIndent(v, "", "  ")
	must(err)

	log.Print(string(out))
}

func must(err error) {
	if err != nil {
		panic(fmt.Sprintf("fatal error: %+v", err))
	}
}
