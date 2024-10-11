package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func Greet(name string) {
	greeting, err := GetGreeting(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greeting)
}

func GreetMany(names ...string) {
	allGreetings, err := GetGreetings(names...)
	if err != nil {
		log.Fatal(err)
	}
	for name, greeting := range allGreetings {
		fmt.Printf("To %v: %v\n", name, greeting)
	}
}

func GetGreeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return fmt.Sprintf(randomGreetTemplate(), name), nil
}

func GetGreetings(names ...string) (map[string]string, error) {
	results := make(map[string]string, len(names))
	for _, name := range names {
		greeting, err := GetGreeting(name)
		if err != nil {
			return nil, fmt.Errorf("getting greeting for %v: %w", name, err)
		}
		results[name] = greeting
	}
	return results, nil
}

func randomGreetTemplate() string {
	randomGreetings := []string{"Wassup %s!", "Howdy doo dee %s.", "Well now %s, helloooo!"}
	return randomGreetings[rand.Intn(len(randomGreetings))]
}
