package fyndiq

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestResourceDeliveryNote(t *testing.T) {
	Convey("Resource Delivery Note", t, func() {
		fapi := NewFyndiqAPI("user", "token", false)

		Convey("should generate correct delivery notes URL", func() {
			url := fapi.deliveryNotesURL(RequestParams{"order_by": "input"})
			expected := "https://fyndiq.se/api/v1/order?order_by=input&token=token&user=user"
			So(url, ShouldEqual, expected)
		})
	})
}
