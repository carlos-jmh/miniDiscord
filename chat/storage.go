package chat

// DAO (data access object) interacts with a database
type DAO struct {
	db map[string][]string
}

func NewDAO() *DAO {
	return &DAO{
		db: make(map[string][]string),
	}
}

func (d *DAO) Put(roomID string, msg string) {
	d.db[roomID] = append(d.db[roomID], msg)
}

func (d *DAO) Get(roomID string) []string {
	return d.db[roomID]
}
