package api

type Toc struct {
}

func GetToc(id int) (Toc, error) {
	toc, err := getHtml[Toc](endPointToc)
	if err != nil {
		return *new(Toc), err
	}

	return toc, nil
}
