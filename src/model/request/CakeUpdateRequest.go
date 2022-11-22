package request

type CakeUpdateRequest struct {
	Id          int     `validate:"required"`
	Title       string  `validate:"min=1,max=255"`
	Description string  `validate:"max=500"`
	Rating      float32 `validate:"numeric"`
	Image       string  `validate:"max=191"`
}
