package template

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		smsOTP := &Sms{}
		o := Otp{
			iOtp: smsOTP,
		}
		o.genAndSendOTP(4)

		fmt.Println("")
		emailOTP := &Email{}
		o = Otp{
			iOtp: emailOTP,
		}
		o.genAndSendOTP(4)
	})
}
