package input

type PhotoCreateInput struct {
	Title    string `json:"title" valid:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" valid:"required"`
}

type PhotoUpdateInput struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

type PhotoUpdateIDUser struct {
	ID int `uri:"id" valid:"required"`
}

type PhotoDeleteIDUser struct {
	ID int `uri:"id" valid:"required"`
}
