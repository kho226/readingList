// table provides the client table interface.
package table

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

type ClientTable struct {
	/*
		ClientTable represents a client table database.
	*/
	lastRecords map[string]interface{}
	clientTable *cache.Cache
}

func New(defaultExpiration, cleanupInterval time.Duration) *ClientTable {
	/*
		New creates a new client table
	*/
	return &ClientTable{
		lastRecords: make(map[string]interface{}),
		clientTable: cache.New(defaultExpiration, cleanupInterval),
	}
}

// Set sets a value for a key.
func (t *ClientTable) Set(k string, x interface{}) {
	/*
		Set a value for a key
	*/
	res, ok := t.Get(k)
	if ok {
		t.lastRecords[k] = res
	}

	t.clientTable.Set(k, x, cache.NoExpiration)
}

func (t *ClientTable) Undo(k string) {
	/*
		Undo the last record for a key
	*/
	t.clientTable.Set(k, t.lastRecords[k], cache.NoExpiration)
}

func (t *ClientTable) Get(k string) (interface{}, bool) {
	/*
		Get returns the record for a key.
	*/
	return t.clientTable.Get(k)
}
