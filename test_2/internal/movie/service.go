package movie

import (
	"github.com/inouttt/stkt/test_2/internal/constant"
	ll "github.com/inouttt/stkt/test_2/internal/log"
	"github.com/inouttt/stkt/test_2/internal/utils"
	"github.com/inouttt/stkt/test_2/pkg/http"
)

type Services interface {
	Search(sr SearchRequest) (string, error)
	Detail(dr DetailRequest) (string, error)
}

type Service struct {
	Url string
	Log ll.Services
}

func NewServices(url string, log ll.Services) Services {
	return &Service{
		Url: url,
		Log: log,
	}
}

func (s *Service) log(req map[string]string, resp string) error {
	el := utils.MapToLog(req)
	el.ResponseBody = resp

	return s.Log.Create(el)
}

func (s *Service) Search(sr SearchRequest) (string, error) {
	// prepare query param
	qry := map[string]string{
		constant.ParamApikey: sr.Apikey,
		constant.ParamSearch: sr.Search,
		constant.ParamPage:   sr.Page,
	}

	// get from url target
	resp, err := http.Get(s.Url, qry)

	// log even error
	go s.log(qry, resp)

	return resp, err
}

func (s *Service) Detail(dr DetailRequest) (string, error) {
	// prepare query param
	qry := map[string]string{
		constant.ParamApikey: dr.Apikey,
		constant.ParamId:     dr.Id,
	}
	// get from url target
	resp, err := http.Get(s.Url, qry)

	// log even error
	go s.log(qry, resp)

	return resp, err
}
