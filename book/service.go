package book

type Service interface {
	FindAll() ([]Book, error)

	FindById(ID int) (Book, error)

	Create(bookRequest BookRequest) (Book, error)

	DeleteById(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
	// return s.repository.FindById(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(price),
		Rating:      int(rating),
		Discount:    int(discount),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) DeleteById(ID int) (Book, error) {
	book := Book{
		ID: ID,
	}
	newBook, err := s.repository.DeleteById(book, ID)
	return newBook, err
}
