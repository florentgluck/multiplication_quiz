package main

import (
    "os"
    "fmt"
    "time"
    "bufio"
    "strings"
    "strconv"
    "math/rand"
    "github.com/fatih/color"
)

const (
    nMax = 10
)

func questionInt(msg string, minValue int, errorMsg string) int {
    var val int
    for {
        val = readInteger(msg)
        if val >= minValue {
            break
        } else {
            fmt.Println(errorMsg)
        }
    }
    return val
}

func waitForEnter(msg string) {
    fmt.Print(msg)
    var s string
    fmt.Scanf("%s", &s)
}

func readInteger(msg string) int {
    for {
        fmt.Print(msg)
        reader := bufio.NewReader(os.Stdin)
        line, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Failed reading input: ", err)
            continue
        }
        line = strings.TrimSpace(line)
        v := strings.Fields(line)
        if len(v) != 1 {
            fmt.Println("Please enter exactly one number!")
            continue
        }
        i, err := strconv.Atoi(v[0])
        if err != nil {
            fmt.Println("Please enter a valid number!")
            continue
        }
        return i
    }
}

func main() {

    rand.Seed(time.Now().UnixNano())

    quit := false

    for !quit {
        fmt.Print("\033[H\033[2J") // clear screen

        multTab := questionInt("Multiplication table up to (2 or higher): ", 2, "Please enter a value larger than 1")
        questionCount := questionInt("Number of questions: ", 1, "At least 1 question must be specified")
        maxTimeInSec := questionInt("Time limit per question (seconds): ", 1, "Please enter a value larger than 0")
        fmt.Println(multTab, "x", nMax, " multiplication table quiz is ready")
        waitForEnter("Press ENTER to start")
        correctAnswers := 0
        correctAnswersInTime := 0
        avgTime := 0.0

        for i := 0; i < questionCount; i++ {
            tab := rand.Intn(multTab-1)+2
            m := rand.Intn(tab-1)+2
            n := rand.Intn(nMax-1)+2

            question := fmt.Sprintf("[Q%d]: %d x %d = ? ", i+1, m, n)
            timeStart := time.Now()
            var answer int
            answer = readInteger(question)
            timeEnd := time.Now()
            timeElapsed := timeEnd.Sub(timeStart)

            correctCalc := false

            if answer == m*n {
                color.Set(color.FgGreen)
                fmt.Print("CORRECT answer 🙂 ")
                correctCalc = true
                correctAnswers += 1
            } else {
                color.Set(color.FgRed)
                fmt.Print("Wrong answer 🙁 ")
            }
            color.Unset()

            timeElapsedInSec := timeElapsed.Seconds()
            if timeElapsedInSec <= float64(maxTimeInSec) {
                color.Set(color.FgGreen)
                fmt.Printf("in %.1f seconds 🙂", timeElapsedInSec)
                if correctCalc {
                    fmt.Println(" Bravo!")
                    correctAnswersInTime += 1
                } else {
                    fmt.Println("")
                }
            } else {
                color.Set(color.FgRed)
                fmt.Printf("in %.1f seconds, too slow 🙁\n", timeElapsedInSec)
            }
            color.Unset()

            if correctCalc {
                avgTime = avgTime+timeElapsedInSec
            }
        }

        fmt.Println("")
        fmt.Println("Correct answers in time: ",correctAnswersInTime, "/", questionCount)
        fmt.Printf("Correct answer average time: %.1f seconds\n", avgTime/float64(correctAnswers))
        if correctAnswersInTime == questionCount {
            fmt.Println("CONGRATULATIONS you achieved a perfect score! 🏆")
        }
        fmt.Println("")
        waitForEnter("Press ENTER to restart")

        // for {
        //     fmt.Println("Restart? (y/n)")
        //     var restart string
        //     fmt.Scanf("%s", &restart)
        //     if restart == "y" {
        //         fmt.Println("")
        //         break
        //     } else if restart == "n" {
        //         quit = true
        //         break
        //     } else {
        //         fmt.Println("Please type \"y\" or \"n\"")
        //     }
        // }
    }
}
