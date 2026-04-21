package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	White = lipgloss.Color("#FFFFFF")
	Red   = lipgloss.Color("#FF0000")
	Grey  = lipgloss.Color("#555555")
)

func GetBunny() string {
	// The new bunny! One eye is a %s for the red dot.
	rawBunny := `
    ⠏⢣ ⠏⢣
  ⢠⡶⠧⠧⠶⠧⠧⠶⢶⡄
  ⡜         ⢣
 ⢸   ⠛   ⠛  ⢣
  ⢣      Y  ⢸
  ⢸      "  ⡜
  ⡜        ⢸
⠺⡜         ⡜
  ⠙⠒⠤⣀⣀⣇⣸⣇⣸
  
 DOT %s DO`

	redEye := lipgloss.NewStyle().Foreground(Red).Render("●")

	// Style the body
	body := lipgloss.NewStyle().Foreground(White).Render(fmt.Sprintf(rawBunny, redEye))

	return body
}
