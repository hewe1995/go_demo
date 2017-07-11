package mp3player

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat int
	progress int
}

func (mp3 *MP3Player) Play(source string)  {
	fmt.Println("play mp3 music", source)

	mp3.progress = 0

	for mp3.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		mp3.progress += 10
	}
	fmt.Println("\n finish play ", source)
}
