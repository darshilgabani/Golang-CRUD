package request

type CreateUsersRequest struct {
	FirstName string `validate:"required,min=1,max=200" json:"firstname"`
}
