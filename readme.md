# TerminalMaze

## How it works

Mazes are generated using randomized DFS(Depth First Search).
After generation, BFS(Breadth First Search) runs from the start to end to find the longest path.

---

## Added in V2

- Turned the project into a playable maze game
- Added movement with arrow keys (`↑`, `↓`, `←`, `→`)
- Added zoom controls with `+` / `-` to resize the maze view
- Improved terminal rendering with lipgloss

---

## Run it

```bash
git clone https://github.com/ChamikaUluwatta/TerminalMaze.git
cd TerminalMaze
go mod tidy
go run main.go
```

---

## Screenshots

![Maze example 1](screenshots/Maze%20example%201.png)

![Maze example 2](screenshots/Maze%20example%202.png)

