package models

type Profile struct {
	User    *ProfileUser    `json:"user"`
	Storage *ProfileStorage `json:"storage"`
}

type ProfileUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ProfileStorage struct {
	Used  int64 `json:"used"`
	Limit int   `json:"limit"`
}

type ProfileUser1 struct {
	Name string `json:"name"`
}

type ProfileStorage1 struct {
	Storage *ProfileStorage `json:"storage"`
	Files   []string        `json:"files"`
	User    *ProfileUser1   `json:"user"`
}
