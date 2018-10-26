package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("damage", "Your devilish little DMG helper")

	infoCmd = app.Command("info", "Print information about a DMG file")
)

func main() {
	app.UsageTemplate(kingpin.CompactUsageTemplate)

	app.HelpFlag.Short('h')
	app.Version("head")
	app.VersionFlag.Short('V')
	app.Author("Amos Wenger <amos@itch.io>")

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
	log.Printf("Should print info!")
}

func must(err error) {
	if err != nil {
		panic(fmt.Sprintf("fatal error: %+v", err))
	}
}
