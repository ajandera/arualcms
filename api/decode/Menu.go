package decode

type Menu struct {
	Name     string
	Url      string
	ParentId string
	SiteId   string
	ID       int    `json:"ID"`
	Id       string `json:"Id"`
}
