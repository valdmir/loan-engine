package http

type Error struct {
	Code        int         `json:"code"`
	Message     interface{} `json:"message"`
	Description interface{} `json:"description"`
}
