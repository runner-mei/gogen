package gengen

import (
	"reflect"
	"testing"
)

func TestAnnotations(t *testing.T) {
	for _, test := range []struct {
		txt   string
		value Annotation
	}{
		{txt: `// @http.GET(path="/concat")`, value: Annotation{Name: "http.GET", Attributes: map[string]string{"path": "/concat"}}},
	} {
		t.Run(test.txt, func(t *testing.T) {
			a := parseAnnotation(test.txt)
			if !reflect.DeepEqual(*a, test.value) {
				t.Error(test.txt)
				t.Error(a)
			}
		})
	}
}

func TestUnderscore(t *testing.T) {
	for key, value := range map[string]string{
		"samplingNode": "sampling_node",
	} {
		if Underscore(key) != value {
			t.Error("want", Underscore(key), "got", value)
		}
	}
}
