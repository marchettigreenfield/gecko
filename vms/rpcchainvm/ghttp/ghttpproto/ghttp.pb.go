// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ghttp.proto

package ghttpproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Userinfo struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PasswordSet          bool     `protobuf:"varint,3,opt,name=passwordSet,proto3" json:"passwordSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Userinfo) Reset()         { *m = Userinfo{} }
func (m *Userinfo) String() string { return proto.CompactTextString(m) }
func (*Userinfo) ProtoMessage()    {}
func (*Userinfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e26bba3d5e69055f, []int{0}
}

func (m *Userinfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Userinfo.Unmarshal(m, b)
}
func (m *Userinfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Userinfo.Marshal(b, m, deterministic)
}
func (m *Userinfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Userinfo.Merge(m, src)
}
func (m *Userinfo) XXX_Size() int {
	return xxx_messageInfo_Userinfo.Size(m)
}
func (m *Userinfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Userinfo.DiscardUnknown(m)
}

var xxx_messageInfo_Userinfo proto.InternalMessageInfo

func (m *Userinfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Userinfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Userinfo) GetPasswordSet() bool {
	if m != nil {
		return m.PasswordSet
	}
	return false
}

type URL struct {
	Scheme               string    `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Opaque               string    `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	User                 *Userinfo `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Host                 string    `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Path                 string    `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	RawPath              string    `protobuf:"bytes,6,opt,name=rawPath,proto3" json:"rawPath,omitempty"`
	ForceQuery           bool      `protobuf:"varint,7,opt,name=forceQuery,proto3" json:"forceQuery,omitempty"`
	RawQuery             string    `protobuf:"bytes,8,opt,name=rawQuery,proto3" json:"rawQuery,omitempty"`
	Fragment             string    `protobuf:"bytes,9,opt,name=fragment,proto3" json:"fragment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *URL) Reset()         { *m = URL{} }
func (m *URL) String() string { return proto.CompactTextString(m) }
func (*URL) ProtoMessage()    {}
func (*URL) Descriptor() ([]byte, []int) {
	return fileDescriptor_e26bba3d5e69055f, []int{1}
}

func (m *URL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URL.Unmarshal(m, b)
}
func (m *URL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URL.Marshal(b, m, deterministic)
}
func (m *URL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URL.Merge(m, src)
}
func (m *URL) XXX_Size() int {
	return xxx_messageInfo_URL.Size(m)
}
func (m *URL) XXX_DiscardUnknown() {
	xxx_messageInfo_URL.DiscardUnknown(m)
}

var xxx_messageInfo_URL proto.InternalMessageInfo

func (m *URL) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *URL) GetOpaque() string {
	if m != nil {
		return m.Opaque
	}
	return ""
}

func (m *URL) GetUser() *Userinfo {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *URL) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *URL) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *URL) GetRawPath() string {
	if m != nil {
		return m.RawPath
	}
	return ""
}

func (m *URL) GetForceQuery() bool {
	if m != nil {
		return m.ForceQuery
	}
	return false
}

func (m *URL) GetRawQuery() string {
	if m != nil {
		return m.RawQuery
	}
	return ""
}

func (m *URL) GetFragment() string {
	if m != nil {
		return m.Fragment
	}
	return ""
}

type Element struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Element) Reset()         { *m = Element{} }
func (m *Element) String() string { return proto.CompactTextString(m) }
func (*Element) ProtoMessage()    {}
func (*Element) Descriptor() ([]byte, []int) {
	return fileDescriptor_e26bba3d5e69055f, []int{2}
}

func (m *Element) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Element.Unmarshal(m, b)
}
func (m *Element) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Element.Marshal(b, m, deterministic)
}
func (m *Element) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Element.Merge(m, src)
}
func (m *Element) XXX_Size() int {
	return xxx_messageInfo_Element.Size(m)
}
func (m *Element) XXX_DiscardUnknown() {
	xxx_messageInfo_Element.DiscardUnknown(m)
}

var xxx_messageInfo_Element proto.InternalMessageInfo

func (m *Element) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Element) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Certificates struct {
	Cert                 [][]byte `protobuf:"bytes,1,rep,name=cert,proto3" json:"cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Certificates) Reset()         { *m = Certificates{} }
func (m *Certificates) String() string { return proto.CompactTextString(m) }
func (*Certificates) ProtoMessage()    {}
func (*Certificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_e26bba3d5e69055f, []int{3}
}

func (m *Certificates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificates.Unmarshal(m, b)
}
func (m *Certificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificates.Marshal(b, m, deterministic)
}
func (m *Certificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificates.Merge(m, src)
}
func (m *Certificates) XXX_Size() int {
	return xxx_messageInfo_Certificates.Size(m)
}
func (m *Certificates) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificates.DiscardUnknown(m)
}

var xxx_messageInfo_Certificates proto.InternalMessageInfo

func (m *Certificates) GetCert() [][]byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

type ConnectionState struct {
	Version                     uint32          `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	HandshakeComplete           bool            `protobuf:"varint,2,opt,name=handshakeComplete,proto3" json:"handshakeComplete,omitempty"`
	DidResume                   bool            `protobuf:"varint,3,opt,name=didResume,proto3" json:"didResume,omitempty"`
	CipherSuite                 uint32          `protobuf:"varint,4,opt,name=cipherSuite,proto3" json:"cipherSuite,omitempty"`
	NegotiatedProtocol          string          `protobuf:"bytes,5,opt,name=negotiatedProtocol,proto3" json:"negotiatedProtocol,omitempty"`
	NegotiatedProtocolIsMutual  bool            `protobuf:"varint,6,opt,name=negotiatedProtocolIsMutual,proto3" json:"negotiatedProtocolIsMutual,omitempty"`
	ServerName                  string          `protobuf:"bytes,7,opt,name=serverName,proto3" json:"serverName,omitempty"`
	PeerCertificates            *Certificates   `protobuf:"bytes,8,opt,name=peerCertificates,proto3" json:"peerCertificates,omitempty"`
	VerifiedChains              []*Certificates `protobuf:"bytes,9,rep,name=verifiedChains,proto3" json:"verifiedChains,omitempty"`
	SignedCertificateTimestamps [][]byte        `protobuf:"bytes,10,rep,name=signedCertificateTimestamps,proto3" json:"signedCertificateTimestamps,omitempty"`
	OcspResponse                []byte          `protobuf:"bytes,11,opt,name=ocspResponse,proto3" json:"ocspResponse,omitempty"`
	TlsUnique                   []byte          `protobuf:"bytes,12,opt,name=tlsUnique,proto3" json:"tlsUnique,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}        `json:"-"`
	XXX_unrecognized            []byte          `json:"-"`
	XXX_sizecache               int32           `json:"-"`
}

func (m *ConnectionState) Reset()         { *m = ConnectionState{} }
func (m *ConnectionState) String() string { return proto.CompactTextString(m) }
func (*ConnectionState) ProtoMessage()    {}
func (*ConnectionState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e26bba3d5e69055f, []int{4}
}

func (m *ConnectionState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionState.Unmarshal(m, b)
}
func (m *ConnectionState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionState.Marshal(b, m, deterministic)
}
func (m *ConnectionState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionState.Merge(m, src)
}
func (m *ConnectionState) XXX_Size() int {
	return xxx_messageInfo_ConnectionState.Size(m)
}
func (m *ConnectionState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionState.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionState proto.InternalMessageInfo

func (m *ConnectionState) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ConnectionState) GetHandshakeComplete() bool {
	if m != nil {
		return m.HandshakeComplete
	}
	return false
}

func (m *ConnectionState) GetDidResume() bool {
	if m != nil {
		return m.DidResume
	}
	return false
}

func (m *ConnectionState) GetCipherSuite() uint32 {
	if m != nil {
		return m.CipherSuite
	}
	return 0
}

func (m *ConnectionState) GetNegotiatedProtocol() string {
	if m != nil {
		return m.NegotiatedProtocol
	}
	return ""
}

func (m *ConnectionState) GetNegotiatedProtocolIsMutual() bool {
	if m != nil {
		return m.NegotiatedProtocolIsMutual
	}
	return false
}

func (m *ConnectionState) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *ConnectionState) GetPeerCertificates() *Certificates {
	if m != nil {
		return m.PeerCertificates
	}
	return nil
}

func (m *ConnectionState) GetVerifiedChains() []*Certificates {
	if m != nil {
		return m.VerifiedChains
	}
	return nil
}

func (m *ConnectionState) GetSignedCertificateTimestamps() [][]byte {
	if m != nil {
		return m.SignedCertificateTimestamps
	}
	return nil
}

func (m *ConnectionState) GetOcspResponse() []byte {
	if m != nil {
		return m.OcspResponse
	}
	return nil
}

func (m *ConnectionState) GetTlsUnique() []byte {
	if m != nil {
		return m.TlsUnique
	}
	return nil
}

type Request struct {
	Method               string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Url                  *URL             `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Proto                string           `protobuf:"bytes,3,opt,name=proto,proto3" json:"proto,omitempty"`
	ProtoMajor           int32            `protobuf:"varint,4,opt,name=protoMajor,proto3" json:"protoMajor,omitempty"`
	ProtoMinor           int32            `protobuf:"varint,5,opt,name=protoMinor,proto3" json:"protoMinor,omitempty"`
	Header               []*Element       `protobuf:"bytes,6,rep,name=header,proto3" json:"header,omitempty"`
	Body                 uint32           `protobuf:"varint,7,opt,name=body,proto3" json:"body,omitempty"`
	ContentLength        int64            `protobuf:"varint,8,opt,name=contentLength,proto3" json:"contentLength,omitempty"`
	TransferEncoding     []string         `protobuf:"bytes,9,rep,name=transferEncoding,proto3" json:"transferEncoding,omitempty"`
	Host                 string           `protobuf:"bytes,10,opt,name=host,proto3" json:"host,omitempty"`
	Form                 []*Element       `protobuf:"bytes,11,rep,name=form,proto3" json:"form,omitempty"`
	PostForm             []*Element       `protobuf:"bytes,12,rep,name=postForm,proto3" json:"postForm,omitempty"`
	TrailerKeys          []string         `protobuf:"bytes,13,rep,name=trailerKeys,proto3" json:"trailerKeys,omitempty"`
	RemoteAddr           string           `protobuf:"bytes,14,opt,name=remoteAddr,proto3" json:"remoteAddr,omitempty"`
	RequestURI           string           `protobuf:"bytes,15,opt,name=requestURI,proto3" json:"requestURI,omitempty"`
	Tls                  *ConnectionState `protobuf:"bytes,16,opt,name=tls,proto3" json:"tls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e26bba3d5e69055f, []int{5}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetUrl() *URL {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *Request) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

func (m *Request) GetProtoMajor() int32 {
	if m != nil {
		return m.ProtoMajor
	}
	return 0
}

func (m *Request) GetProtoMinor() int32 {
	if m != nil {
		return m.ProtoMinor
	}
	return 0
}

func (m *Request) GetHeader() []*Element {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetBody() uint32 {
	if m != nil {
		return m.Body
	}
	return 0
}

func (m *Request) GetContentLength() int64 {
	if m != nil {
		return m.ContentLength
	}
	return 0
}

func (m *Request) GetTransferEncoding() []string {
	if m != nil {
		return m.TransferEncoding
	}
	return nil
}

func (m *Request) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Request) GetForm() []*Element {
	if m != nil {
		return m.Form
	}
	return nil
}

func (m *Request) GetPostForm() []*Element {
	if m != nil {
		return m.PostForm
	}
	return nil
}

func (m *Request) GetTrailerKeys() []string {
	if m != nil {
		return m.TrailerKeys
	}
	return nil
}

func (m *Request) GetRemoteAddr() string {
	if m != nil {
		return m.RemoteAddr
	}
	return ""
}

func (m *Request) GetRequestURI() string {
	if m != nil {
		return m.RequestURI
	}
	return ""
}

func (m *Request) GetTls() *ConnectionState {
	if m != nil {
		return m.Tls
	}
	return nil
}

type HTTPRequest struct {
	ResponseWriter       uint32   `protobuf:"varint,1,opt,name=responseWriter,proto3" json:"responseWriter,omitempty"`
	Request              *Request `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTPRequest) Reset()         { *m = HTTPRequest{} }
func (m *HTTPRequest) String() string { return proto.CompactTextString(m) }
func (*HTTPRequest) ProtoMessage()    {}
func (*HTTPRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e26bba3d5e69055f, []int{6}
}

func (m *HTTPRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPRequest.Unmarshal(m, b)
}
func (m *HTTPRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPRequest.Marshal(b, m, deterministic)
}
func (m *HTTPRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPRequest.Merge(m, src)
}
func (m *HTTPRequest) XXX_Size() int {
	return xxx_messageInfo_HTTPRequest.Size(m)
}
func (m *HTTPRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPRequest proto.InternalMessageInfo

func (m *HTTPRequest) GetResponseWriter() uint32 {
	if m != nil {
		return m.ResponseWriter
	}
	return 0
}

func (m *HTTPRequest) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

type HTTPResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTPResponse) Reset()         { *m = HTTPResponse{} }
func (m *HTTPResponse) String() string { return proto.CompactTextString(m) }
func (*HTTPResponse) ProtoMessage()    {}
func (*HTTPResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e26bba3d5e69055f, []int{7}
}

func (m *HTTPResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPResponse.Unmarshal(m, b)
}
func (m *HTTPResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPResponse.Marshal(b, m, deterministic)
}
func (m *HTTPResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPResponse.Merge(m, src)
}
func (m *HTTPResponse) XXX_Size() int {
	return xxx_messageInfo_HTTPResponse.Size(m)
}
func (m *HTTPResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Userinfo)(nil), "ghttpproto.Userinfo")
	proto.RegisterType((*URL)(nil), "ghttpproto.URL")
	proto.RegisterType((*Element)(nil), "ghttpproto.Element")
	proto.RegisterType((*Certificates)(nil), "ghttpproto.Certificates")
	proto.RegisterType((*ConnectionState)(nil), "ghttpproto.ConnectionState")
	proto.RegisterType((*Request)(nil), "ghttpproto.Request")
	proto.RegisterType((*HTTPRequest)(nil), "ghttpproto.HTTPRequest")
	proto.RegisterType((*HTTPResponse)(nil), "ghttpproto.HTTPResponse")
}

func init() { proto.RegisterFile("ghttp.proto", fileDescriptor_e26bba3d5e69055f) }

var fileDescriptor_e26bba3d5e69055f = []byte{
	// 819 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdf, 0x8f, 0x1b, 0x35,
	0x10, 0x56, 0xba, 0xb9, 0xfc, 0x98, 0x4d, 0xee, 0x0e, 0x53, 0x81, 0x75, 0x45, 0x28, 0xac, 0x10,
	0x44, 0x40, 0x0f, 0x29, 0x7d, 0x44, 0x42, 0x45, 0xa1, 0xa8, 0x15, 0x57, 0x74, 0xf8, 0x2e, 0xe2,
	0xd9, 0x5d, 0x4f, 0xb2, 0xa6, 0xbb, 0xf6, 0xd6, 0xf6, 0xe6, 0x74, 0xff, 0x13, 0xaf, 0xfc, 0x71,
	0xbc, 0x21, 0x7b, 0x77, 0x93, 0x4d, 0xd3, 0xde, 0xdb, 0xcc, 0xf7, 0x8d, 0xc7, 0x9e, 0xf9, 0x66,
	0x0c, 0xf1, 0x26, 0x73, 0xae, 0xbc, 0x2c, 0x8d, 0x76, 0x9a, 0x40, 0x70, 0x82, 0x9d, 0x08, 0x18,
	0xad, 0x2c, 0x1a, 0xa9, 0xd6, 0x9a, 0x5c, 0xc0, 0xa8, 0xb2, 0x68, 0x14, 0x2f, 0x90, 0xf6, 0x66,
	0xbd, 0xf9, 0x98, 0xed, 0x7c, 0xcf, 0x95, 0xdc, 0xda, 0x3b, 0x6d, 0x04, 0x7d, 0x54, 0x73, 0xad,
	0x4f, 0x66, 0x10, 0xb7, 0xf6, 0x0d, 0x3a, 0x1a, 0xcd, 0x7a, 0xf3, 0x11, 0xeb, 0x42, 0xc9, 0x7f,
	0x3d, 0x88, 0x56, 0xec, 0x8a, 0x7c, 0x06, 0x03, 0x9b, 0x66, 0xb8, 0xcb, 0xdf, 0x78, 0x1e, 0xd7,
	0x25, 0x7f, 0x57, 0x61, 0x93, 0xbb, 0xf1, 0xc8, 0x1c, 0xfa, 0xfe, 0x05, 0x21, 0x65, 0xbc, 0x78,
	0x7c, 0xb9, 0x7f, 0xf8, 0x65, 0xfb, 0x6a, 0x16, 0x22, 0x08, 0x81, 0x7e, 0xa6, 0xad, 0xa3, 0xfd,
	0x70, 0x3e, 0xd8, 0x1e, 0x2b, 0xb9, 0xcb, 0xe8, 0x49, 0x8d, 0x79, 0x9b, 0x50, 0x18, 0x1a, 0x7e,
	0x77, 0xed, 0xe1, 0x41, 0x80, 0x5b, 0x97, 0x7c, 0x09, 0xb0, 0xd6, 0x26, 0xc5, 0x3f, 0x2b, 0x34,
	0xf7, 0x74, 0x18, 0x8a, 0xe8, 0x20, 0xbe, 0x03, 0x86, 0xdf, 0xd5, 0xec, 0xa8, 0xee, 0x40, 0xeb,
	0x7b, 0x6e, 0x6d, 0xf8, 0xa6, 0x40, 0xe5, 0xe8, 0xb8, 0xe6, 0x5a, 0x3f, 0x79, 0x06, 0xc3, 0x17,
	0x39, 0x7a, 0x93, 0x9c, 0x43, 0xf4, 0x16, 0xef, 0x9b, 0xda, 0xbd, 0xe9, 0x0b, 0xdf, 0xf2, 0xbc,
	0x42, 0x4b, 0x1f, 0xcd, 0x22, 0x5f, 0x78, 0xed, 0x25, 0x09, 0x4c, 0x96, 0x68, 0x9c, 0x5c, 0xcb,
	0x94, 0x3b, 0xb4, 0xbe, 0x94, 0x14, 0x8d, 0xa3, 0xbd, 0x59, 0x34, 0x9f, 0xb0, 0x60, 0x27, 0xff,
	0xf6, 0xe1, 0x6c, 0xa9, 0x95, 0xc2, 0xd4, 0x49, 0xad, 0x6e, 0x1c, 0x77, 0xe8, 0xcb, 0xdb, 0xa2,
	0xb1, 0x52, 0xab, 0x70, 0xcb, 0x94, 0xb5, 0x2e, 0xf9, 0x01, 0x3e, 0xc9, 0xb8, 0x12, 0x36, 0xe3,
	0x6f, 0x71, 0xa9, 0x8b, 0x32, 0x47, 0x57, 0x77, 0x7b, 0xc4, 0x8e, 0x09, 0xf2, 0x05, 0x8c, 0x85,
	0x14, 0x0c, 0x6d, 0x55, 0x60, 0x23, 0xe8, 0x1e, 0xf0, 0x82, 0xa7, 0xb2, 0xcc, 0xd0, 0xdc, 0x54,
	0xd2, 0x61, 0xe8, 0xf9, 0x94, 0x75, 0x21, 0x72, 0x09, 0x44, 0xe1, 0x46, 0x3b, 0xc9, 0x1d, 0x8a,
	0x6b, 0x2f, 0x58, 0xaa, 0xf3, 0x46, 0x88, 0x0f, 0x30, 0xe4, 0x67, 0xb8, 0x38, 0x46, 0x5f, 0xd9,
	0xd7, 0x95, 0xab, 0x78, 0x1e, 0x94, 0x1a, 0xb1, 0x07, 0x22, 0xbc, 0x78, 0x16, 0xcd, 0x16, 0xcd,
	0x1f, 0x7e, 0x78, 0x87, 0xe1, 0x9e, 0x0e, 0x42, 0x7e, 0x85, 0xf3, 0x12, 0xd1, 0x74, 0x7b, 0x1a,
	0x44, 0x8c, 0x17, 0xb4, 0x3b, 0x54, 0x5d, 0x9e, 0x1d, 0x9d, 0x20, 0xcf, 0xe1, 0x74, 0x8b, 0x46,
	0xae, 0x25, 0x8a, 0x65, 0xc6, 0xa5, 0xb2, 0x74, 0x3c, 0x8b, 0x1e, 0xcc, 0xf1, 0x5e, 0x3c, 0x79,
	0x0e, 0x4f, 0xac, 0xdc, 0x28, 0x14, 0x9d, 0xa8, 0x5b, 0x59, 0xa0, 0x75, 0xbc, 0x28, 0x2d, 0x85,
	0x20, 0xef, 0x43, 0x21, 0x24, 0x81, 0x89, 0x4e, 0x6d, 0xc9, 0xd0, 0x96, 0x5a, 0x59, 0xa4, 0xf1,
	0xac, 0x37, 0x9f, 0xb0, 0x03, 0xcc, 0xab, 0xe7, 0x72, 0xbb, 0x52, 0xd2, 0x6f, 0xd4, 0x24, 0x04,
	0xec, 0x81, 0xe4, 0x9f, 0x3e, 0x0c, 0x19, 0xbe, 0xab, 0xd0, 0x3a, 0x3f, 0x7f, 0x05, 0xba, 0x4c,
	0x8b, 0x76, 0x21, 0x6b, 0x8f, 0x7c, 0x05, 0x51, 0x65, 0xf2, 0x30, 0x1f, 0xf1, 0xe2, 0xec, 0x60,
	0xef, 0xd8, 0x15, 0xf3, 0x1c, 0x79, 0x0c, 0x27, 0x01, 0x09, 0xe3, 0x31, 0x66, 0xb5, 0xe3, 0x85,
	0x08, 0xc6, 0x6b, 0xfe, 0xb7, 0x36, 0x61, 0x32, 0x4e, 0x58, 0x07, 0xd9, 0xf3, 0x52, 0x69, 0x13,
	0x06, 0x62, 0xc7, 0x7b, 0x84, 0x7c, 0x0f, 0x83, 0x0c, 0xb9, 0x40, 0x43, 0x07, 0xa1, 0xb5, 0x9f,
	0x76, 0xef, 0x6e, 0xf6, 0x88, 0x35, 0x21, 0x7e, 0x2b, 0xde, 0x68, 0x51, 0x2f, 0xeb, 0x94, 0x05,
	0x9b, 0x7c, 0x0d, 0xd3, 0x54, 0x2b, 0x87, 0xca, 0x5d, 0xa1, 0xda, 0xb8, 0x2c, 0xc8, 0x1c, 0xb1,
	0x43, 0x90, 0x7c, 0x07, 0xe7, 0xce, 0x70, 0x65, 0xd7, 0x68, 0x5e, 0xa8, 0x54, 0x0b, 0xa9, 0x36,
	0x41, 0xcb, 0x31, 0x3b, 0xc2, 0x77, 0x5f, 0x0b, 0x74, 0xbe, 0x96, 0x6f, 0xa1, 0xbf, 0xd6, 0xa6,
	0xa0, 0xf1, 0xc7, 0x1f, 0x19, 0x02, 0xc8, 0x8f, 0x30, 0x2a, 0xb5, 0x75, 0xbf, 0xf9, 0xe0, 0xc9,
	0xc7, 0x83, 0x77, 0x41, 0x7e, 0xb7, 0x9c, 0xe1, 0x32, 0x47, 0xf3, 0x3b, 0xde, 0x5b, 0x3a, 0x0d,
	0x8f, 0xea, 0x42, 0xbe, 0x85, 0x06, 0x0b, 0xed, 0xf0, 0x17, 0x21, 0x0c, 0x3d, 0xad, 0x67, 0x7d,
	0x8f, 0xd4, 0x7c, 0x90, 0x77, 0xc5, 0x5e, 0xd1, 0xb3, 0x96, 0x6f, 0x11, 0xf2, 0x14, 0x22, 0x97,
	0x5b, 0x7a, 0x1e, 0xb4, 0x7d, 0x72, 0x30, 0xba, 0x87, 0xbf, 0x09, 0xf3, 0x71, 0x89, 0x80, 0xf8,
	0xe5, 0xed, 0xed, 0x75, 0x3b, 0x31, 0xdf, 0xc0, 0xa9, 0x69, 0xe6, 0xec, 0x2f, 0x23, 0x1d, 0x9a,
	0xe6, 0xa3, 0x79, 0x0f, 0x25, 0x4f, 0x61, 0xd8, 0xdc, 0xd9, 0x4c, 0xd1, 0x41, 0xdd, 0x4d, 0x36,
	0xd6, 0xc6, 0x24, 0xa7, 0x30, 0xa9, 0x6f, 0xa9, 0x93, 0x2c, 0x96, 0xd0, 0xf7, 0x3e, 0xf9, 0x09,
	0x06, 0x2f, 0xb9, 0x12, 0x39, 0x92, 0xcf, 0xbb, 0xe7, 0x3b, 0x2f, 0xba, 0xa0, 0xc7, 0x44, 0x9d,
	0xe4, 0xcd, 0x20, 0x60, 0xcf, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x42, 0xa3, 0x7e, 0xab, 0xfe,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HTTPClient is the client API for HTTP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTTPClient interface {
	Handle(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error)
}

type hTTPClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPClient(cc grpc.ClientConnInterface) HTTPClient {
	return &hTTPClient{cc}
}

func (c *hTTPClient) Handle(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error) {
	out := new(HTTPResponse)
	err := c.cc.Invoke(ctx, "/ghttpproto.HTTP/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPServer is the server API for HTTP service.
type HTTPServer interface {
	Handle(context.Context, *HTTPRequest) (*HTTPResponse, error)
}

// UnimplementedHTTPServer can be embedded to have forward compatible implementations.
type UnimplementedHTTPServer struct {
}

func (*UnimplementedHTTPServer) Handle(ctx context.Context, req *HTTPRequest) (*HTTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterHTTPServer(s *grpc.Server, srv HTTPServer) {
	s.RegisterService(&_HTTP_serviceDesc, srv)
}

func _HTTP_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghttpproto.HTTP/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPServer).Handle(ctx, req.(*HTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTTP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ghttpproto.HTTP",
	HandlerType: (*HTTPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _HTTP_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ghttp.proto",
}
