package hook

import (
	"net/http"

        "github.com/ory/kratos/selfservice/strategy/link"
	"github.com/ory/kratos/selfservice/flow/recovery"
	"github.com/ory/kratos/session"
)

var _ recovery.PostHookExecutor = new(TokenDestroyer)

type (
	tokenDestroyerDependencies interface {
		session.ManagementProvider
		session.PersistenceProvider
		link.RecoveryTokenPersistenceProvider
	}
	TokenDestroyer struct {
		r tokenDestroyerDependencies
	}
)

func NewTokenDestroyer(r tokenDestroyerDependencies) *TokenDestroyer {
	return &TokenDestroyer{r: r}
}

func (e *TokenDestroyer) ExecutePostRecoveryHook(_ http.ResponseWriter, r *http.Request, flow *recovery.Flow, s *session.Session) error {
	if _, err := e.r.RecoveryTokenPersister().UseRecoveryToken(r.Context(), flow.CSRFToken); err != nil {
		return err
	}
	return nil
}
