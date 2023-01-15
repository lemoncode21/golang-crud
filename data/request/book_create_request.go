package request

type BookCreateRequest struct {
	Name string `validate:"required min=1,max=100" json:"name"`
}
