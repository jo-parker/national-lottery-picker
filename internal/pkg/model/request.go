package model

type EnterDrawsRequest struct {
	Draws       []Draw      `json:"draws"`
	Credentials Credentials `json:"credentials"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
