package pia

import "github.com/abiriadev/pia/api"

type PiaClient struct {
	token string
}

func NewPiaClient(token string) PiaClient {
	return PiaClient{token}
}

type Novel struct {
	id int
}

func (n Novel) Content() (string, error) {
	s, err := api.GetScript(n.id)
	if err != nil {
		return "", err
	}
	return s.Textify(), nil
}

func (p *PiaClient) Novel(id int) Novel {
	return Novel{id}
}
