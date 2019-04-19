package main

import (
	//im "github.com/Rosalita/ebiten-pkgs/imagemenu"
	lm "github.com/Rosalita/ebiten-pkgs/listmenu"
	//"github.com/Rosalita/my-ebiten/resources/avatars"
	//"github.com/Rosalita/my-ebiten/resources/ui"
)

func initMenus() {



	mainMenuItems := []lm.Item{
		{Name: "play",
			Text:     "PLAY",
			TxtX:     40,
			TxtY:     25,
			BgColour: white},
		{Name: "quit",
			Text:     "QUIT",
			TxtX:     40,
			TxtY:     25,
			BgColour: white},
	}

	mainMenuInput := lm.Input{
		Width:               140,
		Height:              36,
		Tx:                  24,
		Ty:                  24,
		Offy:                40,
		DefaultSelBgColour:  pink,
		DefaultSelTxtColour: white,
		Items:               mainMenuItems,
	}

	mainMenu, _ = lm.NewMenu(mainMenuInput)

}
