package blogs

import "github.com/labstack/echo/v4"

// BlogRepository is
type BlogRepository interface {
	Fetch() (*[]Blog, error)
	Get(string) (*Blog, error)

	Create(*Blog) error
	Update(*Blog, string) error
	Delete(string) error
}

// Usecase is
type Usecase interface {
	Fetch() (*[]Blog, error)
	Get(string) (*Blog, error)

	Create(*Blog) error
	Update(*Blog, string) error
	Delete(string) error
}

// BlogDelivery is
type BlogDelivery interface {
	Fetch(echo.Context) error
	Get(echo.Context) error

	Create(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}