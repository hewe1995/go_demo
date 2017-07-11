package library

import (
	"music/model"
	"errors"
)

type MusicManager struct {
	musics []model.Music
}

/**
构造函数
 */
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]model.Music, 0)}
}

/**
数量
 */
func (m *MusicManager) Len() int {
	return len(m.musics)
}

/**
按索引获取
 */
func (m *MusicManager) Get(index int) (music *model.Music, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

/**
按名称检索
 */
func (m *MusicManager) Find(name string) *model.Music {
	if len(m.musics) == 0 {
		return nil
	}
	for _, v := range m.musics {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

/**
添加
 */
func (m *MusicManager) Add(music *model.Music) {
	m.musics = append(m.musics, *music)
}

/**
删除
 */
func (m *MusicManager) Remove(index int) (removeMusic *model.Music) {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removeMusic = &m.musics[index]
	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		m.musics = make([]model.Music, 0)
	} else {
		m.musics = m.musics[: index-1]
	}
	return
}
