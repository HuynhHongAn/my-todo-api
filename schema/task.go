package schema

type Task struct {
	Id          int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Type        int    `json:"type" bson:"type"`
	IsDone      int    `json:"is_done" bson:"is_done"`
}

type TaskListFilter struct {
	Type int `json:"type"`
}