package structure

type Language map[string]string

func (l Language) Get(text string) string {
	return (l)[text]
}


