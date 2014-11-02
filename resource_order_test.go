package fyndiq

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestResourceOrder(t *testing.T) {
	Convey("Resource Order", t, func() {
		fapi := NewFyndiqAPI("user", "token", false)

		Convey("should generate correct get orders URL", func() {
			url := fapi.ordersURL(RequestParams{"marked": "1"})
			expected := "https://fyndiq.se/api/v1/order?marked=1&token=token&user=user"
			So(url, ShouldEqual, expected)
		})
		Convey("should generate correct update order URL", func() {
			url := fapi.orderURL(1)
			expected := "https://fyndiq.se/api/v1/order/1?token=token&user=user"
			So(url, ShouldEqual, expected)
		})
	})
}
