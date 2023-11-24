package service

import (
	"errors"

	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/entity"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/input"
	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/repository"
)

type commentService struct {
	commentRepository repository.CommentRepository
	photoRepository   repository.PhotoRepository
}

type CommentService interface {
	GetCommentAll() ([]entity.Comment, error)
	CreateComment(input input.CommentInput, idUser int) (entity.Comment, error)
	DeleteComment(id_user int, id_comment int) (entity.Comment, error)
	UpdateComment(id_user int, id_comment int, input input.CommentUpdateInput) (entity.Comment, error)
	GetCommentByID(commentID int) (entity.Comment, error)
	GetCommentsByPhotoID(photoID int) ([]entity.Comment, error)
}

func NewCommentService(commentRepository repository.CommentRepository, photoRepository repository.PhotoRepository) *commentService {
	return &commentService{commentRepository, photoRepository}
}

func (s *commentService) CreateComment(input input.CommentInput, idUser int) (entity.Comment, error) {
	photoData, err := s.photoRepository.FindByID(input.PhotoID)

	if err != nil {
		return entity.Comment{}, err
	}
	if photoData.ID == 0 {
		return entity.Comment{}, errors.New("photo not found")
	}

	newComment := entity.Comment{
		Message: input.Message,
		PhotoID: input.PhotoID,
		UserID:  idUser,
	}

	createNewcomment, err := s.commentRepository.Save(newComment)

	if err != nil {
		return entity.Comment{}, err
	}

	return createNewcomment, nil
}

func (s *commentService) GetCommentAll() ([]entity.Comment, error) {
	comment, err := s.commentRepository.GetAll()

	if err != nil {
		return []entity.Comment{}, err
	}

	return comment, nil
}

func (s *commentService) DeleteComment(id_user int, id_comment int) (entity.Comment, error) {
	comment, err := s.commentRepository.FindByID(id_comment)

	if err != nil {
		return entity.Comment{}, err
	}

	if comment.ID == 0 {
		return entity.Comment{}, errors.New("comment not found")
	}

	if id_user != comment.UserID {
		return entity.Comment{}, errors.New("can't delete other user's comment")
	}

	Deleted, err := s.commentRepository.Delete(id_comment)

	if err != nil {
		return entity.Comment{}, err
	}

	return Deleted, nil
}

func (s *commentService) UpdateComment(id_user int, id_comment int, input input.CommentUpdateInput) (entity.Comment, error) {

	Result, err := s.commentRepository.FindByID(id_comment)

	if err != nil {
		return entity.Comment{}, err
	}

	if Result.ID == 0 {
		return entity.Comment{}, errors.New("comment not found")
	}

	if id_user != Result.UserID {
		return entity.Comment{}, errors.New("can't update other user's comment")
	}

	updated := entity.Comment{
		Message: input.Message,
	}

	commentUpdate, err := s.commentRepository.Update(updated, id_comment)

	if err != nil {
		return entity.Comment{}, err
	}

	return commentUpdate, nil
}

func (s *commentService) GetCommentByID(commentID int) (entity.Comment, error) {
	comment, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (s *commentService) GetCommentsByPhotoID(photoID int) ([]entity.Comment, error) {
	comments, err := s.commentRepository.FindByPhotoID(photoID)

	if err != nil {
		return []entity.Comment{}, err
	}

	return comments, nil
}
