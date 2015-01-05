package service

import (
	db "github.com/s-shin/gobbs/db"
	s "github.com/s-shin/gobbs/service"
	"testing"
)

func init_() {
	err := db.InitDB()
	if err != nil {
		panic(err)
	}
}

func TestCreateThread(t *testing.T) {
	init_()
	_, err := s.CreateThread("thread title", "", "post name", "post content")
	if err != nil {
		t.Error(err)
	}
}

func prepareAssets(n int) {
	init_()
	for i := 0; i < n; i++ {
		_, err := s.CreateThread("foo", "bar", "hoge", "piyo")
		if err != nil {
			panic(err)
		}
	}
}

func TestFetchRecentThreads(t *testing.T) {
	assetNum := 5
	prepareAssets(assetNum)
	num := 3
	threads := s.FetchRecentThreads(0, num)
	if len(threads) != 3 {
		t.Errorf("Should fetch %d threads.", 3)
	}
	if threads[0].Id != int64(assetNum) {
		t.Errorf("Most recent item id should be %d.", assetNum)
	}
}

func TestFetchThreadById(t *testing.T) {
	prepareAssets(1)
	thread := s.FetchThreadById(1)
	if thread == nil {
		t.Errorf("A thread should be fetched.")
	}
}
