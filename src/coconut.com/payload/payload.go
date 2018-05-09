package payload

type Payload struct {
	Title string `json:"title"`
	Url string `json:"url"`
}

type PayloadList struct {
	Project string `json:"project"`
	Apps []Payload `json:"payloads"`
}

var (
	Payloads []Payload
)

func init() {
	Payloads = make([]Payload, 1000, 1000)
}