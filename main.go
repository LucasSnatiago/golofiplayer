package main

import (
	"log"

	"github.com/LucasSnatiago/golofiplayer/internal/videos"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	// Basic controling system
	isPlaying  bool
	isPaused   bool
	lastOption uint
	option     uint

	// Links for the songs
	videos *videos.MusicLinks

	// Audio System
	audioContext *audio.Context
	audioPlayer  *audio.Player
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.Key0) {
		g.option = 0
	} else if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.option = 1
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.option = 2
	} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
		g.option = 3
	} else if inpututil.IsKeyJustPressed(ebiten.Key4) {
		g.option = 4
	} else if inpututil.IsKeyJustPressed(ebiten.Key5) {
		g.option = 5
	} else if inpututil.IsKeyJustPressed(ebiten.Key6) {
		g.option = 6
	}

	if !g.isPaused && g.option < uint(g.videos.Length()) {
		// go Play(g.videos.Link(option))
		g.isPlaying = true
		g.isPaused = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.isPlaying = !g.isPlaying
		g.isPaused = !g.isPaused
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// If there is no inpnut, skip draw.
	if g.lastOption == g.option {
		// As SetScreenClearedEveryFrame(false) is called, the screen is not modified.
		// In this case, Ebitengine optimizes and reduces GPU usages.
		return
	}
	g.lastOption = g.option

	screen.Clear()

	ebitenutil.DebugPrintAt(screen, g.videos.HelpMessage(), 0, 0)
	if g.option != 255 {
		ebitenutil.DebugPrintAt(screen, "->", 0, int(g.option)*16)
	}

	if g.isPaused {
		ebitenutil.DebugPrintAt(screen, "Music Paused!", 0, 16*14)
	} else if g.isPlaying {
		ebitenutil.DebugPrintAt(screen, "Music Playing!", 0, 16*14)
	} else {
		ebitenutil.DebugPrintAt(screen, "There is no Music Playing", 0, 16*14)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Music Player!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetScreenClearedEveryFrame(false)

	game := &Game{
		isPlaying:    false,
		isPaused:     false,
		option:       255,
		videos:       videos.New(),
		audioContext: audio.NewContext(48000),
		audioPlayer:  nil,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
