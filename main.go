package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Bold(true).
			Padding(0, 4).
			MarginBottom(1)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 2).
			MarginTop(1).
			MarginBottom(1)

	alertStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FD5B5B")).
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#25C2A0")).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5B91F5"))

	hintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5D5C61")).
			Italic(true)

	inputLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true)

	guessHigherStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF5F87")).
				Bold(true)

	guessLowerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5F87FF")).
			Bold(true)

	historyBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#5D5C61")).
			Padding(1, 2).
			Width(60).
			Align(lipgloss.Center)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(1)
)

type keyMap struct {
	Enter key.Binding
	Esc   key.Binding
	CtrlH key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Esc, k.CtrlH}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Enter, k.Esc, k.CtrlH},
	}
}

var keys = keyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirmar"),
	),
	Esc: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "sair"),
	),
	CtrlH: key.NewBinding(
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "ajuda"),
	),
}

type model struct {
	textInput     textinput.Model
	progress      progress.Model
	help          help.Model
	numeroSecreto int64
	chutes        []struct {
		valor int64
		hot   bool
	}
	tentativas    int
	tentativasMax int
	mensagem      string
	estadoJogo    int
	erroInput     bool
	width         int
	height        int
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Digite seu chute"
	ti.PromptStyle = inputLabelStyle
	ti.Focus()
	ti.CharLimit = 3
	ti.Width = 20

	p := progress.New(progress.WithDefaultGradient())
	p.Width = 60

	h := help.New()
	h.ShowAll = false

	return model{
		textInput:     ti,
		progress:      p,
		help:          h,
		numeroSecreto: rand.Int64N(101),
		chutes: []struct {
			valor int64
			hot   bool
		}{},
		tentativas:    0,
		tentativasMax: 10,
		mensagem:      "Tente adivinhar o número entre 0 e 100!",
		estadoJogo:    0,
		erroInput:     false,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			if m.estadoJogo != 0 {
				return m, tea.Quit
			}

			chute := m.textInput.Value()
			if chute == "" {
				return m, nil
			}

			chuteInt, err := strconv.ParseInt(strings.TrimSpace(chute), 10, 64)
			if err != nil || chuteInt < 0 || chuteInt > 100 {
				m.erroInput = true
				m.mensagem = "Digite um número válido entre 0 e 100."
				m.textInput.SetValue("")
				return m, nil
			}

			m.erroInput = false

			diff := abs(chuteInt - m.numeroSecreto)
			isHot := diff <= 10

			m.chutes = append(m.chutes, struct {
				valor int64
				hot   bool
			}{chuteInt, isHot})
			m.tentativas++
			m.textInput.SetValue("")

			switch {
			case chuteInt < m.numeroSecreto:
				m.mensagem = fmt.Sprintf("O número é MAIOR que %d", chuteInt)
			case chuteInt > m.numeroSecreto:
				m.mensagem = fmt.Sprintf("O número é MENOR que %d", chuteInt)
			default:
				m.mensagem = fmt.Sprintf("Você acertou o número %d em %d tentativas!",
					m.numeroSecreto, m.tentativas)
				m.estadoJogo = 1
			}

			if m.tentativas >= m.tentativasMax && m.estadoJogo == 0 {
				m.mensagem = fmt.Sprintf("Suas tentativas acabaram! O número era %d", m.numeroSecreto)
				m.estadoJogo = 2
			}

		case tea.KeyCtrlH:
			m.help.ShowAll = !m.help.ShowAll
		}

	}

	tiCmd := m.updateTextInput(msg)
	cmds = append(cmds, tiCmd)

	return m, tea.Batch(cmds...)
}

func (m *model) updateTextInput(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return cmd
}

func (m model) View() string {
	if m.width == 0 {
		return "Carregando..."
	}

	progress := float64(m.tentativas) / float64(m.tentativasMax)
	progressBar := m.progress.ViewAs(progress)

	title := titleStyle.Render("🎮 JOGO DE ADIVINHAÇÃO 🎲")
	title = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, title)

	mainContent := ""

	if m.estadoJogo == 0 {
		gameInfo := fmt.Sprintf("Tentativa %d de %d", m.tentativas, m.tentativasMax)
		gameInfo = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, gameInfo)

		progressBarCentered := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, progressBar)

		msgStyle := infoStyle
		if strings.Contains(m.mensagem, "MAIOR") {
			msgStyle = guessHigherStyle
		} else if strings.Contains(m.mensagem, "MENOR") {
			msgStyle = guessLowerStyle
		} else if m.erroInput {
			msgStyle = alertStyle
		}

		message := msgStyle.Render(m.mensagem)
		message = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, message)

		inputField := lipgloss.JoinHorizontal(
			lipgloss.Center,
			inputLabelStyle.Render("Seu palpite: "),
			m.textInput.View(),
		)
		inputField = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, inputField)

		hint := ""
		if m.tentativas > 0 && m.estadoJogo == 0 {
			lastGuess := m.chutes[len(m.chutes)-1].valor
			diff := abs(lastGuess - m.numeroSecreto)

			if diff <= 5 {
				hint = hintStyle.Render("Muito quente! Você está muito próximo!")
			} else if diff <= 10 {
				hint = hintStyle.Render("Quente! Está se aproximando!")
			} else if diff <= 25 {
				hint = hintStyle.Render("Frio... Tente outro número.")
			} else {
				hint = hintStyle.Render("Muito frio... Você está longe.")
			}

			hint = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, hint)
		}

		history := ""
		if len(m.chutes) > 0 {
			historyItems := []string{}
			for i, c := range m.chutes {
				style := lipgloss.NewStyle()
				if c.hot {
					style = style.Foreground(lipgloss.Color("#FF5F87"))
				} else {
					style = style.Foreground(lipgloss.Color("#5F87FF"))
				}

				arrow := ""
				if c.valor < m.numeroSecreto {
					arrow = "↑"
				} else if c.valor > m.numeroSecreto {
					arrow = "↓"
				} else {
					arrow = "✓"
				}

				historyItems = append(historyItems,
					fmt.Sprintf("%d: %s %s", i+1, style.Render(fmt.Sprintf("%d", c.valor)), arrow))
			}

			history = strings.Join(historyItems, " | ")
			history = historyBoxStyle.Render("Palpites: " + history)
			history = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, history)
		}

		mainContent = lipgloss.JoinVertical(
			lipgloss.Center,
			gameInfo,
			progressBarCentered,
			"",
			message,
			"",
			inputField,
			hint,
			"",
			history,
		)
	} else {
		resultStyle := alertStyle
		if m.estadoJogo == 1 {
			resultStyle = successStyle
		}

		result := resultStyle.Render(m.mensagem)

		emoji := "🎯"
		if m.estadoJogo == 2 {
			emoji = "😢"
		}

		stats := fmt.Sprintf("Número de tentativas: %d/%d", m.tentativas, m.tentativasMax)

		mainContent = lipgloss.JoinVertical(
			lipgloss.Center,
			"",
			emoji,
			"",
			result,
			"",
			infoStyle.Render(stats),
			"",
			hintStyle.Render("Pressione Enter para sair ou Esc para encerrar"),
		)

		mainContent = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, mainContent)
	}

	mainBox := boxStyle.Width(m.width - 4).Render(mainContent)
	mainBox = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, mainBox)

	helpView := m.help.View(keys)
	helpView = helpStyle.Render(helpView)
	helpView = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, helpView)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		mainBox,
		helpView,
	)
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro ao executar o programa: %v\n", err)
	}
}
