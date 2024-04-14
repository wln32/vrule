package vrule

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Regex_Domain_Basic(t *testing.T) {
	type RegeDomainStruct struct {
		Domain  string `v:"domain"`
		Domain1 string `v:"domain"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegeDomainStruct{
			Domain:  "goframe#org",
			Domain1: "1a.2b",
		}
		wants := map[string]string{
			"Domain":  "The Domain value `goframe#org` is not a valid domain format",
			"Domain1": "The Domain1 value `1a.2b` is not a valid domain format",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegeDomainStruct{
			Domain:  "goframe.org",
			Domain1: "99designs.com",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})

}

func Test_Regex_Url_Basic(t *testing.T) {
	type RegexUrlStruct struct {
		URL string `v:"url"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexUrlStruct{
			URL: "ws://goframe.org",
		}
		wants := map[string]string{
			"URL": "The URL value `ws://goframe.org` is not a valid URL address",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexUrlStruct{
			URL: "https://www.bilibili.com/",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_Mac_Basic(t *testing.T) {
	type RegexMacStruct struct {
		MAC string `v:"mac"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexMacStruct{
			MAC: "Z0-CC-6A-D6-B1-1A",
		}
		wants := map[string]string{
			"MAC": "The MAC value `Z0-CC-6A-D6-B1-1A` is not a valid MAC address",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexMacStruct{
			MAC: "4C-CC-6A-D6-B1-1A",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_Ip_Basic(t *testing.T) {
	type RegexIpStruct struct {
		Ip   string `v:"ip"`
		Ipv4 string `v:"ipv4"`
		Ipv6 string `v:"ipv6"`
		Ip2  string `v:"ip"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexIpStruct{
			Ip:   "127.0.0.1",
			Ipv6: "fe80::812b:1158:1f43:f0d1",
			Ipv4: "520.255.255.255", // error >= 10000
			Ip2:  "ze80::812b:1158:1f43:f0d1",
		}
		wants := map[string]string{
			"Ipv4": "The Ipv4 value `520.255.255.255` is not a valid IPv4 address",
			"Ip2":  "The Ip2 value `ze80::812b:1158:1f43:f0d1` is not a valid IP address",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexIpStruct{
			Ip:   "127.0.0.1",
			Ipv6: "fe80::812b:1158:1f43:f0d1",
			Ipv4: "193.255.255.255", // error >= 10000
			Ip2:  "2001:0da8:0207:0000:0000:0000:0000:8207",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_QQ_Basic(t *testing.T) {
	type RegexQQStruct struct {
		QQ string `v:"qq"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexQQStruct{
			QQ: "9999",
		}
		wants := map[string]string{
			"QQ": "The QQ value `9999` is not a valid QQ number",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexQQStruct{
			QQ: "123456789",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_Password_Basic(t *testing.T) {
	type RegexPasswordStruct struct {
		Password  string `v:"password"`
		Password2 string `v:"password2"`
		Password3 string `v:"password3"`
	}
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexPasswordStruct{
			Password:  "9999999",
			Password2: "159789523a",
			Password3: "9851af47953a",
		}
		wants := map[string]string{

			"Password2": "The Password2 value `159789523a` is not a valid password format",
			"Password3": "The Password3 value `9851af47953a` is not a valid password format",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}

	})
	gtest.C(t, func(t *gtest.T) {
		obj := &RegexPasswordStruct{
			Password:  "qwerg_0213",
			Password2: "Aqwehbj13142",
			Password3: "gAtrgfhg1454#0",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_BankCard_Basic(t *testing.T) {
	type Regex_BankCard_Struct struct {
		BankCard1 string `v:"bank-card"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_BankCard_Struct{
			BankCard1: "6225760079930218",
		}
		wants := map[string]string{
			"BankCard1": "The BankCard1 value `6225760079930218` is not a valid bank card number",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_BankCard_Struct{
			BankCard1: "6259650871772098",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_ResidentId_Basic(t *testing.T) {
	type Regex_ResidentId_Struct struct {
		ResidentId string `v:"resident-id"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_ResidentId_Struct{
			ResidentId: "320107199506285482",
		}
		wants := map[string]string{
			"ResidentId": "The ResidentId value `320107199506285482` is not a valid resident id number",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

}

func Test_Regex_Postcode_Basic(t *testing.T) {
	type Regex_Postcode_Struct struct {
		Postcode string `v:"postcode"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Postcode_Struct{
			Postcode: "1000000",
		}
		wants := map[string]string{
			"Postcode": "The Postcode value `1000000` is not a valid postcode format",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Postcode_Struct{
			Postcode: "100000",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		t.Assert(err, nil)
	})
}

func Test_Regex_Passport_Basic(t *testing.T) {
	type Regex_Passport_Struct struct {
		Passport string `v:"passport"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Passport_Struct{
			Passport: "123456",
		}
		wants := map[string]string{
			"Passport": "The Passport value `123456` is not a valid passport format",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Passport_Struct{
			Passport: "a100000",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)
		t.Assert(err, nil)
	})
}

func Test_Regex_Telephone_Basic(t *testing.T) {
	type Regex_Telephone_Struct struct {
		Telephone string `v:"telephone"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Telephone_Struct{
			Telephone: "20-77542145",
		}
		wants := map[string]string{
			"Telephone": "The Telephone value `20-77542145` is not a valid telephone number",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Telephone_Struct{
			Telephone: "010-77542145",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}
func Test_Regex_PhoneLoose_Basic(t *testing.T) {
	type Regex_PhoneLoose_Struct struct {
		Phone string `v:"phone-loose"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_PhoneLoose_Struct{
			Phone: "11578912345",
		}
		wants := map[string]string{
			"Phone": "The Phone value `11578912345` is not a valid phone number",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}

	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_PhoneLoose_Struct{
			Phone: "13578912345",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_Email_Basic(t *testing.T) {
	type Regex_Email_Struct struct {
		Email string `v:"email"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Email_Struct{
			Email: "gf#goframe.org",
		}
		wants := map[string]string{
			"Email": "The Email value `gf#goframe.org` is not a valid email address",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Email_Struct{
			Email: "gf@goframe.org.cn",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}
func Test_Regex_Phone_Basic(t *testing.T) {
	type Regex_Phone_Struct struct {
		Phone string `v:"phone"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Phone_Struct{
			Phone: "11578912345",
		}

		wants := map[string]string{
			"Phone": "The Phone value `11578912345` is not a valid phone number",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Phone_Struct{
			Phone: "13578912345",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_Regex_Basic(t *testing.T) {
	type Regex_Regex_Struct struct {
		Pattern  string `v:"regex:[1-9][0-9]{4,14}"`
		Pattern2 string `v:"regex:[1-9][0-9]{4,14}"`
		Pattern3 string `v:"regex:[1-9][0-9]{4,14}"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Regex_Struct{
			Pattern:  "1234",
			Pattern2: "01234",
			Pattern3: "10000",
		}

		wants := map[string]string{
			"Pattern":  "The Pattern value `1234` must be in regex of: [1-9][0-9]{4,14}",
			"Pattern2": "The Pattern2 value `01234` must be in regex of: [1-9][0-9]{4,14}",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}

	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_Regex_Struct{
			Pattern:  "13578912345",
			Pattern2: "1101234",
			Pattern3: "10233000",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}

func Test_Regex_NotRegex_Basic(t *testing.T) {
	type Regex_NotRegex_Struct struct {
		Regex1 string `v:"regex:\\d{4}"`
		Regex2 string `v:"not-regex:\\d{4}"`
	}

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_NotRegex_Struct{
			Regex1: "1234",
			Regex2: "1234",
		}

		wants := map[string]string{
			"Regex2": "The Regex2 value `1234` should not be in regex of: \\d{4}",
		}
		err := getTestValid().StructNotCache(obj).(*ValidationError)

		for rule, msg := range wants {
			fieldError := err.GetFieldError(rule)
			t.Assert(fieldError.Error(), msg)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		obj := Regex_NotRegex_Struct{
			Regex1: "1234",
			Regex2: "hghghghgh",
		}

		err := getTestValid().StructNotCache(obj).(*ValidationError).Errors()
		t.Assert(err, nil)
	})
}
