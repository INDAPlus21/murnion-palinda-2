package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)
	// TODO: Answer questions.
	go AnswerQuestions(questions, answers)
	// TODO: Make prophecies.
	go MakeProphecies(questions)
	// TODO: Print answers.
	go PrintAnswers(answers)
	return questions
}

func AnswerQuestions(questions chan string, answers chan string) {
	for question := range questions {
		go prophecy(question, answers)
	}
}

func MakeProphecies(answers chan string) {
	questions := []string{
		"What is my ip?",
		"What time is it?",
		"How do I register to vote?",
		"How do I tie a tie?",
		"What song was that?",
		"When is mother's day?",
		"Where am I now?",
		"How to make pancakes?",
		"How to make money?",
		"Why is the sky blue?",
	}

	for {
		time.Sleep(time.Duration(5+rand.Intn(16)) * time.Second)
		question := questions[rand.Intn(10)]
		fmt.Println("\n" + "You need not say it. I know what you want to ask. \"" + question + "\"")
		answers <- question
	}
}

func PrintAnswers(answers chan string) {
	for answer := range answers {
		fmt.Println(star, ": ", answer)
		fmt.Printf(prompt)
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	regWho := regexp.MustCompile(`(?i)who`) // I'm well aware that this will match things that aren't just who, like "show"
	regWhy := regexp.MustCompile(`(?i)why`)
	regWhen := regexp.MustCompile(`(?i)when`)
	regWhat := regexp.MustCompile(`(?i)what`)
	regHow := regexp.MustCompile(`(?i)how`)

	whoCount := len(regWho.FindAllString(question, -1))
	whyCount := len(regWhy.FindAllString(question, -1))
	whenCount := len(regWhen.FindAllString(question, -1))
	whatCount := len(regWhat.FindAllString(question, -1))
	howCount := len(regHow.FindAllString(question, -1))

	chosenVagueness := 0
	if whoCount > whyCount && whoCount > whenCount && whoCount > whatCount && whoCount > howCount {
		chosenVagueness = 1
	} else if whyCount > whoCount && whyCount > whenCount && whyCount > whatCount && whyCount > howCount {
		chosenVagueness = 2
	} else if whenCount > whyCount && whenCount > whoCount && whenCount > whatCount && whenCount > howCount {
		chosenVagueness = 3
	} else if whatCount > whyCount && whatCount > whoCount && whatCount > whenCount && whatCount > howCount {
		chosenVagueness = 4
	} else if howCount > whyCount && howCount > whoCount && howCount > whatCount && howCount > whenCount {
		chosenVagueness = 5
	}

	// Cook up some pointless nonsense.
	vagueNonsense := []string{
		"All I can say is... your lucky numbers are 6 and 11.",
		"The future is shrouded. I do not know.",
		"Truly, that is one of the questions of all time.",
		"Ask Google",
		"One of the enduring mysteries of our age, I'm sure.",
	}
	vaguePersonage := []string{
		"Who indeed?",
		"It could be anyone. It could be you, it could be me, it could even be... no, nevermind.",
		"I don't know much about them, but I heard they rest in the hallowed mountains of Tibet.",
		"Their real name is too terrible for me to speak. You are not strong enough to hear it.",
		"You've met them already. You just need to ask the right person. Not me.",
	}
	vagueReason := []string{
		"Why indeed?",
		"Some reasons are not meant to be found.",
		"Why is the ocean blue, and why is grass green? It simply is.",
		"For the same reasons that the lion hunts, and the rabbit runs. It is simply the nature of things.",
		"I think the reason is obvious. Think.",
	}
	vagueTiming := []string{
		"When indeed?",
		"Tomorrow. The answer is always tomorrow. Never today.",
		"Don't think too much about it. Just know that good things come to those who wait.",
		"The last time was decades ago. It may never happen again.",
		"Nobody knows.",
	}
	vagueObject := []string{
		"What indeed?",
		"Only God knows.",
		"I think it may be better not to dwell on these things.",
		"It is indescribable, unknowable, ineffable.",
		"That's nonsense. You might as well be asking if Pi is a wet or dry number.",
	}
	vaguePossiblity := []string{
		"How indeed?",
		"With great care, and great effort.",
		"It may no longer be possible. We might never find out.",
		"A miracle, nothing less.",
		"I cannot answer. It would ruin the wonder of it all.",
	}
	nonsense := [][]string{
		vagueNonsense,
		vaguePersonage,
		vagueReason,
		vagueTiming,
		vagueObject,
		vaguePossiblity,
	}
	answer <- longestWord + "... " + nonsense[chosenVagueness][rand.Intn(len(nonsense[chosenVagueness]))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
