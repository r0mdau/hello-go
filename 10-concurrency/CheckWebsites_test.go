package _0_concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://futurewe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://romaindauby.fr",
		"waat://futurewe.geds",
	}

	want := map[string]bool{
		"http://google.com":     true,
		"http://romaindauby.fr": true,
		"waat://futurewe.geds":  false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
