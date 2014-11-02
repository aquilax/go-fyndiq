package fyndiq

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestResourceProduct(t *testing.T) {
	Convey("Resource Product", t, func() {
		fapi := NewFyndiqAPI("user", "token", false)

		Convey("should generate correctproducts URL", func() {
			url := fapi.productsURL(RequestParams{"a": "1"})
			expected := "https://fyndiq.se/api/v1/product?a=1&token=token&user=user"
			So(url, ShouldEqual, expected)
		})
		Convey("should generate correct product URL", func() {
			url := fapi.productURL(1)
			expected := "https://fyndiq.se/api/v1/product/1?token=token&user=user"
			So(url, ShouldEqual, expected)
		})
	})
}
