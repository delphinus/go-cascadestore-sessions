package cascadestoreSessions

import (
	"github.com/dsoprea/go-appengine-sessioncascade"
	"github.com/gin-gonic/contrib/sessions"
	gorillaSessions "github.com/gorilla/sessions"
)

// CascadestoreStore is a datastore version of xxxStore
type CascadestoreStore interface {
	sessions.Store
}

// NewCascadestoreStore returns CascadestoreStore instances
func NewCascadestoreStore(keyPairs ...[]byte) CascadestoreStore {
	return &cascadestoreStore{cascadestore.NewCascadeStore(cascadestore.DistributedBackends, keyPairs...)}
}

type cascadestoreStore struct {
	*cascadestore.CascadeStore
}

// Options can set options
func (c *cascadestoreStore) Options(options sessions.Options) {
	c.CascadeStore.Options = &gorillaSessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
