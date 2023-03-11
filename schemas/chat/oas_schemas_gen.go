// Code generated by ogen, DO NOT EDIT.

package chat

import (
	"time"
)

// Contains the client and its api key (only place where the api key is visible).
type APIChatInternalClientPostOKApplicationJSON struct {
	// Client software capeable of sending messages.
	Client OptAPIChatInternalClientPostOKApplicationJSONClient `json:"client"`
	ApiKey OptNilString                                        `json:"apiKey"`
}

// GetClient returns the value of Client.
func (s *APIChatInternalClientPostOKApplicationJSON) GetClient() OptAPIChatInternalClientPostOKApplicationJSONClient {
	return s.Client
}

// GetApiKey returns the value of ApiKey.
func (s *APIChatInternalClientPostOKApplicationJSON) GetApiKey() OptNilString {
	return s.ApiKey
}

// SetClient sets the value of Client.
func (s *APIChatInternalClientPostOKApplicationJSON) SetClient(val OptAPIChatInternalClientPostOKApplicationJSONClient) {
	s.Client = val
}

// SetApiKey sets the value of ApiKey.
func (s *APIChatInternalClientPostOKApplicationJSON) SetApiKey(val OptNilString) {
	s.ApiKey = val
}

// Client software capeable of sending messages.
type APIChatInternalClientPostOKApplicationJSONClient struct {
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
func (s *APIChatInternalClientPostOKApplicationJSONClient) GetName() OptNilString {
	return s.Name
}

// GetQuota returns the value of Quota.
func (s *APIChatInternalClientPostOKApplicationJSONClient) GetQuota() OptInt32 {
	return s.Quota
}

// GetContact returns the value of Contact.
func (s *APIChatInternalClientPostOKApplicationJSONClient) GetContact() OptNilString {
	return s.Contact
}

// GetWebHook returns the value of WebHook.
func (s *APIChatInternalClientPostOKApplicationJSONClient) GetWebHook() OptNilString {
	return s.WebHook
}

// GetWebhookAuth returns the value of WebhookAuth.
func (s *APIChatInternalClientPostOKApplicationJSONClient) GetWebhookAuth() OptNilString {
	return s.WebhookAuth
}

// GetCreated returns the value of Created.
func (s *APIChatInternalClientPostOKApplicationJSONClient) GetCreated() OptDateTime {
	return s.Created
}

// SetName sets the value of Name.
func (s *APIChatInternalClientPostOKApplicationJSONClient) SetName(val OptNilString) {
	s.Name = val
}

// SetQuota sets the value of Quota.
func (s *APIChatInternalClientPostOKApplicationJSONClient) SetQuota(val OptInt32) {
	s.Quota = val
}

// SetContact sets the value of Contact.
func (s *APIChatInternalClientPostOKApplicationJSONClient) SetContact(val OptNilString) {
	s.Contact = val
}

// SetWebHook sets the value of WebHook.
func (s *APIChatInternalClientPostOKApplicationJSONClient) SetWebHook(val OptNilString) {
	s.WebHook = val
}

// SetWebhookAuth sets the value of WebhookAuth.
func (s *APIChatInternalClientPostOKApplicationJSONClient) SetWebhookAuth(val OptNilString) {
	s.WebhookAuth = val
}

// SetCreated sets the value of Created.
func (s *APIChatInternalClientPostOKApplicationJSONClient) SetCreated(val OptDateTime) {
	s.Created = val
}

// Client software capeable of sending messages.
type APIChatInternalClientPostReqApplicationJSON struct {
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
func (s *APIChatInternalClientPostReqApplicationJSON) GetName() OptNilString {
	return s.Name
}

// GetQuota returns the value of Quota.
func (s *APIChatInternalClientPostReqApplicationJSON) GetQuota() OptInt32 {
	return s.Quota
}

// GetContact returns the value of Contact.
func (s *APIChatInternalClientPostReqApplicationJSON) GetContact() OptNilString {
	return s.Contact
}

// GetWebHook returns the value of WebHook.
func (s *APIChatInternalClientPostReqApplicationJSON) GetWebHook() OptNilString {
	return s.WebHook
}

// GetWebhookAuth returns the value of WebhookAuth.
func (s *APIChatInternalClientPostReqApplicationJSON) GetWebhookAuth() OptNilString {
	return s.WebhookAuth
}

// GetCreated returns the value of Created.
func (s *APIChatInternalClientPostReqApplicationJSON) GetCreated() OptDateTime {
	return s.Created
}

// SetName sets the value of Name.
func (s *APIChatInternalClientPostReqApplicationJSON) SetName(val OptNilString) {
	s.Name = val
}

// SetQuota sets the value of Quota.
func (s *APIChatInternalClientPostReqApplicationJSON) SetQuota(val OptInt32) {
	s.Quota = val
}

// SetContact sets the value of Contact.
func (s *APIChatInternalClientPostReqApplicationJSON) SetContact(val OptNilString) {
	s.Contact = val
}

// SetWebHook sets the value of WebHook.
func (s *APIChatInternalClientPostReqApplicationJSON) SetWebHook(val OptNilString) {
	s.WebHook = val
}

// SetWebhookAuth sets the value of WebhookAuth.
func (s *APIChatInternalClientPostReqApplicationJSON) SetWebhookAuth(val OptNilString) {
	s.WebhookAuth = val
}

// SetCreated sets the value of Created.
func (s *APIChatInternalClientPostReqApplicationJSON) SetCreated(val OptDateTime) {
	s.Created = val
}

type APIChatMuteDeleteOKApplicationJSON struct {
	// Uuid of the target user.
	UUID OptNilString `json:"uuid"`
	// Uuid of user performing the mute.
	UnMuter OptNilString `json:"unMuter"`
	// Internal reason.
	Reason OptNilString `json:"reason"`
}

// GetUUID returns the value of UUID.
func (s *APIChatMuteDeleteOKApplicationJSON) GetUUID() OptNilString {
	return s.UUID
}

// GetUnMuter returns the value of UnMuter.
func (s *APIChatMuteDeleteOKApplicationJSON) GetUnMuter() OptNilString {
	return s.UnMuter
}

// GetReason returns the value of Reason.
func (s *APIChatMuteDeleteOKApplicationJSON) GetReason() OptNilString {
	return s.Reason
}

// SetUUID sets the value of UUID.
func (s *APIChatMuteDeleteOKApplicationJSON) SetUUID(val OptNilString) {
	s.UUID = val
}

// SetUnMuter sets the value of UnMuter.
func (s *APIChatMuteDeleteOKApplicationJSON) SetUnMuter(val OptNilString) {
	s.UnMuter = val
}

// SetReason sets the value of Reason.
func (s *APIChatMuteDeleteOKApplicationJSON) SetReason(val OptNilString) {
	s.Reason = val
}

type APIChatMuteDeleteReqApplicationJSON struct {
	// Uuid of the target user.
	UUID OptNilString `json:"uuid"`
	// Uuid of user performing the mute.
	UnMuter OptNilString `json:"unMuter"`
	// Internal reason.
	Reason OptNilString `json:"reason"`
}

// GetUUID returns the value of UUID.
func (s *APIChatMuteDeleteReqApplicationJSON) GetUUID() OptNilString {
	return s.UUID
}

// GetUnMuter returns the value of UnMuter.
func (s *APIChatMuteDeleteReqApplicationJSON) GetUnMuter() OptNilString {
	return s.UnMuter
}

// GetReason returns the value of Reason.
func (s *APIChatMuteDeleteReqApplicationJSON) GetReason() OptNilString {
	return s.Reason
}

// SetUUID sets the value of UUID.
func (s *APIChatMuteDeleteReqApplicationJSON) SetUUID(val OptNilString) {
	s.UUID = val
}

// SetUnMuter sets the value of UnMuter.
func (s *APIChatMuteDeleteReqApplicationJSON) SetUnMuter(val OptNilString) {
	s.UnMuter = val
}

// SetReason sets the value of Reason.
func (s *APIChatMuteDeleteReqApplicationJSON) SetReason(val OptNilString) {
	s.Reason = val
}

type APIChatMutePostOKApplicationJSON struct {
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
	Expires OptDateTime                               `json:"expires"`
	Status  OptAPIChatMutePostOKApplicationJSONStatus `json:"status"`
}

// GetUUID returns the value of UUID.
func (s *APIChatMutePostOKApplicationJSON) GetUUID() OptNilString {
	return s.UUID
}

// GetMuter returns the value of Muter.
func (s *APIChatMutePostOKApplicationJSON) GetMuter() OptNilString {
	return s.Muter
}

// GetUnMuter returns the value of UnMuter.
func (s *APIChatMutePostOKApplicationJSON) GetUnMuter() OptNilString {
	return s.UnMuter
}

// GetMessage returns the value of Message.
func (s *APIChatMutePostOKApplicationJSON) GetMessage() OptNilString {
	return s.Message
}

// GetReason returns the value of Reason.
func (s *APIChatMutePostOKApplicationJSON) GetReason() OptNilString {
	return s.Reason
}

// GetClientId returns the value of ClientId.
func (s *APIChatMutePostOKApplicationJSON) GetClientId() OptInt32 {
	return s.ClientId
}

// GetUnMuteClientId returns the value of UnMuteClientId.
func (s *APIChatMutePostOKApplicationJSON) GetUnMuteClientId() OptInt32 {
	return s.UnMuteClientId
}

// GetTimestamp returns the value of Timestamp.
func (s *APIChatMutePostOKApplicationJSON) GetTimestamp() OptDateTime {
	return s.Timestamp
}

// GetExpires returns the value of Expires.
func (s *APIChatMutePostOKApplicationJSON) GetExpires() OptDateTime {
	return s.Expires
}

// GetStatus returns the value of Status.
func (s *APIChatMutePostOKApplicationJSON) GetStatus() OptAPIChatMutePostOKApplicationJSONStatus {
	return s.Status
}

// SetUUID sets the value of UUID.
func (s *APIChatMutePostOKApplicationJSON) SetUUID(val OptNilString) {
	s.UUID = val
}

// SetMuter sets the value of Muter.
func (s *APIChatMutePostOKApplicationJSON) SetMuter(val OptNilString) {
	s.Muter = val
}

// SetUnMuter sets the value of UnMuter.
func (s *APIChatMutePostOKApplicationJSON) SetUnMuter(val OptNilString) {
	s.UnMuter = val
}

// SetMessage sets the value of Message.
func (s *APIChatMutePostOKApplicationJSON) SetMessage(val OptNilString) {
	s.Message = val
}

// SetReason sets the value of Reason.
func (s *APIChatMutePostOKApplicationJSON) SetReason(val OptNilString) {
	s.Reason = val
}

// SetClientId sets the value of ClientId.
func (s *APIChatMutePostOKApplicationJSON) SetClientId(val OptInt32) {
	s.ClientId = val
}

// SetUnMuteClientId sets the value of UnMuteClientId.
func (s *APIChatMutePostOKApplicationJSON) SetUnMuteClientId(val OptInt32) {
	s.UnMuteClientId = val
}

// SetTimestamp sets the value of Timestamp.
func (s *APIChatMutePostOKApplicationJSON) SetTimestamp(val OptDateTime) {
	s.Timestamp = val
}

// SetExpires sets the value of Expires.
func (s *APIChatMutePostOKApplicationJSON) SetExpires(val OptDateTime) {
	s.Expires = val
}

// SetStatus sets the value of Status.
func (s *APIChatMutePostOKApplicationJSON) SetStatus(val OptAPIChatMutePostOKApplicationJSONStatus) {
	s.Status = val
}

type APIChatMutePostOKApplicationJSONStatus int32

const (
	APIChatMutePostOKApplicationJSONStatus0  APIChatMutePostOKApplicationJSONStatus = 0
	APIChatMutePostOKApplicationJSONStatus1  APIChatMutePostOKApplicationJSONStatus = 1
	APIChatMutePostOKApplicationJSONStatus2  APIChatMutePostOKApplicationJSONStatus = 2
	APIChatMutePostOKApplicationJSONStatus4  APIChatMutePostOKApplicationJSONStatus = 4
	APIChatMutePostOKApplicationJSONStatus8  APIChatMutePostOKApplicationJSONStatus = 8
	APIChatMutePostOKApplicationJSONStatus16 APIChatMutePostOKApplicationJSONStatus = 16
)

type APIChatMutePostReqApplicationJSON struct {
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
	Expires OptDateTime                                `json:"expires"`
	Status  OptAPIChatMutePostReqApplicationJSONStatus `json:"status"`
}

// GetUUID returns the value of UUID.
func (s *APIChatMutePostReqApplicationJSON) GetUUID() OptNilString {
	return s.UUID
}

// GetMuter returns the value of Muter.
func (s *APIChatMutePostReqApplicationJSON) GetMuter() OptNilString {
	return s.Muter
}

// GetUnMuter returns the value of UnMuter.
func (s *APIChatMutePostReqApplicationJSON) GetUnMuter() OptNilString {
	return s.UnMuter
}

// GetMessage returns the value of Message.
func (s *APIChatMutePostReqApplicationJSON) GetMessage() OptNilString {
	return s.Message
}

// GetReason returns the value of Reason.
func (s *APIChatMutePostReqApplicationJSON) GetReason() OptNilString {
	return s.Reason
}

// GetClientId returns the value of ClientId.
func (s *APIChatMutePostReqApplicationJSON) GetClientId() OptInt32 {
	return s.ClientId
}

// GetUnMuteClientId returns the value of UnMuteClientId.
func (s *APIChatMutePostReqApplicationJSON) GetUnMuteClientId() OptInt32 {
	return s.UnMuteClientId
}

// GetTimestamp returns the value of Timestamp.
func (s *APIChatMutePostReqApplicationJSON) GetTimestamp() OptDateTime {
	return s.Timestamp
}

// GetExpires returns the value of Expires.
func (s *APIChatMutePostReqApplicationJSON) GetExpires() OptDateTime {
	return s.Expires
}

// GetStatus returns the value of Status.
func (s *APIChatMutePostReqApplicationJSON) GetStatus() OptAPIChatMutePostReqApplicationJSONStatus {
	return s.Status
}

// SetUUID sets the value of UUID.
func (s *APIChatMutePostReqApplicationJSON) SetUUID(val OptNilString) {
	s.UUID = val
}

// SetMuter sets the value of Muter.
func (s *APIChatMutePostReqApplicationJSON) SetMuter(val OptNilString) {
	s.Muter = val
}

// SetUnMuter sets the value of UnMuter.
func (s *APIChatMutePostReqApplicationJSON) SetUnMuter(val OptNilString) {
	s.UnMuter = val
}

// SetMessage sets the value of Message.
func (s *APIChatMutePostReqApplicationJSON) SetMessage(val OptNilString) {
	s.Message = val
}

// SetReason sets the value of Reason.
func (s *APIChatMutePostReqApplicationJSON) SetReason(val OptNilString) {
	s.Reason = val
}

// SetClientId sets the value of ClientId.
func (s *APIChatMutePostReqApplicationJSON) SetClientId(val OptInt32) {
	s.ClientId = val
}

// SetUnMuteClientId sets the value of UnMuteClientId.
func (s *APIChatMutePostReqApplicationJSON) SetUnMuteClientId(val OptInt32) {
	s.UnMuteClientId = val
}

// SetTimestamp sets the value of Timestamp.
func (s *APIChatMutePostReqApplicationJSON) SetTimestamp(val OptDateTime) {
	s.Timestamp = val
}

// SetExpires sets the value of Expires.
func (s *APIChatMutePostReqApplicationJSON) SetExpires(val OptDateTime) {
	s.Expires = val
}

// SetStatus sets the value of Status.
func (s *APIChatMutePostReqApplicationJSON) SetStatus(val OptAPIChatMutePostReqApplicationJSONStatus) {
	s.Status = val
}

type APIChatMutePostReqApplicationJSONStatus int32

const (
	APIChatMutePostReqApplicationJSONStatus0  APIChatMutePostReqApplicationJSONStatus = 0
	APIChatMutePostReqApplicationJSONStatus1  APIChatMutePostReqApplicationJSONStatus = 1
	APIChatMutePostReqApplicationJSONStatus2  APIChatMutePostReqApplicationJSONStatus = 2
	APIChatMutePostReqApplicationJSONStatus4  APIChatMutePostReqApplicationJSONStatus = 4
	APIChatMutePostReqApplicationJSONStatus8  APIChatMutePostReqApplicationJSONStatus = 8
	APIChatMutePostReqApplicationJSONStatus16 APIChatMutePostReqApplicationJSONStatus = 16
)

type APIChatSendPostOKApplicationJSON struct {
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
func (s *APIChatSendPostOKApplicationJSON) GetUUID() OptNilString {
	return s.UUID
}

// GetName returns the value of Name.
func (s *APIChatSendPostOKApplicationJSON) GetName() OptNilString {
	return s.Name
}

// GetPrefix returns the value of Prefix.
func (s *APIChatSendPostOKApplicationJSON) GetPrefix() OptNilString {
	return s.Prefix
}

// GetMessage returns the value of Message.
func (s *APIChatSendPostOKApplicationJSON) GetMessage() OptNilString {
	return s.Message
}

// GetClientName returns the value of ClientName.
func (s *APIChatSendPostOKApplicationJSON) GetClientName() OptNilString {
	return s.ClientName
}

// SetUUID sets the value of UUID.
func (s *APIChatSendPostOKApplicationJSON) SetUUID(val OptNilString) {
	s.UUID = val
}

// SetName sets the value of Name.
func (s *APIChatSendPostOKApplicationJSON) SetName(val OptNilString) {
	s.Name = val
}

// SetPrefix sets the value of Prefix.
func (s *APIChatSendPostOKApplicationJSON) SetPrefix(val OptNilString) {
	s.Prefix = val
}

// SetMessage sets the value of Message.
func (s *APIChatSendPostOKApplicationJSON) SetMessage(val OptNilString) {
	s.Message = val
}

// SetClientName sets the value of ClientName.
func (s *APIChatSendPostOKApplicationJSON) SetClientName(val OptNilString) {
	s.ClientName = val
}

type APIChatSendPostReqApplicationJSON struct {
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
func (s *APIChatSendPostReqApplicationJSON) GetUUID() OptNilString {
	return s.UUID
}

// GetName returns the value of Name.
func (s *APIChatSendPostReqApplicationJSON) GetName() OptNilString {
	return s.Name
}

// GetPrefix returns the value of Prefix.
func (s *APIChatSendPostReqApplicationJSON) GetPrefix() OptNilString {
	return s.Prefix
}

// GetMessage returns the value of Message.
func (s *APIChatSendPostReqApplicationJSON) GetMessage() OptNilString {
	return s.Message
}

// GetClientName returns the value of ClientName.
func (s *APIChatSendPostReqApplicationJSON) GetClientName() OptNilString {
	return s.ClientName
}

// SetUUID sets the value of UUID.
func (s *APIChatSendPostReqApplicationJSON) SetUUID(val OptNilString) {
	s.UUID = val
}

// SetName sets the value of Name.
func (s *APIChatSendPostReqApplicationJSON) SetName(val OptNilString) {
	s.Name = val
}

// SetPrefix sets the value of Prefix.
func (s *APIChatSendPostReqApplicationJSON) SetPrefix(val OptNilString) {
	s.Prefix = val
}

// SetMessage sets the value of Message.
func (s *APIChatSendPostReqApplicationJSON) SetMessage(val OptNilString) {
	s.Message = val
}

// SetClientName sets the value of ClientName.
func (s *APIChatSendPostReqApplicationJSON) SetClientName(val OptNilString) {
	s.ClientName = val
}

// NewOptAPIChatInternalClientPostOKApplicationJSONClient returns new OptAPIChatInternalClientPostOKApplicationJSONClient with value set to v.
func NewOptAPIChatInternalClientPostOKApplicationJSONClient(v APIChatInternalClientPostOKApplicationJSONClient) OptAPIChatInternalClientPostOKApplicationJSONClient {
	return OptAPIChatInternalClientPostOKApplicationJSONClient{
		Value: v,
		Set:   true,
	}
}

// OptAPIChatInternalClientPostOKApplicationJSONClient is optional APIChatInternalClientPostOKApplicationJSONClient.
type OptAPIChatInternalClientPostOKApplicationJSONClient struct {
	Value APIChatInternalClientPostOKApplicationJSONClient
	Set   bool
}

// IsSet returns true if OptAPIChatInternalClientPostOKApplicationJSONClient was set.
func (o OptAPIChatInternalClientPostOKApplicationJSONClient) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptAPIChatInternalClientPostOKApplicationJSONClient) Reset() {
	var v APIChatInternalClientPostOKApplicationJSONClient
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptAPIChatInternalClientPostOKApplicationJSONClient) SetTo(v APIChatInternalClientPostOKApplicationJSONClient) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptAPIChatInternalClientPostOKApplicationJSONClient) Get() (v APIChatInternalClientPostOKApplicationJSONClient, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptAPIChatInternalClientPostOKApplicationJSONClient) Or(d APIChatInternalClientPostOKApplicationJSONClient) APIChatInternalClientPostOKApplicationJSONClient {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptAPIChatMutePostOKApplicationJSONStatus returns new OptAPIChatMutePostOKApplicationJSONStatus with value set to v.
func NewOptAPIChatMutePostOKApplicationJSONStatus(v APIChatMutePostOKApplicationJSONStatus) OptAPIChatMutePostOKApplicationJSONStatus {
	return OptAPIChatMutePostOKApplicationJSONStatus{
		Value: v,
		Set:   true,
	}
}

// OptAPIChatMutePostOKApplicationJSONStatus is optional APIChatMutePostOKApplicationJSONStatus.
type OptAPIChatMutePostOKApplicationJSONStatus struct {
	Value APIChatMutePostOKApplicationJSONStatus
	Set   bool
}

// IsSet returns true if OptAPIChatMutePostOKApplicationJSONStatus was set.
func (o OptAPIChatMutePostOKApplicationJSONStatus) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptAPIChatMutePostOKApplicationJSONStatus) Reset() {
	var v APIChatMutePostOKApplicationJSONStatus
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptAPIChatMutePostOKApplicationJSONStatus) SetTo(v APIChatMutePostOKApplicationJSONStatus) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptAPIChatMutePostOKApplicationJSONStatus) Get() (v APIChatMutePostOKApplicationJSONStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptAPIChatMutePostOKApplicationJSONStatus) Or(d APIChatMutePostOKApplicationJSONStatus) APIChatMutePostOKApplicationJSONStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptAPIChatMutePostReqApplicationJSONStatus returns new OptAPIChatMutePostReqApplicationJSONStatus with value set to v.
func NewOptAPIChatMutePostReqApplicationJSONStatus(v APIChatMutePostReqApplicationJSONStatus) OptAPIChatMutePostReqApplicationJSONStatus {
	return OptAPIChatMutePostReqApplicationJSONStatus{
		Value: v,
		Set:   true,
	}
}

// OptAPIChatMutePostReqApplicationJSONStatus is optional APIChatMutePostReqApplicationJSONStatus.
type OptAPIChatMutePostReqApplicationJSONStatus struct {
	Value APIChatMutePostReqApplicationJSONStatus
	Set   bool
}

// IsSet returns true if OptAPIChatMutePostReqApplicationJSONStatus was set.
func (o OptAPIChatMutePostReqApplicationJSONStatus) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptAPIChatMutePostReqApplicationJSONStatus) Reset() {
	var v APIChatMutePostReqApplicationJSONStatus
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptAPIChatMutePostReqApplicationJSONStatus) SetTo(v APIChatMutePostReqApplicationJSONStatus) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptAPIChatMutePostReqApplicationJSONStatus) Get() (v APIChatMutePostReqApplicationJSONStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptAPIChatMutePostReqApplicationJSONStatus) Or(d APIChatMutePostReqApplicationJSONStatus) APIChatMutePostReqApplicationJSONStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

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