package main

import (
	"log"
	"time"
)

const M_25MB = 25 * 1000 * 1000

func main() {
	log.Println("esperando quinze segundos...")
	time.Sleep(10 * time.Second)

	var memory []byte
	for {
		log.Println("Alocando")
		memory = append(memory, make([]byte, M_25MB)...)
		log.Println("Esperando 2 segundos...")
		time.Sleep(2 * time.Second)
	}
}
