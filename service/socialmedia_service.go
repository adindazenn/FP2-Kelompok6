package service

import (
	"errors"

	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/entity"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/input"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/repository"
)

type SocialMediaService interface {
	CreateSocialMedia(input input.SocialInput, idUser int) (entity.SocialMedia, error)
	DeleteSocialMedia(id_user int, id_socialmedia int) (entity.SocialMedia, error)
	UpdateSocialMedia(id_user int, id_socialmedia int, input input.SocialInput) (entity.SocialMedia, error)
	GetSocialMedia() ([]entity.SocialMedia, error)
	GetSocialMediaByID(idSocialMedia int) (entity.SocialMedia, error)
}
type socialmediaService struct {
	socialmediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(socialmediaRepository repository.SocialMediaRepository) *socialmediaService {
	return &socialmediaService{socialmediaRepository}
}

func (s *socialmediaService) CreateSocialMedia(input input.SocialInput, idUser int) (entity.SocialMedia, error) {
	newSocialMedia := entity.SocialMedia{
		Name:   input.Name,
		URL:    input.URL,
		UserID: idUser,
	}

	createdSocialmedia, err := s.socialmediaRepository.Save(newSocialMedia)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return createdSocialmedia, nil

}

func (s *socialmediaService) GetSocialMedia() ([]entity.SocialMedia, error) {
	socialmedia, err := s.socialmediaRepository.GetAll()

	if err != nil {
		return []entity.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (s *socialmediaService) DeleteSocialMedia(id_user int, id_socialmedia int) (entity.SocialMedia, error) {
	socialmedia, err := s.socialmediaRepository.FindByID(id_socialmedia)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	if socialmedia.ID == 0 {
		return entity.SocialMedia{}, errors.New("data not found")
	}

	if id_user != socialmedia.UserID {
		return entity.SocialMedia{}, errors.New("can't delete other user's social media")
	}

	socialmediaDeleted, err := s.socialmediaRepository.Delete(id_socialmedia)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return socialmediaDeleted, nil
}

func (s *socialmediaService) UpdateSocialMedia(id_user int, id_socialmedia int, input input.SocialInput) (entity.SocialMedia, error) {

	Result, err := s.socialmediaRepository.FindByID(id_socialmedia)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	if Result.ID == 0 {
		return entity.SocialMedia{}, errors.New("data not found")
	}

	if id_user != Result.UserID {
		return entity.SocialMedia{}, errors.New("can't update other user's social media")
	}

	updated := entity.SocialMedia{
		Name: input.Name,
		URL:  input.URL,
	}

	socialmediaUpdate, err := s.socialmediaRepository.Update(updated, id_socialmedia)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return socialmediaUpdate, nil
}

func (s *socialmediaService) GetSocialMediaByID(idSocialMedia int) (entity.SocialMedia, error) {
	socialmedia, err := s.socialmediaRepository.FindByID(idSocialMedia)

	if err != nil {
		return entity.SocialMedia{}, err
	}

	return socialmedia, nil
}
