package input

type SocialInput struct {
	Name string `json:"name" valid:"required"`
	URL  string `json:"social_media_url" valid:"required"`
}

type DeleteSocialMedia struct {
	ID int `uri:"id" valid:"required"`
}
type UpdateSocialMedia struct {
	ID int `uri:"id" valid:"required"`
}
