package models

import (
	"fmt"
	"time"
)

// GitSource is the model for the git_source table.
type GitSource struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// GitUser is the model for the git_user table.
type GitUser struct {
	ID        int    `db:"id"`
	SourceID  int    `db:"source_id"`
	Username  string `db:"username"`
	URL       string `db:"url"`
	AvatarURL string `db:"avatar_url"`

	Source GitSource `db:"source"`
}

// GitRepo is the model for the git_repo table.
type GitRepo struct {
	ID          int    `db:"id"`
	SourceID    int    `db:"source_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	URL         string `db:"url"`

	Source *GitSource `db:"source"`
}

// GitCommit is the model for the git_commit table.
type GitCommit struct {
	ID       int       `db:"id"`
	SourceID int       `db:"source_id"`
	AuthorID int       `db:"author_id"`
	RepoID   int       `db:"repo_id"`
	Message  string    `db:"message"`
	SHA      string    `db:"sha"`
	URL      string    `db:"url"`
	Date     time.Time `db:"date"`

	Source *GitSource `db:"source"`
	Author *GitUser   `db:"author"`
	Repo   *GitRepo   `db:"repo"`

	ColorBackground string
	ColorForeground string
}

// GetColorTheme ...
func (c *GitCommit) GetColorTheme() {
	// Colors ...
	colors := [][]string{
		// canada
		// {"#1dd1a1", "#000000"}, // wild caribbean green
		{"#ffc93c", "#000000"}, // wild caribbean green
		// {"#ff6b6b", "#ffffff"}, // pastel red
		// {"#222f3e", "#ffffff"}, // imperial primer
		// {"#feca57", "#000000"}, // casandora yellow
		// {"#ff9ff3", "#000000"}, // jigglypuff
		// {"#ff9f43", "#000000"}, // double dragon skin
		{"#0ca9f2", "#000000"}, // jade dust
		{"#9a89f7", "#000000"}, // joust blue

		// // spanish
		// {"#ff793f", "#000000"}, // synthetic pumpkin
		{"#2b5cb7", "#ffffff"}, // c64 purple
		// {"#2c2c54", "#ffffff"}, // lucky point

		// // india
		// {"#2c3a47", "#ffffff"}, // ship's officer
		// {"#b33771", "#ffffff"}, // fiery fuchsia
		// {"#fd7272", "#000000"}, // georgia peach
		// {"#1B9CFC", "#ffffff"}, // clear chill

	}

	commitLength := len(c.Message) + len(c.Author.Username)
	colorIndex := commitLength % len(colors)
	fmt.Printf("Got a colorIndex of %v\n", colorIndex)
	color := colors[colorIndex]

	c.ColorBackground = color[0]
	c.ColorForeground = color[1]

	// if len(c.Author.Username)%6 == 0 {
	// 	c.ColorBackground = "#edf2f7"
	// 	c.ColorForeground = "#000000"
	// }
	// c.ColorBackground = "#edf2f7"
	// c.ColorForeground = "#000000"
}

/*
	colors := [][]string{
		// canada
		// {"#1dd1a1", "#000000"}, // wild caribbean green
		{"#ff6b6b", "#ffffff"}, // pastel red
		{"#222f3e", "#ffffff"}, // imperial primer
		// {"#feca57", "#000000"}, // casandora yellow
		// {"#ff9ff3", "#000000"}, // jigglypuff
		// {"#ff9f43", "#000000"}, // double dragon skin
		{"#00d2d3", "#000000"}, // jade dust
		{"#54a0ff", "#000000"}, // joust blue

		// // spanish
		// {"#ff793f", "#000000"}, // synthetic pumpkin
		{"#706fd3", "#ffffff"}, // c64 purple
		{"#2c2c54", "#ffffff"}, // lucky point

		// // india
		{"#2c3a47", "#ffffff"}, // ship's officer
		{"#b33771", "#ffffff"}, // fiery fuchsia
		// {"#fd7272", "#000000"}, // georgia peach
		{"#1B9CFC", "#ffffff"}, // clear chill

	}
*/
