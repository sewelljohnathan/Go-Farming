package main

import (
	"fmt"
	"strings"
)

const WORLD_HEIGHT = 10
const WORLD_WIDTH = 20

type GameState struct {
	userCharacter, characterInFront string
	userX, userY                    int
	food, wood, stone, iron         int
	quit                            bool
}

func main() {

	// Initial starting world
	startWorld := [WORLD_HEIGHT]string{
		"####################",
		"#................~.#",
		"#.||...............#",
		"#..|......===......#",
		"#.........=....|...#",
		"#.........=....|...#",
		"#.~~......=........#",
		"#.~~............@..#",
		"#..............@@..#",
		"####################",
	}

	// Copy that initial map into the world variable
	var world [WORLD_HEIGHT][]string
	for y := 0; y < WORLD_HEIGHT; y++ {
		world[y] = strings.Split(startWorld[y], "")
	}

	// Global variables
	var gameState *GameState = initGameState()

	// Main game loop
	gameLoop(world, gameState)

	// End game
	fmt.Println("Game has ended!")
}

func gameLoop(world [WORLD_HEIGHT][]string, gameState *GameState) {

	var running bool = true
	for running {

		// Print the board
		print_screen(world, gameState)

		// Get the input
		var userInput string
		fmt.Scanln(&userInput)
		directions := strings.Split(userInput, "") // Split to allow multiple moves combined (e.g. DDSSAW)

		// Loop through all input
		for _, direction := range directions {

			// Perform the cooresponding action
			performAction(world, gameState, strings.ToLower(direction))

			// Quit if needed
			if gameState.quit {
				return
			}
		}
	}
}

func performAction(world [WORLD_HEIGHT][]string, gameState *GameState, direction string) {

	var characterInFront string
	var userX, userY int = gameState.userX, gameState.userY

	switch direction {
	case "a":
		gameState.userCharacter = "<"
		characterInFront = world[userY][userX-1]

		if strings.Compare(characterInFront, ".") == 0 {
			userX -= 1
			characterInFront = world[userY][userX-1]
		}

	case "d":
		gameState.userCharacter = ">"
		characterInFront = world[userY][userX+1]

		if strings.Compare(characterInFront, ".") == 0 {
			userX += 1
			characterInFront = world[userY][userX+1]
		}

	case "w":
		gameState.userCharacter = "^"
		characterInFront = world[userY-1][userX]

		if strings.Compare(characterInFront, ".") == 0 {
			userY -= 1
			characterInFront = world[userY-1][userX]
		}

	case "s":
		gameState.userCharacter = "v"
		characterInFront = world[userY+1][userX]

		if strings.Compare(characterInFront, ".") == 0 {
			userY += 1
			characterInFront = world[userY+1][userX]
		}

	case "c":
		harvest(world, gameState)

	case "q":
		gameState.quit = true
	}

	gameState.userX, gameState.userY = userX, userY
	gameState.characterInFront = characterInFront
}

func initGameState() *GameState {

	var ret *GameState = new(GameState)

	ret.userCharacter = ">"
	ret.characterInFront = "."
	ret.userX = 1
	ret.userY = 1
	ret.food = 10
	ret.wood = 0
	ret.stone = 0
	ret.iron = 0

	return ret
}

func harvest(world [WORLD_HEIGHT][]string, gameState *GameState) {

	// Harvest
	switch gameState.characterInFront {
	case "=":
		// Increase the food amount
		gameState.food += 1
	case "|":
		// Increase the food amount
		gameState.wood += 1
	case "@":
		// Increase the food amount
		gameState.stone += 1
	case "~":
		// Increase the food amount
		gameState.iron += 1
	default:
		return
	}

	// Turn the harvested character to "."
	var userX, userY = gameState.userX, gameState.userY

	switch gameState.userCharacter {
	case "<":
		world[userY][userX-1] = "."
	case ">":
		world[userY][userX+1] = "."
	case "^":
		world[userY-1][userX] = "."
	case "v":
		world[userY+1][userX] = "."
	}
}

func print_screen(world [WORLD_HEIGHT][]string, gameState *GameState) {

	// Print the resources
	fmt.Println("Resources")
	fmt.Printf("| Food: %3d | Wood: %3d | Stone: %3d | Iron: %3d |\n", gameState.food, gameState.wood, gameState.stone, gameState.iron)

	var toPlace string

	// Loop through every position
	for y := 0; y < WORLD_HEIGHT; y++ {
		for x := 0; x < WORLD_WIDTH; x++ {

			// Place character
			if gameState.userY == y && gameState.userX == x {
				toPlace = gameState.userCharacter
			} else {
				toPlace = world[y][x]
			}

			fmt.Print(toPlace)
		}
		fmt.Println()
	}

	// Prompt the player for input
	fmt.Print("WASD to move, C to harvest, or Q to quit: ")
}
