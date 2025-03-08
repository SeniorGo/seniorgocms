package api

type HttpError struct {
	Status      int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h HttpError) Error() string {
	return h.Description
}
