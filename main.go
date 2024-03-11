package main

import (
    "fmt"
)

var PlayerSkin [2]string
var PlayerWins [2]int // {0, 0}




func printMatrix(matrix *[3][3]int) {
    
    fmt.Println("┌───┐") // as 3 is fixed size I can just type 3 "─"s
    for i := 0; i < 3; i++ {
        var line string = "│"
        for j := 0; j < 3; j++ {
            switch matrix[i][j] {
                case 1:
                    line += PlayerSkin[0]
                case 2:
                    line += PlayerSkin[1]
                default:
                    line += " "
            }
        }
        line += "│"
        fmt.Println(line)
    }
    fmt.Println("└───┘") // same as upper
}


func end(matrix *[3][3]int) int { // check if endgame conditions are met and return winner / 0 is nobody / 1 is P1 / 2 is P2
    // some code idk
    return 0
}

func startgame() int {
    var winner int = 0
    var field [3][3]int // 0 is empty / 1 is P1 / 2 is P2

    printMatrix(&field)

    for {
        
        // make prompt for a move and change matrix

        if end(&field) != 0 {
            break
        }
    }



    return winner
}


func main() {
    PlayerSkin[0] = "X"
    PlayerSkin[1] = "O"

    fmt.Print("write some init message\n")

    var tmpstr string
    var winner int

    for i := 0; i < 2; i++ {
        fmt.Printf("Choose Skin (char) for Player %d\ndefault is '%s': ", i, PlayerSkin[i])
        fmt.Scanln(&tmpstr)
        if tmpstr != "" {
            PlayerSkin[i] = tmpstr
        }
        fmt.Printf("You've selected '%s'\n", tmpstr)
    }
    for { 
        winner = startgame()
        PlayerWins[winner]++

        
        fmt.Print("Would you like to play again?\n(Y/n): ")
        fmt.Scan(&tmpstr)

        switch tmpstr {
            case "Y":
            case "y":
            case "" :
            default:
                break
        }
    }

    fmt.Printf("=== Score ===\nP1: %d\n P2:%d\n", PlayerWins[0], PlayerWins[1])
    fmt.Print("write some exit message\n")
}
