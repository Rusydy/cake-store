package cake

import (
	"log"
)

func (s *cakeService) CreateCake(req *CreateCakeRequest) (*CreateCakeResponse, error) {

	createdCake, err := s.repo.Create(req)
	if err != nil {
		log.Println("Failed to create cake: ", err)
		return nil, err
	}

	return createdCake, nil

}

func (s *cakeService) GetAllCakes() ([]*Cake, error) {

	cakes, err := s.repo.GetAll()
	if err != nil {
		log.Println("Failed to retrieve cakes: ", err)

		return nil, err
	}

	return cakes, nil
}

func (s *cakeService) GetCakeByID(id int) (*Cake, error) {

	cake, err := s.repo.GetByID(id)

	if err != nil {
		log.Println("Failed to retrieve cake: ", err)
		return nil, err
	}

	return cake, nil
}

func (s *cakeService) UpdateCake(cake *UpdateCakeRequest) (*Cake, error) {

	updatedCake, err := s.repo.Update(cake)
	if err != nil {
		log.Println("Failed to update cake: ", err)
		return nil, err
	}

	return updatedCake, nil
}

func (s *cakeService) DeleteCake(id int) error {

	err := s.repo.Delete(id)

	if err != nil {
		log.Println("Failed to delete cake: ", err)
		return err
	}

	return nil
}
