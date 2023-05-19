package model

type TrainData struct {
	Categories    []string   `json:"categories"`
	Conversations [][]string `json:"conversations"`
}
