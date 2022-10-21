package model

// RoundInfoResp 获取某场比赛信息
type RoundInfoResp struct {
	Id            int              `json:"id"`
	StartTime     string           `json:"startTime"`
	EndTime       string           `json:"endTime"`
	GroupNames    []string         `json:"groupNames"`
	PerCardInfo   map[string][]int `json:"perCardInfo"`
	FinalCardInfo map[string][]int `json:"finalCardInfo"`
	WinnerGroup   string           `json:"winnerGroup"`
}

// RoundInfoResp 获取全部比赛信息
type CardInfoResp struct {
	RoundInfos []RoundInfo `json:"roundInfos"`
	QueueInfo  []string    `json:"queueInfo"`
	GroupRank  []GroupRank `json:"groupRank"`
}

type RoundInfo struct {
	RoundNum    int      `json:"roundNum"`
	StartTime   string   `json:"startTime"`
	GroupNames  []string `json:"groupNames"`
	WinnerGroup string   `json:"winnerGroup"`
}

type GroupRank struct {
	GroupName string `json:"groupName"`
	Score     int    `json:"score"`
}
