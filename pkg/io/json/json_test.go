package json

import (
	"testing"
)

func TestParse(t *testing.T) {
	data := `
		{
			"title": "Welcome",
			"form": {
				"label": "Name:",
				"placeholder": "John Doe"
			}
		}
	`

	result, err := Parse([]byte(data))

	if err != nil || result == nil {
		t.Fatalf("Error while parsing JSON")
	}

	value := *result

	if value["title"] != "Welcome" {
		t.Fatalf("Did not resolve title key")
	}

	if value["form.label"] != "Name:" {
		t.Fatalf("Did not resolve form.label")
	}
}
