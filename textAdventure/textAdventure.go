package main

// homework
// build your own game
// NPCs - talk to them, fight
// NPC move around the graph
// items that can be picked up or placed down
// accept natural language as input
// verb-noun; attack troll

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choices struct {
	cmd         string
	description string
	nextNode    *storyNode
	nextChoice  *choices
}

type storyNode struct {
	text    string
	choices *choices
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choices{cmd, description, nextNode, nil}

	if node.choices == nil {
		node.choices = choice
	} else {
		currrentChoice := node.choices
		for currrentChoice.nextChoice != nil {
			currrentChoice = currrentChoice.nextChoice
		}
		currrentChoice.nextChoice = choice
	}
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	currrentChoice := node.choices
	for currrentChoice != nil {
		fmt.Println(currrentChoice.cmd, ":", currrentChoice.description)
		currrentChoice = currrentChoice.nextChoice
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	currrentChoice := node.choices
	for currrentChoice != nil {
		if strings.ToLower(currrentChoice.cmd) == strings.ToLower(cmd) {
			return currrentChoice.nextNode
		}
		currrentChoice = currrentChoice.nextChoice
	}
	fmt.Println("Sorry I didn't understand that.")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
	You are in large chamber, deep underground.
	You see three passages leading out. A north passage leads into darkness.
	To the south, a passage appears to head upward. The eastern passages appears
	flat and well traveled`}

	darkRoom := storyNode{text: "It is pitch black. You cannot see a thing."}

	darkRoomLit := storyNode{text: "The dark passage is now lit by your lantern. You can continue north or head back south"}

	grue := storyNode{text: "While stumbling around in the darkness, you are eaten by a grue."}
	trap := storyNode{text: "You head down the well traveled path when suddenly a trap door opens and you fall into a pit."}

	treasure := storyNode{text: "You arrive at a small chamber, filled with treasure!"}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("S", "Try to go back south", &grue)
	darkRoom.addChoice("O", "Turn on lantern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Go South", &start)

	start.play()

	fmt.Println("The End.")
}
