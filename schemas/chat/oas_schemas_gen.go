// Code generated by ogen, DO NOT EDIT.

package chat

import (
	"time"
)

type APIChatMuteDeleteApplicationJSONBadRequest ErrorResponse

func (*APIChatMuteDeleteApplicationJSONBadRequest) aPIChatMuteDeleteRes() {}

type APIChatMuteDeleteApplicationJSONInternalServerError ErrorResponse

func (*APIChatMuteDeleteApplicationJSONInternalServerError) aPIChatMuteDeleteRes() {}

type APIChatMutePostApplicationJSONBadRequest ErrorResponse

func (*APIChatMutePostApplicationJSONBadRequest) aPIChatMutePostRes() {}

type APIChatMutePostApplicationJSONInternalServerError ErrorResponse

func (*APIChatMutePostApplicationJSONInternalServerError) aPIChatMutePostRes() {}

type APIChatSendPostApplicationJSONBadRequest ErrorResponse

func (*APIChatSendPostApplicationJSONBadRequest) aPIChatSendPostRes() {}

type APIChatSendPostApplicationJSONInternalServerError ErrorResponse

func (*APIChatSendPostApplicationJSONInternalServerError) aPIChatSendPostRes() {}

// Ref: #/components/schemas/ChatMessage
type ChatMessage struct {
	// The uuid of the sender.
	UUID OptNilString `json:"uuid"`
	// If available the name of the sender.
	Name OptNilString `json:"name"`
	// What color/prefix the sender has, if empty the color will be white and message should be gray.
	Prefix OptNilString `json:"prefix"`
	// Content of the message.
	Message OptNilString `json:"message"`
	// What client software sent this message.
	ClientName OptNilString `json:"clientName"`
}

// GetUUID returns the value of UUID.
func (s *ChatMessage) GetUUID() OptNilString {
	return s.UUID
}

// GetName returns the value of Name.
func (s *ChatMessage) GetName() OptNilString {
	return s.Name
}

// GetPrefix returns the value of Prefix.
func (s *ChatMessage) GetPrefix() OptNilString {
	return s.Prefix
}

// GetMessage returns the value of Message.
func (s *ChatMessage) GetMessage() OptNilString {
	return s.Message
}

// GetClientName returns the value of ClientName.
func (s *ChatMessage) GetClientName() OptNilString {
	return s.ClientName
}

// SetUUID sets the value of UUID.
func (s *ChatMessage) SetUUID(val OptNilString) {
	s.UUID = val
}

// SetName sets the value of Name.
func (s *ChatMessage) SetName(val OptNilString) {
	s.Name = val
}

// SetPrefix sets the value of Prefix.
func (s *ChatMessage) SetPrefix(val OptNilString) {
	s.Prefix = val
}

// SetMessage sets the value of Message.
func (s *ChatMessage) SetMessage(val OptNilString) {
	s.Message = val
}

// SetClientName sets the value of ClientName.
func (s *ChatMessage) SetClientName(val OptNilString) {
	s.ClientName = val
}

// Client software capeable of sending messages.
// Ref: #/components/schemas/ClientThing
type ClientThing struct {
	// Uuid of the target user.
	Name OptNilString `json:"name"`
	// Per minute send quota.
	Quota OptInt32 `json:"quota"`
	// User who is responsible for this client.
	Contact OptNilString `json:"contact"`
	// Webhook to post new messages too.
	WebHook OptNilString `json:"webHook"`
	// Auth header value for the webhook.
	WebhookAuth OptNilString `json:"webhookAuth"`
	// When this was created.
	Created OptDateTime `json:"created"`
}

// GetName returns the value of Name.
func (s *ClientThing) GetName() OptNilString {
	return s.Name
}

// GetQuota returns the value of Quota.
func (s *ClientThing) GetQuota() OptInt32 {
	return s.Quota
}

// GetContact returns the value of Contact.
func (s *ClientThing) GetContact() OptNilString {
	return s.Contact
}

// GetWebHook returns the value of WebHook.
func (s *ClientThing) GetWebHook() OptNilString {
	return s.WebHook
}

// GetWebhookAuth returns the value of WebhookAuth.
func (s *ClientThing) GetWebhookAuth() OptNilString {
	return s.WebhookAuth
}

// GetCreated returns the value of Created.
func (s *ClientThing) GetCreated() OptDateTime {
	return s.Created
}

// SetName sets the value of Name.
func (s *ClientThing) SetName(val OptNilString) {
	s.Name = val
}

// SetQuota sets the value of Quota.
func (s *ClientThing) SetQuota(val OptInt32) {
	s.Quota = val
}

// SetContact sets the value of Contact.
func (s *ClientThing) SetContact(val OptNilString) {
	s.Contact = val
}

// SetWebHook sets the value of WebHook.
func (s *ClientThing) SetWebHook(val OptNilString) {
	s.WebHook = val
}

// SetWebhookAuth sets the value of WebhookAuth.
func (s *ClientThing) SetWebhookAuth(val OptNilString) {
	s.WebhookAuth = val
}

// SetCreated sets the value of Created.
func (s *ClientThing) SetCreated(val OptDateTime) {
	s.Created = val
}

// Ref: #/components/schemas/ErrorResponse
type ErrorResponse struct {
	Slug    OptNilString `json:"slug"`
	Message OptNilString `json:"message"`
	Trace   OptNilString `json:"trace"`
}

// GetSlug returns the value of Slug.
func (s *ErrorResponse) GetSlug() OptNilString {
	return s.Slug
}

// GetMessage returns the value of Message.
func (s *ErrorResponse) GetMessage() OptNilString {
	return s.Message
}

// GetTrace returns the value of Trace.
func (s *ErrorResponse) GetTrace() OptNilString {
	return s.Trace
}

// SetSlug sets the value of Slug.
func (s *ErrorResponse) SetSlug(val OptNilString) {
	s.Slug = val
}

// SetMessage sets the value of Message.
func (s *ErrorResponse) SetMessage(val OptNilString) {
	s.Message = val
}

// SetTrace sets the value of Trace.
func (s *ErrorResponse) SetTrace(val OptNilString) {
	s.Trace = val
}

// Ref: #/components/schemas/Mute
type Mute struct {
	// Uuid of the target user.
	UUID OptNilString `json:"uuid"`
	// Uuid of user performing the mute.
	Muter OptNilString `json:"muter"`
	// Uuid of user performing the mute.
	UnMuter OptNilString `json:"unMuter"`
	// Message for the user.
	Message OptNilString `json:"message"`
	// Internal reason.
	Reason OptNilString `json:"reason"`
	// What client software added the mute.
	ClientId OptInt32 `json:"clientId"`
	// What client software added the mute.
	UnMuteClientId OptInt32 `json:"unMuteClientId"`
	// When this was created.
	Timestamp OptDateTime `json:"timestamp"`
	// Until when this is active.
	Expires OptDateTime   `json:"expires"`
	Status  OptMuteStatus `json:"status"`
}

// GetUUID returns the value of UUID.
func (s *Mute) GetUUID() OptNilString {
	return s.UUID
}

// GetMuter returns the value of Muter.
func (s *Mute) GetMuter() OptNilString {
	return s.Muter
}

// GetUnMuter returns the value of UnMuter.
func (s *Mute) GetUnMuter() OptNilString {
	return s.UnMuter
}

// GetMessage returns the value of Message.
func (s *Mute) GetMessage() OptNilString {
	return s.Message
}

// GetReason returns the value of Reason.
func (s *Mute) GetReason() OptNilString {
	return s.Reason
}

// GetClientId returns the value of ClientId.
func (s *Mute) GetClientId() OptInt32 {
	return s.ClientId
}

// GetUnMuteClientId returns the value of UnMuteClientId.
func (s *Mute) GetUnMuteClientId() OptInt32 {
	return s.UnMuteClientId
}

// GetTimestamp returns the value of Timestamp.
func (s *Mute) GetTimestamp() OptDateTime {
	return s.Timestamp
}

// GetExpires returns the value of Expires.
func (s *Mute) GetExpires() OptDateTime {
	return s.Expires
}

// GetStatus returns the value of Status.
func (s *Mute) GetStatus() OptMuteStatus {
	return s.Status
}

// SetUUID sets the value of UUID.
func (s *Mute) SetUUID(val OptNilString) {
	s.UUID = val
}

// SetMuter sets the value of Muter.
func (s *Mute) SetMuter(val OptNilString) {
	s.Muter = val
}

// SetUnMuter sets the value of UnMuter.
func (s *Mute) SetUnMuter(val OptNilString) {
	s.UnMuter = val
}

// SetMessage sets the value of Message.
func (s *Mute) SetMessage(val OptNilString) {
	s.Message = val
}

// SetReason sets the value of Reason.
func (s *Mute) SetReason(val OptNilString) {
	s.Reason = val
}

// SetClientId sets the value of ClientId.
func (s *Mute) SetClientId(val OptInt32) {
	s.ClientId = val
}

// SetUnMuteClientId sets the value of UnMuteClientId.
func (s *Mute) SetUnMuteClientId(val OptInt32) {
	s.UnMuteClientId = val
}

// SetTimestamp sets the value of Timestamp.
func (s *Mute) SetTimestamp(val OptDateTime) {
	s.Timestamp = val
}

// SetExpires sets the value of Expires.
func (s *Mute) SetExpires(val OptDateTime) {
	s.Expires = val
}

// SetStatus sets the value of Status.
func (s *Mute) SetStatus(val OptMuteStatus) {
	s.Status = val
}

// Ref: #/components/schemas/MuteStatus
type MuteStatus int32

const (
	MuteStatus0  MuteStatus = 0
	MuteStatus1  MuteStatus = 1
	MuteStatus2  MuteStatus = 2
	MuteStatus4  MuteStatus = 4
	MuteStatus8  MuteStatus = 8
	MuteStatus16 MuteStatus = 16
)

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt32 returns new OptInt32 with value set to v.
func NewOptInt32(v int32) OptInt32 {
	return OptInt32{
		Value: v,
		Set:   true,
	}
}

// OptInt32 is optional int32.
type OptInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if OptInt32 was set.
func (o OptInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMuteStatus returns new OptMuteStatus with value set to v.
func NewOptMuteStatus(v MuteStatus) OptMuteStatus {
	return OptMuteStatus{
		Value: v,
		Set:   true,
	}
}

// OptMuteStatus is optional MuteStatus.
type OptMuteStatus struct {
	Value MuteStatus
	Set   bool
}

// IsSet returns true if OptMuteStatus was set.
func (o OptMuteStatus) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMuteStatus) Reset() {
	var v MuteStatus
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMuteStatus) SetTo(v MuteStatus) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMuteStatus) Get() (v MuteStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMuteStatus) Or(d MuteStatus) MuteStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNilString returns new OptNilString with value set to v.
func NewOptNilString(v string) OptNilString {
	return OptNilString{
		Value: v,
		Set:   true,
	}
}

// OptNilString is optional nullable string.
type OptNilString struct {
	Value string
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilString was set.
func (o OptNilString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilString) Reset() {
	var v string
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilString) SetTo(v string) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o OptNilString) IsNull() bool { return o.Null }

// Get returns value and boolean that denotes whether value was set.
func (o OptNilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}
