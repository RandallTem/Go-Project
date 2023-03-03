package env

type Environment struct {
	values map[string]string
}

var environment Environment

func Add(key string, value string) {
	if environment.values == nil {
		environment.values = make(map[string]string)
	}
	environment.values[key] = value
}

func Find(key string) string {
	if environment.values == nil {
		environment.values = make(map[string]string)
	}
	return environment.values[key]
}
