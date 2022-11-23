package yaml

import (
	"testing"
)

func TestParse(t *testing.T) {
	data := `title: Welcome
form:
  label: 'Name:'
  placeholder: John Doe
`

	result, err := Parse([]byte(data))

	if err != nil || result == nil {
		t.Fatal("Error while parsing YAML", result, err)
	}

	value := *result

	if value["title"] != "Welcome" {
		t.Fatalf("Did not resolve title key")
	}

	if value["form.label"] != "Name:" {
		t.Fatalf("Did not resolve form.label")
	}
}
