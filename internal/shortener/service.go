package shortener

import "planck-lol/pkg/utils"

type LinkService struct {
	Storage *LinkStorage
}

func (s *LinkService) GetLink(code string) (*Link, error) {
	link, err := s.Storage.findByShortCode(code)
	return link, err
}

func (s *LinkService) CreateLink(code *string, url string) {
	if code == nil {
		g := utils.GenerateShortUrl()
		code = &g
	}

	s.Storage.insertLink(code, url, nil)
}
