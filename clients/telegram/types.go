package telegram

type Updates struct {
	ID      int    `json:"update_id"`
	Message string `json:"message"`
}

type UpdateResp struct {
	Ok     bool      `json:"ok"`
	Result []Updates `json:"result"`
}
