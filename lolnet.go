package lolnet

import (
	"fmt"
	"net"
)

// Blood contains some IPs and subnets
type Blood struct {
	Nets []net.IPNet
	IPs  []net.IP
}

// Donor provides list provider
type Donor interface {
	Get() (*string, error)
}

// Recipient implements reciver funcs
type Recipient interface {
	Remove() error
	Add(*Blood) error
	Update(*Blood) error
}

// Doctor filters blood
type Doctor interface {
	Diagnose(*string) (*Blood, error)
}

func (blood Blood) String() string {
	return fmt.Sprintf("subnets: %d; ips: %d", len(blood.Nets), len(blood.IPs))
}
