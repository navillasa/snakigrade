package main

import (
	"log"

	"github.com/navillasa/snakigrade/snakigrade"
)

var (
	game snakigrade.Game
)

func update(screen *ebiten.Image) error {
	if err := game.Update(); err != nil {
		return err
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	game.Draw(screen)
	return nil
}

func main() {
	var err error
	game, err = twenty48.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, twenty48.ScreenWidth, twenty48.ScreenHeight, 1, "2048 (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
