package request

type BookUpdateRequest struct {
	Id   int
	Name string `validate:"required min=1,max=100" json:"name"`
}
