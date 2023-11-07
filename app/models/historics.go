package models

type Historics struct {
	ID       string `json:"id,omitempty"`
	UserID   string `json:"user_id"`
	MovieID  string `json:"movie_id"`
	Duration int    `json:"duration"`
}

func (h Historics) TableName() string {
	return "historics"
}

func (h Historics) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":       h.ID,
		"user_id":  h.UserID,
		"movie_id": h.MovieID,
	}
}

func (h Historics) Validate() bool {
	return h.UserID != "" && h.MovieID != "" && h.Duration >= 0
}

func (h Historics) RequiredFieldsString() string {
	return "[user_id, movie_id, duration]"
}

func (h Historics) ToRichMap(user User, movie Movie) map[string]interface{} {
	return map[string]interface{}{
		"id":    h.ID,
		"user":  user.ToMap(),
		"movie": movie.ToMap(),
	}
}
