package entities

type SuccessResponse struct {
	Status bool   `json:"status"`
	Data   *Movie `json:"data,omitempty"`
}

type MovieList struct {
	Movies []Movie `json:"movies"`
}

type Movie struct {
	Id       string    `json:"id,omitempty"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
