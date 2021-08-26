package testauth

import (
	"encoding/base64"
	auth "google-authenticator/src/authenticator"
	_ "image/jpeg"
	_ "image/png"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func TestGetOneTimePassWd(t *testing.T) {
	tests := []auth.OtpOption {
		{Name: "node36", Secret: "CMBXJLUI6ZKQSWYF"},
		{Name: "jump-01", Secret: "BBGUVDGMG2BORUMPPI6HGJW4N4"},
	}

	for _, tt := range tests {
		code, err := auth.GetOneTimePassword(tt.Secret, tt.Name)
		if err != nil {
			t.Errorf("get code of name error. %v", err.Error())
		} else {
			t.Logf("name: %v, code: %v", tt.Name, code)
		}
	}
}

func TestScanQr(t *testing.T) {

	tests := []struct {
		name, path, want string
	}{
		{name: "node36", path: filepath.Join("examples", "node36.png"), want: "otpauth://totp/anxin@node36?secret=CMBXJLUI6ZKQSWYF"},
		{name: "node36-export", path: filepath.Join("examples", "export03.jpg"), want: "otpauth-migration://offline?data=CiAKChMDdK6I9lUJWwUSDGFueGluQG5vZGUzNiABKAEwAhABGAEgACi7vILO%2FP%2F%2F%2F%2F8B"},
		{name: "node36+jump01-export", path: filepath.Join("examples", "export10.jpg"), want: "otpauth-migration://offline?data=CiAKChMDdK6I9lUJWwUSDGFueGluQG5vZGUzNiABKAEwAgovChAITUqMzDaC6NGPejxzJtxvEgxqdW1wQGp1bXAtMDEaB2p1bXAtMDEgASgBMAIQARgBIAAom6W%2BkP3%2F%2F%2F%2F%2FAQ%3D%3D"},
	}

	for _, tt := range tests {
		file, err := os.Open(tt.path)
		if err != nil {
			t.Errorf("get stream of name error. %v", err.Error())
			continue
		}
		result, err := auth.ScanQr(file)
		if err != nil {
			t.Errorf("get result of name error. %v", err.Error())
		} else {
			if result != tt.want {
				t.Errorf(" test result of %v not equal want", tt.name)
				t.Failed()
			}
			t.Logf("name: %v, code: %v", tt.name, result)
		}
	}
}

func TestBase64(t *testing.T) {
	str := "CiAKChMDdK6I9lUJWwUSDGFueGluQG5vZGUzNiABKAEwAgovChAITUqMzDaC6NGPejxzJtxvEgxqdW1wQGp1bXAtMDEaB2p1bXAtMDEgASgBMAIQARgBIAAom6W%2BkP3%2F%2F%2F%2F%2FAQ%3D%3D"

	escapeStr, _ := url.QueryUnescape(str)
	t.Log(escapeStr)
	basestr, _ := base64.RawURLEncoding.DecodeString(escapeStr)
	t.Log(basestr)

}

func TestParseUrl(t *testing.T) {

	tests := []struct {
		url, want string
	}{
		{url: "otpauth://totp/anxin@node36?secret=CMBXJLUI6ZKQSWYF", want: ""},
		{url: "otpauth-migration://offline?data=CiAKChMDdK6I9lUJWwUSDGFueGluQG5vZGUzNiABKAEwAhABGAEgACi7vILO%2FP%2F%2F%2F%2F8B", want: ""},
		{url: "otpauth-migration://offline?data=CiAKChMDdK6I9lUJWwUSDGFueGluQG5vZGUzNiABKAEwAgovChAITUqMzDaC6NGPejxzJtxvEgxqdW1wQGp1bXAtMDEaB2p1bXAtMDEgASgBMAIQARgBIAAom6W%2BkP3%2F%2F%2F%2F%2FAQ%3D%3D", want: ""},
	}

	for _, tt := range tests {
		t.Log("===============================================================")
		t.Log(tt.url)
		secrets, err := auth.ParseUrl(tt.url)
		if err != nil {
			t.Error("failed", err)
			return
		}
		t.Logf("exist %v secrets.", len(secrets))
		for index, secret := range secrets {
			t.Logf("index: %v     name: %v   secret:%v  ", index, secret.Name, secret.Secret)
		}
	}
}

func TestReg(t *testing.T) {
	text := `123+341/314++asd/fa====`
	reg := regexp.MustCompile(`\+\+|\=|/`)
	str := reg.ReplaceAllString(text, " ")
	t.Log(str)

	newstr := auth.Replace(text, []string{`\+`, "=", "/"}, " ")
	t.Log(newstr)

}
