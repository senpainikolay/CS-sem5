package inMemoryDB

// In Memory database

type Database struct {
	mapDB     map[string][]byte
	UserToken map[string][]byte
}

func GetDB() *Database {
	return &Database{
		make(map[string][]byte),
		make(map[string][]byte),
	}
}

func (db *Database) GetUserPassword(name string) []byte {
	return db.mapDB[name]
}

func (db *Database) RegisterUser(name string, pw []byte) {
	db.mapDB[name] = pw
}

func (db *Database) Delete(name string) {
	delete(db.mapDB, name)
}

func (db *Database) Update(name string, pw []byte) {
	db.mapDB[name] = pw
}
