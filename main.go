package main

import (
	"bytes"
	_ "embed" // ファイルから自動登録する
	"image"
	_ "image/png" // 画像のエンコードに必要
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenX  = 640
	screenY  = 480
	fontSize = 10
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

// 初期処理
func init() {
	log.Println("Init func.")
	// シード値の生成
	rand.Seed(time.Now().UnixNano())

	// ファイルをバイトで読み込み、imageとしてデコードする
	// byteDinosaur1Imgの読み込み
	img, _, err := image.Decode(bytes.NewReader(byteDinosaur1Img))
	if err != nil {
		log.Fatal(err)
	}
	dinosaur1Img = ebiten.NewImageFromImage(img)

	// byteDinosaur2Imgの読み込み
	img, _, err = image.Decode(bytes.NewReader(byteDinosaur2Img))
	if err != nil {
		log.Fatal(err)
	}
	dinosaur2Img = ebiten.NewImageFromImage(img)

	// byteTreeSmallImgの読み込み
	img, _, err = image.Decode(bytes.NewReader(byteTreeSmallImg))
	if err != nil {
		log.Fatal(err)
	}
	treeSmallImg = ebiten.NewImageFromImage(img)

	// byteTreeBigImgの読み込み
	img, _, err = image.Decode(bytes.NewReader(byteTreeBigImg))
	if err != nil {
		log.Fatal(err)
	}
	treeBigImg = ebiten.NewImageFromImage(img)

	// byteGroundImgの読み込み
	img, _, err = image.Decode(bytes.NewReader(byteGroundImg))
	if err != nil {
		log.Fatal(err)
	}
	groundImg = ebiten.NewImageFromImage(img)

	// フォントの読み込み
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

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
