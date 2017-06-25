package cascadestoreSessions

import (
	"github.com/dsoprea/go-appengine-sessioncascade"
	"github.com/gin-contrib/sessions"
	gorillaSessions "github.com/gorilla/sessions"
)

// CascadestoreStore means a wrapper of CascadeStore
type CascadestoreStore struct {
	*cascadestore.CascadeStore
}

// Options can set options
func (c *CascadestoreStore) Options(options sessions.Options) {
	c.CascadeStore.Options = &gorillaSessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}

// NewCascadestoreStore returns CascadestoreStore instances
func NewCascadestoreStore(keyPairs ...[]byte) *CascadestoreStore {
	return &CascadestoreStore{cascadestore.NewCascadeStore(cascadestore.DistributedBackends, keyPairs...)}
}
