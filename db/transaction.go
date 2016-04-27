package db

func NewTransaction(parent Transaction) Transaction {
	return &transaction{UnsetStorage{}, Storage{}, parent}
}

type Transaction interface {
	Get(string) (string, bool)
	Set(string, string)
	Unset(string)
	Rollback() (Transaction, error)
	Commit() Transaction
	KeysWithValue(string) map[string]bool
}

type transaction struct {
	unsets  UnsetStorage
	storage Storage
	parent  Transaction
}

func (t *transaction) Set(key, val string) {
	t.storage[key] = val
	delete(t.unsets, key)
}

func (t *transaction) Get(key string) (string, bool) {

	if unseted, ok := t.unsets[key]; ok && unseted {
		return "", false
	}

	if val, ok := t.storage[key]; ok {
		return val, true
	}

	return t.parent.Get(key)
}

func (t *transaction) Unset(key string) {
	delete(t.storage, key)
	t.unsets[key] = true
}

func (t *transaction) Rollback() (Transaction, error) {
	return t.parent, nil
}

func (t *transaction) Commit() Transaction {

	for k, _ := range t.unsets {
		t.parent.Unset(k)
	}

	for k, v := range t.storage {
		t.parent.Set(k, v)
	}

	return t.parent.Commit()
}

func (t *transaction) KeysWithValue(val string) map[string]bool {
	keys := t.parent.KeysWithValue(val)

	for k, v := range t.storage {
		keys[k] = v == val
	}

	for k, _ := range t.unsets {
		keys[k] = false
	}

	return keys
}
