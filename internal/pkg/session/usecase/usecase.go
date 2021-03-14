package usecase

import (
	"github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/models"
	"github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/session"
	"github.com/go-park-mail-ru/2021_1_DuckLuck/internal/server/errors"
)

type SessionUseCase struct {
	SessionRepo session.Repository
}

func NewUseCase(SessionRepo session.Repository) session.UseCase {
	return &SessionUseCase{
		SessionRepo: SessionRepo,
	}
}

func (h *SessionUseCase) Check(sessionCookieValue string) (*models.Session, error) {
	sess, err := h.SessionRepo.GetByValue(sessionCookieValue)
	if err == errors.ErrSessionNotFound {
		return nil, errors.ErrUserUnauthorized
	}

	return sess, nil
}

func (h *SessionUseCase) Create(userId uint64) (*models.Session, error) {
	sess := models.NewSession(userId)
	err := h.SessionRepo.Add(sess)
	if err != nil {
		return nil, errors.ErrInternalError
	}

	return sess, nil
}

func (h *SessionUseCase) DestroyCurrent(sessionCookieValue string) error {
	err := h.SessionRepo.DestroyByValue(sessionCookieValue)
	if err != nil {
		return errors.ErrInternalError
	}

	return nil
}
