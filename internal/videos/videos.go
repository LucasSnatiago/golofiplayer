package videos

const LOFI = "https://www.youtube.com/watch?v=jfKfPfyJRdk"
const LOFI_PIANO = "https://www.youtube.com/watch?v=ZXzTfdP5dDs"
const LOFI_DARK = "https://www.youtube.com/watch?v=S_MOd40zlYU"
const LOFI_SYNCWAVE = "https://www.youtube.com/watch?v=4xDzrJKXOOY"
const LOFI_HIPHOP = "https://www.youtube.com/watch?v=rUxyKA_-grg"
const BOYCE_AVENUE = "https://www.youtube.com/watch?v=FeluT-nU0Qk&list=RD1sioip9Uc4o"

const HELP_MESSAGE = "Please choose your prefered song:\n 1- Lofi\n 2- Lofi Piano\n 3- Dark Lofi\n 4- SyncWave Lofi\n 5- Lofi Hip Hop\n 6- Boyce Avenue\n"

// Get number of music links available
func GetNumMusicLinks() int {
	return 6
}
