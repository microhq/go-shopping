// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shipping.proto

/*
Package shipping is a generated protocol buffer package.

It is generated from these files:
	shipping.proto

It has these top-level messages:
	ShippingCostRequest
	ShippingCostResponse
	MarkShippedRequest
	MarkShippedResponse
	ShippingStatusRequest
	ShippingStatusResponse
	ShippingStatus
	ShippingCost
*/
package shipping

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ShippingMethod int32

const (
	ShippingMethod_SM_UNKNOWN ShippingMethod = 0
	ShippingMethod_SM_USPS    ShippingMethod = 1
	ShippingMethod_SM_UPS     ShippingMethod = 2
	ShippingMethod_SM_FEDEX   ShippingMethod = 3
	ShippingMethod_SM_RAVEN   ShippingMethod = 4
)

var ShippingMethod_name = map[int32]string{
	0: "SM_UNKNOWN",
	1: "SM_USPS",
	2: "SM_UPS",
	3: "SM_FEDEX",
	4: "SM_RAVEN",
}
var ShippingMethod_value = map[string]int32{
	"SM_UNKNOWN": 0,
	"SM_USPS":    1,
	"SM_UPS":     2,
	"SM_FEDEX":   3,
	"SM_RAVEN":   4,
}

func (x ShippingMethod) String() string {
	return proto.EnumName(ShippingMethod_name, int32(x))
}
func (ShippingMethod) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ShippingCostRequest struct {
	Sku     string `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	ZipCode string `protobuf:"bytes,2,opt,name=zip_code,json=zipCode" json:"zip_code,omitempty"`
}

func (m *ShippingCostRequest) Reset()                    { *m = ShippingCostRequest{} }
func (m *ShippingCostRequest) String() string            { return proto.CompactTextString(m) }
func (*ShippingCostRequest) ProtoMessage()               {}
func (*ShippingCostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ShippingCostRequest) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *ShippingCostRequest) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

type ShippingCostResponse struct {
	ShippingCosts []*ShippingCost `protobuf:"bytes,1,rep,name=shipping_costs,json=shippingCosts" json:"shipping_costs,omitempty"`
}

func (m *ShippingCostResponse) Reset()                    { *m = ShippingCostResponse{} }
func (m *ShippingCostResponse) String() string            { return proto.CompactTextString(m) }
func (*ShippingCostResponse) ProtoMessage()               {}
func (*ShippingCostResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ShippingCostResponse) GetShippingCosts() []*ShippingCost {
	if m != nil {
		return m.ShippingCosts
	}
	return nil
}

type MarkShippedRequest struct {
	Sku            string         `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	OrderId        uint64         `protobuf:"varint,2,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	Note           string         `protobuf:"bytes,3,opt,name=note" json:"note,omitempty"`
	ShippingMethod ShippingMethod `protobuf:"varint,4,opt,name=shipping_method,json=shippingMethod,enum=shipping.ShippingMethod" json:"shipping_method,omitempty"`
}

func (m *MarkShippedRequest) Reset()                    { *m = MarkShippedRequest{} }
func (m *MarkShippedRequest) String() string            { return proto.CompactTextString(m) }
func (*MarkShippedRequest) ProtoMessage()               {}
func (*MarkShippedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MarkShippedRequest) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *MarkShippedRequest) GetOrderId() uint64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *MarkShippedRequest) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *MarkShippedRequest) GetShippingMethod() ShippingMethod {
	if m != nil {
		return m.ShippingMethod
	}
	return ShippingMethod_SM_UNKNOWN
}

type MarkShippedResponse struct {
	Success        bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	TrackingNumber string `protobuf:"bytes,2,opt,name=tracking_number,json=trackingNumber" json:"tracking_number,omitempty"`
}

func (m *MarkShippedResponse) Reset()                    { *m = MarkShippedResponse{} }
func (m *MarkShippedResponse) String() string            { return proto.CompactTextString(m) }
func (*MarkShippedResponse) ProtoMessage()               {}
func (*MarkShippedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MarkShippedResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *MarkShippedResponse) GetTrackingNumber() string {
	if m != nil {
		return m.TrackingNumber
	}
	return ""
}

type ShippingStatusRequest struct {
	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
}

func (m *ShippingStatusRequest) Reset()                    { *m = ShippingStatusRequest{} }
func (m *ShippingStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*ShippingStatusRequest) ProtoMessage()               {}
func (*ShippingStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ShippingStatusRequest) GetOrderId() uint64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type ShippingStatusResponse struct {
	ShippingStatus *ShippingStatus `protobuf:"bytes,1,opt,name=shipping_status,json=shippingStatus" json:"shipping_status,omitempty"`
}

func (m *ShippingStatusResponse) Reset()                    { *m = ShippingStatusResponse{} }
func (m *ShippingStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*ShippingStatusResponse) ProtoMessage()               {}
func (*ShippingStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ShippingStatusResponse) GetShippingStatus() *ShippingStatus {
	if m != nil {
		return m.ShippingStatus
	}
	return nil
}

type ShippingStatus struct {
	TrackingNumber string         `protobuf:"bytes,1,opt,name=tracking_number,json=trackingNumber" json:"tracking_number,omitempty"`
	ShippingMethod ShippingMethod `protobuf:"varint,2,opt,name=shipping_method,json=shippingMethod,enum=shipping.ShippingMethod" json:"shipping_method,omitempty"`
}

func (m *ShippingStatus) Reset()                    { *m = ShippingStatus{} }
func (m *ShippingStatus) String() string            { return proto.CompactTextString(m) }
func (*ShippingStatus) ProtoMessage()               {}
func (*ShippingStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ShippingStatus) GetTrackingNumber() string {
	if m != nil {
		return m.TrackingNumber
	}
	return ""
}

func (m *ShippingStatus) GetShippingMethod() ShippingMethod {
	if m != nil {
		return m.ShippingMethod
	}
	return ShippingMethod_SM_UNKNOWN
}

type ShippingCost struct {
	Method ShippingMethod `protobuf:"varint,1,opt,name=method,enum=shipping.ShippingMethod" json:"method,omitempty"`
	Price  int64          `protobuf:"varint,2,opt,name=price" json:"price,omitempty"`
}

func (m *ShippingCost) Reset()                    { *m = ShippingCost{} }
func (m *ShippingCost) String() string            { return proto.CompactTextString(m) }
func (*ShippingCost) ProtoMessage()               {}
func (*ShippingCost) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ShippingCost) GetMethod() ShippingMethod {
	if m != nil {
		return m.Method
	}
	return ShippingMethod_SM_UNKNOWN
}

func (m *ShippingCost) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*ShippingCostRequest)(nil), "shipping.ShippingCostRequest")
	proto.RegisterType((*ShippingCostResponse)(nil), "shipping.ShippingCostResponse")
	proto.RegisterType((*MarkShippedRequest)(nil), "shipping.MarkShippedRequest")
	proto.RegisterType((*MarkShippedResponse)(nil), "shipping.MarkShippedResponse")
	proto.RegisterType((*ShippingStatusRequest)(nil), "shipping.ShippingStatusRequest")
	proto.RegisterType((*ShippingStatusResponse)(nil), "shipping.ShippingStatusResponse")
	proto.RegisterType((*ShippingStatus)(nil), "shipping.ShippingStatus")
	proto.RegisterType((*ShippingCost)(nil), "shipping.ShippingCost")
	proto.RegisterEnum("shipping.ShippingMethod", ShippingMethod_name, ShippingMethod_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Shipping service

type ShippingClient interface {
	GetShippingCost(ctx context.Context, in *ShippingCostRequest, opts ...client.CallOption) (*ShippingCostResponse, error)
	MarkItemShipped(ctx context.Context, in *MarkShippedRequest, opts ...client.CallOption) (*MarkShippedResponse, error)
	GetShippingStatus(ctx context.Context, in *ShippingStatusRequest, opts ...client.CallOption) (*ShippingStatusResponse, error)
}

type shippingClient struct {
	c           client.Client
	serviceName string
}

func NewShippingClient(serviceName string, c client.Client) ShippingClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "shipping"
	}
	return &shippingClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingClient) GetShippingCost(ctx context.Context, in *ShippingCostRequest, opts ...client.CallOption) (*ShippingCostResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Shipping.GetShippingCost", in)
	out := new(ShippingCostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingClient) MarkItemShipped(ctx context.Context, in *MarkShippedRequest, opts ...client.CallOption) (*MarkShippedResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Shipping.MarkItemShipped", in)
	out := new(MarkShippedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingClient) GetShippingStatus(ctx context.Context, in *ShippingStatusRequest, opts ...client.CallOption) (*ShippingStatusResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Shipping.GetShippingStatus", in)
	out := new(ShippingStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Shipping service

type ShippingHandler interface {
	GetShippingCost(context.Context, *ShippingCostRequest, *ShippingCostResponse) error
	MarkItemShipped(context.Context, *MarkShippedRequest, *MarkShippedResponse) error
	GetShippingStatus(context.Context, *ShippingStatusRequest, *ShippingStatusResponse) error
}

func RegisterShippingHandler(s server.Server, hdlr ShippingHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Shipping{hdlr}, opts...))
}

type Shipping struct {
	ShippingHandler
}

func (h *Shipping) GetShippingCost(ctx context.Context, in *ShippingCostRequest, out *ShippingCostResponse) error {
	return h.ShippingHandler.GetShippingCost(ctx, in, out)
}

func (h *Shipping) MarkItemShipped(ctx context.Context, in *MarkShippedRequest, out *MarkShippedResponse) error {
	return h.ShippingHandler.MarkItemShipped(ctx, in, out)
}

func (h *Shipping) GetShippingStatus(ctx context.Context, in *ShippingStatusRequest, out *ShippingStatusResponse) error {
	return h.ShippingHandler.GetShippingStatus(ctx, in, out)
}

func init() { proto.RegisterFile("shipping.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x4d, 0x69, 0xc3, 0xeb, 0x48, 0xc3, 0xdb, 0x98, 0xc2, 0xc4, 0xa0, 0xca, 0x85, 0x8a,
	0xc3, 0x84, 0xca, 0x99, 0xc3, 0x18, 0x05, 0x4d, 0xa8, 0xa1, 0x4a, 0xd4, 0x32, 0x89, 0x43, 0xd4,
	0x25, 0x16, 0x8b, 0xaa, 0xc6, 0x21, 0x76, 0x2e, 0x13, 0x47, 0x7e, 0x04, 0x3f, 0x17, 0xc5, 0x71,
	0xd2, 0x84, 0x25, 0xa0, 0xdd, 0xfc, 0xf9, 0xbd, 0xf7, 0xf9, 0xfb, 0xec, 0x2f, 0x01, 0x83, 0xdf,
	0x44, 0x49, 0x12, 0xc5, 0xdf, 0xcf, 0x92, 0x94, 0x09, 0x86, 0x7a, 0x89, 0xed, 0xf7, 0x70, 0xe8,
	0xa9, 0xf5, 0x05, 0xe3, 0xc2, 0xa5, 0x3f, 0x32, 0xca, 0x05, 0x9a, 0xa0, 0xf1, 0x6d, 0x66, 0x91,
	0x09, 0x99, 0x3e, 0x72, 0xf3, 0x25, 0x3e, 0x03, 0xfd, 0x36, 0x4a, 0xfc, 0x80, 0x85, 0xd4, 0xea,
	0xc9, 0xed, 0xe1, 0x6d, 0x94, 0x5c, 0xb0, 0x90, 0xda, 0x2b, 0x38, 0x6a, 0x72, 0xf0, 0x84, 0xc5,
	0x9c, 0xe2, 0xbb, 0xfd, 0xb9, 0x7e, 0xc0, 0xb8, 0xe0, 0x16, 0x99, 0x68, 0xd3, 0xd1, 0xec, 0xf8,
	0xac, 0x92, 0xd3, 0x98, 0x7b, 0xcc, 0x6b, 0x88, 0xdb, 0xbf, 0x09, 0xe0, 0x62, 0x93, 0x6e, 0x65,
	0x0f, 0x0d, 0xff, 0x29, 0x8d, 0xa5, 0x21, 0x4d, 0xfd, 0x28, 0x94, 0xd2, 0xfa, 0xee, 0x50, 0xe2,
	0xcb, 0x10, 0x11, 0xfa, 0x31, 0x13, 0xd4, 0xd2, 0x64, 0xb7, 0x5c, 0xe3, 0x39, 0x8c, 0x2b, 0x59,
	0x3b, 0x2a, 0x6e, 0x58, 0x68, 0xf5, 0x27, 0x64, 0x6a, 0xcc, 0xac, 0xbb, 0xba, 0x16, 0xb2, 0xee,
	0x56, 0x3e, 0x0a, 0x6c, 0x5f, 0xc1, 0x61, 0x43, 0x99, 0x32, 0x6c, 0xc1, 0x90, 0x67, 0x41, 0x40,
	0x39, 0x97, 0xf2, 0x74, 0xb7, 0x84, 0xf8, 0x0a, 0xc6, 0x22, 0xdd, 0x04, 0xdb, 0xfc, 0xcc, 0x38,
	0xdb, 0x5d, 0xd3, 0x54, 0x5d, 0xa2, 0x51, 0x6e, 0x3b, 0x72, 0xd7, 0x9e, 0xc1, 0xd3, 0xf2, 0x6c,
	0x4f, 0x6c, 0x44, 0xc6, 0x4b, 0xdb, 0x75, 0x93, 0xa4, 0x61, 0xd2, 0xfe, 0x06, 0xc7, 0x7f, 0xcf,
	0x28, 0x41, 0x75, 0xab, 0x5c, 0x96, 0xe4, 0xec, 0xa8, 0xcd, 0xaa, 0x1a, 0xad, 0xac, 0x16, 0xd8,
	0xfe, 0x09, 0x46, 0xb3, 0xa3, 0xcd, 0x0b, 0x69, 0xf3, 0xd2, 0x76, 0xd1, 0xbd, 0x7b, 0x5e, 0xf4,
	0x1a, 0x0e, 0xea, 0x11, 0xc1, 0x37, 0x30, 0x50, 0x4c, 0xe4, 0x3f, 0x4c, 0xaa, 0x0f, 0x8f, 0xe0,
	0x61, 0x92, 0x46, 0x41, 0x11, 0x5a, 0xcd, 0x2d, 0xc0, 0xeb, 0xd5, 0xde, 0x55, 0xd1, 0x8f, 0x06,
	0x80, 0xb7, 0xf0, 0x57, 0xce, 0x67, 0xe7, 0xcb, 0x57, 0xc7, 0x7c, 0x80, 0x23, 0x18, 0xe6, 0xd8,
	0x5b, 0x7a, 0x26, 0x41, 0x80, 0x41, 0x0e, 0x96, 0x9e, 0xd9, 0xc3, 0x03, 0xd0, 0xbd, 0x85, 0xff,
	0x71, 0xfe, 0x61, 0x7e, 0x65, 0x6a, 0x0a, 0xb9, 0xe7, 0xeb, 0xb9, 0x63, 0xf6, 0x67, 0xbf, 0x7a,
	0xa0, 0x97, 0xbc, 0xb8, 0x84, 0xf1, 0x27, 0x2a, 0x1a, 0xf2, 0x4f, 0x3b, 0x92, 0x5f, 0xbc, 0xf1,
	0xc9, 0x8b, 0xae, 0xb2, 0x7a, 0x4e, 0x07, 0xc6, 0x79, 0xec, 0x2e, 0x05, 0xdd, 0xa9, 0xe8, 0xe1,
	0xf3, 0xfd, 0xc8, 0xdd, 0x6f, 0xe5, 0xe4, 0xb4, 0xa3, 0xaa, 0xf8, 0xd6, 0xf0, 0xa4, 0xa6, 0x50,
	0x3d, 0xef, 0xcb, 0xce, 0x68, 0x28, 0xd2, 0x49, 0x77, 0x43, 0xc1, 0x7b, 0x3d, 0x90, 0x7f, 0x99,
	0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x65, 0x87, 0x1d, 0x61, 0x77, 0x04, 0x00, 0x00,
}
