package fyndiq

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestResourceAccount(t *testing.T) {
	Convey("Resource Account", t, func() {
		fapi := NewFyndiqAPI("user", "token", false)

		Convey("should generate correct get account URL", func() {
			url := fapi.getAccountURL()
			expected := "https://fyndiq.se/api/v1/account?token=token&user=user"
			So(url, ShouldEqual, expected)
		})
		Convey("should generate correct update account URL", func() {
			url := fapi.updateAccountURL(1)
			expected := "https://fyndiq.se/api/v1/account/1?token=token&user=user"
			So(url, ShouldEqual, expected)
		})
		Convey("should generate correct create account URL", func() {
			url := fapi.createAccountURL()
			expected := "https://fyndiq.se/api/v1/account"
			So(url, ShouldEqual, expected)
		})
	})
}
