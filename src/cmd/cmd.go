package cmd

import (
	"github.com/VivaTS-Drop/CHAOS/src/ui"
	"github.com/VivaTS-Drop/CHAOS/src/util"
)

func Start() {
	util.DetectOS()
	util.ClearScreen()
	ui.StartMenu()
}
