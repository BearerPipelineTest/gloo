// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *Gateway) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Gateway)
	if !ok {
		that2, ok := that.(Gateway)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetSsl() != target.GetSsl() {
		return false
	}

	if strings.Compare(m.GetBindAddress(), target.GetBindAddress()) != 0 {
		return false
	}

	if m.GetBindPort() != target.GetBindPort() {
		return false
	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(equality.Equalizer); ok {
		if !h.Equal(target.GetNamespacedStatuses()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetNamespacedStatuses(), target.GetNamespacedStatuses()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetUseProxyProto()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUseProxyProto()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUseProxyProto(), target.GetUseProxyProto()) {
			return false
		}
	}

	if len(m.GetProxyNames()) != len(target.GetProxyNames()) {
		return false
	}
	for idx, v := range m.GetProxyNames() {

		if strings.Compare(v, target.GetProxyNames()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetRouteOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRouteOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRouteOptions(), target.GetRouteOptions()) {
			return false
		}
	}

	switch m.GatewayType.(type) {

	case *Gateway_HttpGateway:
		if _, ok := target.GatewayType.(*Gateway_HttpGateway); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttpGateway()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpGateway()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttpGateway(), target.GetHttpGateway()) {
				return false
			}
		}

	case *Gateway_TcpGateway:
		if _, ok := target.GatewayType.(*Gateway_TcpGateway); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTcpGateway()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpGateway()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTcpGateway(), target.GetTcpGateway()) {
				return false
			}
		}

	case *Gateway_HybridGateway:
		if _, ok := target.GatewayType.(*Gateway_HybridGateway); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHybridGateway()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHybridGateway()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHybridGateway(), target.GetHybridGateway()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.GatewayType != target.GatewayType {
			return false
		}
	}

	return true
}

// Equal function
func (m *TcpGateway) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpGateway)
	if !ok {
		that2, ok := that.(TcpGateway)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetTcpHosts()) != len(target.GetTcpHosts()) {
		return false
	}
	for idx, v := range m.GetTcpHosts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpHosts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetTcpHosts()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *HybridGateway) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HybridGateway)
	if !ok {
		that2, ok := that.(HybridGateway)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetMatchedGateways()) != len(target.GetMatchedGateways()) {
		return false
	}
	for idx, v := range m.GetMatchedGateways() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMatchedGateways()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMatchedGateways()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetDelegatedHttpGateways()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDelegatedHttpGateways()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDelegatedHttpGateways(), target.GetDelegatedHttpGateways()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *DelegatedHttpGateway) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DelegatedHttpGateway)
	if !ok {
		that2, ok := that.(DelegatedHttpGateway)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.SelectionType.(type) {

	case *DelegatedHttpGateway_Ref:
		if _, ok := target.SelectionType.(*DelegatedHttpGateway_Ref); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRef()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRef(), target.GetRef()) {
				return false
			}
		}

	case *DelegatedHttpGateway_Selector:
		if _, ok := target.SelectionType.(*DelegatedHttpGateway_Selector); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSelector()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelector()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSelector(), target.GetSelector()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.SelectionType != target.SelectionType {
			return false
		}
	}

	return true
}

// Equal function
func (m *MatchedGateway) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MatchedGateway)
	if !ok {
		that2, ok := that.(MatchedGateway)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMatcher()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMatcher()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMatcher(), target.GetMatcher()) {
			return false
		}
	}

	switch m.GatewayType.(type) {

	case *MatchedGateway_HttpGateway:
		if _, ok := target.GatewayType.(*MatchedGateway_HttpGateway); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttpGateway()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpGateway()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttpGateway(), target.GetHttpGateway()) {
				return false
			}
		}

	case *MatchedGateway_TcpGateway:
		if _, ok := target.GatewayType.(*MatchedGateway_TcpGateway); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTcpGateway()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpGateway()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTcpGateway(), target.GetTcpGateway()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.GatewayType != target.GatewayType {
			return false
		}
	}

	return true
}

// Equal function
func (m *Matcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Matcher)
	if !ok {
		that2, ok := that.(Matcher)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetSslConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSslConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSslConfig(), target.GetSslConfig()) {
			return false
		}
	}

	if len(m.GetSourcePrefixRanges()) != len(target.GetSourcePrefixRanges()) {
		return false
	}
	for idx, v := range m.GetSourcePrefixRanges() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSourcePrefixRanges()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSourcePrefixRanges()[idx]) {
				return false
			}
		}

	}

	return true
}
