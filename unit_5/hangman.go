package main

import "fmt"
import "os"
import "bufio"
import "math/rand"
import "strings"
import "time"

var words = []string{"food", "rubber", "elevator", "wombat", "rotary", "heuristic", "orange", "doorhinge"}
const maxMisses = 10  // change this to make the game harder or easier

func main() {
	// seed the defaut random source (otherwise we'd get same random word every time we run the program)
	rand.Seed(time.Now().UnixNano())
	// pick a random word from the the words list
	word := words[rand.Intn(len(words))]
	// init the reader
	stdin := bufio.NewReader(os.Stdin)
    // Init miss counter
    missCounter := 0
    // Init unguessed word
    ungWord := setUngWord(word)

    for missCounter < maxMisses {

        // Player won
        if !(strings.Contains(string(ungWord), "_")) {
            fmt.Println("Congratulations you won!")
            return
        }

	    // Prompt player to enter a single letter.
        fmt.Print("Enter a letter: ")
        char, err := stdin.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading from STDIN: ", err)
        }
        char = strings.TrimSpace(char)

        // If bad input ask again...
        if len(char) > 1 {
            fmt.Println("Only one character!")
            continue
        }

        // After each guess, tell player how many guess they have left and display
        // the word with yet unguessed letters displayed as
        // underscores, e.g. elep_a_t for the word elephant without h and n.
        if strings.Contains(word, char) {
            for i := 0; i < len(word); i+=1 {
                if string(word[i]) == char {
                    ungWord[i] = rune(word[i])
                }
            }
        } else {
            missCounter += 1
            fmt.Printf("Bad guess, you have only %d more tries.\n", (maxMisses - missCounter))
        }
        fmt.Println(string(ungWord))
    }
    fmt.Println("You lost! :( ")
}

// Init unguessed word state
func setUngWord(word string) []rune {
    ungWord := []rune(word)
    for i := 0; i < len(ungWord); i+=1 {
        ungWord[i] = '_'
    }
    return ungWord
}
