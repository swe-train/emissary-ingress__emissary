// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/config/route/v3/scoped_route.proto

package envoy_config_route_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Specifies a routing scope, which associates a
// :ref:`Key<envoy_api_msg_config.route.v3.ScopedRouteConfiguration.Key>` to a
// :ref:`envoy_api_msg_config.route.v3.RouteConfiguration` (identified by its resource name).
//
// The HTTP connection manager builds up a table consisting of these Key to
// RouteConfiguration mappings, and looks up the RouteConfiguration to use per
// request according to the algorithm specified in the
// :ref:`scope_key_builder<envoy_api_field_extensions.filters.network.http_connection_manager.v3.ScopedRoutes.scope_key_builder>`
// assigned to the HttpConnectionManager.
//
// For example, with the following configurations (in YAML):
//
// HttpConnectionManager config:
//
// .. code::
//
//	...
//	scoped_routes:
//	  name: foo-scoped-routes
//	  scope_key_builder:
//	    fragments:
//	      - header_value_extractor:
//	          name: X-Route-Selector
//	          element_separator: ,
//	          element:
//	            separator: =
//	            key: vip
//
// ScopedRouteConfiguration resources (specified statically via
// :ref:`scoped_route_configurations_list<envoy_api_field_extensions.filters.network.http_connection_manager.v3.ScopedRoutes.scoped_route_configurations_list>`
// or obtained dynamically via SRDS):
//
// .. code::
//
//	(1)
//	 name: route-scope1
//	 route_configuration_name: route-config1
//	 key:
//	    fragments:
//	      - string_key: 172.10.10.20
//
//	(2)
//	 name: route-scope2
//	 route_configuration_name: route-config2
//	 key:
//	   fragments:
//	     - string_key: 172.20.20.30
//
// A request from a client such as:
//
// .. code::
//
//	GET / HTTP/1.1
//	Host: foo.com
//	X-Route-Selector: vip=172.10.10.20
//
// would result in the routing table defined by the `route-config1`
// RouteConfiguration being assigned to the HTTP request/stream.
type ScopedRouteConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the RouteConfiguration should be loaded on demand.
	OnDemand bool `protobuf:"varint,4,opt,name=on_demand,json=onDemand,proto3" json:"on_demand,omitempty"`
	// The name assigned to the routing scope.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource name to use for a :ref:`envoy_api_msg_service.discovery.v3.DiscoveryRequest` to an
	// RDS server to fetch the :ref:`envoy_api_msg_config.route.v3.RouteConfiguration` associated
	// with this scope.
	RouteConfigurationName string `protobuf:"bytes,2,opt,name=route_configuration_name,json=routeConfigurationName,proto3" json:"route_configuration_name,omitempty"`
	// The key to match against.
	Key *ScopedRouteConfiguration_Key `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ScopedRouteConfiguration) Reset() {
	*x = ScopedRouteConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_route_v3_scoped_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopedRouteConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopedRouteConfiguration) ProtoMessage() {}

func (x *ScopedRouteConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_route_v3_scoped_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopedRouteConfiguration.ProtoReflect.Descriptor instead.
func (*ScopedRouteConfiguration) Descriptor() ([]byte, []int) {
	return file_envoy_config_route_v3_scoped_route_proto_rawDescGZIP(), []int{0}
}

func (x *ScopedRouteConfiguration) GetOnDemand() bool {
	if x != nil {
		return x.OnDemand
	}
	return false
}

func (x *ScopedRouteConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScopedRouteConfiguration) GetRouteConfigurationName() string {
	if x != nil {
		return x.RouteConfigurationName
	}
	return ""
}

func (x *ScopedRouteConfiguration) GetKey() *ScopedRouteConfiguration_Key {
	if x != nil {
		return x.Key
	}
	return nil
}

// Specifies a key which is matched against the output of the
// :ref:`scope_key_builder<envoy_api_field_extensions.filters.network.http_connection_manager.v3.ScopedRoutes.scope_key_builder>`
// specified in the HttpConnectionManager. The matching is done per HTTP
// request and is dependent on the order of the fragments contained in the
// Key.
type ScopedRouteConfiguration_Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ordered set of fragments to match against. The order must match the
	// fragments in the corresponding
	// :ref:`scope_key_builder<envoy_api_field_extensions.filters.network.http_connection_manager.v3.ScopedRoutes.scope_key_builder>`.
	Fragments []*ScopedRouteConfiguration_Key_Fragment `protobuf:"bytes,1,rep,name=fragments,proto3" json:"fragments,omitempty"`
}

func (x *ScopedRouteConfiguration_Key) Reset() {
	*x = ScopedRouteConfiguration_Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_route_v3_scoped_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopedRouteConfiguration_Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopedRouteConfiguration_Key) ProtoMessage() {}

func (x *ScopedRouteConfiguration_Key) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_route_v3_scoped_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopedRouteConfiguration_Key.ProtoReflect.Descriptor instead.
func (*ScopedRouteConfiguration_Key) Descriptor() ([]byte, []int) {
	return file_envoy_config_route_v3_scoped_route_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ScopedRouteConfiguration_Key) GetFragments() []*ScopedRouteConfiguration_Key_Fragment {
	if x != nil {
		return x.Fragments
	}
	return nil
}

type ScopedRouteConfiguration_Key_Fragment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*ScopedRouteConfiguration_Key_Fragment_StringKey
	Type isScopedRouteConfiguration_Key_Fragment_Type `protobuf_oneof:"type"`
}

func (x *ScopedRouteConfiguration_Key_Fragment) Reset() {
	*x = ScopedRouteConfiguration_Key_Fragment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_route_v3_scoped_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopedRouteConfiguration_Key_Fragment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopedRouteConfiguration_Key_Fragment) ProtoMessage() {}

func (x *ScopedRouteConfiguration_Key_Fragment) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_route_v3_scoped_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopedRouteConfiguration_Key_Fragment.ProtoReflect.Descriptor instead.
func (*ScopedRouteConfiguration_Key_Fragment) Descriptor() ([]byte, []int) {
	return file_envoy_config_route_v3_scoped_route_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (m *ScopedRouteConfiguration_Key_Fragment) GetType() isScopedRouteConfiguration_Key_Fragment_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ScopedRouteConfiguration_Key_Fragment) GetStringKey() string {
	if x, ok := x.GetType().(*ScopedRouteConfiguration_Key_Fragment_StringKey); ok {
		return x.StringKey
	}
	return ""
}

type isScopedRouteConfiguration_Key_Fragment_Type interface {
	isScopedRouteConfiguration_Key_Fragment_Type()
}

type ScopedRouteConfiguration_Key_Fragment_StringKey struct {
	// A string to match against.
	StringKey string `protobuf:"bytes,1,opt,name=string_key,json=stringKey,proto3,oneof"`
}

func (*ScopedRouteConfiguration_Key_Fragment_StringKey) isScopedRouteConfiguration_Key_Fragment_Type() {
}

var File_envoy_config_route_v3_scoped_route_proto protoreflect.FileDescriptor

var file_envoy_config_route_v3_scoped_route_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x5f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x04, 0x0a,
	0x18, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6e, 0x5f,
	0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x6e,
	0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x18, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x16,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x1a, 0x92, 0x02, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12,
	0x64, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x2e, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x09, 0x66, 0x72, 0x61, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x73, 0x0a, 0x08, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4b,
	0x65, 0x79, 0x3a, 0x39, 0x9a, 0xc5, 0x88, 0x1e, 0x34, 0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4b, 0x65, 0x79, 0x2e, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0b, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x3a, 0x30, 0x9a, 0xc5, 0x88, 0x1e,
	0x2b, 0x0a, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x3a, 0x2c, 0x9a, 0xc5,
	0x88, 0x1e, 0x27, 0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x41, 0x0a, 0x23, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x33, 0x42, 0x10, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_route_v3_scoped_route_proto_rawDescOnce sync.Once
	file_envoy_config_route_v3_scoped_route_proto_rawDescData = file_envoy_config_route_v3_scoped_route_proto_rawDesc
)

func file_envoy_config_route_v3_scoped_route_proto_rawDescGZIP() []byte {
	file_envoy_config_route_v3_scoped_route_proto_rawDescOnce.Do(func() {
		file_envoy_config_route_v3_scoped_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_route_v3_scoped_route_proto_rawDescData)
	})
	return file_envoy_config_route_v3_scoped_route_proto_rawDescData
}

var file_envoy_config_route_v3_scoped_route_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_config_route_v3_scoped_route_proto_goTypes = []interface{}{
	(*ScopedRouteConfiguration)(nil),              // 0: envoy.config.route.v3.ScopedRouteConfiguration
	(*ScopedRouteConfiguration_Key)(nil),          // 1: envoy.config.route.v3.ScopedRouteConfiguration.Key
	(*ScopedRouteConfiguration_Key_Fragment)(nil), // 2: envoy.config.route.v3.ScopedRouteConfiguration.Key.Fragment
}
var file_envoy_config_route_v3_scoped_route_proto_depIdxs = []int32{
	1, // 0: envoy.config.route.v3.ScopedRouteConfiguration.key:type_name -> envoy.config.route.v3.ScopedRouteConfiguration.Key
	2, // 1: envoy.config.route.v3.ScopedRouteConfiguration.Key.fragments:type_name -> envoy.config.route.v3.ScopedRouteConfiguration.Key.Fragment
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_route_v3_scoped_route_proto_init() }
func file_envoy_config_route_v3_scoped_route_proto_init() {
	if File_envoy_config_route_v3_scoped_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_route_v3_scoped_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopedRouteConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_route_v3_scoped_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopedRouteConfiguration_Key); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_route_v3_scoped_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopedRouteConfiguration_Key_Fragment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_config_route_v3_scoped_route_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ScopedRouteConfiguration_Key_Fragment_StringKey)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_route_v3_scoped_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_route_v3_scoped_route_proto_goTypes,
		DependencyIndexes: file_envoy_config_route_v3_scoped_route_proto_depIdxs,
		MessageInfos:      file_envoy_config_route_v3_scoped_route_proto_msgTypes,
	}.Build()
	File_envoy_config_route_v3_scoped_route_proto = out.File
	file_envoy_config_route_v3_scoped_route_proto_rawDesc = nil
	file_envoy_config_route_v3_scoped_route_proto_goTypes = nil
	file_envoy_config_route_v3_scoped_route_proto_depIdxs = nil
}
