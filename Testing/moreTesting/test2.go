package main

import (
	"fmt"
)

type RGB struct {
	R, G, B uint8
}

func linInterp(s, e uint8, t float64) uint8 {
	return uint8(float64(s)*(1-t) + float64(e)*t)
}

func generateGradient(s RGB, e RGB, length int) []RGB {
	var gradient []RGB
	for i := 0; i < length; i++ {
		t := float64(i) / float64(length-1)

		redChan := linInterp(s.R, e.R, t)
		blueChan := linInterp(s.B, e.B, t)
		greenChan := linInterp(s.G, e.G, t)
		gradient = append(gradient, RGB{redChan, greenChan, blueChan})
	}
	return gradient
}

// Define a simple color map for different characters
// func applyGradient(g []byte) {
/*
Test function for our generated gradient
func rgbToAnsi(r, g, b uint8) string {
	// ANSI escape code format: \033[38;2;R;G;Bm for RGB colors
	return fmt.Sprintf("\033[38;2;%d;%d;%dm█\033[0m", r, g, b)
}
*/
func foo() {
	// ASCII Art for VANTA
	logo := []string{
		"@@@  @@@   @@@@@@   @@@  @@@  @@@@@@@   @@@@@@   ",
		"@@@  @@@  @@@@@@@@  @@@@ @@@  @@@@@@@  @@@@@@@@  ",
		"@@!  @@@  @@!  @@@  @@!@!@@@    @@!    @@!  @@@  ",
		"!@!  @!@  !@!  @!@  !@!!@!@!    !@!    !@!  @!@  ",
		"@!@  !@!  @!@!@!@!  @!@ !!@!    @!!    @!@!@!@!  ",
		"!@!  !!!  !!!@!!!!  !@!  !!!    !!!    !!!@!!!!  ",
		":!:  !!:  !!:  !!!  !!:  !!!    !!:    !!:  !!!  ",
		" ::!!:!   :!:  !:!  :!:  !:!    :!:    :!:  !:!  ",
		"  ::::    ::   :::   ::   ::     ::    ::   :::  ",
		"   :       :   : :  ::    :      :      :   : :  ",
		"	   *******                 *                                         ",
		"    *       ***             **                                          ",
		"   *         **             **                                          ",
		"   **        *              **                                          ",
		"    ***             ****    **                  ***  ****               ",
		"   ** ***          * ***  * **  ***      ***     **** **** *    ***     ",
		"    *** ***       *   ****  ** * ***    * ***     **   ****    * ***    ",
		"      *** ***    **    **   ***   ***  *   ***    **          *   ***   ",
		"        *** ***  **    **   **     ** **    ***   **         **    ***  ",
		"          ** *** **    **   **     ** ********    **         ********   ",
		"           ** ** **    **   **     ** *******     **         *******    ",
		"            * *  **    **   **     ** **          **         **         ",
		"  ***        *   *******    **     ** ****    *   ***        ****    *  ",
		" *  *********    ******     **     **  *******     ***        *******   ",
		"*     *****      **          **    **   *****                  *****    ",
		"*                **                *                                    ",
		" **              **               *                                     ",
		"                  **             *                                      ",
		"                                *   									 ",
	}

	// Apply character-based colors
	//coloredText := applyCharacterColors(logo)

	length := len(logo)
	fmt.Println(length)
	startRgb := RGB{0, 255, 255}
	endRgb := RGB{255, 0, 0}
	gradient := generateGradient(startRgb, endRgb, length)
	for i := range logo {
		fmt.Printf("%v  RGB = %v\n", logo[i], gradient[i])
	}

	// Print the result
	//fmt.Println(coloredText)
}
func main() {
	// Create a gradient from blue (0, 0, 255) to purple (128, 0, 128)
	startRgb := RGB{0, 0, 255}
	endRgb := RGB{128, 0, 128}
	length := 20
	gradient := generateGradient(startRgb, endRgb, length)

	// Print out the gradient using ANSI escape codes
	for _, color := range gradient {
		fmt.Print(rgbToAnsi(color.R, color.G, color.B))
	}
	fmt.Println() // New line after the gradient
}
