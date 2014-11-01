package fyndiq

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewFyndiqApi(t *testing.T) {
	Convey("Constructor should not return nil", t, func() {
		fapi := NewFyndiqAPI("user", "token", true)
		So(fapi, ShouldNotEqual, nil)
	})
}
