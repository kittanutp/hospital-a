package service

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kittanutp/hospital-app/config"
	"github.com/kittanutp/hospital-app/database"
	"github.com/kittanutp/hospital-app/repository"
	"github.com/kittanutp/hospital-app/schema"
	"golang.org/x/crypto/bcrypt"
)

type StaffService struct {
	staffRepository repository.StaffRepository
	cfg             config.Server
}

func NewStaffService(repo repository.StaffRepository, cfg config.Server) StaffServiceInterface {
	return &StaffService{
		staffRepository: repo,
		cfg:             cfg,
	}
}

func (s *StaffService) ProcessNewStaff(data schema.CreateStaffSchema) repository.StaffResponse {
	salt, err := gensalt()
	if err != nil {
		return repository.StaffResponse{
			Staff: nil,
			Err:   err,
		}
	}

	hashPassword, err := EncryptPassword(data.Password, salt)
	if err != nil {
		return repository.StaffResponse{
			Staff: nil,
			Err:   err,
		}
	}
	staff := database.Staff{
		Username:     data.Username,
		Password:     hashPassword,
		Salt:         salt,
		HospitalName: data.HospitalName,
	}
	return s.staffRepository.CreateStaff(&staff)
}

func (s *StaffService) ProcessLogIn(data schema.LogInSchema) (schema.TokenResponseSchema, error) {
	resp := s.staffRepository.GetStaffByUsername(data.Username)
	if resp.Err != nil {
		return schema.TokenResponseSchema{
			TokenType: "error",
			Token:     "error",
		}, resp.Err
	}

	submitPassword := data.Password + resp.Staff.Salt
	err := bcrypt.CompareHashAndPassword([]byte(resp.Staff.Password), []byte(submitPassword))
	if err != nil {
		return schema.TokenResponseSchema{
			TokenType: "error",
			Token:     "error",
		}, err
	}
	return schema.TokenResponseSchema{
		TokenType: "Bearer",
		Token:     generateToken(resp.Staff.ID, s.cfg.OAuthKey),
	}, nil

}

func gensalt() (string, error) {
	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(salt), nil
}

func EncryptPassword(password string, salt string) (string, error) {
	pwdSalt := password + salt
	hash, err := bcrypt.GenerateFromPassword([]byte(pwdSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

type authCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

func generateToken(id uint, oAuthKey string) string {
	claims := &authCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(oAuthKey))
	if err != nil {
		panic(err)
	}
	return t
}
