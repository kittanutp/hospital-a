package service

import (
	"github.com/kittanutp/hospital-app/app/config"
	"github.com/kittanutp/hospital-app/app/database"
	"github.com/kittanutp/hospital-app/app/repository"
	"github.com/kittanutp/hospital-app/app/schema"
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

	hashPassword := encryptPassword(data.Password, salt)
	staff := database.Staff{
		Username:     data.Username,
		Password:     hashPassword,
		Salt:         salt,
		HospitalName: data.HospitalName,
	}
	return s.staffRepository.CreateStaff(&staff)
}

func (s *StaffService) ProcessLogIn(data schema.LogInSchema) (schema.TokenResponseSchema, error) {
	resp := s.staffRepository.GetStaff(data.Username)
	if resp.Err != nil {
		return schema.TokenResponseSchema{
			TokenType: "error",
			Token:     "error",
		}, resp.Err
	}
	submitPassword := encryptPassword(data.Password, resp.Staff.Salt)
	err := bcrypt.CompareHashAndPassword([]byte(submitPassword), []byte(resp.Staff.Password))
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
