package authenticator

type OneTimePassword struct {
	OtpOption
	Password  uint32 `json:"password"`
	Remaining int64  `json:"remaining"`

	algorithm string
	digits    int32
	otpType   string
}

type OtpOption struct {
	Secret string `json:"secret"`
	Name   string `json:"name"`
}

var OtpMap = map[string]OtpOption{
	"CMBXJLUI6ZKQSWYF": {
		Name:   "node36",
		Secret: "CMBXJLUI6ZKQSWYF",
	},
	"BBGUVDGMG2BORUMPPI6HGJW4N4": {
		Name:   "jump-01",
		Secret: "BBGUVDGMG2BORUMPPI6HGJW4N4",
	},
}
