package request

type UpdateTagsRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required, min=1, max=10" json:"name"`
}
