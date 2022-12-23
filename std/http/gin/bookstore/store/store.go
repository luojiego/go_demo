package store

type Book struct {
	Id      string
	Name    string
	Authors []string
	Press   string
}

type Store interface {
	Create(*Book) error
	Update(*Book) error
	Get(string) (Book, error)
	GetAll() ([]Book, error)
	Delete(string) error
}
