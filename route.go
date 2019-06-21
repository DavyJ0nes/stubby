package stubby

// Route describes a configurable stub route
type Route struct {
	Path     string
	Response string
	Status   int
	Queries  []string
}
