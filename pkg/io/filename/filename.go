package filename

import (
	"bytes"
	"errors"
	templating "html/template"
)

type Variables struct {
	LanguageLocale  string
	LanguageName    string
	FormatExtension string
}

func getNoopVariables() Variables {
	return Variables{
		LanguageLocale:  "en",
		LanguageName:    "English",
		FormatExtension: "json",
	}
}

func newTemplateFromString(value string) (*templating.Template, error) {
	template := templating.New("filename")
	template = template.Option("missingkey=error")

	return template.Parse(value)
}

func IsValid(value string) bool {
	template, err := newTemplateFromString(value)

	if err != nil {
		return false
	}

	var result bytes.Buffer
	err = template.Execute(&result, getNoopVariables())

	if err != nil {
		return false
	}

	return true
}

func Render(value string, variables Variables) (string, error) {
	template, err := newTemplateFromString(value)

	if err != nil {
		return "", errors.New("File name is invalid")
	}

	var result bytes.Buffer

	err = template.Execute(&result, variables)

	if err != nil {
		return "", err
	}

	return result.String(), nil
}
