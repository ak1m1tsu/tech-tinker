package auth

import (
	"context"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/insan1a/tech-tinker/internal/domain/interfaces"
	"github.com/insan1a/tech-tinker/internal/domain/model"
)

var _ interfaces.AuthService = &Service{}

type Service struct {
	cfg     *Config
	emprepo interfaces.EmployeeRepo
	cache   *expirable.LRU[string, *model.Employee]
}

func New(cfg *Config, repo interfaces.EmployeeRepo) *Service {
	return &Service{
		cfg:     cfg,
		emprepo: repo,
		cache:   expirable.NewLRU[string, *model.Employee](cfg.Cache.Size, nil, cfg.Cache.TTL),
	}
}

// CreateToken implements interfaces.AuthService.
func (s *Service) CreateToken(ctx context.Context, e *model.Employee) (string, error) {
	panic("unimplemented")
}

// GetByEmail implements interfaces.AuthService.
func (s *Service) GetByEmail(ctx context.Context, email string) (*model.Employee, error) {
	employee, ok := s.cache.Get(email)
	if ok {
		return employee, nil
	}

	employee, err := s.emprepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	s.cache.Add(email, employee)

	return employee, nil
}
