package schema

type Task struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}
