package service

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kittanutp/hospital-app/config"
	"github.com/kittanutp/hospital-app/repository"
)

type StaffAuthService struct {
	staffAuthRepository repository.StaffAuthRepository
	cfg                 config.Server
}

func NewStaffAuthService(repo repository.StaffAuthRepository, cfg config.Server) StaffAuthServiceInterface {
	return &StaffAuthService{
		staffAuthRepository: repo,
		cfg:                 cfg,
	}
}

func (s *StaffAuthService) ProcessStaffToken(headerToken string) repository.StaffResponse {
	const BEARER_SCHEMA = "Bearer "
	tokenString := headerToken[len(BEARER_SCHEMA):]
	token, err := validateToken(tokenString, s.cfg.OAuthKey)

	if err != nil {
		return repository.StaffResponse{
			Staff: nil,
			Err:   err,
		}
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		var id uint
		if idFloat, ok := claims["id"].(float64); ok {
			id = uint(idFloat)
		} else {
			return repository.StaffResponse{
				Staff: nil,
				Err:   fmt.Errorf("invalid id value expect float64, got=%T", idFloat),
			}
		}
		return s.staffAuthRepository.GetStaffById(id)

	} else {
		return repository.StaffResponse{
			Staff: nil,
			Err:   errors.New("invalid token"),
		}
	}

}

func validateToken(encodedToken string, oAuthKey string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("invalid token")

		}
		return []byte(oAuthKey), nil
	})

}
