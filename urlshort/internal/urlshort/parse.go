package urlshort

type pair struct {
	Path string `json:"path" yaml:"path"`
	URL  string `json:"url" yaml:"url"`
}

type unmarshalFunction func(in []byte, out any) (err error)

func parseStructured(in []byte, unmarshalFunc unmarshalFunction) (map[string]string, error) {
	var pairs []pair

	err := unmarshalFunc(in, &pairs)

	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)

	for _, v := range pairs {
		pathsToUrls[v.Path] = v.URL
	}

	return pathsToUrls, nil
}
