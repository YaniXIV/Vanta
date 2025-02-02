package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var logo = []string{
	"@@@  @@@   @@@@@@   @@@  @@@  @@@@@@@   @@@@@@   ",
	"@@@  @@@  @@@@@@@@  @@@@ @@@  @@@@@@@  @@@@@@@@  ",
	"@@!  @@@  @@!  @@@  @@!@!@@@    @@!    @@!  @@@  ",
	"!@!  @!@  !@!  @!@  !@!!@!@!    !@!    !@!  @!@  ",
	"@!@  !@!  @!@!@!@!  @!@ !!@!    @!!    @!@!@!@!  ",
	"!@!  !!!  !!!@!!!!  !@!  !!!    !!!    !!!@!!!!  ",
	":!:  !!:  !!:  !!!  !!:  !!!    !!:    !!:  !!!  ",
	" ::!!:!   :!:  !:!  :!:  !:!    :!:    :!:  !:!  ",
	"  ::::    ::   :::   ::   ::     ::    ::   :::  ",
	"   :       :   : :  ::    :      :      :   : :  ",
}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
	isPressed bool
}
type RGB struct {
	R, G, B uint8
}

func linInterp(s, e uint8, t float64) uint8 {
	return uint8(float64(s)*(1-t) + float64(e)*t)
}

func generateGradient(s RGB, e RGB, length int) []string {
	var gradient []string
	for i := 0; i < length; i++ {
		t := float64(i) / float64(length-1)

		redChan := linInterp(s.R, e.R, t)
		blueChan := linInterp(s.B, e.B, t)
		greenChan := linInterp(s.G, e.G, t)

		hex := fmt.Sprintf("#%02x%02x%02x", redChan, greenChan, blueChan)
		gradient = append(gradient, hex)
	}
	return gradient
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "< Send Message >"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			m.isPressed = true
		}

	case errMsg:
		m.err = msg
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf(
		buildlogo(),
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func buildlogo() string {
	var result strings.Builder
	length := len(logo)
	fmt.Println(length)
	//blue
	startRgb := RGB{0, 0, 255}
	//purple
	endRgb := RGB{128, 0, 128}
	gradient := generateGradient(startRgb, endRgb, length)
	for i := range logo {
		styledLine := lipgloss.NewStyle().Foreground(lipgloss.Color(gradient[i])).Render(logo[i])
		result.WriteString(styledLine)
		result.WriteString("\n")
	}
	/*
		for i := range logo {
			fmt.Printf("%v  RGB = %v\n", logo[i], gradient[i])
		}
	*/
	return result.String()
}

func main() {
	//InitWebsocketClient()
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
		tea.WithFPS(120),
	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Could not start program: %s\n", err)
	}
	fmt.Println("End of Program reached")

}
