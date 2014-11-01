package fyndiq

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUtils(t *testing.T) {
	Convey("Test utils", t, func() {
		fapi := NewFyndiqApi("user", "token")

		Convey("getUrl without params", func() {
			url := fapi.getUrl("test", RequestParams{})
			expected := URL + "test?token=token&user=user"
			So(url, ShouldEqual, expected)
		})

		Convey("getUrl with params", func() {
			url := fapi.getUrl("test", RequestParams{
				"marked": "true",
			})
			expected := URL + "test?marked=true&token=token&user=user"
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
