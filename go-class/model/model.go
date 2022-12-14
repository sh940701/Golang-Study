package model

type model struct {}

func NewModel() *model {
	mod := &model{}
	return mod
}

func (m *model) Run(s string) string {
	return "user " + s
}

func (m *model) Jump(s string) string {
	return "user " + s
}

func (m *model) Sleep(s string) string {
	return "user " + s
}

func (m *model) Walk(s string) string {
	return "user " + s
}

func (m *model) Fly(s string) string {
	return "user " + s
}