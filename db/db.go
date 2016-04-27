package db

type Storage map[string]string
type UnsetStorage map[string]bool

func New() *DB {
	return &DB{NewBaseTransaction()}
}

type DB struct {
	currentTransaction Transaction
}

func (d *DB) Get(key string) (string, bool) {
	return d.currentTransaction.Get(key)
}

func (d *DB) Set(key, val string) {
	d.currentTransaction.Set(key, val)
}

func (d *DB) Unset(key string) {
	d.currentTransaction.Unset(key)
}

func (d *DB) NumEqualTo(val string) int {
	keys := d.currentTransaction.KeysWithValue(val)

	var count int
	for _, v := range keys {
		if v {
			count++
		}
	}

	return count
}

func (d *DB) Begin() {
	d.currentTransaction = NewTransaction(d.currentTransaction)
}

func (d *DB) Rollback() error {
	t, err := d.currentTransaction.Rollback()

	if err == nil {
		d.currentTransaction = t
	}

	return err
}

func (d *DB) Commit() {
	d.currentTransaction = d.currentTransaction.Commit()
}
