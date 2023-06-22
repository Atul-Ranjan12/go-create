package forms

type errors map[string][]string

func (e errors) Add(key string, message string) {
	e[key] = append(e[key], message)
}

func (e errors) Get(key string) string {
	if len(e[key]) == 0 {
		return ""
	} else {
		return e[key][0]
	}
}
