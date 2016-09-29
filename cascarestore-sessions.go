package cascadestoreSessions

import (
	"github.com/delphinus/go-appengine-sessioncascade"
	"github.com/gin-gonic/contrib/sessions"
	gorillaSessions "github.com/gorilla/sessions"
)

// DatastoreStore is a datastore version of xxxStore
type DatastoreStore interface {
	sessions.Store
}

// NewDatastoreStore returns DatastoreStore instances
func NewDatastoreStore(secret []byte) DatastoreStore {
	return &datastoreStore{cascadestore.NewCascadeStore(cascadestore.DistributedBackends, secret)}
}

type datastoreStore struct {
	*cascadestore.CascadeStore
}

// Options can set options
func (c *datastoreStore) Options(options sessions.Options) {
	c.CascadeStore.Options = &gorillaSessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
