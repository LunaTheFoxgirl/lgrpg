package main

import (
	"math/rand"
	"time"
	"strings"
)

var consBlends []string = []string{
	"bl",
	"br",
	"ch",
	"cl",
	"cr",
	"dr",
	"fl",
	"fr",
	"gl",
	"gr",
	"pl",
	"pr",
	"sc",
	"sh",
	"sk",
	"si",
	"sm",
	"sn",
	"sp",
	"st",
	"sw",
	"th",
	"tr",
	"tw",
	"wh",
	"wr",
	"sch",
	"scr",
	"shr",
	"sph",
	"spl",
	"spr",
	"squ",
	"str",
	"thr",
}

var consonants []string = []string{
	"b",
	"c",
	"d",
	"f",
	"g",
	"h",
	"j",
	"k",
	"l",
	"m",
	"n",
	"p",
	"q",
	"r",
	"s",
	"t",
	"v",
	"w",
	"x",
	"z",
}

var vowels []string = []string {
	"a",
	"e",
	"i",
	"o",
	"u",
	"y",
}

var letters []string = []string {
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
	"i",
	"j",
	"k",
	"l",
	"m",
	"n",
	"o",
	"p",
	"q",
	"r",
	"s",
	"t",
	"u",
	"v",
	"w",
	"x",
	"y",
	"z",
}

func GenerateRandomWord(length int) string {
	curr := 0
	isVowel := false
	ofCons := 0
	out := ""
	rand.Seed(time.Now().Unix())
	if (rand.Int()%100 == 50) {
		// Start with consonant blend
		out += consBlends[rand.Int()%len(consBlends)]
		curr += len(out)
		isVowel = true
	} else {
		// start with either vowel or consonant
		if (rand.Int()%100 == 50) {
			out += consonants[rand.Int()%len(consonants)]
			curr++
			isVowel = true
		} else {
			out += vowels[rand.Int()%len(vowels)]
			curr++
		}
	}
	for curr < length {
		if (!isVowel) {
			out += consonants[rand.Int()%len(consonants)]
			ofCons++
			curr++
		} else {
			ofCons = 0
			out += vowels[rand.Int()%len(vowels)]
			curr++
		}
		if (!isVowel) {
			if (ofCons < 2) {
				if (rand.Int()%100 == 50) {
					isVowel = !isVowel
				}
				continue
			}
		}
		isVowel = !isVowel
	}
	return out
}

func GenerateRandomSpacename(iteration int) string {
	rand.Seed(time.Now().Unix())
	spaceName := GenerateRandomWord(rand.Int()%12+4)
	spaceName = strings.ToUpper(spaceName[0:1]) + spaceName[1:]
	spaceName += " " + strings.ToUpper(letters[iteration%len(letters)])
	return spaceName
}