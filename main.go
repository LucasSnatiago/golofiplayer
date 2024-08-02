package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/LucasSnatiago/golofiplayer/internal/videos"
)

func ChooseMusic() string {
	music_option := 0

	videos := videos.New()

	for music_option < 1 || music_option > videos.Length() {
		fmt.Print(videos.HelpMessage())
		fmt.Scanf("%d", &music_option)
	}

	return videos.Links[music_option-1].Link
}

func main() {
	song := ChooseMusic()
	ytdlp := exec.Command("yt-dlp", "-f", "bestaudio", "--yes-playlist", song, "-o", "-")
	ffplay := exec.Command("ffplay", "-nodisp", "-volume", "25", "-i", "-")

	ytdlpOut, err := ytdlp.StdoutPipe()
	if err != nil {
		log.Fatalf("failed to create pipe for yt-dlp stdout: %v", err)
	}
	defer ytdlpOut.Close()

	ffplay.Stdin = ytdlpOut

	if err := ytdlp.Start(); err != nil {
		log.Fatalf("failed to start yt-dlp: %v", err)
	}

	defer func() {
		if err := ytdlp.Process.Kill(); err != nil {
			log.Printf("failed to kill yt-dlp process: %v", err)
		}
		ytdlp.Wait()
	}()

	if err := ffplay.Start(); err != nil {
		log.Fatalf("failed to start ffplay: %v", err)
	}

	defer func() {
		if err := ffplay.Process.Kill(); err != nil {
			log.Printf("failed to kill ffplay process: %v", err)
		}
		ffplay.Wait()
	}()

	// Esperar até que um sinal seja recebido para finalizar a execução
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
