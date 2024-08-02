package main

import (
	"log"
	"os/exec"
)

func (c *ControlMusic) Play(song string) {
	ytdlp := exec.Command("yt-dlp", "-f", "bestaudio", "--yes-playlist", song, "-o", "-")

	ytdlpOut, err := ytdlp.StdoutPipe()
	if err != nil {
		log.Fatalf("failed to create pipe for yt-dlp stdout: %v", err)
	}
	defer ytdlpOut.Close()

	if err := ytdlp.Start(); err != nil {
		log.Fatalf("failed to start yt-dlp: %v", err)
	}

	defer func() {
		if err := ytdlp.Process.Kill(); err != nil {
			log.Printf("failed to kill yt-dlp process: %v", err)
		}
		ytdlp.Wait()
	}()

	// Play ebiten Music

	// Esperar até que um sinal seja recebido para finalizar a execução
	<-c.shouldStop
}
