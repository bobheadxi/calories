package facebook

// UserProfile : Response from Facebook after a request for a user's profile
type UserProfile struct {
	FirstName string `json:"first_name"`
	Timezone  int    `json:"timezone,omitempty"`
	Gender    string `json:"gender,omitempty"`
}

// appProfile : Used to set a welcome screen and greeting for new users
type appProfile struct {
	GetStarted getStarted `json:"get_started"`
}

type getStarted struct {
	Payload string `json:"payload"`
}
