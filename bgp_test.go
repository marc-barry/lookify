package main

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	NotifgicationShutdown = `{"exabgp":"3.4.6","time":1426275092,"host":"host1","pid":"3891","ppid":"993","counter":1,"type":"notification","notification":"shutdown"}`
	StateConnected        = `{"exabgp":"3.4.6","time":1426275029,"host":"host1","pid":"3891","ppid":"993","type":"state","neighbor":{"ip":"10.0.50.1","address":{"local":"10.0.50.2","peer":"10.0.50.1"},"asn":{"local":"65534","peer":"65101"},"state":"connected"}}`
	StateUp               = `{"exabgp":"3.4.6","time":1426275030,"host":"host1","pid":"3891","ppid":"993","type":"state","neighbor":{"ip":"10.0.50.1","address":{"local":"10.0.50.2","peer":"10.0.50.1"},"asn":{"local":"65534","peer":"65101"},"state":"up"}}`
	StateDown             = `{"exabgp":"3.4.6","time":1426275026,"host":"host1","pid":"3891","ppid":"993","type":"state","neighbor":{"ip":"10.0.50.1","address":{"local":"10.0.50.2","peer":"10.0.50.1"},"asn":{"local":"65534","peer":"65101"},"state":"down","reason":"out loop, peer reset, message [closing connection] error[the TCP connection was closed by the remote end]"}}`
	UpdateAnnounce        = `{"exabgp":"3.4.6","time":1427907444,"host":"host1","pid":"5676","ppid":"5486","type":"update","neighbor":{"ip":"10.0.50.1","address":{"local":"10.0.50.2","peer":"10.0.50.1"},"asn":{"local":"65534","peer":"65101"},"message":{"update":{"attribute":{"origin":"igp","as-path":[65101,65000,65102],"confederation-path":[],"atomic-aggregate":false},"announce":{"ipv4 unicast":{"10.0.50.1":{"10.0.60.0/24":{}}}}}}}}`
	UpdateWithdraw        = `{"exabgp":"3.4.6","time":1426593594,"host":"host1","pid":"5676","ppid":"5486","type":"update","neighbor":{"ip":"10.0.50.1","address":{"local":"10.0.50.2","peer":"10.0.50.1"},"asn":{"local":"65534","peer":"65101"},"message":{"update":{"withdraw":{"ipv4 unicast":{"192.168.99.2/32":{}}}}}}}`
)

func TestBGP(t *testing.T) {
	Convey("Testing BGP message reading", t, func() {
		buffer := bytes.NewBuffer([]byte{})
		_, _ = buffer.WriteString(NotifgicationShutdown + "\n")
		_, _ = buffer.WriteString(StateConnected + "\n")
		_, _ = buffer.WriteString(StateUp + "\n")
		_, _ = buffer.WriteString(StateDown + "\n")
		_, _ = buffer.WriteString(UpdateAnnounce + "\n")
		_, _ = buffer.WriteString(UpdateWithdraw + "\n")
		bgp, err := NewBGP(buffer)
		So(err, ShouldBeNil)
		So(bgp, ShouldNotBeNil)

		bgp.ReadMessages()

		n := bgp.Notifications()
		So(len(n), ShouldEqual, 1)

		s := bgp.States()
		So(len(s), ShouldEqual, 3)

		u := bgp.Updates()
		So(len(u), ShouldEqual, 2)
	})
}
