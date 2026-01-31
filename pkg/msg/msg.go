// Copyright 2016 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package msg

import (
	"net"
	"reflect"
)

const (
	TypeLogin              = 'o'
	TypeLoginResp          = '1'
	TypeNewProxy           = 'p'
	TypeNewProxyResp       = '2'
	TypeCloseProxy         = 'c'
	TypeNewWorkConn        = 'w'
	TypeReqWorkConn        = 'r'
	TypeStartWorkConn      = 's'
	TypeNewVisitorConn     = 'v'
	TypeNewVisitorConnResp = '3'
	TypePing               = 'h'
	TypePong               = '4'
	TypeUDPPacket          = 'u'
	TypeNatHoleVisitor     = 'i'
	TypeNatHoleClient      = 'n'
	TypeNatHoleResp        = 'm'
	TypeNatHoleSid         = '5'
	TypeNatHoleReport      = '6'
)

var msgTypeMap = map[byte]any{
	TypeLogin:              Login{},
	TypeLoginResp:          LoginResp{},
	TypeNewProxy:           NewProxy{},
	TypeNewProxyResp:       NewProxyResp{},
	TypeCloseProxy:         CloseProxy{},
	TypeNewWorkConn:        NewWorkConn{},
	TypeReqWorkConn:        ReqWorkConn{},
	TypeStartWorkConn:      StartWorkConn{},
	TypeNewVisitorConn:     NewVisitorConn{},
	TypeNewVisitorConnResp: NewVisitorConnResp{},
	TypePing:               Ping{},
	TypePong:               Pong{},
	TypeUDPPacket:          UDPPacket{},
	TypeNatHoleVisitor:     NatHoleVisitor{},
	TypeNatHoleClient:      NatHoleClient{},
	TypeNatHoleResp:        NatHoleResp{},
	TypeNatHoleSid:         NatHoleSid{},
	TypeNatHoleReport:      NatHoleReport{},
}

var TypeNameNatHoleResp = reflect.TypeOf(&NatHoleResp{}).Elem().Name()

type ClientSpec struct {
	// Due to the support of VirtualClient, frps needs to know the client type in order to
	// differentiate the processing logic.
	// Optional values: ssh-tunnel
	Type string `json:"tp,omitempty"`
	// If the value is true, the client will not require authentication.
	AlwaysAuthPass bool `json:"a_a_p,omitempty"`
}

// When frpc start, client send this message to login to server.
type Login struct {
	Version      string            `json:"a,omitempty"`
	Hostname     string            `json:"b,omitempty"`
	Os           string            `json:"c,omitempty"`
	Arch         string            `json:"d,omitempty"`
	User         string            `json:"e,omitempty"`
	PrivilegeKey string            `json:"f,omitempty"`
	Timestamp    int64             `json:"t,omitempty"`
	RunID        string            `json:"r,omitempty"`
	ClientID     string            `json:"c_i,omitempty"`
	Metas        map[string]string `json:"m,omitempty"`

	// Currently only effective for VirtualClient.
	ClientSpec ClientSpec `json:"cs,omitempty"`

	// Some global configures.
	PoolCount int `json:"pc,omitempty"`
}

type LoginResp struct {
	Version string `json:"a,omitempty"`
	RunID   string `json:"r,omitempty"`
	Error   string `json:"g,omitempty"`
}

// When frpc login success, send this message to frps for running a new proxy.
type NewProxy struct {
	ProxyName          string            `json:"pn,omitempty"`
	ProxyType          string            `json:"pt,omitempty"`
	UseEncryption      bool              `json:"ue,omitempty"`
	UseCompression     bool              `json:"uc,omitempty"`
	BandwidthLimit     string            `json:"bl,omitempty"`
	BandwidthLimitMode string            `json:"blm,omitempty"`
	Group              string            `json:"gg,omitempty"`
	GroupKey           string            `json:"gk,omitempty"`
	Metas              map[string]string `json:"m,omitempty"`
	Annotations        map[string]string `json:"an,omitempty"`

	// tcp and udp only
	RemotePort int `json:"remote_port,omitempty"`

	// http and https only
	CustomDomains     []string          `json:"cd,omitempty"`
	SubDomain         string            `json:"subd,omitempty"`
	Locations         []string          `json:"loc,omitempty"`
	HTTPUser          string            `json:"h_u,omitempty"`
	HTTPPwd           string            `json:"h_p,omitempty"`
	HostHeaderRewrite string            `json:"h_h_r,omitempty"`
	Headers           map[string]string `json:"hdrs,omitempty"`
	ResponseHeaders   map[string]string `json:"r_h,omitempty"`
	RouteByHTTPUser   string            `json:"r_b_h_u,omitempty"`

	// stcp, sudp, xtcp
	Sk         string   `json:"sk,omitempty"`
	AllowUsers []string `json:"a_us,omitempty"`

	// tcpmux
	Multiplexer string `json:"mpl,omitempty"`
}

type NewProxyResp struct {
	ProxyName  string `json:"pn,omitempty"`
	RemoteAddr string `json:"r_a,omitempty"`
	Error      string `json:"g,omitempty"`
}

type CloseProxy struct {
	ProxyName string `json:"pn,omitempty"`
}

type NewWorkConn struct {
	RunID        string `json:"r,omitempty"`
	PrivilegeKey string `json:"f,omitempty"`
	Timestamp    int64  `json:"t,omitempty"`
}

type ReqWorkConn struct{}

type StartWorkConn struct {
	ProxyName string `json:"pn,omitempty"`
	SrcAddr   string `json:"s_a,omitempty"`
	DstAddr   string `json:"d_a,omitempty"`
	SrcPort   uint16 `json:"s_p,omitempty"`
	DstPort   uint16 `json:"d_p,omitempty"`
	Error     string `json:"g,omitempty"`
}

type NewVisitorConn struct {
	RunID          string `json:"r,omitempty"`
	ProxyName      string `json:"pn,omitempty"`
	SignKey        string `json:"s_k,omitempty"`
	Timestamp      int64  `json:"t,omitempty"`
	UseEncryption  bool   `json:"ue,omitempty"`
	UseCompression bool   `json:"uc,omitempty"`
}

type NewVisitorConnResp struct {
	ProxyName string `json:"pn,omitempty"`
	Error     string `json:"g,omitempty"`
}

type Ping struct {
	PrivilegeKey string `json:"f,omitempty"`
	Timestamp    int64  `json:"t,omitempty"`
}

type Pong struct {
	Error string `json:"g,omitempty"`
}

type UDPPacket struct {
	Content    string       `json:"c,omitempty"`
	LocalAddr  *net.UDPAddr `json:"l,omitempty"`
	RemoteAddr *net.UDPAddr `json:"r,omitempty"`
}

type NatHoleVisitor struct {
	TransactionID string   `json:"tr_i,omitempty"`
	ProxyName     string   `json:"pn,omitempty"`
	PreCheck      bool     `json:"p_c,omitempty"`
	Protocol      string   `json:"pr,omitempty"`
	SignKey       string   `json:"s_k,omitempty"`
	Timestamp     int64    `json:"t,omitempty"`
	MappedAddrs   []string `json:"m_a,omitempty"`
	AssistedAddrs []string `json:"a_a,omitempty"`
}

type NatHoleClient struct {
	TransactionID string   `json:"tr_i,omitempty"`
	ProxyName     string   `json:"pn,omitempty"`
	Sid           string   `json:"sid,omitempty"`
	MappedAddrs   []string `json:"m_a,omitempty"`
	AssistedAddrs []string `json:"a_a,omitempty"`
}

type PortsRange struct {
	From int `json:"from,omitempty"`
	To   int `json:"to,omitempty"`
}

type NatHoleDetectBehavior struct {
	Role              string       `json:"rl,omitempty"` // sender or receiver
	Mode              int          `json:"md,omitempty"` // 0, 1, 2...
	TTL               int          `json:"ttl,omitempty"`
	SendDelayMs       int          `json:"s_d_m,omitempty"`
	ReadTimeoutMs     int          `json:"r_t,omitempty"`
	CandidatePorts    []PortsRange `json:"c_p,omitempty"`
	SendRandomPorts   int          `json:"s_r_p,omitempty"`
	ListenRandomPorts int          `json:"l_r_p,omitempty"`
}

type NatHoleResp struct {
	TransactionID  string                `json:"tr_i,omitempty"`
	Sid            string                `json:"sid,omitempty"`
	Protocol       string                `json:"pr,omitempty"`
	CandidateAddrs []string              `json:"c_a,omitempty"`
	AssistedAddrs  []string              `json:"a_a,omitempty"`
	DetectBehavior NatHoleDetectBehavior `json:"d_b,omitempty"`
	Error          string                `json:"g,omitempty"`
}

type NatHoleSid struct {
	TransactionID string `json:"tr_i,omitempty"`
	Sid           string `json:"sid,omitempty"`
	Response      bool   `json:"rsp,omitempty"`
	Nonce         string `json:"nnc,omitempty"`
}

type NatHoleReport struct {
	Sid     string `json:"sid,omitempty"`
	Success bool   `json:"suc,omitempty"`
}
