package key

import (
	p "github.com/senpainikolay/CS-sem5/block/permutations"
)

type StringArr []string

var ShiftPerRound = p.GetShiftArr()
var KC = p.GetKeyCompressionArr()

type Key struct {
	StringVal string
	// array of bit values
	InitKey56BitValue []string
	// Array of arrays of bit values per round.
	KeyPerRoundArr []StringArr
}

func GetKeysPerRound(keyString string) []StringArr {
	key := InitKey(keyString)
	key.Init56Key()
	key.GenerateKeysPerRound()
	return key.KeyPerRoundArr

}

func InitKey(val string) Key {
	return Key{val, make([]string, 0), make([]StringArr, 0)}
}

func ReverseKeysArray(arr []StringArr) []StringArr {
	temp := arr
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return temp

}
func (c *Key) Init56Key() {
	keyBinArr := p.HexToBinArray(p.StringToHex(c.StringVal))

	keyp := []int{57, 49, 41, 33, 25, 17, 9,
		1, 58, 50, 42, 34, 26, 18,
		10, 2, 59, 51, 43, 35, 27,
		19, 11, 3, 60, 52, 44, 36,
		63, 55, 47, 39, 31, 23, 15,
		7, 62, 54, 46, 38, 30, 22,
		14, 6, 61, 53, 45, 37, 29,
		21, 13, 5, 28, 20, 12, 4}

	finalKey := p.Permute(keyBinArr, keyp, 56)
	c.InitKey56BitValue = finalKey

}

func (c *Key) GenerateKeysPerRound() {
	left, right := c.InitKey56BitValue[0:28], c.InitKey56BitValue[28:56]
	var combine_str []string

	for i := 0; i < 16; i++ {
		// Shifting the bits by 1 or 2 bits
		left = ShiftLeft(left, ShiftPerRound[i])
		right = ShiftLeft(right, ShiftPerRound[i])
		combine_str = append(left, right...)

		//  56 to 48 bits key compression
		round_key := p.Permute(combine_str, KC, 48)

		c.KeyPerRoundArr = append(c.KeyPerRoundArr, round_key)
	}

}

func ShiftLeft(key []string, round int) []string {
	var s []string
	for i := 0; i < round; i++ {
		for j := 1; j < len(key); j++ {
			s = append(s, key[j])
		}
		s = append(s, key[0])
		key = s
		s = make([]string, 0)
	}
	return key

}
