# BOTW Save Editor

A terminal-based save editor for The Legend of Zelda: Breath of the Wild, built with [Bubble Tea](https://github.com/charmbracelet/bubbletea). Edit rupees, inventory, and more in `.sav` files via a TUI.

## Features

- Open and edit BOTW `.sav` files
- Modify rupees and inventory items
- Inventory category filtering and editing
- Backup save files before changes
- User-friendly keyboard navigation

## Usage

1. Build and run:

```sh
go build -o save-editor
./save-editor
```

2. Select a `.sav` file to edit.
3. Use arrow keys to navigate, `e` to edit, `i` to view inventory, `c` to change category, `b` to go back, and `q` to quit.
4. Confirm changes as prompted. Backups are created in the `backups/` directory.

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Bubbles](https://github.com/charmbracelet/bubbles)
- [Lipgloss](https://github.com/charmbracelet/lipgloss)
- See `go.mod` for full list.

## Project Structure

- `main.go` — Main TUI application logic
- `components/` — Custom UI and data components
- `game_data.sav` — Example save file
- `backups/` — Auto-created for save file backups

## .gitignore

```
save-editor
debug.log
test.sav
.aider*
game_data.sav
backups
```
