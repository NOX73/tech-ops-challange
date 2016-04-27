package db

import "errors"

var (
	NoTransactionError = errors.New("NO TRANSACTION")
)

func NewBaseTransaction() Transaction {
	return &baseTransaction{Storage{}}
}

type baseTransaction struct {
	storage Storage
}

func (t *baseTransaction) Set(key, val string) {
	t.storage[key] = val
}

func (t *baseTransaction) Get(key string) (string, bool) {
	val, ok := t.storage[key]
	return val, ok
}

func (t *baseTransaction) Unset(key string) {
	delete(t.storage, key)
}

func (t *baseTransaction) Rollback() (Transaction, error) {
	return nil, NoTransactionError
}

func (t *baseTransaction) Commit() Transaction {
	return t
}

func (t *baseTransaction) KeysWithValue(val string) map[string]bool {
	var keys = map[string]bool{}

	for k, v := range t.storage {
		if v == val {
			keys[k] = true
		}
	}

	return keys
}
