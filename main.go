package main

import (
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	filepicker    filepicker.Model
	selectedFile  string
	valueList     list.Model
	choice        string
	newValue      string
	fileSelected  bool
	valueSelected bool
}

func newModel() *Model {
	items := []list.Item{
		ListItem{Name: "Rupees", Hash: 0x23149bf8, value: 0},
	}

	l := list.New(items, ItemDelegate{}, 0, 0)
	l.Title = "Select a value to edit"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	fp := filepicker.New()
	fp.AllowedTypes = []string{".sav"}
	fp.CurrentDirectory, _ = os.Getwd() // Set the initial directory to current working directory
	return &Model{
		filepicker: fp,
		valueList:  l,
	}
}

func (m Model) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}

	ok, file := m.filepicker.DidSelectFile(msg)
	if ok {
		m.selectedFile = file
	}

	if m.selectedFile != "" {
		updateListSelection(msg, m)
	}

	var cmd tea.Cmd
	var cmdList []tea.Cmd

	m.filepicker, cmd = m.filepicker.Update(msg)
	cmdList = append(cmdList, cmd)

	return m, tea.Batch(cmdList...)
}

func (m Model) View() string {
	var s strings.Builder

	s.WriteString("Welcome to the BOTW Save Editor!\n")

	if m.selectedFile == "" {
		s.WriteString("Please select a save file (.sav) to edit.\n")
	} else {
		s.WriteString("Selected file: " + m.selectedFile + "\n")
	}

	s.WriteString(m.filepicker.View())

	if m.selectedFile != "" {
		return listSelectionView(m)
	}

	return s.String()
}

func updateListSelection(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}
	return m, nil
}

func listSelectionView(m Model) string {
	var s strings.Builder

	s.WriteString("Select a value to edit:\n")
	s.WriteString(m.valueList.View())

	if m.choice != "" {
		s.WriteString("\nSelected: " + m.choice)
	}

	if m.newValue != "" {
		s.WriteString("\nNew Value: " + m.newValue)
	}

	s.WriteString("\nPress 'q' to quit.")

	return s.String()
}

func main() {
	m := newModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

/* func edit() {
	inputFile := flag.String("file", "", "Path to the BOTW save file (game_data.sav)")
	newRupees := flag.Int("rupees", -1, "New rupee value (-1 to read only)")
	outputFile := flag.String("output", "", "Output file path (defaults to overwriting input file)")
	flag.Parse()

	// Validate arguments
	if *inputFile == "" {
		fmt.Println("Error: Input file path is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Read save file
	saveData, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading save file: %v\n", err)
		os.Exit(1)
	}

	rupeesOffset, err := findRupeesOffset(saveData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	currentRupees := binary.LittleEndian.Uint32(saveData[rupeesOffset : rupeesOffset+4])
	fmt.Printf("Current rupees: %d\n", currentRupees)

	// If no new value provided, exit
	if *newRupees < 0 {
		return
	}

	// Validate rupee value
	if *newRupees > 999999 {
		fmt.Println("Warning: Values over 999999 may cause issues in-game")
	}

	// Write new rupee value
	binary.LittleEndian.PutUint32(saveData[rupeesOffset:rupeesOffset+4], uint32(*newRupees))
	fmt.Printf("Updated rupees to: %d\n", *newRupees)

	// Determine output path
	outPath := *inputFile
	if *outputFile != "" {
		outPath = *outputFile
	}

	// Write modified save file
	err = os.WriteFile(outPath, saveData, 0644)
	if err != nil {
		fmt.Printf("Error writing save file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Save file successfully updated and written to %s\n", outPath)
}

func findRupeesOffset(data []byte) (int, error) {
	// Start searching from offset 0x0c
	for i := 0x0c; i < len(data)-8; i += 8 {
		// Check if current 4 bytes match rupees hash
		hash := binary.LittleEndian.Uint32(data[i : i+4])
		if hash == RUPEES_HASH {
			// Found it - value is 4 bytes after hash
			return i + 4, nil
		}
	}

	// Try Switch endianness (big endian)
	for i := 0x0c; i < len(data)-8; i += 8 {
		hash := binary.BigEndian.Uint32(data[i : i+4])
		if hash == RUPEES_HASH {
			return i + 4, nil
		}
	}

	return 0, fmt.Errorf("rupees value not found in save file")
} */
