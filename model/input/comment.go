package input

type CommentInput struct {
	Message string `json:"message" valid:"required"`
	PhotoID int    `json:"photo_id" valid:"required"`
}

type CommentUpdateInput struct {
	Message string `json:"message" valid:"required"`
}

type DeleteComment struct {
	ID int `uri:"id" valid:"required"`
}

type UpdateComment struct {
	ID int `uri:"id" valid:"required"`
}
