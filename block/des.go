package block

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	k "github.com/senpainikolay/CS-sem5/block/key"
	p "github.com/senpainikolay/CS-sem5/block/permutations"
)

var Sbox = p.GetSBox()
var EXP_DBOX_Perm = p.GetExpansionDBoxArr()
var PBOX_Perm = p.GetPBoxPermutationArr()
var FP = p.GetFPArr()

type DES struct {
	KeysPerRound []k.StringArr
}

func GetDES() DES {
	return DES{make([]k.StringArr, 0)}
}

func (bc *DES) KeyInit(key string) {
	bc.KeysPerRound = k.GetKeysPerRound(key)

}

func (bc *DES) Encrypt(hexString string) (string, string) {

	binStArr := p.HexToBinArray(hexString)
	binStArr = InitialPerm(binStArr)

	LPT, RPT := binStArr[:32], binStArr[32:]

	for i := 0; i < 16; i++ {

		dbox := p.Permute(RPT, EXP_DBOX_Perm, 48)
		xored := XorStringBitArr(dbox, bc.KeysPerRound[i])
		sbox_strArr := SBoxSub(xored)
		sbox_strArr = p.Permute(sbox_strArr, PBOX_Perm, 32)
		result := XorStringBitArr(LPT, sbox_strArr)
		LPT = result
		if i != 15 {
			LPT, RPT = RPT, LPT

		}

	}
	combine_strArr := append(LPT, RPT...)
	cipher := p.Permute(combine_strArr, FP, 64)

	return ConvertBinArrToHex(cipher)

}

func (bc *DES) Decrypt(pt string) (string, string) {

	bc.KeysPerRound = k.ReverseKeysArray(bc.KeysPerRound)
	return bc.Encrypt(pt)

}

func ConvertBinArrToHex(binStr []string) (string, string) {

	var hexSt string
	temp := strings.Join(binStr, "")
	k := 0
	for i := 8; i < len(temp)+8; i += 8 {
		kek := parseBinToHex(temp[k:i])
		if len(kek) == 1 {
			kek = "0" + kek
		}
		k += 8
		hexSt += kek
	}

	// Printing the actual string
	bs, err := hex.DecodeString(hexSt)
	if err != nil {
		panic(err)
	}

	return hexSt, string(bs)

}

func SBoxSub(bitXorStr []string) []string {
	var newArr []string
	for i := 0; i < 8; i++ {
		row := bitXorStr[i*6] + bitXorStr[i*6+5]
		rowa, _ := strconv.ParseInt(row, 2, 64)
		col := bitXorStr[i*6+1] + bitXorStr[i*6+2] + bitXorStr[i*6+3] + bitXorStr[i*6+4]
		cola, _ := strconv.ParseInt(col, 2, 64)
		val := Sbox[i][rowa][cola]
		binary := fmt.Sprint(strconv.FormatInt(int64(val), 2))
		for len(binary) < 4 {
			binary = "0" + binary
		}
		for l := range binary {
			newArr = append(newArr, string(binary[l]))
		}
	}
	return newArr

}

func parseBinToHex(s string) string {
	ui, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return "error"
	}

	return fmt.Sprintf("%x", ui)
}

func XorStringBitArr(a []string, b []string) []string {

	var arr []string

	for i := 0; i < len(b); i++ {

		b1, b2 := a[i], b[i]
		bc1, _ := strconv.ParseBool(b1)
		bc2, _ := strconv.ParseBool(b2)

		// XOR
		if (bc1 || bc2) && !(bc1 && bc2) {
			arr = append(arr, "1")
			continue
		}
		arr = append(arr, "0")

	}

	return arr

}

func InitialPerm(binSt []string) []string {
	initPerm := []int{58, 50, 42, 34, 26, 18, 10, 2,
		60, 52, 44, 36, 28, 20, 12, 4,
		62, 54, 46, 38, 30, 22, 14, 6,
		64, 56, 48, 40, 32, 24, 16, 8,
		57, 49, 41, 33, 25, 17, 9, 1,
		59, 51, 43, 35, 27, 19, 11, 3,
		61, 53, 45, 37, 29, 21, 13, 5,
		63, 55, 47, 39, 31, 23, 15, 7}

	var newBinArr []string

	for i := 0; i < 64; i++ {
		newBinArr = append(newBinArr, binSt[initPerm[i]-1])
	}
	return newBinArr
}
