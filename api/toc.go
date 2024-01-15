package api

type Toc struct {
	Chapters []string `goquery:"#episode_list_viewer > table > tr > .ion-bookmark"`
}

func GetToc(id int) (Toc, error) {
	toc, err := getHtml[Toc](endPointToc)
	if err != nil {
		return *new(Toc), err
	}

	return toc, nil
}
