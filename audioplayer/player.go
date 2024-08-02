package audioplayer

import (
	"context"
	"fmt"

	"github.com/lrstanley/go-ytdlp"
)

type AudioPlayer struct {
	player  *ytdlp.Command
	context *context.Context
}

func NewAudioPlayer() *AudioPlayer {
	// If yt-dlp isn't installed yet, download and cache it for further use.
	ctx := context.TODO()
	ytdlp.MustInstall(ctx, nil)

	return &AudioPlayer{
		player: ytdlp.New().
			Format("bestaudio").
			DefaultSearch("ytsearch").
			YesPlaylist().
			Output("-"),
		context: &ctx,
	}
}

func (v *AudioPlayer) YtdlpStdout(video_name string) string {
	result, err := v.player.Run(*v.context, video_name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.Stdout)

	return result.Stdout
}
