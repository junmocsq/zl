package access


type Accessor interface {
	SetDb(db string) Accessor
	SetTag(tag string) Accessor
	SetKey(key string) Accessor
	ClearTag(tag string) Accessor
	ClearKey(tag, key string) Accessor
	PrepareSql(sql string, params ...interface{}) Accessor
	FetchOne(interface{}) bool
	FetchAll(interface{}) bool
	InsertOne(id int)
	BatchInsert(ids int)
	Update() bool
	Delete() bool
	BeginTx() bool
	Commit() bool
	RollBack() bool
}
