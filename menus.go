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
		{Name: "options",
			Text:     "OPTIONS",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "quit",
			Text:     "QUIT",
			TxtX:     40,
			TxtY:     25,
			BgColour: white},
	}

	mainMenuInput := lm.Input{
		Width:               180,
		Height:              36,
		Tx:                  24,
		Ty:                  24,
		Offy:                40,
		DefaultSelBgColour:  pink,
		DefaultSelTxtColour: white,
		Items:               mainMenuItems,
	}

	mainMenu, _ = lm.NewMenu(mainMenuInput)

	optionsMenuItems := []lm.Item{
		{Name: "screensize",
			Text:     "SCREEN SIZE",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
	}

	optionsMenuInput := lm.Input{
		Width:               180,
		Height:              36,
		Tx:                  24,
		Ty:                  24,
		Offy:                40,
		DefaultSelBgColour:  pink,
		DefaultSelTxtColour: white,
		Items:               optionsMenuItems,
	}

	optionsMenu, _ = lm.NewMenu(optionsMenuInput)

	screensizeMenuItems := []lm.Item{
		{Name: "400x300",
			Text:     "400 x 300",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "600x400",
			Text:     "600 x 400",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "800x600",
			Text:     "800 x 600",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
	}

	screensizeMenuInput := lm.Input{
		Width:               180,
		Height:              36,
		Tx:                  24,
		Ty:                  24,
		Offy:                40,
		DefaultSelBgColour:  pink,
		DefaultSelTxtColour: white,
		Items:               screensizeMenuItems,
	}

	screensizeMenu, _ = lm.NewMenu(screensizeMenuInput)

}
