package main

import (
	"fmt"
//	term "github.com/nsf/termbox-go"
//	gc "code.google.com/p/goncurses"
//	"os"
)

//Resets
const R_EVERYTHING = "\033[0m" //Resets everything within the escape character
const R_COLOR = "\033[0;00m" //Resets the color of the escape characters

//Clear
const CLEAR_SCREEN = "\033[2J"

var boxIndex []BoarderBox = make([]BoarderBox, 10)
var activeBox int = 0

//Colors
const BLACK = "\033[m"
const BOLD_BLACK = "\033[1m"
const RED = "\033[31m"
const BOLD_RED = "\033[1;31m"
const GREEN = "\033[32m"
const BOLD_GREEN = "\033[1;32m"
const YELLOW = "\033[33m"
const BOLD_YELLOW = "\033[1;33m"
const BLUE = "\033[34m"
const BOLD_BLUE = "\033[1;34m"
const PURPLE = "\033[35m"
const BOLD_PURPLE = "\033[1;35m"
const CYAN = "\033[36m"
const BOLD_CYAN = "\033[1;36m"
const WHITE = "\033[37m"
const BOLD_WHITE = "\033[1;37m"

//Background colors
const BACK_BLACK = "\033[40m"
const BACK_RED = "\033[41m"
const BACK_GREEN = "\033[42m"
const BACK_YELLOW = "\033[43m"
const BACK_BLUE = "\033[44m"
const BACK_PURPLE = "\033[45m"
const BACK_CYAN = "\033[46m"
const BACK_WHITE = "\033[47m"


type BoarderBox struct {
	x int
	y int
	width int
	height int
	forcolor string
	backcolor string
}

func NewBoarderBox(x, y, w, h int, fc string, bc string) *BoarderBox {
	bX := BoarderBox {
		x: x,
		y: y,
		width: 0,
		height: 0,
		forcolor: fc,
		backcolor: bc,
	}
	bX.width = bX.x+w
	bX.height = bX.y+h

	return &bX
}

func PrntWPos(x, y int, str string) {
	fmt.Printf("\033[%d;%dH%s", x, y, str)
}

func PrntPosWClr(x, y int, str, fg, bg string) {
	fmt.Printf(fg + bg + "\033[%d;%dH%s", x, y, str)
}

func ClearScreen() {
	fmt.Printf(CLEAR_SCREEN)
}

//Boarder box funtions 
func (bX *BoarderBox) DrawBox() {
	ClearScreen()
//	w, h := term.Size()
	for i := bX.x; i < bX.width; i++ {
		for j := bX.y; j < bX.height; j++ {
			  //Just the walls of the room
      			if i == bX.width-1 || i == bX.x {
        			PrntPosWClr(i, j, "═", bX.forcolor, bX.backcolor)
			}else if i > bX.x && i < bX.width-1 && j == bX.y || j == bX.height-1{
 				PrntPosWClr(i, j, "║", bX.forcolor, bX.backcolor)
			}

			if i == bX.x && j == bX.y {
        			PrntWPos(i, j, "╔")
      			}else if i == bX.width-1 && j == bX.y {
        			PrntWPos(i, j, "╚")
      			}else if i == bX.x && j == bX.height-1 {
        			PrntWPos(i, j, "╗")
     			}else if i == bX.width-1 && j == bX.height-1 {
        			PrntWPos(i, j, "╝")
      			}
		}
	}
	fmt.Printf(R_COLOR)
}


