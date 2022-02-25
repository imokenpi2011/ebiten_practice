package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// ゲームのインターフェース
type Game struct{}

// ゲームの状態を60fpsで更新する
func (g *Game) Update() error {
	// ここにゲームの状態を書き込む
	return nil
}

// ゲームの状態を描画する
func (g *Game) Draw(screen *ebiten.Image) {
	// ここにゲームのレンダリングを描きます

}

const (
	screenX = 640
	screenY = 480
)

// 画面サイズを返す
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenX, screenY
}

func main() {
	ebiten.SetWindowSize(screenX, screenY)
	ebiten.SetWindowTitle("Dinosaur Jump")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
