package models

type File struct {
	ID           int    `db:"id" json:"id"`
	UserID       int    `db:"user_id" json:"user_id"`
	Filename     string `db:"filename" json:"filename"`
	OriginalName string `db:"original_name" json:"original_name"`
	Size         int64  `db:"size" json:"size"`
	MimeType     string `db:"mime_type" json:"mime_type"`
	CreatedAt    string `db:"created_at" json:"created_at"`
}

type FileUpload struct {
	Message  string `json:"message"`
	FileName string `json:"filename"`
}

type ListFile struct {
	Files   []string      `json:"files"`
	Storage *ListFileUsed `json:"storage"`
}

type ListFileUsed struct {
	Used  int64 `json:"used"`
	Limit int   `json:"limit"`
}
