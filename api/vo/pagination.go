package vo

type Pagination struct {
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
}

func (p *Pagination) CheckOrResetBounds() {
	if p.Limit >= 100 {
		p.Limit = 20
	}
}
