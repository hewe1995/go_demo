package player

import "fmt"

/**
播放接口
 */
type Player interface {
	Play(source string)
}

func Play(source string, mtype string)  {
	var p Player
	switch mtype {
	case "MP3":
		p = &MP3Player{}
	default:
		fmt.Println("unsupport !")
		return
	}
	p.Play(source)
}
