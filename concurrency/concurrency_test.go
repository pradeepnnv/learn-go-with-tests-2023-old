package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebSiteChecker(url string) bool {
	return url != "waat://furhurterwe.gedds"
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}

}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.gedds",
	}
	want := map[string]bool{
		"waat://furhurterwe.gedds":   false,
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
	}

	got := CheckWebsites(mockWebSiteChecker, websites)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v,got %v", want, got)
	}

}
