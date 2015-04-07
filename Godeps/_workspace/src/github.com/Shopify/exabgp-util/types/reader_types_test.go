package types

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExaBGPReaderMessage(t *testing.T) {
	Convey("Testing ability to unmarshal JSON test files into ExaBGPReaderMessage types", t, func() {
		Convey("Testing state type", func() {
			Convey("Testing connected state", func() {
				b, err := ioutil.ReadFile("../files/json/state_connected.json")
				So(err, ShouldBeNil)
				So(b, ShouldNotBeEmpty)

				m, err := UnmarshalExaBGPReaderMessage(b)
				So(err, ShouldBeNil)
				So(m, ShouldNotBeNil)

				So(m.Type, ShouldEqual, StateMessageType)

				t, err := UnmarshalStateType(m.Neighbor)
				So(err, ShouldBeNil)
				So(t, ShouldNotBeNil)

				u, err := json.Marshal(t)
				So(err, ShouldBeNil)
				So(u, ShouldNotBeEmpty)

				compacted := bytes.Buffer{}
				err = json.Compact(&compacted, m.Neighbor)
				So(err, ShouldBeNil)

				So(compacted.String(), ShouldEqual, string(u))
			})
			Convey("Testing up state", func() {
				b, err := ioutil.ReadFile("../files/json/state_up.json")
				So(err, ShouldBeNil)
				So(b, ShouldNotBeEmpty)

				m, err := UnmarshalExaBGPReaderMessage(b)
				So(err, ShouldBeNil)
				So(m, ShouldNotBeNil)

				So(m.Type, ShouldEqual, StateMessageType)

				t, err := UnmarshalStateType(m.Neighbor)
				So(err, ShouldBeNil)
				So(t, ShouldNotBeNil)

				u, err := json.Marshal(t)
				So(err, ShouldBeNil)
				So(u, ShouldNotBeEmpty)

				compacted := bytes.Buffer{}
				err = json.Compact(&compacted, m.Neighbor)
				So(err, ShouldBeNil)

				So(compacted.String(), ShouldEqual, string(u))
			})
			Convey("Testing down state", func() {
				b, err := ioutil.ReadFile("../files/json/state_down.json")
				So(err, ShouldBeNil)
				So(b, ShouldNotBeEmpty)

				m, err := UnmarshalExaBGPReaderMessage(b)
				So(err, ShouldBeNil)
				So(m, ShouldNotBeNil)

				So(m.Type, ShouldEqual, StateMessageType)

				t, err := UnmarshalStateType(m.Neighbor)
				So(err, ShouldBeNil)
				So(t, ShouldNotBeNil)

				u, err := json.Marshal(t)
				So(err, ShouldBeNil)
				So(u, ShouldNotBeEmpty)

				compacted := bytes.Buffer{}
				err = json.Compact(&compacted, m.Neighbor)
				So(err, ShouldBeNil)

				So(compacted.String(), ShouldEqual, string(u))
			})
		})

		Convey("Testing notification type", func() {
			Convey("Testing shutdown notification", func() {
				b, err := ioutil.ReadFile("../files/json/notification_shutdown.json")
				So(err, ShouldBeNil)
				So(b, ShouldNotBeEmpty)

				m, err := UnmarshalExaBGPReaderMessage(b)
				So(err, ShouldBeNil)
				So(m, ShouldNotBeNil)

				So(m.Type, ShouldEqual, NotificationMessageType)

				t, err := UnmarshalNotificationType(m.Notification)
				So(err, ShouldBeNil)
				So(t, ShouldNotBeNil)

				u, err := json.Marshal(t)
				So(err, ShouldBeNil)
				So(u, ShouldNotBeEmpty)

				compacted := bytes.Buffer{}
				err = json.Compact(&compacted, m.Notification)
				So(err, ShouldBeNil)

				So(compacted.String(), ShouldEqual, string(u))
			})
		})

		Convey("Testing update type", func() {
			Convey("Testing announce update", func() {
				b, err := ioutil.ReadFile("../files/json/update_announce.json")
				So(err, ShouldBeNil)
				So(b, ShouldNotBeEmpty)

				m, err := UnmarshalExaBGPReaderMessage(b)
				So(err, ShouldBeNil)
				So(m, ShouldNotBeNil)

				So(m.Type, ShouldEqual, UpdateMessageType)

				t, err := UnmarshalUpdateType(m.Neighbor)
				So(err, ShouldBeNil)
				So(t, ShouldNotBeNil)

				u, err := json.Marshal(t)
				So(err, ShouldBeNil)
				So(u, ShouldNotBeEmpty)

				compacted := bytes.Buffer{}
				err = json.Compact(&compacted, m.Neighbor)
				So(err, ShouldBeNil)

				So(compacted.String(), ShouldEqual, string(u))
			})
			Convey("Testing withdraw update", func() {
				b, err := ioutil.ReadFile("../files/json/update_withdraw.json")
				So(err, ShouldBeNil)
				So(b, ShouldNotBeEmpty)

				m, err := UnmarshalExaBGPReaderMessage(b)
				So(err, ShouldBeNil)
				So(m, ShouldNotBeNil)

				So(m.Type, ShouldEqual, UpdateMessageType)

				t, err := UnmarshalUpdateType(m.Neighbor)
				So(err, ShouldBeNil)
				So(t, ShouldNotBeNil)

				u, err := json.Marshal(t)
				So(err, ShouldBeNil)
				So(u, ShouldNotBeEmpty)

				compacted := bytes.Buffer{}
				err = json.Compact(&compacted, m.Neighbor)
				So(err, ShouldBeNil)

				So(compacted.String(), ShouldEqual, string(u))
			})
		})
	})
}
