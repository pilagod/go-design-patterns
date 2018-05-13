package template

import (
	"strings"
	"testing"
)

type TestStruct struct {
	Template
}

func (ts *TestStruct) Message() string {
	return "world"
}

func TestTemplateExecuteAlgorithm(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T) {
		s := &TestStruct{
			Template: &TemplateImpl{},
		}
		res := s.ExecuteAlgorithm(s)
		expected := "world"

		if !strings.Contains(res, expected) {
			t.Errorf("expected: %s, but got: %s", expected, res)
		}
	})

	t.Run("Using anonymous functions", func(t *testing.T) {
		a := new(AnonymousTemplate)
		res := a.ExecuteAlgorithm(func() string {
			return "world"
		})
		expected := "world"

		if !strings.Contains(res, expected) {
			t.Errorf("expected: %s, but got: %s", expected, res)
		}
	})

	t.Run("Using anonymous functions adapted to an interface", func(t *testing.T) {
		messageRetriever := MessageRetrieverAdapter(func() string {
			return "world"
		})
		if messageRetriever == nil {
			t.Fatal("Can not continue with a nil MessageRetriever")
		}
		template := TemplateImpl{}
		res := template.ExecuteAlgorithm(messageRetriever)
		expected := "world"

		if !strings.Contains(res, expected) {
			t.Errorf("expected: %s, but got: %s", expected, res)
		}
	})
}
