package model

// RoundInfoResp 获取某场比赛信息
type RoundInfoResp struct {
	PreCardsInfo   map[string][]int `json:"preCardsInfo"`
	FinalCardsInfo map[string][]int `json:"finalCardsInfo"`
	WinnerGroup    string           `json:"winnerGroup"`
}

// RoundInfoResp 获取全部比赛信息
type CardInfoResp struct {
	RoundInfos []RoundInfo `json:"roundInfos"`
	QueueInfo  []string    `json:"queueInfo"`
	GroupRank  []GroupRank `json:"groupRank"`
}

type RoundInfo struct {
	RoundNum    int      `json:"roundNum"`
	StartTime   int64    `json:"startTime"`
	GroupNames  []string `json:"groupNames"`
	WinnerGroup string   `json:"winnerGroup"`
}

type GroupRank struct {
	GroupName string `json:"groupName"`
	Score     int    `json:"score"`
}
