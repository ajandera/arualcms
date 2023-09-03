package decode

type Menu struct {
	Name     string
	Url      string
	ParentId string `json:"ParentId"`
	SiteId   string
	ID       int    `json:"ID"`
	Id       string `json:"Id"`
}
