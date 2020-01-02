package dao

type Competition struct {
	Name  string `json:"name"`
	Date  string `json:"eventDate"`
	City  string `json:"city"`
	State string `json:"state"`
}
