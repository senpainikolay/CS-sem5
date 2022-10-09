package stream

import (
	"math"
)

type RC4 struct {
	NR_OF_BYTES int
	kS          []byte
	plaintText  []byte
	s           []byte
}

func RC4_Init(NR_OF_BYTES int) *RC4 {
	rc := RC4{
		NR_OF_BYTES,
		make([]byte, 0),
		make([]byte, 0),
		make([]byte, 0),
	}
	return &rc

}

func (c *RC4) Encrypt(keyScheduleKey string, Text string) string {
	c.streamGenerator(keyScheduleKey, Text)
	en := make([]byte, len(c.kS))

	for i := range c.kS {
		// XOR
		en[i] = c.plaintText[i] ^ c.kS[i]
	}

	return string(en)

}

func (c *RC4) Decrypt(keyScheduleKey string, Text string) string {
	c.streamGenerator(keyScheduleKey, Text)
	de := make([]byte, len(c.kS))

	for i := range c.kS {
		// XOR
		de[i] = c.plaintText[i] ^ c.kS[i]
	}

	return string(de)

}

func (c *RC4) streamGenerator(keyScheduleKey string, Text string) {

	c.keySchedule(keyScheduleKey)
	c.plaintText = []byte(Text)
	c.kS = make([]byte, 0)
	j := 0
	for i := 1; i <= len(c.plaintText); i++ {
		j = int(math.Mod(float64(j)+float64(c.s[i]), float64(c.NR_OF_BYTES)))
		c.s[i], c.s[j] = c.s[j], c.s[i]
		t := int(math.Mod(float64(c.s[i])+float64(c.s[j]), float64(c.NR_OF_BYTES)))
		c.kS = append(c.kS, c.s[t])
	}

}

func (c *RC4) keySchedule(Key string) {
	c.s = make([]byte, c.NR_OF_BYTES)
	for i := 0; i < c.NR_OF_BYTES; i++ {
		c.s[i] = byte(i)
	}
	K := make([]byte, 0)

	// Init Keystream byte array for key scheduling
	remainder := math.Mod(float64(c.NR_OF_BYTES), float64(len([]byte(Key))))
	// add the key to the Keystream for key scheduling
	for {
		if len(K) >= c.NR_OF_BYTES {
			break
		}
		K = append(K, []byte(Key)...)
	}

	// adjust Keystream lenght for key scheduling
	if remainder != 0.0 {
		adjustedSlice := K[:len(K)-len([]byte(Key))+int(remainder)]
		K = adjustedSlice
	}

	//key scheuling
	j := 0
	for i := 0; i < c.NR_OF_BYTES; i++ {
		j = int(math.Mod(float64(j)+float64(c.s[i])+float64(K[i]), float64(c.NR_OF_BYTES)))
		c.s[i], c.s[j] = c.s[j], c.s[i]
	}

}
