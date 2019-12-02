package idgenerator

import "math/rand"

func GenerateUniqueID(lengthOfID int) string{
	alphabets := []rune("1234567890qwertyuiopasdfghjklzxcvbnm")
	b := make([]rune, lengthOfID)
	for i := range b{
		b[i] = alphabets[rand.Intn(len(alphabets))]
	}
	return string(b)
}
