// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vessel.proto

package vessel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0c0e45ee319d513, []int{0}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0c0e45ee319d513, []int{1}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0c0e45ee319d513, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Specification)(nil), "vessel.Specification")
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
	proto.RegisterType((*Response)(nil), "vessel.Response")
}

func init() { proto.RegisterFile("vessel.proto", fileDescriptor_e0c0e45ee319d513) }

var fileDescriptor_e0c0e45ee319d513 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x4b, 0xc3, 0x50,
	0x10, 0x34, 0x69, 0x9b, 0x26, 0xab, 0x29, 0xb2, 0x20, 0x3c, 0x8b, 0x42, 0xc8, 0x41, 0x72, 0xea,
	0xa1, 0xde, 0xbc, 0x79, 0x11, 0xd4, 0x5b, 0x0a, 0x7a, 0x11, 0xca, 0x6b, 0xb2, 0xea, 0x42, 0xbe,
	0xc8, 0x0b, 0x69, 0xfd, 0x37, 0xfe, 0x54, 0x61, 0x93, 0x54, 0x2a, 0xd2, 0xdb, 0xce, 0x4c, 0x66,
	0x98, 0xcc, 0x83, 0xb3, 0x96, 0x8c, 0xa1, 0x6c, 0x51, 0xd5, 0x65, 0x53, 0xa2, 0xd3, 0xa1, 0xf0,
	0x09, 0xfc, 0x55, 0x45, 0x09, 0xbf, 0x73, 0xa2, 0x1b, 0x2e, 0x0b, 0x9c, 0x83, 0x9b, 0xe8, 0x4a,
	0x27, 0xdc, 0x7c, 0x29, 0x2b, 0xb0, 0xa2, 0x49, 0xbc, 0xc7, 0x78, 0x0d, 0x90, 0xeb, 0xdd, 0x7a,
	0x4b, 0xfc, 0xf1, 0xd9, 0x28, 0x5b, 0x54, 0x2f, 0xd7, 0xbb, 0x57, 0x21, 0xc2, 0x6f, 0x0b, 0x9c,
	0x17, 0x89, 0xc5, 0x19, 0xd8, 0x9c, 0x8a, 0xdf, 0x8b, 0x6d, 0x4e, 0x0f, 0x52, 0xed, 0xa3, 0xa9,
	0xa3, 0x3f, 0xa9, 0x88, 0x30, 0x2e, 0x74, 0x4e, 0x6a, 0x2c, 0x61, 0x72, 0xe3, 0x15, 0x78, 0xba,
	0xd5, 0x9c, 0xe9, 0x4d, 0x46, 0x6a, 0x12, 0x58, 0x91, 0x1b, 0xff, 0x12, 0x78, 0x09, 0x6e, 0xb9,
	0x2d, 0xa8, 0x5e, 0x73, 0xaa, 0x1c, 0x71, 0x4d, 0x05, 0x3f, 0xa6, 0xe1, 0x1b, 0xb8, 0x31, 0x99,
	0xaa, 0x2c, 0x0c, 0xe1, 0x0d, 0xf4, 0x23, 0x48, 0xcf, 0xd3, 0xe5, 0x6c, 0xd1, 0x2f, 0xd4, 0xfd,
	0x43, 0xdc, 0xab, 0x18, 0xc1, 0xb4, 0xbb, 0x8c, 0xb2, 0x83, 0xd1, 0x3f, 0x1f, 0x0e, 0xf2, 0xf2,
	0x19, 0xfc, 0x8e, 0x5a, 0x51, 0xdd, 0x72, 0x42, 0x78, 0x07, 0xfe, 0x03, 0x17, 0xe9, 0xfd, 0xbe,
	0xda, 0xc5, 0x60, 0x3d, 0x18, 0x7d, 0x7e, 0x3e, 0xd0, 0x43, 0xb9, 0xf0, 0x64, 0xe3, 0xc8, 0x43,
	0xdd, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x81, 0x5c, 0xb1, 0xb8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}
