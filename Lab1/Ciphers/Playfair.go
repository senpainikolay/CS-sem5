package ciphers

import (
	"math"
	"strings"
)

type Playfair struct {
	grid       [5][5]string
	Diagraphs  []string
	Msg        string
	SecretWord string
	Alphabet   string
}

//global coords of i/j in alphabet
var t1, t2 int

func (c *Playfair) Init() {
	// alphabet init
	c.Msg = strings.ToLower(c.Msg)
	c.SecretWord = strings.ToLower(c.SecretWord)

	tempAlphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, ch := range c.SecretWord {
		if ContainsElem(c.Alphabet, string(ch)) {
			continue
		}
		c.Alphabet += string(ch)
	}

	for _, ch := range tempAlphabet {
		if ContainsElem(c.Alphabet, string(ch)) {
			continue
		}
		c.Alphabet += string(ch)
	}

	// grid init
	aC := 0
	t1, t2 = -1, -1
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if c.Alphabet[aC] == 'j' || c.Alphabet[aC] == 'i' {
				if t1 == -1 || t2 == -1 {
					t1, t2 = i, j
				} else {
					c.grid[t1][t2] += string(c.Alphabet[aC])
					aC += 1
					j -= 1
					continue
				}
			}
			c.grid[i][j] = string(c.Alphabet[aC])
			aC += 1
		}
	}

	// two letter pair elimination
	for {
		if checkDublicates(c.Msg) {
			break
		}
		for i := 1; i < len(c.Msg); i++ {
			if c.Msg[i] == c.Msg[i-1] {
				c.Msg = Insert(c.Msg, "x", i)
				break
			}
		}
	}
	// adding extra bogus letter if there is a non paired diagraph
	if math.Mod(float64(len(c.Msg)), 2) == 1 {
		c.Msg += "z"

	}
	// forming the actual Diagraphs
	for i := 0; i < len(c.Msg); i += 2 {
		c.Diagraphs = append(c.Diagraphs, c.Msg[i:i+2])
	}

}

func (c *Playfair) Encrypt() string {

	var encrypted string

	for _, elem := range c.Diagraphs {
		i1, j1 := c.GetGridPos(string(elem[0]))
		i2, j2 := c.GetGridPos(string(elem[1]))

		// case 1 : in the same column
		if j1 == j2 {
			en1 := int(math.Mod(float64(i1+1), 5))
			en2 := int(math.Mod(float64(i2+1), 5))
			encrypted += c.grid[en1][j1]
			encrypted += c.grid[en2][j2]
			continue
		}
		// case 2 : in the same row
		if i1 == i2 {
			en1 := int(math.Mod(float64(j1+1), 5))
			en2 := int(math.Mod(float64(j2+1), 5))
			encrypted += c.grid[i1][en1]
			encrypted += c.grid[i2][en2]
			continue
		}

		// case 3: rectangle
		encrypted += c.grid[i1][j2]
		encrypted += c.grid[i2][j1]

	}

	return encrypted

}

func ContainsElem(arr string, ch string) bool {
	for _, elem := range arr {
		if string(elem) == ch {
			return true
		}

	}
	return false
}

func (c *Playfair) GetGridPos(l string) (int, int) {
	if l == "l" || l == "i" {
		return t1, t2
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if string(c.grid[i][j]) == l {
				return i, j
			}

		}
	}

	return -1, -1
}

func Insert(msg string, w string, index int) string {
	return string(msg[:index]) + w + string(msg[index:])
}

func checkDublicates(msg string) bool {
	for i := 1; i < len(msg); i++ {
		if msg[i] == msg[i-1] {
			return false
		}
	}
	return true
}
