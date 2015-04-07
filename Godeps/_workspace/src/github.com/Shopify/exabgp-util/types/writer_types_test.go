package types

import (
	"net"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCommunityAttributeType(t *testing.T) {
	Convey("Testing community attribute type", t, func() {
		ca := CommunityAttribute{65534, 100}
		So(ca.String(), ShouldEqual, "65534:100")
	})
}

func TestCommunityAttributeListString(t *testing.T) {
	Convey("Testing community attribute list string", t, func() {
		cas := []CommunityAttribute{CommunityAttribute{65534, 100}, CommunityAttribute{65534, 200}, CommunityAttribute{65534, 300}}
		So(communityAttributeListString(cas), ShouldEqual, "[65534:100 65534:200 65534:300]")
	})
}

func TestAnnounceRouteType(t *testing.T) {
	Convey("Testing announce route type", t, func() {
		_, ipnet, err := net.ParseCIDR("192.168.99.2/32")
		So(err, ShouldBeNil)

		ipNH := net.ParseIP("192.168.99.1")

		ca := CommunityAttribute{65534, 100}

		t := AnnounceRouteType{IPNet: ipnet, NextHopIP: ipNH, CommunityAttributes: []CommunityAttribute{ca}}
		So(t.String(), ShouldEqual, "announce route 192.168.99.2/32 next-hop 192.168.99.1 community [65534:100]")
	})
}

func TestWithdrawRouteType(t *testing.T) {
	Convey("Testing withdraw route type", t, func() {
		_, ipnet, err := net.ParseCIDR("192.168.99.2/32")
		So(err, ShouldBeNil)

		ipNH := net.ParseIP("192.168.99.1")

		t := WithdrawRouteType{IPNet: ipnet, NextHopIP: ipNH}
		So(t.String(), ShouldEqual, "withdraw route 192.168.99.2/32 next-hop 192.168.99.1")
	})
}
