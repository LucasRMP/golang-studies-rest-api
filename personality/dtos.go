package personality

type Personality struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

type CreatePersonalityDTO struct {
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

type UpdatePersonalityDTO struct {
	Name      string `json:"name"`
	Biography string `json:"biography"`
}
