package model

type Todo struct {
	ID        string `json:"id" bson:"_id"`
	Todo      string `json:"text" bson:"text"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	IsDone    bool   `json:"is_done" bson:"is_done"`
}
