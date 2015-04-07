package types

import (
	"bytes"
	"fmt"
	"net"
	"strings"
)

// Represents a BGP community attribute in AA:NN format.
// AA is a 16-bit AS number and NN is a 2-byte interger.
// Ultimately, a community attribute is just a 32-bit value.
type CommunityAttribute struct {
	AA uint16
	NN uint16
}

type AnnounceRouteType struct {
	IPNet               *net.IPNet
	NextHopIP           net.IP
	CommunityAttributes []CommunityAttribute
}

type WithdrawRouteType struct {
	IPNet     *net.IPNet
	NextHopIP net.IP
}

func (t *CommunityAttribute) String() string {
	return strings.Join([]string{fmt.Sprintf("%d", t.AA), fmt.Sprintf("%d", t.NN)}, ":")
}

func communityAttributeListString(cas []CommunityAttribute) string {
	buffer := bytes.NewBufferString("[")

	count := len(cas)
	for i, attr := range cas {
		if _, err := buffer.WriteString(attr.String()); err != nil {
			break
		}
		if i != count-1 {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}

func (t *AnnounceRouteType) String() string {
	return strings.Join([]string{
		"announce route",
		t.IPNet.String(),
		"next-hop",
		t.NextHopIP.String(),
		"community",
		communityAttributeListString(t.CommunityAttributes)}, " ")
}

func (t *WithdrawRouteType) String() string {
	return strings.Join([]string{
		"withdraw route",
		t.IPNet.String(),
		"next-hop",
		t.NextHopIP.String()}, " ")
}
