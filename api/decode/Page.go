package decode

type Page struct {
	Key         string
	Body        string
	Title       string
	MetaTitle   string
	Keywords    string
	Description string
	SiteId      string
	ID          int    `json:"ID"`
	Id          string `json:"Id"`
}
