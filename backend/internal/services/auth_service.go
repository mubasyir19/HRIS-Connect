package services

import (
	"backend/internal/dto/requests"
	"backend/internal/dto/response"
	"backend/internal/models"
	"backend/internal/repositories"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	GenerateToken(user *models.Employee) (string, error)
	GenerateRefreshToken(user *models.Employee) (string, error)
	// Login -> gunakan cookies jika bisa
	Login(req requests.LoginRequest) (*response.LoginResponse, error)
	// Register -> sudah di handle di employee service (transaction tambah employee & user)
	Update(id uuid.UUID, req requests.UpdateUserRequest) (*models.User, error)
	// GetProfile
	GetProfile(email string) (*response.DataUserResponse, error)
	ValidateRefreshToken(token string) (jwt.MapClaims, error)
	// Logout
	Logout(id uuid.UUID) error
}

type authService struct {
	employeeRepository repositories.EmployeeRepository
	userRepository     repositories.UserRepository
	jwtSecret          string
}

func NewAuthService(employeeRepository repositories.EmployeeRepository, userRepository repositories.UserRepository, jwtSecret string) AuthService {
	return &authService{employeeRepository, userRepository, jwtSecret}
}

func (s *authService) GenerateToken(user *models.Employee) (string, error) {
	if s.jwtSecret == "" {
		return "", errors.New("missing secret key")
	}

	role := ""
	if user.User != nil {
		role = string(user.User.Role)
	}

	claims := jwt.MapClaims{
		"userId":     user.ID.String(),
		"email":      user.Email,
		"employeeId": user.ID.String(),
		"role":       role,
		"type":       "access",
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *authService) GenerateRefreshToken(user *models.Employee) (string, error) {
	if s.jwtSecret == "" {
		return "", errors.New("missing secret key")
	}

	role := ""
	if user.User != nil {
		role = string(user.User.Role)
	}

	claims := jwt.MapClaims{
		"userId":     user.ID.String(),
		"email":      user.Email,
		"employeeId": user.ID.String(),
		"role":       role,
		"type":       "refresh",
		"exp":        time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *authService) Login(req requests.LoginRequest) (*response.LoginResponse, error) {
	// cek field kosong atau tidak
	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	// cek user ada atau tidak, cek dengan email
	user, err := s.employeeRepository.GetByEmail(req.Email)
	// return errors.New() jika tidak ada
	if err != nil {
		return nil, errors.New("user not found")
	}

	// cek compare request password dengan password yang tersimpan di database, pakai bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.User.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("password doesn't match")
	}

	// generate access token
	accessToken, err := s.GenerateToken(user)
	if err != nil {
		return nil, fmt.Errorf("generate access token: %w", err)
	}

	// generate refresh token
	refreshToken, err := s.GenerateRefreshToken(user)
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	if user.User == nil {
		return nil, errors.New("user data is incomplete")
	}

	// update refresh token in database
	if err := s.userRepository.UpdateRefreshToken(user.User.ID.String(), refreshToken); err != nil {
		return nil, fmt.Errorf("update refresh token: %w", err)
	}

	user.User.Password = ""
	return &response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Data:         *response.ToAuthResponse(user.User, user),
	}, nil
}

func (s *authService) GetProfile(email string) (*response.DataUserResponse, error) {
	// cek user ada atau tidak, cek dengan email
	user, err := s.employeeRepository.GetByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.User == nil {
		return nil, errors.New("user data is incomplete")
	}

	user.User.Password = ""
	return response.ToAuthResponse(user.User, user), nil
}

func (s *authService) ValidateRefreshToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Verifikasi tipe token
	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		return nil, errors.New("invalid token type")
	}

	// Verifikasi belum expired
	exp, ok := claims["exp"].(float64)
	if ok && int64(exp) < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func (s *authService) Update(id uuid.UUID, req requests.UpdateUserRequest) (*models.User, error) {
	// find user from employee repo
	existingEmployee, err := s.employeeRepository.GetByID(id)
	if err != nil {
		return nil, errors.New("failed get user")
	}

	if existingEmployee.User == nil {
		return nil, errors.New("user not found")
	}

	// get object user from existing employee
	user := existingEmployee.User

	// validate request update
	if req.EmployeeID != nil {
		user.EmployeeID = *req.EmployeeID
	}

	if req.Email != nil {
		user.Email = *req.Email
	}

	if req.Password != nil && *req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("hash password: %w", err)
		}
		user.Password = string(hashedPassword)
	}

	if req.Role != "" {
		user.Role = req.Role
	}

	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if req.LastLogin != nil {
		user.LastLogin = req.LastLogin
	}

	if req.RefreshToken != nil {
		user.RefreshToken = *req.RefreshToken
	}

	// update user
	if err := s.userRepository.UpdateUser(user); err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}

	// hide password response
	user.Password = ""

	// return
	return user, nil
}

func (s *authService) Logout(id uuid.UUID) error {
	existingEmployee, err := s.employeeRepository.GetByID(id)
	if err != nil {
		return errors.New("failed get user")
	}

	if existingEmployee.User == nil {
		return errors.New("user not found")
	}

	if err := s.userRepository.UpdateRefreshToken(existingEmployee.User.ID.String(), ""); err != nil {
		return fmt.Errorf("logout: %w", err)
	}

	return nil
}
