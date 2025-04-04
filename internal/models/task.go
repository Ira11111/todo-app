package models

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StatDate    string `json:"stat_date"`
	EndDate     string `json:"end_date"`
}

type TaskList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}

type ListItem struct {
	Id     int
	TaskId int
	ListId int
}
