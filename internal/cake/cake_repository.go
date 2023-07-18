package cake

import (
	"time"
)

func (r *cakeRepository) Create(req *CreateCakeRequest) (*CreateCakeResponse, error) {
	query := "INSERT INTO cakes (title, description, rating, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, req.Title, req.Description, req.Rating, req.Image, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &CreateCakeResponse{
		ID:          int(id),
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (r *cakeRepository) GetAll() ([]*Cake, error) {
	query := "SELECT * FROM cakes WHERE deleted_at IS NULL"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cakes := []*Cake{}
	for rows.Next() {
		cake := &Cake{}
		var createdAt, updatedAt, deletedAt []uint8
		err := rows.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return nil, err
		}
		cake.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
		cake.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		cakes = append(cakes, cake)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cakes, nil
}

func (r *cakeRepository) GetByID(id int) (*Cake, error) {
	query := "SELECT * FROM cakes WHERE id = ? AND deleted_at IS NULL"
	row := r.db.QueryRow(query, id)

	cake := &Cake{}
	var createdAt, updatedAt, deletedAt []uint8
	err := row.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &createdAt, &updatedAt, &deletedAt)
	if err != nil {
		return nil, err
	}
	cake.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
	cake.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))

	return cake, nil
}

func (r *cakeRepository) Update(cake *UpdateCakeRequest) (*Cake, error) {
	query := "UPDATE cakes SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, cake.Title, cake.Description, cake.Rating, cake.Image, time.Now(), cake.ID)
	if err != nil {
		return nil, err
	}

	return r.GetByID(cake.ID)
}

func (r *cakeRepository) Delete(id int) error {
	query := "UPDATE cakes SET deleted_at = ? WHERE id = ?"

	_, err := r.db.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
