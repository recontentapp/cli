package fileformat

import "errors"

type Fileformat string

const (
	FileformatJSON       Fileformat = "json"
	FileformatNestedJSON Fileformat = "json_nested"
	FileformatYAML       Fileformat = "yaml"
	FileformatNestedYAML Fileformat = "yaml_nested"
)

func New(raw string) (*Fileformat, error) {
	possibleValues := []string{
		string(FileformatJSON),
		string(FileformatNestedJSON),
		string(FileformatYAML),
		string(FileformatNestedYAML),
	}

	for _, possibleValue := range possibleValues {
		if possibleValue == raw {
			fileFormat := Fileformat(raw)
			return &fileFormat, nil
		}
	}

	return nil, errors.New("Fileformat is invalid")
}
