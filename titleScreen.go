package main

import(
	"os"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil" // required for isKeyJustPressed
)

func updateTitleScreen(screen *ebiten.Image) error{
	activeMenu = mainMenu
		activeMenu.SetScale(bestRatio)

		activeMenu.Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			activeMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			activeMenu.IncrementSelected()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch activeMenu.GetSelectedItem() {
			case "play":
				initialisePlay()
				state = play
			case "options":
				state = options
			case "quit":
				os.Exit(0)
			}
			return nil
		}
	return nil
}