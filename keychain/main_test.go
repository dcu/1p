package keychain

/*
Generate with salt:
openssl enc -aes-128-cbc -p -k "<password>" -S "<hex salt>" -salt -a -p <<<data

Generate without salt:
openssl enc -aes-128-cbc -p -k "<password>" -nosalt -a -p <<<data

*/
var (
	AES128_Expectations = []map[string]string{
		map[string]string{
			"password":  "password",
			"salt":      "AAAAAAAAAAAAAAAA",
			"key":       "0471219B5CE9A0F64C27093DEA07B362",
			"iv":        "8070AC35A7FD31ACB828A4FA828F566E",
			"plain":     "data",
			"encrypted": "0DVZmV3ZuhHYEZJnwISvGw==",
			"openssl":   "U2FsdGVkX1+qqqqqqqqqqtA1WZld2boR2BGSZ8CErxs=",
		},
		map[string]string{
			"password":  "password",
			"salt":      "ABABABABABABABAB",
			"key":       "95D1F52CECEBA04EC6F499DBB9A0080F",
			"iv":        "170CE82107B5D230B708BF30FCC0FE52",
			"plain":     "data",
			"encrypted": "hEApoliqnfX3VcqejgGhWQ==",
			"openssl":   "U2FsdGVkX1+rq6urq6urq4RAKaJYqp3191XKno4BoVk=",
		},
		map[string]string{
			"password":  "#4Sz{[.4SksZ!",
			"salt":      "A000000000000000",
			"key":       "238973C6764117397E19A9EEA574BFC2",
			"iv":        "B9DC36CBEFAFC11A90D0391ECC4E7C4B",
			"plain":     "datadatadatadatadata",
			"encrypted": "VX335LsxKympV6uz5lFrpJ2TkppdcM92vZd2IbmdQDc=",
			"openssl":   "U2FsdGVkX1+gAAAAAAAAAFV99+S7MSspqVers+ZRa6Sdk5KaXXDPdr2XdiG5nUA3",
		},
		map[string]string{
			"password":  "#4Sz{[.4SksZ!",
			"salt":      "",
			"key":       "8FE4EFD29E3DB807021C20396FF0DE67",
			"iv":        "B1B5120580211898B1B4526186459B0A",
			"plain":     "datadatadatadatadata",
			"encrypted": "xyoDM+1e+XGE1JQQGpJexrpSHuHI4eDEuGu4rLcYK0Q=",
			"openssl":   "xyoDM+1e+XGE1JQQGpJexrpSHuHI4eDEuGu4rLcYK0Q=",
		},
		map[string]string{
			"password":  "",
			"salt":      "",
			"key":       "D41D8CD98F00B204E9800998ECF8427E",
			"iv":        "59ADB24EF3CDBE0297F05B395827453F",
			"plain":     "datadatadatadatadata",
			"encrypted": "ioN2sJknj6LHRg86Jfa2caTTk2RF4z99FqgX7GUkTQo=",
			"openssl":   "ioN2sJknj6LHRg86Jfa2caTTk2RF4z99FqgX7GUkTQo=",
		},
	}
)
