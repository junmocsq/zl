package access

type dao struct {
	tag    string
	key    string
	db     string
	sql    string
	params []interface{}
	isKey  bool
}

func (d *dao) SetDb(db string) Accessor {
	d.db = db
	return d
}

func (d *dao) SetTag(tag string) Accessor {
	d.tag = tag
	return d
}

func (d *dao) SetKey(key string) Accessor {
	d.key = key
	d.isKey = true
	return d
}

func (d *dao) ClearTag(tag string) Accessor {
	d.SetTag(tag)
	return d
}

func (d *dao) ClearKey(tag, key string) Accessor {
	d.SetTag(tag)
	d.SetKey(key)
	return d
}

func (d *dao) PrepareSql(sql string, params ...interface{}) Accessor {
	return d
}

func (d *dao) FetchOne(interface{}) bool {
	return false
}

func (d *dao) InsertOne(id int) {
	return
}
func (d *dao) FetchAll(interface{}) bool {
	return false
}

func (d *dao) BatchInsert(ids int) {
	return
}

func (d *dao) Update() bool {
	return false
}

func (d *dao) Delete() bool {
	return false
}

func (d *dao) BeginTx() bool {
	return false
}

func (d *dao) Commit() bool {
	return false
}

func (d *dao) RollBack() bool {
	return false
}

func (d *dao) clear() error {
	d.tag = ""
	d.key = ""
	d.db = ""
	d.sql = ""
	d.params = nil
	d.isKey = false
	return nil
}
