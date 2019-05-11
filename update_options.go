package main

import(
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil" // required for isKeyJustPressed
)

func updateOptions(screen *ebiten.Image) error{
	activeMenu = optionsMenu
		activeMenu.SetScale(bestRatio)

		activeMenu.Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			activeMenu.DecrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			activeMenu.IncrementSelected()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			state = titleScreen
			return nil
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			switch activeMenu.GetSelectedItem() {
			case "screensize":
				state = screensize
				return nil
			}
			return nil
		}
	return nil
}