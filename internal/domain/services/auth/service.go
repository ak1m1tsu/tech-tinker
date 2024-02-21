package auth

import (
	"context"
	"github.com/patrickmn/go-cache"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/insan1a/tech-tinker/internal/domain/interfaces"
	"github.com/insan1a/tech-tinker/internal/domain/model"
	"github.com/insan1a/tech-tinker/internal/lib/jwt"
	"github.com/pkg/errors"
)

var _ interfaces.AuthService = &Service{}

type Service struct {
	cfg     *Config
	emprepo interfaces.EmployeeRepo
	tokens  *cache.Cache
	cache   *expirable.LRU[string, *model.Employee]
}

func New(cfg *Config, repo interfaces.EmployeeRepo) *Service {
	return &Service{
		cfg:     cfg,
		emprepo: repo,
		tokens:  cache.New(cache.NoExpiration, cache.NoExpiration),
		cache:   expirable.NewLRU[string, *model.Employee](cfg.Cache.Size, nil, cfg.Cache.TTL),
	}
}

// CreateToken implements interfaces.AuthService.
func (s *Service) CreateToken(_ context.Context, e *model.Employee) (string, error) {
	item, ok := s.tokens.Get(e.Email)
	if ok {
		if token, ok := item.(string); ok {
			return token, nil
		}
	}

	token, err := jwt.GenerateToken(&jwt.Employee{
		ID:   e.ID,
		Role: e.Role.String(),
	}, s.cfg.JWT.TTL, s.cfg.JWT.PrivateKey)
	if err != nil {
		return "", errors.WithMessagef(err, "failed to generate token for user %s", e.Email)
	}

	s.tokens.Set(e.Email, token, s.cfg.JWT.TTL)

	return token, nil
}

// GetByEmail implements interfaces.AuthService.
func (s *Service) GetByEmail(ctx context.Context, email string) (*model.Employee, error) {
	employee, ok := s.cache.Get(email)
	if ok {
		return employee, nil
	}

	employee, err := s.emprepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to find user %s", email)
	}

	s.cache.Add(email, employee)

	return employee, nil
}
