package main

import (
	"image/color"

	lm "github.com/Rosalita/ebiten-pkgs/listmenu"
	"github.com/hajimehoshi/ebiten"
)

type gameState int

const (
	titleScreen gameState = iota
	options
	screensize
	play
)

var (
	state          gameState
	activeMenu     lm.ListMenu
	mainMenu       lm.ListMenu
	optionsMenu    lm.ListMenu
	screensizeMenu lm.ListMenu
	myDeck         cardStack
	myDiscardPile  cardStack
	myHand         hand
	white          = &color.NRGBA{0xff, 0xff, 0xff, 0xff}
	pink           = &color.NRGBA{0xff, 0x69, 0xb4, 0xff}
	bestRatio      = 1.0
)

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	switch state {
	case titleScreen:
		err := updateTitleScreen(screen)
		if err != nil {
			return err
		}

	case options:
		err := updateOptions(screen)
		if err != nil {
			return err
		}

	case screensize:
		err := updateScreensize(screen)
		if err != nil {
			return err
		}

	case play:
		err := updatePlay(screen)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	initMenus()
	state = titleScreen

	if err := ebiten.Run(update, 1024, 768, 1, "Card Game"); err != nil {
		panic(err)
	}
}
