package schema

type Response struct {
	Status int `json:"status" bson:"status"`
	Code int `json:"code" bson:"code"`
	Message string `json:"message" bson:"message"`
	Data interface{} `json:"data" bson:"data"`
}
