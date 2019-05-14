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
			TxtX:     60,
			TxtY:     25,
			BgColour: white},
		{Name: "options",
			Text:     "OPTIONS",
			TxtX:     40,
			TxtY:     25,
			BgColour: white},
		{Name: "quit",
			Text:     "QUIT",
			TxtX:     60,
			TxtY:     25,
			BgColour: white},
	}

	mainMenuInput := lm.Input{
		Width:               180,
		ItemHeight:          36,
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
		ItemHeight:          36,
		Offy:                40,
		DefaultSelBgColour:  pink,
		DefaultSelTxtColour: white,
		Items:               optionsMenuItems,
	}

	optionsMenu, _ = lm.NewMenu(optionsMenuInput)

	screensizeMenuItems := []lm.Item{
		{Name: "640x480",
			Text:     "640 x 480",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "800x600",
			Text:     "800 x 600",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "1024x768",
			Text:     "1024 x 768",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "1280x720",
			Text:     "1280 x 720",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "1336x768",
			Text:     "1336 x 768",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "1440x1080",
			Text:     "1440 x 1080",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "1600x900",
			Text:     "1600 x 900",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "1600x1200",
			Text:     "1600 x 1200",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "1920x1080",
			Text:     "1920 x 1080",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
		{Name: "1920x1200",
			Text:     "1920 x 1200",
			TxtX:     20,
			TxtY:     25,
			BgColour: white},
	}

	screensizeMenuInput := lm.Input{
		Width:               180,
		ItemHeight:          36,
		Offy:                40,
		DefaultSelBgColour:  pink,
		DefaultSelTxtColour: white,
		Items:               screensizeMenuItems,
	}

	screensizeMenu, _ = lm.NewMenu(screensizeMenuInput)

}
