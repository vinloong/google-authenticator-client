package authenticator

import (
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	stProto "google-authenticator/src/protobuf"
	"google.golang.org/protobuf/proto"
	"image"
	"io"
	"net/url"
)

const (
	URL_PARAM_KEY_SECRET = "secret"
	URL_PARAM_KEY_DATA   = "data"
	URL_TOTP_HOST        = "totp"
	URL_OFFLINE_HOST     = "offline"
)

func ScanQr(file io.Reader) (string, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", err
	}
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	return If(result != nil, result.GetText(), "").(string), err
}

func ParseUrl(rawurl string) ([]OneTimePassword, error) {
	urlParsed, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	urlParams, err := url.ParseQuery(urlParsed.RawQuery)
	if err != nil {
		return nil, err
	}
	if urlParsed.Host == URL_TOTP_HOST {
		name := urlParsed.Path
		name = TrimCOMM(name)
		secret := OneTimePassword{
			OtpOption: OtpOption{
				Secret: urlParams.Get(URL_PARAM_KEY_SECRET),
				Name:   name,
			},
		}
		return []OneTimePassword{secret}, nil
	} else if urlParsed.Host == URL_OFFLINE_HOST {
		dataStr := urlParams.Get(URL_PARAM_KEY_DATA)
		base64str, err := base64.StdEncoding.DecodeString(dataStr)
		if err != nil {
			return nil, err
		}
		base64strbytes := bytes.NewBuffer(base64str)
		msg := &stProto.MigrationPayload{}
		err = proto.Unmarshal(base64strbytes.Bytes(), msg)
		if err != nil {
			return nil, err
		}
		secrets := make([]OneTimePassword, len(msg.OtpParameters))
		for index, otp := range msg.OtpParameters {
			name := otp.Name
			secret := base32.StdEncoding.EncodeToString(otp.Secret)
			secret = TrimCOMM(secret)
			secrets[index] = OneTimePassword{
				OtpOption: OtpOption{
					Secret: secret,
					Name:   name,
				},
				algorithm: otp.Algorithm.String(),
				digits:    otp.Digits,
				otpType:   otp.Type.String(),
			}
		}
		return secrets, nil

	} else {
		return nil, errors.New("unresolved host.")
	}
}
