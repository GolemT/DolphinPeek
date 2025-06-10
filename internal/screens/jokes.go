package screens

import (
	"fmt"
	"math/rand"
)

func GetRandomJoke() string {
	jokes := []string{
		"Why don't dolphins pay taxes?\nBecause they live in international waters!",
		"What do you call a dolphin magician?\nA dolphin-dini!",
		"Why did the dolphin cross the ocean?\nTo get to the other tide!",
		"What's a dolphin's favorite TV show?\nWhale of Fortune!",
		"Why are dolphins so smart?\nBecause they're always thinking outside the tank!",
	}

	facts := []string{
		"Dolphins sleep with one eye open!\nPerfect for monitoring system resources 24/7",
		"Dolphins can hold their breath for 15 minutes!\nLonger than most coding sessions without coffee ☕",
		"Dolphins use echolocation to navigate!\nBetter than GPS for finding bugs in code",
	}

	// Mix jokes and facts randomly
	if rand.Intn(2) == 0 {
		joke := jokes[rand.Intn(len(jokes))]
		return fmt.Sprintf(`
🐬 DOLPHIN JOKE 🐬

%s

        🌊🐬🌊
     .-""""""-.
   .'          '.
  /   O      O   \
 :                :
 |                |
 :       <        :
  \  .-""""""-.  /
   '.'        '.'
     '-.......-'
     
Press nothing to continue...`, joke)
	} else {
		fact := facts[rand.Intn(len(facts))]
		return fmt.Sprintf(`
🧠 DOLPHIN WISDOM 🧠

%s

     .-.   .-.
    (  o\ /o  )
     \   '   /
      ) ___ (
     /  \_/  \
    /    _    \
   /    (_)    \
  
GolemT's System Powered by Dolphin Logic™`, fact)
	}
}
