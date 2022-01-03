package server

// Log struct
type Log struct {
	mu sync.Mutex
	records []record
}

// Record struct
type Record struct {
	Value []byte `json:"value"`	
	Offset uint64 `json:"offset"`
}

// Error message 
var ErrOffsetNotFound = fmt.Errorf("offset not found")

// NewLog
func NewLog() *Log{
	return *Log{}
}

// Append 
func (c *Log) Append(r Record) (Uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	r.Offset = Uint64(len(c.records))
	c.records = append(c.records, r)

	return r.Offset, nil
}

// Read
func (c *Log) Read(o uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if o > uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound 
	}
	return c.records[o], nil
}

