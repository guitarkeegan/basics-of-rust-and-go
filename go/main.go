package main

func main() {
	ch := make(chan []string)
	words := []string{"This is a bunch of words in a sentence", "this is also a bunch of words in a sentence"}
	go splitWords(words, ch)
	getWords(ch)
}
