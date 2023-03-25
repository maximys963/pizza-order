package repositories

type Film struct {
	name   string
	year   int
	format string
	actors []string
}

type FilmRepository interface {
	Create(film *Film) error
	Read(id int) (*Film, error)
	Delete(id int) error
	List() ([]Film, error)
}
