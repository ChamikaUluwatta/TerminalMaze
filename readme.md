# TerminalMaze

A maze generator that runs in your terminal, written in Go.

---

## How it works

Mazes are generated using randomized DFS(Depth First Search).
After generation, BFS(Breadth First Search) runs from the start to end to find the longest path.

---

## Run it

```bash
git clone https://github.com/ChamikaUluwatta/TerminalMaze.git
cd TerminalMaze
go mod tidy
go run main.go
```

---

## Example

```
+-+ +-+-+    +-+-+ +-+    + +-+-+-+
| |   | |    |     | |    | |     |
+ +-+ + +    + +-+-+ +    + +-+ + +
|   | | |    |   |   |    |   | | |
+-+ + + +    +-+ +-+ +    +-+ + + +
|   |   |    | | |   |    | | | | |
+ +-+-+ +    + + + + +    + + + + +
|       |    |     | |    |     | |
+ +-+-+-+    +-+-+-+ +    +-+-+-+ +


```

