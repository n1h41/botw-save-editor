package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/n1h41/save-editor/components"
)

type FileState struct {
	selectedFile string
	data         []byte
	fileSelected bool
}

type EditingState struct {
	choice            *components.ListItem
	editing           bool
	newValueTextInput textinput.Model
	readValue         string
	confirmSave       bool
}

type Model struct {
	filepicker    filepicker.Model
	valueList     list.Model
	inventoryList list.Model // New list for inventory items
	file          FileState
	edit          EditingState
	inventory     components.InventorySection // New section for inventory management
	mode          string                      // "main", "inventory", etc.
	statusMsg     string
}

func newModel() *Model {
	items := []list.Item{
		components.ListItem{Name: "Rupees", Hash: 0x23149bf8, Value: 0},
	}

	width, height := 80, 20

	ti := textinput.New()
	ti.Placeholder = "Enter new value"
	ti.Focus()
	ti.CharLimit = 10
	ti.Width = 20

	l := list.New(items, components.ItemDelegate{}, width, height)
	l.Title = "Select a value to edit"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = components.TitleStyle
	l.Styles.PaginationStyle = components.PaginationStyle
	l.Styles.HelpStyle = components.HelpStyle

	// Initialize empty inventory list with the same dimensions
	invList := list.New([]list.Item{}, components.InventoryItemDelegate{}, width, height)
	invList.Title = "Inventory Items"
	invList.SetShowStatusBar(false)
	invList.SetFilteringEnabled(true)
	invList.Styles.Title = components.TitleStyle
	invList.Styles.PaginationStyle = components.PaginationStyle
	invList.Styles.HelpStyle = components.HelpStyle

	fp := filepicker.New()
	fp.AllowedTypes = []string{".sav"}
	fp.CurrentDirectory, _ = os.Getwd()

	// Initialize inventory section with constants
	inventory := components.InventorySection{
		ItemsHash:         0x5f283289, // ITEMS hash
		ItemsQuantityHash: 0x6a09fc59, // ITEMS_QUANTITY hash
		Items:             []components.InventoryListItem{},
		Categories:        []string{"All", "Weapons", "Bows", "Shields", "Armor", "Materials", "Food", "Key Items", "Other"},
		CurrentCategory:   "All",
	}

	return &Model{
		filepicker:    fp,
		valueList:     l,
		inventoryList: invList,
		file:          FileState{},
		edit: EditingState{
			newValueTextInput: ti,
		},
		inventory: inventory,
		mode:      "main",
	}
}

func (m Model) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "b":
			if m.file.fileSelected {
				// If in inventory mode, go back to main mode
				if m.mode == "inventory" {
					m.mode = "main"
					m.statusMsg = "Returned to main view."
					return m, nil
				}
				// Otherwise go back to file selection
				m.file.fileSelected = false
				m.edit.choice = nil
				m.edit.editing = false
				m.edit.confirmSave = false
				m.statusMsg = ""
				return m, nil
			}
		case "i":
			// Toggle inventory view
			if m.mode != "inventory" && m.file.fileSelected {
				m.mode = "inventory"
				// Initialize inventory if needed
				if len(m.inventory.Items) == 0 {
					inventory, err := readInventoryData(m.file.data)
					if err != nil {
						m.statusMsg = fmt.Sprintf("Error reading inventory: %v", err)
					} else {
						m.inventory = inventory
						m.inventoryList = initInventoryList(m.inventory, 80, 20)
					}
				}
				m.statusMsg = "Viewing inventory. Press 'c' to change category, 'e' to edit item."
				return m, nil
			} else if m.mode == "inventory" {
				m.mode = "main"
				m.statusMsg = "Returned to main view."
				return m, nil
			}
		case "c":
			// When in inventory mode, cycle through categories
			if m.mode == "inventory" {
				currentIndex := -1
				for i, cat := range m.inventory.Categories {
					if cat == m.inventory.CurrentCategory {
						currentIndex = i
						break
					}
				}

				// Cycle to next category
				nextIndex := (currentIndex + 1) % len(m.inventory.Categories)
				m.inventory.CurrentCategory = m.inventory.Categories[nextIndex]

				// Refresh the list with the new category
				m.inventoryList = initInventoryList(m.inventory, 80, 20)
				m.statusMsg = fmt.Sprintf("Category: %s", m.inventory.CurrentCategory)
				return m, nil
			}
		}
	}

	ok, file := m.filepicker.DidSelectFile(msg)
	if ok {
		m.file.selectedFile = file
		m.file.fileSelected = true
		m.statusMsg = "File loaded successfully."

		data, err := readFile(m.file.selectedFile)
		if err != nil {
			m.statusMsg = fmt.Sprintf("Error reading file: %v", err)
			m.file.fileSelected = false
			return m, nil
		}
		m.file.data = data

		// Reset any inventory data when loading new file
		m.inventory.Items = nil
	}

	var cmd tea.Cmd
	var cmdList []tea.Cmd

	// Update the filepicker
	m.filepicker, cmd = m.filepicker.Update(msg)
	cmdList = append(cmdList, cmd)

	// Update the appropriate list based on mode
	if m.file.fileSelected {
		if m.mode == "inventory" {
			m, cmd = updateInventorySelection(msg, m)
		} else {
			m, cmd = updateListSelection(msg, m)
		}
		cmdList = append(cmdList, cmd)
	}

	return m, tea.Batch(cmdList...)
}

func (m Model) View() string {
	var s strings.Builder

	s.WriteString("🎮 BOTW Save Editor 🎮\n\n")

	if !m.file.fileSelected {
		s.WriteString("Please select a save file (.sav) to edit.\n\n")
		s.WriteString(m.filepicker.View())
		s.WriteString("\n\nPress 'q' to quit.")
	} else {
		return listSelectionView(m)
	}

	return s.String()
}

func readFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func updateListSelection(msg tea.Msg, m Model) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmdList []tea.Cmd

	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "esc":
			if m.edit.editing {
				m.edit.editing = false
				m.edit.newValueTextInput.Reset()
				m.statusMsg = "Editing canceled."
			} else if m.edit.confirmSave {
				m.edit.confirmSave = false
				m.statusMsg = "Save canceled."
			}
		case "enter":
			if m.edit.editing {
				if m.edit.choice != nil {
					newValue, err := strconv.ParseUint(m.edit.newValueTextInput.Value(), 10, 32)
					if err != nil {
						m.statusMsg = "Invalid value. Please enter a valid number."
						return m, nil
					}
					m.edit.readValue = strconv.FormatUint(newValue, 10)
					m.edit.editing = false
					m.edit.confirmSave = true
					m.statusMsg = "Press 'y' to confirm saving this value, or 'esc' to cancel."
				}
			} else if m.edit.confirmSave {
				// Already handled by 'y' key
			} else {
				// If not editing, select the item
				if i, ok := m.valueList.SelectedItem().(components.ListItem); ok {
					m.edit.choice = &i
					updatedItem, found := readItemValue(*m.edit.choice, m.file.data)
					if !found {
						m.statusMsg = "Hash not found in save file."
						m.edit.readValue = "N/A"
					} else {
						m.edit.readValue = strconv.FormatUint(uint64(updatedItem.Value), 10)
						m.statusMsg = fmt.Sprintf("Selected %s. Press 'e' to edit or 'x' to deselect.", updatedItem.Name)
					}
					log.Printf("Selected item: %s, Value: %s\n", updatedItem.Name, m.edit.readValue)
				}
			}
		case "y":
			if m.edit.confirmSave && m.edit.choice != nil {
				newValue, _ := strconv.ParseUint(m.edit.readValue, 10, 32)
				offset, found := findOffset(m.edit.choice.Hash, m.file.data)
				if !found {
					m.statusMsg = "Error: Hash not found in save file."
					m.edit.confirmSave = false
					return m, nil
				}

				// Create backup before writing
				err := createBackup(m.file.selectedFile)
				if err != nil {
					m.statusMsg = fmt.Sprintf("Warning: Failed to create backup: %v. Changes will still be saved.", err)
				}

				err = updateValue(uint32(newValue), m.file.data, offset, 4, m.file.selectedFile)
				if err != nil {
					m.statusMsg = fmt.Sprintf("Error saving changes: %v", err)
				} else {
					m.statusMsg = fmt.Sprintf("Successfully updated %s to %s", m.edit.choice.Name, m.edit.readValue)

					// Update the item in the list
					updatedItems := make([]list.Item, len(m.valueList.Items()))
					for i, item := range m.valueList.Items() {
						li := item.(components.ListItem)
						if li.Hash == m.edit.choice.Hash {
							li.Value = uint32(newValue)
						}
						updatedItems[i] = li
					}
					m.valueList.SetItems(updatedItems)
				}
				m.edit.confirmSave = false
			}
		case "e":
			if m.edit.choice != nil && !m.edit.editing && !m.edit.confirmSave {
				m.edit.editing = true
				m.edit.newValueTextInput.SetValue("")
				m.edit.newValueTextInput.Focus()
				m.statusMsg = "Enter a new value and press Enter."
				return m, textinput.Blink
			}
		case "x":
			if !m.edit.editing && !m.edit.confirmSave {
				m.edit.choice = nil
				m.statusMsg = "Selection cleared."
			}
		}
	}

	m.valueList, cmd = m.valueList.Update(msg)
	cmdList = append(cmdList, cmd)

	m.edit.newValueTextInput, cmd = m.edit.newValueTextInput.Update(msg)
	cmdList = append(cmdList, cmd)

	return m, tea.Batch(cmdList...)
}

// updateInventorySelection handles user interaction with the inventory list
func updateInventorySelection(msg tea.Msg, m Model) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmdList []tea.Cmd

	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "esc":
			if m.edit.editing {
				m.edit.editing = false
				m.edit.newValueTextInput.Reset()
				m.statusMsg = "Editing canceled."
			} else if m.edit.confirmSave {
				m.edit.confirmSave = false
				m.statusMsg = "Save canceled."
			}
		case "enter":
			if m.edit.editing {
				// Process edited value
				newValue, err := strconv.ParseUint(m.edit.newValueTextInput.Value(), 10, 32)
				if err != nil {
					m.statusMsg = "Invalid value. Please enter a valid number."
					return m, nil
				}
				m.edit.readValue = strconv.FormatUint(newValue, 10)
				m.edit.editing = false
				m.edit.confirmSave = true
				m.statusMsg = "Press 'y' to confirm saving this value, or 'esc' to cancel."
			} else if m.edit.confirmSave {
				// Already handled by 'y' key
			}
		case "y":
			if m.edit.confirmSave {
				item, ok := m.inventoryList.SelectedItem().(components.InventoryListItem)
				if !ok {
					m.statusMsg = "Error: No item selected."
					m.edit.confirmSave = false
					return m, nil
				}

				newValue, _ := strconv.ParseUint(m.edit.readValue, 10, 32)

				// Create backup before writing
				err := createBackup(m.file.selectedFile)
				if err != nil {
					m.statusMsg = fmt.Sprintf("Warning: Failed to create backup: %v. Changes will still be saved.", err)
				}

				// Update the item in memory and file
				err = updateInventoryItem(m.file.data, m.inventory.QuantitiesOffset, item, uint32(newValue))
				if err != nil {
					m.statusMsg = fmt.Sprintf("Error updating item: %v", err)
					m.edit.confirmSave = false
					return m, nil
				}

				// Write changes to file
				err = os.WriteFile(m.file.selectedFile, m.file.data, 0644)
				if err != nil {
					m.statusMsg = fmt.Sprintf("Error saving file: %v", err)
					m.edit.confirmSave = false
					return m, nil
				}

				// Update the UI
				m.statusMsg = fmt.Sprintf("Successfully updated %s to %s", item.Name, m.edit.readValue)

				// Update quantity in the inventory items slice
				for i := range m.inventory.Items {
					if m.inventory.Items[i].Index == item.Index {
						m.inventory.Items[i].Quantity = uint32(newValue)
						break
					}
				}

				// Refresh the inventory list to show updated values
				m.inventoryList = initInventoryList(m.inventory, 80, 20)

				m.edit.confirmSave = false
			}
		case "e":
			// Start editing the selected item's quantity
			if !m.edit.editing && !m.edit.confirmSave {
				item, ok := m.inventoryList.SelectedItem().(components.InventoryListItem)
				if ok {
					m.edit.editing = true
					m.edit.newValueTextInput.SetValue(fmt.Sprintf("%d", item.Quantity))
					m.edit.newValueTextInput.Focus()
					m.statusMsg = "Enter a new value and press Enter."
					return m, textinput.Blink
				} else {
					m.statusMsg = "No item selected."
				}
			}
		}
	}

	// Update the inventory list
	m.inventoryList, cmd = m.inventoryList.Update(msg)
	cmdList = append(cmdList, cmd)

	// Update the text input if editing
	m.edit.newValueTextInput, cmd = m.edit.newValueTextInput.Update(msg)
	cmdList = append(cmdList, cmd)

	return m, tea.Batch(cmdList...)
}

func readItemValue(item components.ListItem, data []byte) (components.ListItem, bool) {
	offset, found := findOffset(item.Hash, data)
	if !found {
		return item, false
	}
	value := binary.LittleEndian.Uint32(data[offset : offset+4])
	item.Value = value
	return item, true
}

func findOffset(item_hash uint32, data []byte) (int, bool) {
	// Try little endian first
	for i := 0x0c; i < len(data)-8; i += 8 {
		if i+4 >= len(data) {
			break
		}
		hash := binary.LittleEndian.Uint32(data[i : i+4])
		if hash == item_hash {
			return i + 4, true
		}
	}

	// Try big endian if little endian fails
	for i := 0x0c; i < len(data)-8; i += 8 {
		if i+4 >= len(data) {
			break
		}
		hash := binary.BigEndian.Uint32(data[i : i+4])
		if hash == item_hash {
			return i + 4, true
		}
	}
	return 0, false
}

func updateValue(newValue uint32, data []byte, offset int, length int, filename string) error {
	binary.LittleEndian.PutUint32(data[offset:offset+length], newValue)
	return os.WriteFile(filename, data, 0644)
}

func createBackup(filepath string) error {
	backupDir := "backups"
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return err
	}

	// Get the base filename without path
	baseFilename := filepath
	// Remove path if present
	if lastSlash := strings.LastIndex(baseFilename, "/"); lastSlash >= 0 {
		baseFilename = baseFilename[lastSlash+1:]
	}

	// Separate name and extension
	extension := ".sav"
	baseName := baseFilename
	if dotIndex := strings.LastIndex(baseFilename, "."); dotIndex >= 0 {
		extension = baseFilename[dotIndex:]
		baseName = baseFilename[:dotIndex]
	}

	// Create backup filename with timestamp
	timestamp := time.Now().Format("20060102_150405")
	backupFilename := fmt.Sprintf("%s_%s%s", baseName, timestamp, extension)
	backupPath := fmt.Sprintf("%s/%s", backupDir, backupFilename)

	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	return os.WriteFile(backupPath, data, 0644)
}

// readBOTWString reads a string from BOTW's special string format
// BOTW stores strings in 8-byte chunks, with actual data in first 4 bytes
func readBOTWString(data []byte, offset int) string {
	var result strings.Builder

	// BOTW strings can span up to 16 chunks
	for i := range 16 {
		chunkOffset := offset + (i * 8)
		if chunkOffset+4 > len(data) {
			break
		}

		// Process 4 bytes of character data
		for j := range 4 {
			b := data[chunkOffset+j]
			if b == 0 {
				// Null terminator
				return result.String()
			}
			result.WriteByte(b)
		}
		// Skip 4 bytes of padding
	}

	return result.String()
}

// writeBOTWString writes a string in BOTW's special format
/* func writeBOTWString(data []byte, offset int, str string) {
	// Clear the space first
	for i := range 16 * 8 {
		if offset+i < len(data) {
			data[offset+i] = 0
		}
	}

	// Write the string in chunks of 4 bytes with 4 bytes padding
	for i := range len(str) {
		chunkIndex := i / 4
		byteIndex := i % 4

		if offset+(chunkIndex*8)+byteIndex < len(data) {
			data[offset+(chunkIndex*8)+byteIndex] = str[i]
		}
	}
} */

// readInventoryData reads the inventory from the save file
func readInventoryData(data []byte) (components.InventorySection, error) {
	var inventory components.InventorySection

	// Define hash constants
	inventory.ItemsHash = 0x5f283289
	inventory.ItemsQuantityHash = 0x6a09fc59

	// Find section offsets
	itemsOffset, foundItems := findOffset(inventory.ItemsHash, data)
	quantitiesOffset, foundQuantities := findOffset(inventory.ItemsQuantityHash, data)

	if !foundItems || !foundQuantities {
		return inventory, fmt.Errorf("could not find inventory sections in save file")
	}

	// Store offsets (subtract 4 since itemsOffset points to value after hash)
	inventory.ItemsOffset = itemsOffset - 4
	inventory.QuantitiesOffset = quantitiesOffset - 4

	// Initialize categories
	inventory.Categories = []string{
		"All", "Weapons", "Bows", "Shields", "Armor", "Materials", "Food", "Key Items", "Other",
	}
	inventory.CurrentCategory = "All"

	// Read inventory items (limit to MAX_ITEMS, typically 420)
	inventory.Items = make([]components.InventoryListItem, 0)

	const ITEM_NAME_SIZE = 0x80 // 128 bytes per item entry
	const MAX_ITEMS = 420

	for i := range MAX_ITEMS {
		// Calculate offsets
		nameOffset := itemsOffset + (i * ITEM_NAME_SIZE)
		quantityOffset := quantitiesOffset + (i * 8)

		if nameOffset+ITEM_NAME_SIZE > len(data) || quantityOffset+4 > len(data) {
			break
		}

		// Read item name
		itemName := readBOTWString(data, nameOffset)
		if itemName == "" {
			// Empty slot, we've reached the end of inventory
			break
		}

		// Read quantity/durability
		quantity := binary.LittleEndian.Uint32(data[quantityOffset : quantityOffset+4])

		// Get category
		category := components.GetItemCategory(itemName)

		// Get display name
		displayName := components.GetItemDisplayName(itemName)

		inventory.Items = append(inventory.Items, components.InventoryListItem{
			Name:     displayName,
			RawName:  itemName,
			Quantity: quantity,
			Index:    i,
			Category: category,
		})
	}

	return inventory, nil
}

// initInventoryList initializes the inventory list with current category filter
func initInventoryList(inventory components.InventorySection, width, height int) list.Model {
	items := make([]list.Item, 0)

	// Convert inventory items to list items, filtering by category
	for _, item := range inventory.Items {
		if inventory.CurrentCategory == "All" || item.Category == inventory.CurrentCategory {
			items = append(items, item)
		}
	}

	// Create the list with inventory delegate
	l := list.New(items, components.InventoryItemDelegate{}, width, height)
	l.Title = fmt.Sprintf("Inventory Items - %s", inventory.CurrentCategory)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = components.TitleStyle
	l.Styles.PaginationStyle = components.PaginationStyle
	l.Styles.HelpStyle = components.HelpStyle

	return l
}

// updateInventoryItem updates an item's quantity in memory and the save file
func updateInventoryItem(data []byte, quantitiesOffset int, item components.InventoryListItem, newValue uint32) error {
	// Calculate offset for the quantity value
	quantityValueOffset := quantitiesOffset + 4 + (item.Index * 8)

	if quantityValueOffset+4 > len(data) {
		return fmt.Errorf("invalid offset for item quantity")
	}

	// Update the value in memory
	binary.LittleEndian.PutUint32(data[quantityValueOffset:quantityValueOffset+4], newValue)
	return nil
}

func listSelectionView(m Model) string {
	var s strings.Builder

	s.WriteString("📂 " + m.file.selectedFile + "\n\n")

	// Different view based on mode
	if m.mode == "inventory" {
		// Inventory view
		s.WriteString(m.inventoryList.View())

		s.WriteString("\n\n")

		// Show selected inventory item details
		if item, ok := m.inventoryList.SelectedItem().(components.InventoryListItem); ok {
			s.WriteString(fmt.Sprintf("Selected: %s\n", item.Name))

			if item.Category == "Weapons" || item.Category == "Bows" || item.Category == "Shields" {
				s.WriteString(fmt.Sprintf("Durability: %d\n", item.Quantity))
			} else {
				s.WriteString(fmt.Sprintf("Quantity: %d\n", item.Quantity))
			}

			s.WriteString(fmt.Sprintf("ID: %s\n", item.RawName))
			s.WriteString(fmt.Sprintf("Category: %s\n", item.Category))
		}

		// Show editing UI if applicable
		if m.edit.editing {
			s.WriteString("\nEnter new value: ")
			s.WriteString(m.edit.newValueTextInput.View())
			s.WriteString("\n(Press Enter to confirm, Esc to cancel)")
		} else if m.edit.confirmSave {
			item, _ := m.inventoryList.SelectedItem().(components.InventoryListItem)
			s.WriteString(fmt.Sprintf("\nConfirm changing %s to %s?\n", item.Name, m.edit.readValue))
			s.WriteString("Press 'y' to save, Esc to cancel")
		} else {
			// Regular commands
			s.WriteString("\nCommands: [c] Change category  [e] Edit value  ")
			s.WriteString("[i] Back to main view  [b] Back to file selection  [q] Quit")
		}
	} else {
		// Main regular view
		s.WriteString(m.valueList.View())

		s.WriteString("\n\n")

		if m.edit.choice != nil {
			s.WriteString(fmt.Sprintf("Selected: %s\n", m.edit.choice.Name))
			s.WriteString(fmt.Sprintf("Current value: %s\n\n", m.edit.readValue))
		}

		if m.edit.editing {
			s.WriteString("Enter new value: ")
			s.WriteString(m.edit.newValueTextInput.View())
			s.WriteString("\n(Press Enter to confirm, Esc to cancel)")
		} else if m.edit.confirmSave {
			s.WriteString(fmt.Sprintf("Confirm changing %s to %s?\n", m.edit.choice.Name, m.edit.readValue))
			s.WriteString("Press 'y' to save, Esc to cancel")
		} else {
			s.WriteString("Commands: ")
			if m.edit.choice != nil {
				s.WriteString("[e] Edit value  [x] Clear selection  ")
			}
			s.WriteString("[i] View inventory  [b] Back to file selection  [q] Quit")
		}
	}

	// Status message
	if m.statusMsg != "" {
		s.WriteString("\n\n")
		s.WriteString(m.statusMsg)
	}

	return s.String()
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Printf("Error creating log file: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()

	log.SetOutput(f)

	log.Println("Starting BOTW Save Editor...")

	// Preload translations
	err = components.LoadTranslations()
	if err != nil {
		log.Printf("Warning: Could not load translations: %v. Will use fallback item names.", err)
	} else {
		log.Println("Successfully loaded translations")
	}

	m := newModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
