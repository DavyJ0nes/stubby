package stubby

// Route describes a configurable stub route
type Route struct {
	Path     string
	Response string
	Status   int
	Headers  map[string]string
	Queries  []string
}
