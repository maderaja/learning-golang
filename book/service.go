package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest CreateBookRequest) (Book, error)
	Update(ID int, bookRequest UpdateBookRequest) (Book, error)
	Delete(ID int) (Book, error)
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
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest CreateBookRequest) (Book, error) {
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

func (s *service) Update(ID int, bookRequest UpdateBookRequest) (Book, error) {
	book, _ := s.repository.FindByID(ID)
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = int(price)
	book.Rating = int(rating)
	book.Discount = int(discount)

	updatedBook, err := s.repository.Update(book)
	return updatedBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.repository.FindByID(ID)
	deletedBook, err := s.repository.Delete(book)

	return deletedBook, err
}
