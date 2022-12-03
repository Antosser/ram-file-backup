package main

import (
	"io/fs"
	"os"

	"github.com/sqweek/dialog"
)

func main() {
	if len(os.Args) >= 2 {

		data, _ := os.ReadFile(os.Args[1])

		if dialog.Message("%s", "Restore file now?").Title("RAM File Backup").YesNo() {
			os.WriteFile(os.Args[1], data, fs.ModePerm)
		}
	} else {
		dialog.Message("Thanks for using RAM File Backup!\nRight-Click a file and click \"Save to RAM\" to back it up.").Title("RAM File Backup").Info()
	}
}
