package stubby

// Route describes a configurable stub route
type Route struct {
	Methods  []string
	Path     string
	Response string
	Status   int
}
