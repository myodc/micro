// Code generated by protoc-gen-go. DO NOT EDIT.
// source: signup/signup.proto

package go_micro_service_signup

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

type SendVerificationEmailRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendVerificationEmailRequest) Reset()         { *m = SendVerificationEmailRequest{} }
func (m *SendVerificationEmailRequest) String() string { return proto.CompactTextString(m) }
func (*SendVerificationEmailRequest) ProtoMessage()    {}
func (*SendVerificationEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{0}
}

func (m *SendVerificationEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendVerificationEmailRequest.Unmarshal(m, b)
}
func (m *SendVerificationEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendVerificationEmailRequest.Marshal(b, m, deterministic)
}
func (m *SendVerificationEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendVerificationEmailRequest.Merge(m, src)
}
func (m *SendVerificationEmailRequest) XXX_Size() int {
	return xxx_messageInfo_SendVerificationEmailRequest.Size(m)
}
func (m *SendVerificationEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendVerificationEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendVerificationEmailRequest proto.InternalMessageInfo

func (m *SendVerificationEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type SendVerificationEmailResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendVerificationEmailResponse) Reset()         { *m = SendVerificationEmailResponse{} }
func (m *SendVerificationEmailResponse) String() string { return proto.CompactTextString(m) }
func (*SendVerificationEmailResponse) ProtoMessage()    {}
func (*SendVerificationEmailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{1}
}

func (m *SendVerificationEmailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendVerificationEmailResponse.Unmarshal(m, b)
}
func (m *SendVerificationEmailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendVerificationEmailResponse.Marshal(b, m, deterministic)
}
func (m *SendVerificationEmailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendVerificationEmailResponse.Merge(m, src)
}
func (m *SendVerificationEmailResponse) XXX_Size() int {
	return xxx_messageInfo_SendVerificationEmailResponse.Size(m)
}
func (m *SendVerificationEmailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendVerificationEmailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendVerificationEmailResponse proto.InternalMessageInfo

type VerifyRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// Email token that was received in an email.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{2}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *VerifyRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyResponse struct {
	// Auth token to be saved into `~/.micro`
	// For users who are already registered and paid,
	// the flow stops here.
	// For users who are yet to be registered
	// the token will be acquired in the `FinishSignup` step.
	AuthToken *AuthToken `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	// Payment provider custommer id that can be used to
	// acquire a payment method, see `micro login` flow for more.
	// @todo this is likely not needed
	CustomerID string `protobuf:"bytes,2,opt,name=customerID,proto3" json:"customerID,omitempty"`
	// Namespace to use
	// @todod deprecated since we no longer support OTP logins
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Message to display to the user
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// Whether payment is required or not
	PaymentRequired bool `protobuf:"varint,5,opt,name=payment_required,json=paymentRequired,proto3" json:"payment_required,omitempty"`
	// Namespaces one has access to based on previous invites
	// Currently only 1 is supported
	Namespaces           []string `protobuf:"bytes,6,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyResponse) Reset()         { *m = VerifyResponse{} }
func (m *VerifyResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyResponse) ProtoMessage()    {}
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{3}
}

func (m *VerifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyResponse.Unmarshal(m, b)
}
func (m *VerifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyResponse.Marshal(b, m, deterministic)
}
func (m *VerifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResponse.Merge(m, src)
}
func (m *VerifyResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyResponse.Size(m)
}
func (m *VerifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResponse proto.InternalMessageInfo

func (m *VerifyResponse) GetAuthToken() *AuthToken {
	if m != nil {
		return m.AuthToken
	}
	return nil
}

func (m *VerifyResponse) GetCustomerID() string {
	if m != nil {
		return m.CustomerID
	}
	return ""
}

func (m *VerifyResponse) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *VerifyResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *VerifyResponse) GetPaymentRequired() bool {
	if m != nil {
		return m.PaymentRequired
	}
	return false
}

func (m *VerifyResponse) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

type SetPaymentMethodRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	PaymentMethod        string   `protobuf:"bytes,2,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetPaymentMethodRequest) Reset()         { *m = SetPaymentMethodRequest{} }
func (m *SetPaymentMethodRequest) String() string { return proto.CompactTextString(m) }
func (*SetPaymentMethodRequest) ProtoMessage()    {}
func (*SetPaymentMethodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{4}
}

func (m *SetPaymentMethodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetPaymentMethodRequest.Unmarshal(m, b)
}
func (m *SetPaymentMethodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetPaymentMethodRequest.Marshal(b, m, deterministic)
}
func (m *SetPaymentMethodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPaymentMethodRequest.Merge(m, src)
}
func (m *SetPaymentMethodRequest) XXX_Size() int {
	return xxx_messageInfo_SetPaymentMethodRequest.Size(m)
}
func (m *SetPaymentMethodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPaymentMethodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetPaymentMethodRequest proto.InternalMessageInfo

func (m *SetPaymentMethodRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SetPaymentMethodRequest) GetPaymentMethod() string {
	if m != nil {
		return m.PaymentMethod
	}
	return ""
}

type SetPaymentMethodResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetPaymentMethodResponse) Reset()         { *m = SetPaymentMethodResponse{} }
func (m *SetPaymentMethodResponse) String() string { return proto.CompactTextString(m) }
func (*SetPaymentMethodResponse) ProtoMessage()    {}
func (*SetPaymentMethodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{5}
}

func (m *SetPaymentMethodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetPaymentMethodResponse.Unmarshal(m, b)
}
func (m *SetPaymentMethodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetPaymentMethodResponse.Marshal(b, m, deterministic)
}
func (m *SetPaymentMethodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPaymentMethodResponse.Merge(m, src)
}
func (m *SetPaymentMethodResponse) XXX_Size() int {
	return xxx_messageInfo_SetPaymentMethodResponse.Size(m)
}
func (m *SetPaymentMethodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPaymentMethodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetPaymentMethodResponse proto.InternalMessageInfo

type HasPaymentMethodRequest struct {
	// We can't read by email because that would be too easy to guess.
	// The token is already used for identification purposes during the signup
	// so we will use that too to pull for the payment method.
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasPaymentMethodRequest) Reset()         { *m = HasPaymentMethodRequest{} }
func (m *HasPaymentMethodRequest) String() string { return proto.CompactTextString(m) }
func (*HasPaymentMethodRequest) ProtoMessage()    {}
func (*HasPaymentMethodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{6}
}

func (m *HasPaymentMethodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasPaymentMethodRequest.Unmarshal(m, b)
}
func (m *HasPaymentMethodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasPaymentMethodRequest.Marshal(b, m, deterministic)
}
func (m *HasPaymentMethodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasPaymentMethodRequest.Merge(m, src)
}
func (m *HasPaymentMethodRequest) XXX_Size() int {
	return xxx_messageInfo_HasPaymentMethodRequest.Size(m)
}
func (m *HasPaymentMethodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HasPaymentMethodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HasPaymentMethodRequest proto.InternalMessageInfo

func (m *HasPaymentMethodRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type HasPaymentMethodResponse struct {
	Has                  bool     `protobuf:"varint,1,opt,name=has,proto3" json:"has,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HasPaymentMethodResponse) Reset()         { *m = HasPaymentMethodResponse{} }
func (m *HasPaymentMethodResponse) String() string { return proto.CompactTextString(m) }
func (*HasPaymentMethodResponse) ProtoMessage()    {}
func (*HasPaymentMethodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{7}
}

func (m *HasPaymentMethodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasPaymentMethodResponse.Unmarshal(m, b)
}
func (m *HasPaymentMethodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasPaymentMethodResponse.Marshal(b, m, deterministic)
}
func (m *HasPaymentMethodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasPaymentMethodResponse.Merge(m, src)
}
func (m *HasPaymentMethodResponse) XXX_Size() int {
	return xxx_messageInfo_HasPaymentMethodResponse.Size(m)
}
func (m *HasPaymentMethodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HasPaymentMethodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HasPaymentMethodResponse proto.InternalMessageInfo

func (m *HasPaymentMethodResponse) GetHas() bool {
	if m != nil {
		return m.Has
	}
	return false
}

type CompleteSignupRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// The token has to be passed here too for identification purposes.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// This payment method ID is the one we got back from Stripe on the frontend (ie. `m3o.com/subscribe.html`)
	// deprecated: signup service now knows the payment method due to the
	// SetPaymentMethod call issued by the frontend.
	PaymentMethodID string `protobuf:"bytes,3,opt,name=paymentMethodID,proto3" json:"paymentMethodID,omitempty"`
	// The secret/password to use for the account
	Secret string `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	// Which namespace to sign up to based on previous invite
	Namespace            string   `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteSignupRequest) Reset()         { *m = CompleteSignupRequest{} }
func (m *CompleteSignupRequest) String() string { return proto.CompactTextString(m) }
func (*CompleteSignupRequest) ProtoMessage()    {}
func (*CompleteSignupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{8}
}

func (m *CompleteSignupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteSignupRequest.Unmarshal(m, b)
}
func (m *CompleteSignupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteSignupRequest.Marshal(b, m, deterministic)
}
func (m *CompleteSignupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteSignupRequest.Merge(m, src)
}
func (m *CompleteSignupRequest) XXX_Size() int {
	return xxx_messageInfo_CompleteSignupRequest.Size(m)
}
func (m *CompleteSignupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteSignupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteSignupRequest proto.InternalMessageInfo

func (m *CompleteSignupRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CompleteSignupRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CompleteSignupRequest) GetPaymentMethodID() string {
	if m != nil {
		return m.PaymentMethodID
	}
	return ""
}

func (m *CompleteSignupRequest) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *CompleteSignupRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type CompleteSignupResponse struct {
	AuthToken            *AuthToken `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	Namespace            string     `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CompleteSignupResponse) Reset()         { *m = CompleteSignupResponse{} }
func (m *CompleteSignupResponse) String() string { return proto.CompactTextString(m) }
func (*CompleteSignupResponse) ProtoMessage()    {}
func (*CompleteSignupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{9}
}

func (m *CompleteSignupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteSignupResponse.Unmarshal(m, b)
}
func (m *CompleteSignupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteSignupResponse.Marshal(b, m, deterministic)
}
func (m *CompleteSignupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteSignupResponse.Merge(m, src)
}
func (m *CompleteSignupResponse) XXX_Size() int {
	return xxx_messageInfo_CompleteSignupResponse.Size(m)
}
func (m *CompleteSignupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteSignupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteSignupResponse proto.InternalMessageInfo

func (m *CompleteSignupResponse) GetAuthToken() *AuthToken {
	if m != nil {
		return m.AuthToken
	}
	return nil
}

func (m *CompleteSignupResponse) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// lifted from https://github.com/micro/go-micro/blob/master/auth/service/proto/auth.proto
type AuthToken struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Expiry               int64    `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthToken) Reset()         { *m = AuthToken{} }
func (m *AuthToken) String() string { return proto.CompactTextString(m) }
func (*AuthToken) ProtoMessage()    {}
func (*AuthToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{10}
}

func (m *AuthToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthToken.Unmarshal(m, b)
}
func (m *AuthToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthToken.Marshal(b, m, deterministic)
}
func (m *AuthToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthToken.Merge(m, src)
}
func (m *AuthToken) XXX_Size() int {
	return xxx_messageInfo_AuthToken.Size(m)
}
func (m *AuthToken) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthToken.DiscardUnknown(m)
}

var xxx_messageInfo_AuthToken proto.InternalMessageInfo

func (m *AuthToken) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *AuthToken) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *AuthToken) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *AuthToken) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type RecoverRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecoverRequest) Reset()         { *m = RecoverRequest{} }
func (m *RecoverRequest) String() string { return proto.CompactTextString(m) }
func (*RecoverRequest) ProtoMessage()    {}
func (*RecoverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{11}
}

func (m *RecoverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecoverRequest.Unmarshal(m, b)
}
func (m *RecoverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecoverRequest.Marshal(b, m, deterministic)
}
func (m *RecoverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoverRequest.Merge(m, src)
}
func (m *RecoverRequest) XXX_Size() int {
	return xxx_messageInfo_RecoverRequest.Size(m)
}
func (m *RecoverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecoverRequest proto.InternalMessageInfo

func (m *RecoverRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type RecoverResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecoverResponse) Reset()         { *m = RecoverResponse{} }
func (m *RecoverResponse) String() string { return proto.CompactTextString(m) }
func (*RecoverResponse) ProtoMessage()    {}
func (*RecoverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d51c8882b3c28c6, []int{12}
}

func (m *RecoverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecoverResponse.Unmarshal(m, b)
}
func (m *RecoverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecoverResponse.Marshal(b, m, deterministic)
}
func (m *RecoverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoverResponse.Merge(m, src)
}
func (m *RecoverResponse) XXX_Size() int {
	return xxx_messageInfo_RecoverResponse.Size(m)
}
func (m *RecoverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecoverResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SendVerificationEmailRequest)(nil), "go.micro.service.signup.SendVerificationEmailRequest")
	proto.RegisterType((*SendVerificationEmailResponse)(nil), "go.micro.service.signup.SendVerificationEmailResponse")
	proto.RegisterType((*VerifyRequest)(nil), "go.micro.service.signup.VerifyRequest")
	proto.RegisterType((*VerifyResponse)(nil), "go.micro.service.signup.VerifyResponse")
	proto.RegisterType((*SetPaymentMethodRequest)(nil), "go.micro.service.signup.SetPaymentMethodRequest")
	proto.RegisterType((*SetPaymentMethodResponse)(nil), "go.micro.service.signup.SetPaymentMethodResponse")
	proto.RegisterType((*HasPaymentMethodRequest)(nil), "go.micro.service.signup.HasPaymentMethodRequest")
	proto.RegisterType((*HasPaymentMethodResponse)(nil), "go.micro.service.signup.HasPaymentMethodResponse")
	proto.RegisterType((*CompleteSignupRequest)(nil), "go.micro.service.signup.CompleteSignupRequest")
	proto.RegisterType((*CompleteSignupResponse)(nil), "go.micro.service.signup.CompleteSignupResponse")
	proto.RegisterType((*AuthToken)(nil), "go.micro.service.signup.AuthToken")
	proto.RegisterType((*RecoverRequest)(nil), "go.micro.service.signup.RecoverRequest")
	proto.RegisterType((*RecoverResponse)(nil), "go.micro.service.signup.RecoverResponse")
}

func init() { proto.RegisterFile("signup/signup.proto", fileDescriptor_2d51c8882b3c28c6) }

var fileDescriptor_2d51c8882b3c28c6 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x1b, 0x92, 0x34, 0xd3, 0xe6, 0xc2, 0x42, 0x1b, 0xcb, 0x2a, 0x10, 0x8c, 0x28, 0x41,
	0x42, 0x0e, 0x94, 0xcb, 0x0b, 0x2f, 0x20, 0x8a, 0x44, 0x1f, 0x90, 0x90, 0x83, 0x2a, 0x21, 0x21,
	0x45, 0xc6, 0x99, 0x26, 0x16, 0xb5, 0xd7, 0xd9, 0xdd, 0x94, 0xe6, 0x03, 0xe0, 0x47, 0xf8, 0x00,
	0xbe, 0x8e, 0x77, 0xe4, 0xdd, 0xcd, 0xc5, 0x6e, 0xec, 0x52, 0xc4, 0x53, 0x32, 0xb3, 0x67, 0xce,
	0xec, 0x9e, 0xb3, 0xb3, 0x86, 0x1b, 0x3c, 0x18, 0x45, 0xd3, 0xb8, 0xa7, 0x7e, 0x9c, 0x98, 0x51,
	0x41, 0x49, 0x7b, 0x44, 0x9d, 0x30, 0xf0, 0x19, 0x75, 0x38, 0xb2, 0xb3, 0xc0, 0x47, 0x47, 0x2d,
	0xdb, 0xcf, 0x60, 0xaf, 0x8f, 0xd1, 0xf0, 0x18, 0x59, 0x70, 0x12, 0xf8, 0x9e, 0x08, 0x68, 0xf4,
	0x36, 0xf4, 0x82, 0x53, 0x17, 0x27, 0x53, 0xe4, 0x82, 0xdc, 0x84, 0x32, 0x26, 0xb1, 0x69, 0x74,
	0x8c, 0x6e, 0xcd, 0x55, 0x81, 0x7d, 0x07, 0x6e, 0xe5, 0x54, 0xf1, 0x98, 0x46, 0x1c, 0xed, 0x97,
	0x50, 0x97, 0x8b, 0xb3, 0x42, 0x9e, 0x24, 0x2b, 0xe8, 0x57, 0x8c, 0xcc, 0x0d, 0x95, 0x95, 0x81,
	0xfd, 0xdb, 0x80, 0xc6, 0xbc, 0x5a, 0xf1, 0x91, 0x57, 0x50, 0xf3, 0xa6, 0x62, 0xfc, 0x51, 0x82,
	0x13, 0x8a, 0xad, 0x03, 0xdb, 0xc9, 0x39, 0x93, 0xf3, 0x7a, 0x8e, 0x74, 0x97, 0x45, 0xe4, 0x36,
	0x80, 0x3f, 0xe5, 0x82, 0x86, 0xc8, 0x8e, 0x0e, 0x75, 0xbf, 0x95, 0x0c, 0xd9, 0x83, 0x5a, 0xe4,
	0x85, 0xc8, 0x63, 0xcf, 0x47, 0xb3, 0x24, 0x97, 0x97, 0x09, 0x62, 0x42, 0x35, 0x44, 0xce, 0xbd,
	0x11, 0x9a, 0xd7, 0xe4, 0xda, 0x3c, 0x24, 0x0f, 0xa1, 0x15, 0x7b, 0xb3, 0x10, 0x23, 0x31, 0x60,
	0x38, 0x99, 0x06, 0x0c, 0x87, 0x66, 0xb9, 0x63, 0x74, 0x37, 0xdd, 0xa6, 0xce, 0xbb, 0x3a, 0x9d,
	0x6c, 0x61, 0xc1, 0xc8, 0xcd, 0x4a, 0xa7, 0x94, 0x6c, 0x61, 0x99, 0xb1, 0x8f, 0xa1, 0xdd, 0x47,
	0xf1, 0x41, 0x55, 0xbd, 0x47, 0x31, 0xa6, 0xc3, 0x62, 0xf9, 0xee, 0x43, 0x63, 0xde, 0x3b, 0x94,
	0x70, 0x7d, 0xae, 0x7a, 0xbc, 0xca, 0x61, 0x5b, 0x60, 0x5e, 0xe4, 0xd5, 0x46, 0xf5, 0xa0, 0xfd,
	0xce, 0xe3, 0x79, 0x3d, 0xc5, 0x42, 0xef, 0x85, 0x39, 0x8f, 0xc0, 0xbc, 0x58, 0xa0, 0x5d, 0x6a,
	0x41, 0x69, 0xec, 0x71, 0x89, 0xdf, 0x74, 0x93, 0xbf, 0xf6, 0x4f, 0x03, 0x76, 0xde, 0xd0, 0x30,
	0x3e, 0x45, 0x81, 0x7d, 0xe9, 0xce, 0x3f, 0x5c, 0x08, 0xd2, 0x85, 0x66, 0xea, 0x44, 0x47, 0x87,
	0xda, 0xa1, 0x6c, 0x9a, 0xec, 0x42, 0x85, 0xa3, 0xcf, 0x50, 0x68, 0x9b, 0x74, 0x94, 0x76, 0xb7,
	0x9c, 0x71, 0xd7, 0x3e, 0x87, 0xdd, 0xec, 0x26, 0xff, 0xdb, 0xbd, 0x4b, 0x75, 0xde, 0xc8, 0x76,
	0xfe, 0x6e, 0x40, 0x6d, 0x51, 0x46, 0xee, 0xc2, 0xb6, 0xe7, 0xfb, 0xc8, 0xf9, 0x60, 0x55, 0xf8,
	0x2d, 0x95, 0x53, 0x90, 0x7b, 0x50, 0x67, 0x78, 0xc2, 0x90, 0x8f, 0x07, 0xab, 0x42, 0x6d, 0xeb,
	0xa4, 0x02, 0x99, 0x50, 0xf5, 0x19, 0x7a, 0x02, 0x87, 0x52, 0xa7, 0x92, 0x3b, 0x0f, 0x13, 0x7d,
	0xf0, 0x3c, 0x0e, 0xd8, 0x4c, 0xea, 0x53, 0x72, 0x75, 0x64, 0xef, 0x43, 0xc3, 0x45, 0x9f, 0x9e,
	0x21, 0x2b, 0x1e, 0xfc, 0xeb, 0xd0, 0x5c, 0xe0, 0x94, 0x44, 0x07, 0xbf, 0xca, 0x50, 0x51, 0xaa,
	0x91, 0x1f, 0x06, 0xec, 0xac, 0x7d, 0x17, 0xc8, 0xf3, 0x5c, 0xd1, 0x8a, 0x5e, 0x1f, 0xeb, 0xc5,
	0x55, 0xcb, 0xb4, 0x6d, 0x9f, 0xa0, 0xa2, 0x1e, 0x10, 0xb2, 0x9f, 0xcb, 0x90, 0x7a, 0x9f, 0xac,
	0x07, 0x97, 0xe2, 0x34, 0xf5, 0x37, 0x68, 0x65, 0x87, 0x89, 0x3c, 0x2e, 0xd8, 0xe6, 0xda, 0x79,
	0xb6, 0x9e, 0x5c, 0xa1, 0x62, 0xd9, 0x38, 0x3b, 0x78, 0x05, 0x8d, 0x73, 0x86, 0xba, 0xa0, 0x71,
	0xee, 0x54, 0x4f, 0xa0, 0x91, 0x9e, 0x0e, 0xe2, 0xe4, 0x92, 0xac, 0x9d, 0x75, 0xab, 0xf7, 0xd7,
	0x78, 0xdd, 0xf2, 0x33, 0x54, 0xf5, 0x35, 0x23, 0xf9, 0xc6, 0xa4, 0x2f, 0xac, 0xd5, 0xbd, 0x1c,
	0xa8, 0xd8, 0xbf, 0x54, 0xe4, 0x37, 0xf1, 0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x60,
	0xec, 0x98, 0x2a, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SignupClient is the client API for Signup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignupClient interface {
	// Sends the verification email to the user
	SendVerificationEmail(ctx context.Context, in *SendVerificationEmailRequest, opts ...grpc.CallOption) (*SendVerificationEmailResponse, error)
	// Verify kicks off the process of verification
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
	SetPaymentMethod(ctx context.Context, in *SetPaymentMethodRequest, opts ...grpc.CallOption) (*SetPaymentMethodResponse, error)
	HasPaymentMethod(ctx context.Context, in *HasPaymentMethodRequest, opts ...grpc.CallOption) (*HasPaymentMethodResponse, error)
	// Creates a subscription and an account
	CompleteSignup(ctx context.Context, in *CompleteSignupRequest, opts ...grpc.CallOption) (*CompleteSignupResponse, error)
	Recover(ctx context.Context, in *RecoverRequest, opts ...grpc.CallOption) (*RecoverResponse, error)
}

type signupClient struct {
	cc *grpc.ClientConn
}

func NewSignupClient(cc *grpc.ClientConn) SignupClient {
	return &signupClient{cc}
}

func (c *signupClient) SendVerificationEmail(ctx context.Context, in *SendVerificationEmailRequest, opts ...grpc.CallOption) (*SendVerificationEmailResponse, error) {
	out := new(SendVerificationEmailResponse)
	err := c.cc.Invoke(ctx, "/go.micro.service.signup.Signup/SendVerificationEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/go.micro.service.signup.Signup/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupClient) SetPaymentMethod(ctx context.Context, in *SetPaymentMethodRequest, opts ...grpc.CallOption) (*SetPaymentMethodResponse, error) {
	out := new(SetPaymentMethodResponse)
	err := c.cc.Invoke(ctx, "/go.micro.service.signup.Signup/SetPaymentMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupClient) HasPaymentMethod(ctx context.Context, in *HasPaymentMethodRequest, opts ...grpc.CallOption) (*HasPaymentMethodResponse, error) {
	out := new(HasPaymentMethodResponse)
	err := c.cc.Invoke(ctx, "/go.micro.service.signup.Signup/HasPaymentMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupClient) CompleteSignup(ctx context.Context, in *CompleteSignupRequest, opts ...grpc.CallOption) (*CompleteSignupResponse, error) {
	out := new(CompleteSignupResponse)
	err := c.cc.Invoke(ctx, "/go.micro.service.signup.Signup/CompleteSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupClient) Recover(ctx context.Context, in *RecoverRequest, opts ...grpc.CallOption) (*RecoverResponse, error) {
	out := new(RecoverResponse)
	err := c.cc.Invoke(ctx, "/go.micro.service.signup.Signup/Recover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignupServer is the server API for Signup service.
type SignupServer interface {
	// Sends the verification email to the user
	SendVerificationEmail(context.Context, *SendVerificationEmailRequest) (*SendVerificationEmailResponse, error)
	// Verify kicks off the process of verification
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	SetPaymentMethod(context.Context, *SetPaymentMethodRequest) (*SetPaymentMethodResponse, error)
	HasPaymentMethod(context.Context, *HasPaymentMethodRequest) (*HasPaymentMethodResponse, error)
	// Creates a subscription and an account
	CompleteSignup(context.Context, *CompleteSignupRequest) (*CompleteSignupResponse, error)
	Recover(context.Context, *RecoverRequest) (*RecoverResponse, error)
}

// UnimplementedSignupServer can be embedded to have forward compatible implementations.
type UnimplementedSignupServer struct {
}

func (*UnimplementedSignupServer) SendVerificationEmail(ctx context.Context, req *SendVerificationEmailRequest) (*SendVerificationEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerificationEmail not implemented")
}
func (*UnimplementedSignupServer) Verify(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (*UnimplementedSignupServer) SetPaymentMethod(ctx context.Context, req *SetPaymentMethodRequest) (*SetPaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPaymentMethod not implemented")
}
func (*UnimplementedSignupServer) HasPaymentMethod(ctx context.Context, req *HasPaymentMethodRequest) (*HasPaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasPaymentMethod not implemented")
}
func (*UnimplementedSignupServer) CompleteSignup(ctx context.Context, req *CompleteSignupRequest) (*CompleteSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteSignup not implemented")
}
func (*UnimplementedSignupServer) Recover(ctx context.Context, req *RecoverRequest) (*RecoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recover not implemented")
}

func RegisterSignupServer(s *grpc.Server, srv SignupServer) {
	s.RegisterService(&_Signup_serviceDesc, srv)
}

func _Signup_SendVerificationEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVerificationEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServer).SendVerificationEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.service.signup.Signup/SendVerificationEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServer).SendVerificationEmail(ctx, req.(*SendVerificationEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signup_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.service.signup.Signup/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signup_SetPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServer).SetPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.service.signup.Signup/SetPaymentMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServer).SetPaymentMethod(ctx, req.(*SetPaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signup_HasPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasPaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServer).HasPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.service.signup.Signup/HasPaymentMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServer).HasPaymentMethod(ctx, req.(*HasPaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signup_CompleteSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServer).CompleteSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.service.signup.Signup/CompleteSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServer).CompleteSignup(ctx, req.(*CompleteSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signup_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupServer).Recover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.service.signup.Signup/Recover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupServer).Recover(ctx, req.(*RecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Signup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.service.signup.Signup",
	HandlerType: (*SignupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendVerificationEmail",
			Handler:    _Signup_SendVerificationEmail_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Signup_Verify_Handler,
		},
		{
			MethodName: "SetPaymentMethod",
			Handler:    _Signup_SetPaymentMethod_Handler,
		},
		{
			MethodName: "HasPaymentMethod",
			Handler:    _Signup_HasPaymentMethod_Handler,
		},
		{
			MethodName: "CompleteSignup",
			Handler:    _Signup_CompleteSignup_Handler,
		},
		{
			MethodName: "Recover",
			Handler:    _Signup_Recover_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signup/signup.proto",
}
