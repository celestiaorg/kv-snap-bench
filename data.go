package main

import (
	"crypto/sha256"
	"log"
	"math/rand"
	"strconv"
)

var hasher = sha256.New()

func hash(x int) []byte {
	hasher.Write([]byte(strconv.Itoa(x)))
	return hasher.Sum(nil)
}

func randomValue() []byte {
	l := rand.Intn(maxValSize-minValSize) + minValSize
	data := make([]byte, l)
	_, err := rand.Read(data)
	if err != nil {
		log.Fatalln("error while creating random value", err)
	}
	return data
}

func randomValues(n int) [][]byte {
	data := make([][]byte, n)
	for i := 0; i < n; i++ {
		data[i] = randomValue()
	}
	return data
}

func insertInitialData(store KV, keyRange int, data [][]byte) {
	for i := 0; i < keyRange; i++ {
		store.Set(hash(i), data[i])
	}
}
