package models

type Session struct {
	SessionCode string `json:session_id`
	IsStarted   bool   `json:is_started`
	Word        string `json:word`
	SpyGamerId  string `spy_gamer_id`
}
