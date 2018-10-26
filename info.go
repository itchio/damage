package damage

func Info(dmgpath string) error {
	l, err := hdiutil.GetPlist(
		"imageinfo",
		"-plist",
		dmgpath
	)
}