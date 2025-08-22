package models

type ProfileUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ProfileStorage struct {
	Used  int64 `json:"used"`
	Limit int   `json:"limit"`
}

type Profile struct {
	Storage *ProfileStorage `json:"storage"`
	Files   []string        `json:"files"`
	User    *ProfileUser    `json:"user"`
}
