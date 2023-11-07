package handler

type LoginResponse struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
}

type CreateResponse struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
}
