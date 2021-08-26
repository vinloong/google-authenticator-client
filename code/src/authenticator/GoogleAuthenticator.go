package authenticator

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"strings"
	"time"
)

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

func oneTimePassword(key []byte, value []byte) uint32 {
	// sign the value using HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F

	// get a 32-bit (4-byte) chunk from the hash starting at offset
	hashParts := hash[offset : offset+4]

	// ignore the most significant bit as per RFC 4226
	hashParts[0] = hashParts[0] & 0x7F

	number := toUint32(hashParts)

	// size to 6 digits
	// one million is the first number with 7 digits so the remainder
	// of the division will always return < 7 digits
	pwd := number % 1000000

	return pwd
}

func GetOneTimePassword(secretKey, name string) (*OneTimePassword, error) {
	inputNoSpaces := strings.Replace(secretKey, " ", "", -1)
	inputNoSpacesUpper := strings.ToUpper(inputNoSpaces)
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(inputNoSpacesUpper)
	if err != nil {

		return nil, err
	}
	epochSeconds := time.Now().Unix()
	pwd := oneTimePassword(key, toBytes(epochSeconds/30))
	secondsRemaining := 30 - (epochSeconds % 30)

	timePwd := OneTimePassword{
		OtpOption: OtpOption{
			Secret: secretKey,
			Name:   name,
		},
		Password:  pwd,
		Remaining: secondsRemaining,
	}
	//fmt.Sprintf("%06d (%d second(s) remaining)\n", pwd, secondsRemaining)
	return &timePwd, nil
}
