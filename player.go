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

	for music_option < 1 || music_option > videos.GetNumMusicLinks() {
		fmt.Print(videos.HELP_MESSAGE)
		fmt.Scanf("%d", &music_option)
	}

	var music_link string
	switch music_option {
	case 1:
		music_link = videos.LOFI
	case 2:
		music_link = videos.LOFI_PIANO
	case 3:
		music_link = videos.LOFI_DARK
	case 4:
		music_link = videos.LOFI_SYNCWAVE
	case 5:
		music_link = videos.LOFI_HIPHOP
	case 6:
		music_link = videos.BOYCE_AVENUE
	}

	return music_link
}

func Play() error {
	song := ChooseMusic()
	ytdlp := exec.Command("yt-dlp", "-f", "bestaudio", "--yes-playlist", song, "-o", "-")
	mpv := exec.Command("mpv", "--volume=50", "-")

	ytdlpOut, err := ytdlp.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create pipe for yt-dlp stdout: %v", err)
	}
	defer ytdlpOut.Close()

	mpv.Stdin = ytdlpOut

	if err := ytdlp.Start(); err != nil {
		return fmt.Errorf("failed to start yt-dlp: %v", err)
	}

	defer func() {
		if err := ytdlp.Process.Kill(); err != nil {
			log.Printf("failed to kill yt-dlp process: %v", err)
		}
		ytdlp.Wait()
	}()

	if err := mpv.Start(); err != nil {
		return fmt.Errorf("failed to start mpv: %v", err)
	}

	defer func() {
		if err := mpv.Process.Kill(); err != nil {
			log.Printf("failed to kill mpv process: %v", err)
		}
		mpv.Wait()
	}()

	// Esperar até que um sinal seja recebido para finalizar a execução
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	return nil
}
