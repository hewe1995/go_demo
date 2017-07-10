package library

import (
	"testing"
	"music/model"
)

func TestOpt(t *testing.T) {
	mm := NewMusicManager()
	m0 := &model.Music{"1", "My love", "tom", "http://www.hewe.vip", "MP3"}
	mm.Add(m0)
	if mm.Len() != 1 {
		t.Error("add faild")
	}
	m := mm.Find("My love")
	if m == nil {
		t.Error("find faild")
	}
	m, err := mm.Get(0)
	if err != nil {
		t.Error("get faild")
	}
	m = mm.Remove(0)
	if m == nil {
		t.Error("remove faild")
	}
}
