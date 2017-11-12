package facebook

// UserProfile : Response from Facebook after a request for a user's profile
type UserProfile struct {
	FirstName string `json:"first_name"`
	Timezone  int    `json:"timezone,omitempty"`
	Gender    string `json:"gender,omitempty"`
}

// WelcomeScreen : Used to set a welcome screen for new users
type WelcomeScreen struct {
	GetStarted string     `json:"get_started"`
	Greeting   []Greeting `json:"greeting"`
}

// Greeting : Specifies a greeting message for a locale for the welcome screen
type Greeting struct {
	Locale string `json:"locale"`
	Text   string `json:"string"`
}
