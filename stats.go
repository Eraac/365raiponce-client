package raiponce

type GeneralStats struct {
	CountRemarksPublish     int `json:"count_remarks_publish"`
	CountRemarksUnpublish   int `json:"count_remarks_unpublish"`
	CountResponsesPublish   int `json:"count_responses_publish"`
	CountResponsesUnpublish int `json:"count_responses_unpublish"`
	CountUsers              int `json:"count_users"`
	CountVotesRemarksSexist int `json:"count_votes_remarks_sexist"`
	CountVotesRemarksLived  int `json:"count_votes_remarks_lived"`
	CountVotesResponses     int `json:"count_votes_responses"`
}

type Stats struct {
	Count          int `json:"nb"`
	CreatedYear    int `json:"created_year"`
	CreatedMonth   int `json:"created_month"`
	CreatedDay     int `json:"created_day"`
	PostedYear     int `json:"posted_year"`
	PostedMonth    int `json:"posted_month"`
	PostedDay      int `json:"posted_day"`
	Emotion        int `json:"emotion_id"`
	Theme          int `json:"theme_id"`
	ScaleEmotion   int `json:"scale_emotion"`
	Remark         int `json:"remark_id"`
	Author         int `json:"author_id"`
	User           int `json:"user_id"`
	Action         int `json:"action"`
	IsUsedForScore int `json:"is_used_for_score"`
	Type           int `json:"type"`
	Voter          int `json:"voter_id"`
	Receiver       int `json:"receiver_id"`
	Response       int `json:"response_id"`
}

type CollectionStats []Stats
type ByCreated struct{ CollectionStats }
type ByPosted struct{ CollectionStats }

const (
	statsURI               = "/stats"
	StatsRemarksURI        = statsURI + "/remarks"
	StatsResponsesURI      = statsURI + "/responses"
	StatsScoresURI         = statsURI + "/scores"
	StatsUsersURI          = statsURI + "/users"
	StatsVotesRemarksURI   = statsURI + "/votes/remarks"
	StatsVotesResponsesURI = statsURI + "/votes/responses"

	FilterFrom           = "from"
	FilterTo             = "to"
	FilterCreatedBefore  = "created_before"
	FilterCreatedAfter   = "created_after"
	FilterPostedBefore   = "posted_before"
	FilterPostedAfter    = "posted_after"
	FilterCreatedYear    = "created_year"
	FilterCreatedMonth   = "created_month"
	FilterCreatedDay     = "created_day"
	FilterPostedYear     = "posted_year"
	FilterPostedMonth    = "posted_month"
	FilterPostedDay      = "posted_day"
	FilterEmotion        = "emotion"
	FilterTheme          = "theme"
	FilterScaleEmotion   = "scale_emotion"
	FilterRemark         = "remark"
	FilterResponse       = "response"
	FilterAuthor         = "author"
	FilterUser           = "user"
	FilterAction         = "action"
	FilterIsUsedForScore = "is_used_for_score"
	FilterType           = "type"
	FilterVoter          = "voter"
	FilterReceiver       = "receiver"

	OrderCreatedYear    = "created_year"
	OrderCreatedMonth   = "created_month"
	OrderCreatedDay     = "created_day"
	OrderPostedYear     = "posted_year"
	OrderPostedMonth    = "posted_month"
	OrderPostedDay      = "posted_day"
	OrderEmotion        = "emotion"
	OrderTheme          = "theme"
	OrderScaleEmotion   = "scale_emotion"
	OrderCount          = "count"
	OrderRemark         = "remark"
	OrderResponse       = "response"
	OrderAuthor         = "author"
	OrderUser           = "user"
	OrderAction         = "action"
	OrderIsUsedForScore = "is_used_for_score"
	OrderType           = "type"
	OrderVoter          = "voter"
	OrderReceiver       = "receiver"

	GroupCreatedYear    = "created_year"
	GroupCreatedMonth   = "created_month"
	GroupCreatedDay     = "created_day"
	GroupPostedYear     = "posted_year"
	GroupPostedMonth    = "posted_month"
	GroupPostedDay      = "posted_day"
	GroupEmotion        = "emotion"
	GroupTheme          = "theme"
	GroupScaleEmotion   = "scale_emotion"
	GroupRemark         = "remark"
	GroupResponse       = "response"
	GroupAuthor         = "author"
	GroupUser           = "user"
	GroupAction         = "action"
	GroupIsUsedForScore = "is_used_for_score"
	GroupType           = "type"
	GroupVoter          = "voter"
	GroupReceiver       = "receiver"
)

// === Sort implementation ===
func (stats CollectionStats) Len() int {
	return len(stats)
}

func (stats CollectionStats) Swap(i, j int) {
	stats[i], stats[j] = stats[j], stats[i]
}

func (stats ByCreated) Less(i, j int) bool {
	if stats.CollectionStats[i].CreatedYear != stats.CollectionStats[j].CreatedYear {
		return stats.CollectionStats[i].CreatedYear < stats.CollectionStats[j].CreatedYear
	}

	if stats.CollectionStats[i].CreatedMonth != stats.CollectionStats[j].CreatedMonth {
		return stats.CollectionStats[i].CreatedMonth < stats.CollectionStats[j].CreatedMonth
	}

	if stats.CollectionStats[i].CreatedDay != stats.CollectionStats[j].CreatedDay {
		return stats.CollectionStats[i].CreatedDay < stats.CollectionStats[j].CreatedDay
	}

	return false
}

func (stats ByPosted) Less(i, j int) bool {
	if stats.CollectionStats[i].PostedYear != stats.CollectionStats[j].PostedYear {
		return stats.CollectionStats[i].PostedYear < stats.CollectionStats[j].PostedYear
	}

	if stats.CollectionStats[i].PostedMonth != stats.CollectionStats[j].PostedMonth {
		return stats.CollectionStats[i].PostedMonth < stats.CollectionStats[j].PostedMonth
	}

	if stats.CollectionStats[i].PostedDay != stats.CollectionStats[j].PostedDay {
		return stats.CollectionStats[i].PostedDay < stats.CollectionStats[j].PostedDay
	}

	return false
}

// ====

func (client *Client) GetGeneralStats(query *QueryFilter) *GeneralStats {
	stats := &GeneralStats{}

	client.cget(statsURI, stats, query)

	return stats
}

func (client *Client) GetStats(uri string, query *QueryFilter) *CollectionStats {
	stats := &CollectionStats{}

	client.cget(uri, stats, query)

	return stats
}
