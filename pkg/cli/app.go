package cli

import (
	"context"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/Tech-Preta/aws-resources/pkg/services"
)

// Screen represents different screens in the app
type Screen int

const (
	MainMenu Screen = iota
	S3Menu
	EC2Menu
	S3CreateBucket
	EC2CreateInstances
	ResultScreen
)

// Model represents the main application model
type Model struct {
	screen       Screen
	cursor       int
	choices      []string
	selected     map[int]struct{}
	width        int
	height       int
	result       *services.ResourceResult
	errorMsg     string
	
	// Form fields
	bucketName    string
	region        string
	imageID       string
	instanceType  string
	keyName       string
	count         string
	inputField    int
	inputActive   bool
}

// Styling
var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1)

	itemStyle = lipgloss.NewStyle().
		PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(lipgloss.Color("#7D56F4")).
		Bold(true)

	inputStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1).
		Width(50)

	successStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#04B575")).
		Bold(true)

	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF5F87")).
		Bold(true)
)

// initialModel creates the initial model
func initialModel() Model {
	return Model{
		screen:    MainMenu,
		choices:   []string{"S3 - Manage Buckets", "EC2 - Manage Instances", "Exit"},
		selected:  make(map[int]struct{}),
		region:    "us-east-1", // default region
		count:     "1",         // default count
	}
}

// Init is the first function that will be called. It returns an optional initial command.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles all the I/O and updates the model accordingly.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case resultMsg:
		return m.handleResult(msg)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "esc":
			if m.inputActive {
				m.inputActive = false
				return m, nil
			}
			// Navigate back
			switch m.screen {
			case S3Menu, EC2Menu:
				m.screen = MainMenu
				m.cursor = 0
			case S3CreateBucket:
				m.screen = S3Menu
				m.cursor = 0
				m.inputField = 0
			case EC2CreateInstances:
				m.screen = EC2Menu
				m.cursor = 0
				m.inputField = 0
			case ResultScreen:
				m.screen = MainMenu
				m.cursor = 0
			}
			return m, nil

		case "enter":
			return m.handleEnter()

		case "up", "k":
			if !m.inputActive {
				m.cursor--
				if m.cursor < 0 {
					m.cursor = len(m.getChoices()) - 1
				}
			}

		case "down", "j":
			if !m.inputActive {
				m.cursor++
				if m.cursor >= len(m.getChoices()) {
					m.cursor = 0
				}
			}

		case "tab":
			if m.screen == S3CreateBucket {
				m.inputField = (m.inputField + 1) % 2
			} else if m.screen == EC2CreateInstances {
				m.inputField = (m.inputField + 1) % 5
			}

		default:
			if m.inputActive {
				return m.handleInput(msg.String())
			}
		}
	}

	return m, nil
}

// getChoices returns the current menu choices based on screen
func (m Model) getChoices() []string {
	switch m.screen {
	case MainMenu:
		return []string{"S3 - Manage Buckets", "EC2 - Manage Instances", "Exit"}
	case S3Menu:
		return []string{"Create Bucket", "Back to Main Menu"}
	case EC2Menu:
		return []string{"Create Instances", "Back to Main Menu"}
	case S3CreateBucket:
		return []string{"Create Bucket", "Back to S3 Menu"}
	case EC2CreateInstances:
		return []string{"Launch Instances", "Back to EC2 Menu"}
	default:
		return []string{}
	}
}

// handleEnter processes the enter key based on current screen and cursor position
func (m Model) handleEnter() (tea.Model, tea.Cmd) {
	if m.inputActive {
		m.inputActive = false
		return m, nil
	}

	switch m.screen {
	case MainMenu:
		switch m.cursor {
		case 0: // S3
			m.screen = S3Menu
			m.cursor = 0
		case 1: // EC2
			m.screen = EC2Menu
			m.cursor = 0
		case 2: // Exit
			return m, tea.Quit
		}

	case S3Menu:
		switch m.cursor {
		case 0: // Create Bucket
			m.screen = S3CreateBucket
			m.cursor = 0
			m.inputField = 0
		case 1: // Back
			m.screen = MainMenu
			m.cursor = 0
		}

	case EC2Menu:
		switch m.cursor {
		case 0: // Create Instances
			m.screen = EC2CreateInstances
			m.cursor = 0
			m.inputField = 0
		case 1: // Back
			m.screen = MainMenu
			m.cursor = 0
		}

	case S3CreateBucket:
		switch m.cursor {
		case 0: // Create Bucket
			return m, m.createS3Bucket()
		case 1: // Back
			m.screen = S3Menu
			m.cursor = 0
		}

	case EC2CreateInstances:
		switch m.cursor {
		case 0: // Launch Instances
			return m, m.createEC2Instances()
		case 1: // Back
			m.screen = EC2Menu
			m.cursor = 0
		}

	default:
		if m.screen == S3CreateBucket || m.screen == EC2CreateInstances {
			m.inputActive = true
		}
	}

	return m, nil
}

// handleInput handles text input for form fields
func (m Model) handleInput(key string) (tea.Model, tea.Cmd) {
	switch key {
	case "backspace":
		return m.removeChar(), nil
	case "enter":
		m.inputActive = false
		return m, nil
	default:
		if len(key) == 1 {
			return m.addChar(key), nil
		}
	}
	return m, nil
}

// addChar adds a character to the current input field
func (m Model) addChar(char string) Model {
	switch m.screen {
	case S3CreateBucket:
		switch m.inputField {
		case 0:
			m.bucketName += char
		case 1:
			m.region += char
		}
	case EC2CreateInstances:
		switch m.inputField {
		case 0:
			m.imageID += char
		case 1:
			m.instanceType += char
		case 2:
			m.keyName += char
		case 3:
			m.count += char
		case 4:
			m.region += char
		}
	}
	return m
}

// removeChar removes the last character from the current input field
func (m Model) removeChar() Model {
	switch m.screen {
	case S3CreateBucket:
		switch m.inputField {
		case 0:
			if len(m.bucketName) > 0 {
				m.bucketName = m.bucketName[:len(m.bucketName)-1]
			}
		case 1:
			if len(m.region) > 0 {
				m.region = m.region[:len(m.region)-1]
			}
		}
	case EC2CreateInstances:
		switch m.inputField {
		case 0:
			if len(m.imageID) > 0 {
				m.imageID = m.imageID[:len(m.imageID)-1]
			}
		case 1:
			if len(m.instanceType) > 0 {
				m.instanceType = m.instanceType[:len(m.instanceType)-1]
			}
		case 2:
			if len(m.keyName) > 0 {
				m.keyName = m.keyName[:len(m.keyName)-1]
			}
		case 3:
			if len(m.count) > 0 {
				m.count = m.count[:len(m.count)-1]
			}
		case 4:
			if len(m.region) > 0 {
				m.region = m.region[:len(m.region)-1]
			}
		}
	}
	return m
}

// createS3Bucket creates an S3 bucket
func (m Model) createS3Bucket() tea.Cmd {
	return func() tea.Msg {
		if m.bucketName == "" {
			return resultMsg{
				result: &services.ResourceResult{
					Success: false,
					Error:   "ValidationError",
					Message: "Bucket name is required",
				},
			}
		}

		s3Service, err := services.NewS3Service(m.region)
		if err != nil {
			return resultMsg{
				result: &services.ResourceResult{
					Success: false,
					Error:   "ServiceError",
					Message: fmt.Sprintf("Failed to create S3 service: %v", err),
				},
			}
		}

		params := map[string]interface{}{
			"bucket_name": m.bucketName,
			"region":      m.region,
		}

		result, err := s3Service.CreateResource(context.TODO(), params)
		if err != nil {
			return resultMsg{
				result: &services.ResourceResult{
					Success: false,
					Error:   "CreationError",
					Message: fmt.Sprintf("Failed to create bucket: %v", err),
				},
			}
		}

		return resultMsg{result: result}
	}
}

// createEC2Instances creates EC2 instances
func (m Model) createEC2Instances() tea.Cmd {
	return func() tea.Msg {
		if m.imageID == "" || m.instanceType == "" || m.keyName == "" {
			return resultMsg{
				result: &services.ResourceResult{
					Success: false,
					Error:   "ValidationError",
					Message: "Image ID, instance type, and key name are required",
				},
			}
		}

		ec2Service, err := services.NewEC2Service(m.region)
		if err != nil {
			return resultMsg{
				result: &services.ResourceResult{
					Success: false,
					Error:   "ServiceError",
					Message: fmt.Sprintf("Failed to create EC2 service: %v", err),
				},
			}
		}

		params := map[string]interface{}{
			"image_id":      m.imageID,
			"instance_type": m.instanceType,
			"key_name":      m.keyName,
			"count":         m.count,
			"region":        m.region,
		}

		result, err := ec2Service.CreateResource(context.TODO(), params)
		if err != nil {
			return resultMsg{
				result: &services.ResourceResult{
					Success: false,
					Error:   "CreationError",
					Message: fmt.Sprintf("Failed to create instances: %v", err),
				},
			}
		}

		return resultMsg{result: result}
	}
}

// resultMsg represents a result message
type resultMsg struct {
	result *services.ResourceResult
}

// Update handles result messages
func (m Model) handleResult(msg resultMsg) (tea.Model, tea.Cmd) {
	m.result = msg.result
	m.screen = ResultScreen
	m.cursor = 0
	return m, nil
}

// View renders the application
func (m Model) View() string {
	switch m.screen {
	case MainMenu:
		return m.renderMainMenu()
	case S3Menu:
		return m.renderS3Menu()
	case EC2Menu:
		return m.renderEC2Menu()
	case S3CreateBucket:
		return m.renderS3CreateBucket()
	case EC2CreateInstances:
		return m.renderEC2CreateInstances()
	case ResultScreen:
		return m.renderResult()
	}
	return ""
}

func (m Model) renderMainMenu() string {
	s := titleStyle.Render("AWS Resources CLI") + "\n\n"
	s += "Choose a service to manage:\n\n"

	for i, choice := range m.getChoices() {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			choice = selectedItemStyle.Render(choice)
		} else {
			choice = itemStyle.Render(choice)
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n" + lipgloss.NewStyle().Faint(true).Render("Use ↑/↓ to navigate, Enter to select, q to quit")
	return s
}

func (m Model) renderS3Menu() string {
	s := titleStyle.Render("S3 - Simple Storage Service") + "\n\n"
	s += "Choose an action:\n\n"

	for i, choice := range m.getChoices() {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			choice = selectedItemStyle.Render(choice)
		} else {
			choice = itemStyle.Render(choice)
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n" + lipgloss.NewStyle().Faint(true).Render("Use ↑/↓ to navigate, Enter to select, Esc to go back")
	return s
}

func (m Model) renderEC2Menu() string {
	s := titleStyle.Render("EC2 - Elastic Compute Cloud") + "\n\n"
	s += "Choose an action:\n\n"

	for i, choice := range m.getChoices() {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			choice = selectedItemStyle.Render(choice)
		} else {
			choice = itemStyle.Render(choice)
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n" + lipgloss.NewStyle().Faint(true).Render("Use ↑/↓ to navigate, Enter to select, Esc to go back")
	return s
}

func (m Model) renderS3CreateBucket() string {
	s := titleStyle.Render("Create S3 Bucket") + "\n\n"

	// Bucket Name field
	bucketLabel := "Bucket Name:"
	if m.inputField == 0 {
		bucketLabel = selectedItemStyle.Render("→ " + bucketLabel)
	} else {
		bucketLabel = itemStyle.Render(bucketLabel)
	}
	
	bucketInput := m.bucketName
	if m.inputField == 0 && m.inputActive {
		bucketInput += "_"
	}
	
	s += bucketLabel + "\n"
	s += inputStyle.Render(bucketInput) + "\n\n"

	// Region field
	regionLabel := "Region:"
	if m.inputField == 1 {
		regionLabel = selectedItemStyle.Render("→ " + regionLabel)
	} else {
		regionLabel = itemStyle.Render(regionLabel)
	}
	
	regionInput := m.region
	if m.inputField == 1 && m.inputActive {
		regionInput += "_"
	}
	
	s += regionLabel + "\n"
	s += inputStyle.Render(regionInput) + "\n\n"

	// Action buttons
	for i, choice := range m.getChoices() {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			choice = selectedItemStyle.Render(choice)
		} else {
			choice = itemStyle.Render(choice)
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n" + lipgloss.NewStyle().Faint(true).Render("Use Tab to switch fields, Enter to edit/confirm, Esc to go back")
	return s
}

func (m Model) renderEC2CreateInstances() string {
	s := titleStyle.Render("Create EC2 Instances") + "\n\n"

	fields := []struct {
		label string
		value string
		field int
	}{
		{"Image ID (AMI):", m.imageID, 0},
		{"Instance Type:", m.instanceType, 1},
		{"Key Name:", m.keyName, 2},
		{"Count:", m.count, 3},
		{"Region:", m.region, 4},
	}

	for _, field := range fields {
		label := field.label
		if m.inputField == field.field {
			label = selectedItemStyle.Render("→ " + label)
		} else {
			label = itemStyle.Render(label)
		}

		value := field.value
		if m.inputField == field.field && m.inputActive {
			value += "_"
		}

		s += label + "\n"
		s += inputStyle.Render(value) + "\n\n"
	}

	// Action buttons
	for i, choice := range m.getChoices() {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			choice = selectedItemStyle.Render(choice)
		} else {
			choice = itemStyle.Render(choice)
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n" + lipgloss.NewStyle().Faint(true).Render("Use Tab to switch fields, Enter to edit/confirm, Esc to go back")
	return s
}

func (m Model) renderResult() string {
	if m.result == nil {
		return "No result to display"
	}

	var s strings.Builder
	
	if m.result.Success {
		s.WriteString(successStyle.Render("✅ Success!") + "\n\n")
		s.WriteString(m.result.Message + "\n\n")
		
		if m.result.Data != nil {
			s.WriteString("Details:\n")
			for key, value := range m.result.Data {
				s.WriteString(fmt.Sprintf("  %s: %v\n", key, value))
			}
		}
	} else {
		s.WriteString(errorStyle.Render("❌ Error!") + "\n\n")
		s.WriteString(m.result.Message + "\n")
		if m.result.Error != "" {
			s.WriteString(fmt.Sprintf("Error Code: %s\n", m.result.Error))
		}
	}

	s.WriteString("\n" + lipgloss.NewStyle().Faint(true).Render("Press Esc to return to main menu"))
	return s.String()
}

// Run starts the Bubble Tea application
func Run() error {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}