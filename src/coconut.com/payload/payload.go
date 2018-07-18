package payload

type Payload struct {
	Title string `json:"title"`
	Note string `json:"note"`
	ManifestUrl string `json:"url"`
}

type List struct {
	Target string `json:"target"`
	Payloads []Payload `json:"payloads"`
}
