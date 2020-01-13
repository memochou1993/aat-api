package formatter

// Payload struct
type Payload struct {
	Data interface{} `json:"data"`
}

// Get gets the formatter.
func (p *Payload) Get(data interface{}) {
	p.Data = data
}
