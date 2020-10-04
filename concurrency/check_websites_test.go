package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://invaliduri.com/"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://www.google.com/",
		"https://www.yahoo.com/",
		"waat://invaliduri.com/",
	}

	want := map[string]bool{
		"https://www.google.com/": true,
		"https://www.yahoo.com/":  true,
		"waat://invaliduri.com/":  false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
