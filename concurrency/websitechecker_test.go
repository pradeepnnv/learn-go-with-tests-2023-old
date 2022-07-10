package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "dummyurl"
}

func mockSlowWebsitechecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
		"dummyurl",
	}
	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    true,
		"dummyurl":                   false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v but got %v", want, got)
	}
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(mockSlowWebsitechecker, urls)
	}
}
