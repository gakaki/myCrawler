package bilibili

import (
	"fmt"
	"testing"
)

func TestPrintAllBilibiliPlayLists(t *testing.T) {
	item := BilbiliAllPlayList()
	fmt.Println(item)
}

func TestPrintAllBilibiliVideos(t *testing.T) {
	items := SimpleBiliBiliVideos(BilbiliAllPlayList())
	fmt.Println("共有", len(items), "个视频，从老到新")
}
