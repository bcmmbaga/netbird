package uspfilter

import (
	"net"

	"github.com/google/gopacket"

	fw "github.com/netbirdio/netbird/client/firewall"
)

// Rule to handle management of rules
type Rule struct {
	id         string
	ip         net.IP
	ipLayer    gopacket.LayerType
	matchByIP  bool
	protoLayer gopacket.LayerType
	direction  fw.RuleDirection
	sPort      uint16
	dPort      uint16
	drop       bool
	comment    string

	udpHook func([]byte) bool
}

// GetRuleID returns the rule id
func (r *Rule) GetRuleID() string {
	return r.id
}
