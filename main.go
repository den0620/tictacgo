package main

import (
    "fmt"
)

var PlayerSkin [2]string
var PlayerWins [2]int // {0, 0}




func printMatrix(matrix *[3][3]int) {
    
    fmt.Println("┌───────┐" ) // as 3 is fixed size I can just type 3 "─"s
    for i := 0; i < 3; i++ {
        var line string = "│ "
        for j := 0; j < 3; j++ {
            switch matrix[i][j] {
                case 1:
                    line += PlayerSkin[0] + " "
                case 2:
                    line += PlayerSkin[1] + " "
                default:
                    line += "  "
            }
        }
        line += "│"
        fmt.Println(line)
    }
    fmt.Println("└───────┘") // same as upper
}


func end(matrix *[3][3]int) int { // check if endgame conditions are met and return winner / 0 is space / 1 is P1 / 2 is P2 / 3 is draw
    // All win cases
    for i := 0; i < 3; i++ { // search all rows for wins
        if matrix[i][0] == matrix[i][1] && matrix[i][1] == matrix[i][2] {
            return matrix[i][0]
        }
    }
    for j := 0; j < 3; j++ { // search all columns for wins
        if matrix[0][j] == matrix[1][j] && matrix[1][j] == matrix[2][j] {
            return matrix[0][j]
        }
    }
    if matrix[0][0] == matrix[1][1] && matrix[1][1] == matrix[2][2] { // search skews for wins
        return matrix[0][0]
    } else if matrix[0][2] == matrix[1][1] && matrix[1][1] == matrix[2][0] {
        return matrix[0][2]
    }
    // Noone won so far; check for space
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if matrix[i][j] == 0 { // still some space
                return 0
            }
        }
    }
    
    return 3 // then draw
}

func startgame() int {
    var winner int = 0  // 0 is continue / 1 is P1 / 2 is P2 / 3 is draw
    var turn int = 1    // Px
    var field [3][3]int // 0 is empty / 1 is P1 / 2 is P2
    var numkey, i, j int


    for { // make prompt for a move and change matrix
        printMatrix(&field)
        for { // repeat 'til proper num
            fmt.Print("Use numpad to chose (number)\nYour move: ")
            fmt.Scanln(&numkey)
            if 0 < numkey && numkey < 10 { // numkey is numkey and convert to index
                // so num is in range but is this cell free?
                i = 2 - (numkey-1)/3
                j = (numkey-1)%3
                if field[i][j] == 0 { // eg empty
                    break
                }
            }
            fmt.Println("Move is invalid")
        }
        field[i][j] = turn
        switch turn {
            case 1:
                turn = 2
            case 2:
                turn = 1
        }
		winner = end(&field)
        if winner != 0 {
            break
        }
    }
    printMatrix(&field)

    return winner
}


func main() {
    PlayerSkin[0] = "X"
    PlayerSkin[1] = "O"

	fmt.Println("===== Tic Tac Go =====")
	fmt.Println("||   CLI  edition   ||")
	fmt.Println("======================\n")

    var tmpstr string
    var winner int

    for i := 0; i < 2; i++ {
        fmt.Printf("Choose Skin (prefer char) for Player %d\ndefault is '%s': ", i+1, PlayerSkin[i])
        fmt.Scanln(&tmpstr)
        if tmpstr != "" {
            PlayerSkin[i] = tmpstr
            tmpstr = ""
        }
    }
    fmt.Print("\nNow we're gaming!\n\n")
    var EC bool = false
    for EC != true { 
        winner = startgame()
		if winner != 3 {
        	PlayerWins[winner-1]++

        	fmt.Printf("\nPlayer %d '%s' won!\n\n", winner, PlayerSkin[winner-1])
		} else {
			fmt.Printf("\nThat's a draw!\n\n")
		}
        
        fmt.Print("Would you like to play again?\n(Y/n): ")
        fmt.Scanln(&tmpstr)

        switch tmpstr {
            case "Y":
                fallthrough
            case "y":
                fallthrough
            case "" :
                fmt.Println("Ok playing again")
            default:
                EC = true
        }
    }
	
    fmt.Println("\n======= Score ========")
	fmt.Printf("Player 1 '%s': %d\n", PlayerSkin[0], PlayerWins[0])
	fmt.Printf("Player 2 '%s': %d\n", PlayerSkin[1], PlayerWins[1])
    fmt.Print("===== Goodbye :3 =====\n")
}
