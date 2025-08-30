package usecase

import (
	"os"
	"time"

	"github.com/MingPV/EventService/internal/entities"
	"github.com/MingPV/EventService/internal/user/repository"
	"github.com/MingPV/EventService/pkg/apperror"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// EventService struct
type EventService struct {
	repo repository.UserRepository
}

// Init EventService
func NewEventService(repo repository.UserRepository) UserUseCase {
	return &EventService{repo: repo}
}

// EventService Methods - 1 Register user (hash password)
func (s *EventService) Register(user *entities.User) error {
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser != nil {
		return apperror.ErrAlreadyExists
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPwd)

	return s.repo.Save(user)
}

// EventService Methods - 2 Login user (check email + password)
func (s *EventService) Login(email string, password string) (string, *entities.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user == nil {
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, err
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}

// EventService Methods - 3 Get user by id
func (s *EventService) FindUserByID(id string) (*entities.User, error) {
	return s.repo.FindByID(id)
}

// EventService Methods - 4 Get all users
func (s *EventService) FindAllUsers() ([]*entities.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// EventService Methods - 5 Get user by email
func (s *EventService) GetUserByEmail(email string) (*entities.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// EventService Methods - 6 Patch
func (s *EventService) PatchUser(id string, user *entities.User) (*entities.User, error) {
	if err := s.repo.Patch(id, user); err != nil {
		return nil, err
	}
	updatedUser, _ := s.repo.FindByID(id)

	return updatedUser, nil
}

// EventService Methods - 7 Delete
func (s *EventService) DeleteUser(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
