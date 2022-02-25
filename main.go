package main

import (
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

const (
	screenX = 640
	screenY = 480
)

//go:embed resources/images/dinosaur_01.png
var byteDinosaur1Img []byte

//go:embed resources/images/dinosaur_02.png
var byteDinosaur2Img []byte

//go:embed resources/images/tree_small.png
var byteTreeSmallImg []byte

//go:embed resources/images/tree_big.png
var byteTreeBigImg []byte

//go:embed resources/images/ground.png
var byteGroundImg []byte

var (
	dinosaur1Img *ebiten.Image
	dinosaur2Img *ebiten.Image
	treeSmallImg *ebiten.Image
	treeBigImg   *ebiten.Image
	groundImg    *ebiten.Image
	arcadeFont   font.Face
)

// ゲームのインターフェース
type Game struct{}

// ゲームの初期化を行う
func NewGame() *Game {
	g := &Game{}
	return g
}

// ゲームの状態を60fpsで更新する
func (g *Game) Update() error {
	// ここにゲームの状態を書き込む
	return nil
}

// ゲームの状態を描画する
func (g *Game) Draw(screen *ebiten.Image) {
	// ここにゲームのレンダリングを描きます

}

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
