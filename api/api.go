package api

import (
	"sync"
	"time"

	"github.com/DiheChen/PixivAPI/client"
	"github.com/gin-gonic/gin"
)

var Router = gin.Default()
var Client client.PixivClient

func GetIllustDetail(c *gin.Context) {
	illustID := c.Query("illust_id")
	illust, err := Client.Get("/v1/illust/detail", map[string]string{"illust_id": illustID})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, illust)
}

func GetIllustRanking(c *gin.Context) {
	date := c.DefaultQuery("date", time.Now().AddDate(0, 0, -1).Format("2006-1-2"))
	// 默认设为昨天
	mode := c.DefaultQuery("mode", "day")
	// 可选 week, month, day_male, day_female, week_original, week_rookie, day_r18, day_male_r18, day_female_r18
	offset := c.DefaultQuery("offset", "30")
	ranking, err := Client.Get("/v1/illust/ranking",
		map[string]string{"date": date, "mode": mode, "offset": offset})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, ranking)
}

func GetIllustFollow(c *gin.Context) {
	restrict := c.DefaultQuery("restrict", "public")
	follow, err := Client.Get("/v2/illust/follow", map[string]string{"restrict": restrict})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, follow)
}

func GetUserDetail(c *gin.Context) {
	userID := c.Query("user_id")
	user, err := Client.Get("/v1/user/detail", map[string]string{"user_id": userID})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}

func GetUserIllusts(c *gin.Context) {
	userID := c.Query("user_id")
	illustType := c.DefaultQuery("type", "illust") // 可选 illust, manga, ugoira
	offset := c.DefaultQuery("offset", "30")
	userIllusts, err := Client.Get("/v1/user/illusts",
		map[string]string{"user_id": userID, "type": illustType, "offset": offset})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, userIllusts)
}

func GetUserBookmarksIllust(c *gin.Context) {
	userID := c.Query("user_id")
	restrict := c.DefaultQuery("restrict", "public")
	userBookmarksIllust, err := Client.Get("/v1/user/bookmarks/illust",
		map[string]string{"user_id": userID, "restrict": restrict})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, userBookmarksIllust)
}

func GetUserFollowing(c *gin.Context) {
	userID := c.Query("user_id")
	offset := c.DefaultQuery("offset", "30")
	userFollowing, err := Client.Get("/v1/user/following",
		map[string]string{"user_id": userID, "offset": offset})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, userFollowing)
}

func GetUserFollowers(c *gin.Context) {
	userID := c.Query("user_id")
	offset := c.DefaultQuery("offset", "30")
	userFollower, err := Client.Get("/v1/user/follower",
		map[string]string{"user_id": userID, "offset": offset})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, userFollower)
}

func SearchIllust(c *gin.Context) {
	word := c.Query("word")
	searchTarget := c.DefaultQuery("search_target", "partial_match_for_tags")
	// 可选 partial_match_for_tags, exact_match_for_tags, title_and_caption
	sort := c.DefaultQuery("sort", "date_desc")
	// 可选 date_asc, popular_desc, 其中 popular_desc 需要高级会员
	duration := c.DefaultQuery("duration", "within_last_day")
	// 可选 within_last_week, within_last_month
	offset := c.DefaultQuery("offset", "30")
	searchIllust, err := Client.Get("/v1/search/illust",
		map[string]string{"word": word, "search_target": searchTarget, "sort": sort, "duration": duration, "offset": offset})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, searchIllust)
}

func GetTrendingTagsIllust(c *gin.Context) {
	trendingTagsIllust, err := Client.Get("/v1/trending-tags/illust", nil)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, trendingTagsIllust)
}

func GetUgoiraMetadata(c *gin.Context) {
	illustID := c.Query("illust_id")
	ugoiraMetadata, err := Client.Get("/v1/ugoira/metadata", map[string]string{"illust_id": illustID})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, ugoiraMetadata)
}

func GetUserNovels(c *gin.Context) {
	userID := c.Query("user_id")
	offset := c.DefaultQuery("offset", "30")
	userNovels, err := Client.Get("/v1/user/novels",
		map[string]string{"user_id": userID, "offset": offset})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, userNovels)
}

func GetNovelSeries(c *gin.Context) {
	seriesID := c.Query("series_id")
	novelSeries, err := Client.Get("/v1/novel/series", map[string]string{"series_id": seriesID})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, novelSeries)
}

func GetNovelDetail(c *gin.Context) {
	novelID := c.Query("novel_id")
	novelDetail, err := Client.Get("/v1/novel/detail", map[string]string{"novel_id": novelID})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, novelDetail)
}

func GetNovelText(c *gin.Context) {
	novelID := c.Query("novel_id")
	novelText, err := Client.Get("/v1/novel/text", map[string]string{"novel_id": novelID})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, novelText)
}

func GetNovelNew(c *gin.Context) {
	maxNovelID := c.Query("max_novel_id")
	novelNew, err := Client.Get("/v1/novel/new", map[string]string{"max_novel_id": maxNovelID})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, novelNew)
}

func SearchNovel(c *gin.Context) {
	word := c.Query("word")
	searchTarget := c.DefaultQuery("search_target", "partial_match_for_tags")
	sort := c.DefaultQuery("sort", "date_desc")
	mergePlainKeywordResults := c.DefaultQuery("merge_plain_keyword_results", "true")
	includeTranslatedTagResults := c.DefaultQuery("include_translated_tag_results", "true")
	duration := c.DefaultQuery("duration", "within_last_day")
	offset := c.DefaultQuery("offset", "30")
	searchNovel, err := Client.Get("/v1/search/novel",
		map[string]string{"word": word,
			"search_target":                  searchTarget,
			"sort":                           sort,
			"merge_plain_keyword_results":    mergePlainKeywordResults,
			"include_translated_tag_results": includeTranslatedTagResults,
			"duration":                       duration,
			"offset":                         offset,
		})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, searchNovel)
}

func AddIllustBookmark(c *gin.Context) {
	illustID := c.PostForm("illust_id")
	restrict := c.DefaultPostForm("restrict", "public")
	addIllustBookmark, err := Client.Post("/v1/illust/bookmark/add",
		map[string]string{"illust_id": illustID, "restrict": restrict})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, addIllustBookmark)
}

func DeleteIllustBookmark(c *gin.Context) {
	illustID := c.PostForm("illust_id")
	deleteIllustBookmark, err := Client.Post("/v1/illust/bookmark/delete",
		map[string]string{"illust_id": illustID})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, deleteIllustBookmark)
}

func AddUserFollow(c *gin.Context) {
	userID := c.PostForm("user_id")
	restrict := c.DefaultPostForm("restrict", "public")
	addUserFollow, err := Client.Post("/v1/user/follow/add",
		map[string]string{"user_id": userID, "restrict": restrict})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, addUserFollow)
}

func DeleteUserFollow(c *gin.Context) {
	userID := c.PostForm("user_id")
	deleteUserFollow, err := Client.Post("/v1/user/follow/delete",
		map[string]string{"user_id": userID})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, deleteUserFollow)
}

func refreshClient() {
	Client, _ = client.Refresh(GetRefreshToken())
}

func Run() {
	var mutex sync.Mutex
	mutex.Lock()
	refreshClient()
	mutex.Unlock()
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		for {
			time.Sleep(time.Minute * 30)
			refreshClient() // refresh client every 30 minutes
		}
	}()
	Router.GET("/v1/illust/detail", GetIllustDetail)
	Router.GET("/v1/illust/ranking", GetIllustRanking)
	Router.GET("/v2/illust/follow", GetIllustFollow)
	Router.GET("/v1/user/detail", GetUserDetail)
	Router.GET("/v1/user/illusts", GetUserIllusts)
	Router.GET("/v1/user/bookmarks/illust", GetUserBookmarksIllust)
	Router.GET("/v1/user/following", GetUserFollowing)
	Router.GET("/v1/user/followers", GetUserFollowers)
	Router.GET("/v1/search/illust", SearchIllust)
	Router.GET("/v1/trending-tags/illust", GetTrendingTagsIllust)
	Router.GET("/v1/ugoira/metadata", GetUgoiraMetadata)
	Router.GET("/v1/user/novels", GetUserNovels)
	Router.GET("/v2/novel/series", GetNovelSeries)
	Router.GET("/v2/novel/detail", GetNovelDetail)
	Router.GET("/v2/novel/text", GetNovelText)
	Router.GET("/v1/novel/new", GetNovelNew)
	Router.GET("/v1/search/novel", SearchNovel)
	// Router.POST("/v2/illust/bookmark/add", AddIllustBookmark)
	// Router.POST("/v1/illust/bookmark/delete", DeleteIllustBookmark)
	// Router.POST("/v1/user/follow/add", AddUserFollow)
	// Router.POST("/v1/user/follow/delete", DeleteUserFollow)
	_ = Router.Run(":8080")
}
