package template

import "strings"

type MessageRetriever interface {
	Message() string
}

type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

/* TemplateImpl */

type TemplateImpl struct{}

func (ti *TemplateImpl) first() string {
	return "hello"
}

func (tl *TemplateImpl) third() string {
	return "template"
}

func (tl *TemplateImpl) ExecuteAlgorithm(mr MessageRetriever) string {
	return strings.Join([]string{tl.first(), mr.Message(), tl.third()}, " ")
}

/* AnonymousTemplate */

type AnonymousTemplate struct{}

func (at *AnonymousTemplate) first() string {
	return "hello"
}

func (at *AnonymousTemplate) third() string {
	return "template"
}

func (at *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{at.first(), f(), at.third()}, " ")
}

/* TemplateAdapter */

type TemplateAdapter struct {
	myFunc func() string
}

func (ta *TemplateAdapter) Message() string {
	return ta.myFunc()
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &TemplateAdapter{
		myFunc: f,
	}
}
