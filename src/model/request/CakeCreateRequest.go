package request

type CakeCreateRequest struct {
	Title       string  `validate:"required,min=1,max=255"`
	Description string  `validate:"max=500"`
	Rating      float32 `validate:"numeric"`
	Image       string  `validate:"max=191"`
}
