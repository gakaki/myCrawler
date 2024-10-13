package bilibili

import (
	"fmt"
	"myCrawler/utils"
)

func BilbiliAllPlayList() BilibiliPlayList {
	item := utils.ReadJSON[BilibiliPlayList]("bilibili_json/api_full.json")
	return item
}

func SimpleBiliBiliVideos(bilbiliAllPlayList BilibiliPlayList) []*BilibiliSimpleVideo {
	var videos = make([]*BilibiliSimpleVideo, 0)
	for index, item := range bilbiliAllPlayList.Data.UgcSeason.Sections[0].Episodes {
		v := BilibiliSimpleVideo{}
		v.Link = fmt.Sprintf("https://www.bilibili.com/video/%s/?spm_id_from=333.337.search-card.all.click", item.Bvid)
		v.Title = item.Title
		v.Bvid = item.Bvid
		fmt.Println(index, item.Title, v.Link, item.Bvid)
		videos = append(videos, &v)
	}
	return videos
}
