# Pentago (Go Hackathon Project)

This is a simple implementation of the game **Pentago**, written in Go. It was created as part of the boot.dev hackathon, and it's my **first time working with Go**! Had a blast learning the language and creating this fun game. I also used AI to speed up my ideas a couple times, but the majority of the code was written by me.

## What is Pentago?

Pentago is a two-player abstract strategy game. Players take turns placing a marble on a 6x6 board and then rotating one of four 3x3 quadrants. The first player to align five of their pieces in a row (horizontally, vertically) wins. Here, I only check for horizontal and vertical, and it's for two local players.

## How to Run

1. Make sure you have Go installed. You can download it at [https://golang.org/dl/](https://golang.org/dl/)
2. Clone this repository or download the source files.
3. Open a terminal in the project folder.
4. Run the game: `go run main.go` or the binary `./pentago`

```
morales0  » ./pentago
Penta-Go!

○ ○ ○ | ● ○ ○
● ○ ○ | ○ ○ ○
○ ○ ○ | ○ ○ ○
-------------
○ ○ ○ | ○ ○ ○
○ ○ ● | ○ ○ ○
○ ○ ○ | ○ ○ ○

Player 1: Choose a quadrant then type direction (hjkl, r for right, e for left)
```

##  Ideas and Improvements

There’s a lot of room to grow this project. If I had more time, here’s what I’d explore:

-   **Organize code**: Break the code into smaller files (e.g., `board.go`, `game.go`, `ai.go`) to improve readability and structure.
    
-   **Add an AI opponent**: Implement a simple AI using minimax or even just random/legal move selection.
    
-   **Play online**: Use sockets or a lightweight web framework to let players compete online.
    
-   **Write tests**: Add unit tests for move validation, win detection, and board rotation.
