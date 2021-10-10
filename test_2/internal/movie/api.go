package movie

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inouttt/stkt/test_2/internal/constant"
	ll "github.com/inouttt/stkt/test_2/internal/log"
	"github.com/inouttt/stkt/test_2/pkg/response"
)

type Handler interface {
	Handle(app *fiber.App) error
}

func (m *movie) Handle(app *fiber.App) {
	app.Get("/", m.HandleGet)
}

type movie struct {
	S Services
}

func NewMovie(url string, ls ll.Services) movie {
	s := NewServices(url, ls)
	return movie{
		S: s,
	}
}

func (m *movie) HandleGet(c *fiber.Ctx) error {
	if key := c.Query(constant.ParamApikey); key == "" {
		c.Context().Error(constant.GetError(constant.CodeBadRequest))
		return c.Next()
	}

	var resp string

	var sr SearchRequest
	c.QueryParser(&sr)
	if sr.Search != "" {
		resp, _ = m.S.Search(sr)
		c = response.Generate(c, constant.CodeSuccess, resp)
		return nil
	}

	var dr DetailRequest
	c.QueryParser(&dr)
	if dr.Id != "" {
		resp, _ = m.S.Detail(dr)
		c = response.Generate(c, constant.CodeSuccess, resp)
		return nil
	}

	c.Context().Error(constant.GetError(constant.CodeBadRequest))
	return nil
}
