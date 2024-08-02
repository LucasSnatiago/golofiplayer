package videos

import "fmt"

type MusicInfo struct {
	Name string
	Link string
}

type MusicLinks struct {
	Links []MusicInfo
}

func New() *MusicLinks {
	var m MusicLinks

	m.Links = []MusicInfo{
		{"Lofi", "https://www.youtube.com/watch?v=jfKfPfyJRdk"},
		{"Lofi Piano", "https://www.youtube.com/watch?v=ZXzTfdP5dDs"},
		{"Dark Lofi", "https://www.youtube.com/watch?v=S_MOd40zlYU"},
		{"Lofi SyncWave", "https://www.youtube.com/watch?v=4xDzrJKXOOY"},
		{"Lofi HipHop", "https://www.youtube.com/watch?v=jfKfPfyJRdk"},
		{"Boyce Avenue Playlist", "https://youtube.com/playlist?list=PLtd--8s9Fp4XniKoIcLOD1RNTWwAtnCja"},
	}

	return &m
}

// Get number of music links available
func (m *MusicLinks) Length() int {
	return len(m.Links)
}

// Help message
func (m *MusicLinks) HelpMessage() string {
	help := "Please choose your prefered song:\n"

	for i := range m.Length() {
		help += fmt.Sprintf("   %d- %s\n", i+1, m.Links[i].Name)
	}

	return help
}

func (m *MusicLinks) Link(index uint) string {
	return m.Links[index].Link
}
