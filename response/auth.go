package response

import "github.com/jihanlugas/calendar/model"

type Init struct {
	User model.UserView `json:"user,omitempty"`
}
