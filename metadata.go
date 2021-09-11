package multidag

import "encoding/json"

func (m Metadata) MarshalJSON() ([]byte, error) {
	metadata := struct {
		URI         string `json:"uri"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		URI:         m.URI.String(),
		Title:       m.Title,
		Description: m.Description,
	}
	return json.Marshal(metadata)
}
