package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	poss := map[string]string{
		// First Letter is Person choose and second is computer choose
		"RP": "C",
		"PR": "P",
		"SR": "C",
		"RS": "P",
		"PS": "C",
		"SP": "P",
	}
	var l string
	f := false
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	fmt.Println("ROCK(R), PAPER(P) OR SCISSORS(S): ")
	for i := 0; i < 10; i++ {
		k, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		switch k {
		case 'R':
			l += "R"
			f = true
		case 'P':
			l += "P"
			f = true
		case 'S':
			l += "S"
			f = true
		default:
			time.Sleep(time.Millisecond * 10)
		}
		if f {
			break
		}
	}
	cAns := rand.Intn(3)
	switch cAns {
	case 1:
		l += "R"
	case 2:
		l += "P"
	case 3:
		l += "S"
	default:
		panic(cAns)
	}

	switch poss[l] {
	case "C":
		fmt.Printf("You: %v | PC: %v\nPC Won", string(l[0]), string(l[1]))
	case "P":
		fmt.Printf("You: %v | PC: %v\nYou Won", string(l[0]), string(l[1]))
	default:
		fmt.Printf("You: %v | PC: %v\nTIE", string(l[0]), string(l[1]))
	}
}
