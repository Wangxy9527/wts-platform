package serializer

import "time"

type ProblemSerialize struct {
}

type FightSerialize struct {
	ID        int    `json:"id"`
	UserBlue  int    `json:"user_blue"`
	UserRed   int    `json:"user_red"`
	ProblemId string `json:"problem_id"`

	ScoreBlue int `json:"score_blue"`
	ScoreRed  int `json:"score_red"`
	result    int `json:"result"`

	CreatedAt time.Time `json:"create_time"`
	UpdatedAt time.Time `json:"update_time"`
}
