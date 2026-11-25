package urlshort

import "go.yaml.in/yaml/v4"

type pair struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYaml(yml []byte) (map[string]string, error) {
	var yamlMap []pair

	err := yaml.Unmarshal(yml, &yamlMap)

	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)

	for _, v := range yamlMap {
		pathsToUrls[v.Path] = v.URL
	}

	return pathsToUrls, nil
}
