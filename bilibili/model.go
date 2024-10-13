package bilibili

type BilibiliSimpleVideo struct {
	Bvid  string `json:"bvid"`
	Title string `json:"title"`
	Link  string `json:"link"`
}
type BilibiliPlayList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		DescV2    any    `json:"desc_v2"`
		State     int    `json:"state"`
		Duration  int    `json:"duration"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
			ArcPay        int `json:"arc_pay"`
			FreeWatch     int `json:"free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			Vt         int    `json:"vt"`
		} `json:"stat"`
		ArgueInfo struct {
			ArgueMsg  string `json:"argue_msg"`
			ArgueType int    `json:"argue_type"`
			ArgueLink string `json:"argue_link"`
		} `json:"argue_info"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		SeasonID           int    `json:"season_id"`
		Premiere           any    `json:"premiere"`
		TeenageMode        int    `json:"teenage_mode"`
		IsChargeableSeason bool   `json:"is_chargeable_season"`
		IsStory            bool   `json:"is_story"`
		IsUpowerExclusive  bool   `json:"is_upower_exclusive"`
		IsUpowerPlay       bool   `json:"is_upower_play"`
		IsUpowerPreview    bool   `json:"is_upower_preview"`
		EnableVt           int    `json:"enable_vt"`
		VtDisplay          string `json:"vt_display"`
		NoCache            bool   `json:"no_cache"`
		Pages              []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
			FirstFrame string `json:"first_frame"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool  `json:"allow_submit"`
			List        []any `json:"list"`
		} `json:"subtitle"`
		UgcSeason struct {
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Cover     string `json:"cover"`
			Mid       int    `json:"mid"`
			Intro     string `json:"intro"`
			SignState int    `json:"sign_state"`
			Attribute int    `json:"attribute"`
			Sections  []struct {
				SeasonID int    `json:"season_id"`
				ID       int    `json:"id"`
				Title    string `json:"title"`
				Type     int    `json:"type"`
				Episodes []struct {
					SeasonID  int    `json:"season_id"`
					SectionID int    `json:"section_id"`
					ID        int    `json:"id"`
					Aid       int    `json:"aid"`
					Cid       int    `json:"cid"`
					Title     string `json:"title"`
					Attribute int    `json:"attribute"`
					Arc       struct {
						Aid       int    `json:"aid"`
						Videos    int    `json:"videos"`
						TypeID    int    `json:"type_id"`
						TypeName  string `json:"type_name"`
						Copyright int    `json:"copyright"`
						Pic       string `json:"pic"`
						Title     string `json:"title"`
						Pubdate   int    `json:"pubdate"`
						Ctime     int    `json:"ctime"`
						Desc      string `json:"desc"`
						State     int    `json:"state"`
						Duration  int    `json:"duration"`
						Rights    struct {
							Bp            int `json:"bp"`
							Elec          int `json:"elec"`
							Download      int `json:"download"`
							Movie         int `json:"movie"`
							Pay           int `json:"pay"`
							Hd5           int `json:"hd5"`
							NoReprint     int `json:"no_reprint"`
							Autoplay      int `json:"autoplay"`
							UgcPay        int `json:"ugc_pay"`
							IsCooperation int `json:"is_cooperation"`
							UgcPayPreview int `json:"ugc_pay_preview"`
							ArcPay        int `json:"arc_pay"`
							FreeWatch     int `json:"free_watch"`
						} `json:"rights"`
						Author struct {
							Mid  int    `json:"mid"`
							Name string `json:"name"`
							Face string `json:"face"`
						} `json:"author"`
						Stat struct {
							Aid        int    `json:"aid"`
							View       int    `json:"view"`
							Danmaku    int    `json:"danmaku"`
							Reply      int    `json:"reply"`
							Fav        int    `json:"fav"`
							Coin       int    `json:"coin"`
							Share      int    `json:"share"`
							NowRank    int    `json:"now_rank"`
							HisRank    int    `json:"his_rank"`
							Like       int    `json:"like"`
							Dislike    int    `json:"dislike"`
							Evaluation string `json:"evaluation"`
							ArgueMsg   string `json:"argue_msg"`
							Vt         int    `json:"vt"`
							Vv         int    `json:"vv"`
						} `json:"stat"`
						Dynamic   string `json:"dynamic"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
						DescV2             any    `json:"desc_v2"`
						IsChargeableSeason bool   `json:"is_chargeable_season"`
						IsBlooper          bool   `json:"is_blooper"`
						EnableVt           int    `json:"enable_vt"`
						VtDisplay          string `json:"vt_display"`
					} `json:"arc"`
					Page struct {
						Cid       int    `json:"cid"`
						Page      int    `json:"page"`
						From      string `json:"from"`
						Part      string `json:"part"`
						Duration  int    `json:"duration"`
						Vid       string `json:"vid"`
						Weblink   string `json:"weblink"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
					} `json:"page"`
					Bvid  string `json:"bvid"`
					Pages []struct {
						Cid       int    `json:"cid"`
						Page      int    `json:"page"`
						From      string `json:"from"`
						Part      string `json:"part"`
						Duration  int    `json:"duration"`
						Vid       string `json:"vid"`
						Weblink   string `json:"weblink"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
					} `json:"pages"`
				} `json:"episodes"`
			} `json:"sections"`
			Stat struct {
				SeasonID int `json:"season_id"`
				View     int `json:"view"`
				Danmaku  int `json:"danmaku"`
				Reply    int `json:"reply"`
				Fav      int `json:"fav"`
				Coin     int `json:"coin"`
				Share    int `json:"share"`
				NowRank  int `json:"now_rank"`
				HisRank  int `json:"his_rank"`
				Like     int `json:"like"`
				Vt       int `json:"vt"`
				Vv       int `json:"vv"`
			} `json:"stat"`
			EpCount     int  `json:"ep_count"`
			SeasonType  int  `json:"season_type"`
			IsPaySeason bool `json:"is_pay_season"`
			EnableVt    int  `json:"enable_vt"`
		} `json:"ugc_season"`
		IsSeasonDisplay bool `json:"is_season_display"`
		UserGarb        struct {
			URLImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		HonorReply struct {
		} `json:"honor_reply"`
		LikeIcon          string `json:"like_icon"`
		NeedJumpBv        bool   `json:"need_jump_bv"`
		DisableShowUpInfo bool   `json:"disable_show_up_info"`
		IsStoryPlay       int    `json:"is_story_play"`
		IsViewSelf        bool   `json:"is_view_self"`
	} `json:"data"`
}

type BilibiliVideo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		DescV2    any    `json:"desc_v2"`
		State     int    `json:"state"`
		Duration  int    `json:"duration"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
			ArcPay        int `json:"arc_pay"`
			FreeWatch     int `json:"free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			Vt         int    `json:"vt"`
		} `json:"stat"`
		ArgueInfo struct {
			ArgueMsg  string `json:"argue_msg"`
			ArgueType int    `json:"argue_type"`
			ArgueLink string `json:"argue_link"`
		} `json:"argue_info"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		SeasonID           int    `json:"season_id"`
		Premiere           any    `json:"premiere"`
		TeenageMode        int    `json:"teenage_mode"`
		IsChargeableSeason bool   `json:"is_chargeable_season"`
		IsStory            bool   `json:"is_story"`
		IsUpowerExclusive  bool   `json:"is_upower_exclusive"`
		IsUpowerPlay       bool   `json:"is_upower_play"`
		IsUpowerPreview    bool   `json:"is_upower_preview"`
		EnableVt           int    `json:"enable_vt"`
		VtDisplay          string `json:"vt_display"`
		NoCache            bool   `json:"no_cache"`
		Pages              []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
			FirstFrame string `json:"first_frame"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool  `json:"allow_submit"`
			List        []any `json:"list"`
		} `json:"subtitle"`
		UgcSeason struct {
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Cover     string `json:"cover"`
			Mid       int    `json:"mid"`
			Intro     string `json:"intro"`
			SignState int    `json:"sign_state"`
			Attribute int    `json:"attribute"`
			Sections  []struct {
				SeasonID int    `json:"season_id"`
				ID       int    `json:"id"`
				Title    string `json:"title"`
				Type     int    `json:"type"`
				Episodes []struct {
					SeasonID  int    `json:"season_id"`
					SectionID int    `json:"section_id"`
					ID        int    `json:"id"`
					Aid       int    `json:"aid"`
					Cid       int    `json:"cid"`
					Title     string `json:"title"`
					Attribute int    `json:"attribute"`
					Arc       struct {
						Aid       int    `json:"aid"`
						Videos    int    `json:"videos"`
						TypeID    int    `json:"type_id"`
						TypeName  string `json:"type_name"`
						Copyright int    `json:"copyright"`
						Pic       string `json:"pic"`
						Title     string `json:"title"`
						Pubdate   int    `json:"pubdate"`
						Ctime     int    `json:"ctime"`
						Desc      string `json:"desc"`
						State     int    `json:"state"`
						Duration  int    `json:"duration"`
						Rights    struct {
							Bp            int `json:"bp"`
							Elec          int `json:"elec"`
							Download      int `json:"download"`
							Movie         int `json:"movie"`
							Pay           int `json:"pay"`
							Hd5           int `json:"hd5"`
							NoReprint     int `json:"no_reprint"`
							Autoplay      int `json:"autoplay"`
							UgcPay        int `json:"ugc_pay"`
							IsCooperation int `json:"is_cooperation"`
							UgcPayPreview int `json:"ugc_pay_preview"`
							ArcPay        int `json:"arc_pay"`
							FreeWatch     int `json:"free_watch"`
						} `json:"rights"`
						Author struct {
							Mid  int    `json:"mid"`
							Name string `json:"name"`
							Face string `json:"face"`
						} `json:"author"`
						Stat struct {
							Aid        int    `json:"aid"`
							View       int    `json:"view"`
							Danmaku    int    `json:"danmaku"`
							Reply      int    `json:"reply"`
							Fav        int    `json:"fav"`
							Coin       int    `json:"coin"`
							Share      int    `json:"share"`
							NowRank    int    `json:"now_rank"`
							HisRank    int    `json:"his_rank"`
							Like       int    `json:"like"`
							Dislike    int    `json:"dislike"`
							Evaluation string `json:"evaluation"`
							ArgueMsg   string `json:"argue_msg"`
							Vt         int    `json:"vt"`
							Vv         int    `json:"vv"`
						} `json:"stat"`
						Dynamic   string `json:"dynamic"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
						DescV2             any    `json:"desc_v2"`
						IsChargeableSeason bool   `json:"is_chargeable_season"`
						IsBlooper          bool   `json:"is_blooper"`
						EnableVt           int    `json:"enable_vt"`
						VtDisplay          string `json:"vt_display"`
					} `json:"arc"`
					Page struct {
						Cid       int    `json:"cid"`
						Page      int    `json:"page"`
						From      string `json:"from"`
						Part      string `json:"part"`
						Duration  int    `json:"duration"`
						Vid       string `json:"vid"`
						Weblink   string `json:"weblink"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
					} `json:"page"`
					Bvid  string `json:"bvid"`
					Pages []struct {
						Cid       int    `json:"cid"`
						Page      int    `json:"page"`
						From      string `json:"from"`
						Part      string `json:"part"`
						Duration  int    `json:"duration"`
						Vid       string `json:"vid"`
						Weblink   string `json:"weblink"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
					} `json:"pages"`
				} `json:"episodes"`
			} `json:"sections"`
			Stat struct {
				SeasonID int `json:"season_id"`
				View     int `json:"view"`
				Danmaku  int `json:"danmaku"`
				Reply    int `json:"reply"`
				Fav      int `json:"fav"`
				Coin     int `json:"coin"`
				Share    int `json:"share"`
				NowRank  int `json:"now_rank"`
				HisRank  int `json:"his_rank"`
				Like     int `json:"like"`
				Vt       int `json:"vt"`
				Vv       int `json:"vv"`
			} `json:"stat"`
			EpCount     int  `json:"ep_count"`
			SeasonType  int  `json:"season_type"`
			IsPaySeason bool `json:"is_pay_season"`
			EnableVt    int  `json:"enable_vt"`
		} `json:"ugc_season"`
		IsSeasonDisplay bool `json:"is_season_display"`
		UserGarb        struct {
			URLImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		HonorReply struct {
		} `json:"honor_reply"`
		LikeIcon          string `json:"like_icon"`
		NeedJumpBv        bool   `json:"need_jump_bv"`
		DisableShowUpInfo bool   `json:"disable_show_up_info"`
		IsStoryPlay       int    `json:"is_story_play"`
		IsViewSelf        bool   `json:"is_view_self"`
	} `json:"data"`
}
