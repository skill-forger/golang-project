package contract

type ProfileResponse struct {
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
