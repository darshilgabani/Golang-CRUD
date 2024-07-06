package request

type UpdateUsersRequest struct {
	Id        int    `validate:"required"`
	FirstName string `validate:"required,max=200,min=1" json:"firstname"`
}
