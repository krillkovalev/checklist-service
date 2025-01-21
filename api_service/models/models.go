package models

type RequestBody struct {
	ID	string `json:"id"`
}

type Messsage struct {
    Timestamp   string  `json:"timestamp"`
    Action      string  `json:"action"`
}