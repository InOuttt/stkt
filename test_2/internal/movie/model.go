package movie

type SearchRequest struct {
	Apikey string `query:"apikey" validate:"required"`
	Search string `query:"s"`
	Page   string `query:"page"`
}

type DetailRequest struct {
	Apikey string `query:"apikey" validate:"required"`
	Id     string `query:"i"`
}
