package fyndiq

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUtils(t *testing.T) {
	Convey("Test utils", t, func() {
		fapi := NewFyndiqAPI("user", "token", true)

		Convey("getURL without params", func() {
			url := fapi.getURL("test", RequestParams{})
			expected := apiURL + "test?test=1&token=token&user=user"
			So(url, ShouldEqual, expected)
		})

		Convey("getURL with params", func() {
			url := fapi.getURL("test", RequestParams{
				"marked": "true",
			})
			expected := apiURL + "test?marked=true&test=1&token=token&user=user"
			So(url, ShouldEqual, expected)
		})

		Convey("getPath returns correct path with empty params", func() {
			path := getPath([]string{})
			expected := "api/v1"
			So(path, ShouldEqual, expected)
		})

		Convey("getPath returns correct path", func() {
			path := getPath([]string{"new", "test", "1"})
			expected := "api/v1/new/test/1"
			So(path, ShouldEqual, expected)
		})
	})
}
