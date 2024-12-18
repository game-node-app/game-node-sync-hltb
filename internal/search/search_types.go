package search

type HLTBResponse struct {
	Color           string             `json:"color"`
	Title           string             `json:"title"`
	Category        string             `json:"category"`
	Count           int                `json:"count"`
	PageCurrent     int                `json:"pageCurrent"`
	PageTotal       int                `json:"pageTotal"`
	PageSize        int                `json:"pageSize"`
	Data            []HLTBResponseItem `json:"data"`
	UserData        []interface{}      `json:"userData"`
	DisplayModifier interface{}        `json:"displayModifier"`
}

type HLTBResponseItem struct {
	GameId          int    `json:"game_id"`
	GameName        string `json:"game_name"`
	GameNameDate    int    `json:"game_name_date"`
	GameAlias       string `json:"game_alias"`
	GameType        string `json:"game_type"`
	GameImage       string `json:"game_image"`
	CompLvlCombine  int    `json:"comp_lvl_combine"`
	CompLvlSp       int    `json:"comp_lvl_sp"`
	CompLvlCo       int    `json:"comp_lvl_co"`
	CompLvlMp       int    `json:"comp_lvl_mp"`
	CompMain        int    `json:"comp_main"`
	CompPlus        int    `json:"comp_plus"`
	Comp100         int    `json:"comp_100"`
	CompAll         int    `json:"comp_all"`
	CompMainCount   int    `json:"comp_main_count"`
	CompPlusCount   int    `json:"comp_plus_count"`
	Comp100Count    int    `json:"comp_100_count"`
	CompAllCount    int    `json:"comp_all_count"`
	InvestedCo      int    `json:"invested_co"`
	InvestedMp      int    `json:"invested_mp"`
	InvestedCoCount int    `json:"invested_co_count"`
	InvestedMpCount int    `json:"invested_mp_count"`
	CountComp       int    `json:"count_comp"`
	CountSpeedrun   int    `json:"count_speedrun"`
	CountBacklog    int    `json:"count_backlog"`
	CountReview     int    `json:"count_review"`
	ReviewScore     int    `json:"review_score"`
	CountPlaying    int    `json:"count_playing"`
	CountRetired    int    `json:"count_retired"`
	ProfilePopular  int    `json:"profile_popular"`
	ReleaseWorld    int    `json:"release_world"`
}

type HLTBSearchRequestOptionsUsers struct {
	//Id           string `json:"id"`
	SortCategory string `json:"sortCategory"`
}

type HLTBSearchRequestOptions struct {
	Users HLTBSearchRequestOptionsUsers `json:"users"`
}

type HLTBSearchRequest struct {
	SearchType    string                   `json:"searchType"`
	SearchTerms   []string                 `json:"searchTerms"`
	SearchPage    int                      `json:"searchPage"`
	Size          int                      `json:"size"`
	SearchOptions HLTBSearchRequestOptions `json:"searchOptions"`
}
