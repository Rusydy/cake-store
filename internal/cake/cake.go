package cake

import (
	"database/sql"
	"time"
)

type Cake struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Rating      int        `json:"rating"`
	Image       string     `json:"image"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type CreateCakeRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Rating      int    `json:"rating" validate:"gte=1,lte=5"`
	Image       string `json:"image" validate:"required"`
}

type CreateCakeResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateCakeRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Rating      int    `json:"rating" validate:"gte=1,lte=5"`
	Image       string `json:"image" validate:"required"`
}

type GetCakeByIDRequest struct {
	ID int `json:"id" validate:"required"`
}

type CakeRepository interface {
	Create(req *CreateCakeRequest) (*CreateCakeResponse, error)
	GetAll() ([]*Cake, error)
	GetByID(id int) (*Cake, error)
	Update(req *UpdateCakeRequest) (*Cake, error)
	Delete(id int) error
}

type CakeService interface {
	CreateCake(req *CreateCakeRequest) (*CreateCakeResponse, error)
	GetAllCakes() ([]*Cake, error)
	GetCakeByID(int) (*Cake, error)
	UpdateCake(req *UpdateCakeRequest) (*Cake, error)
	DeleteCake(id int) error
}

type cakeRepository struct {
	db *sql.DB
}

type cakeService struct {
	repo CakeRepository
}

func NewCakeRepository(db *sql.DB) CakeRepository {
	return &cakeRepository{
		db: db,
	}
}

func NewCakeService(repo CakeRepository) CakeService {
	return &cakeService{
		repo: repo,
	}
}
