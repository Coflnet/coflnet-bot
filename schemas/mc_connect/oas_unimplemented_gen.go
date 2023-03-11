// Code generated by ogen, DO NOT EDIT.

package mc_connect

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// ConnectMinecraftMcUuidGet implements GET /Connect/minecraft/{mcUuid} operation.
//
// GET /Connect/minecraft/{mcUuid}
func (UnimplementedHandler) ConnectMinecraftMcUuidGet(ctx context.Context, params ConnectMinecraftMcUuidGetParams) (r *User, _ error) {
	return r, ht.ErrNotImplemented
}

// ConnectUserUserIdGet implements GET /Connect/user/{userId} operation.
//
// GET /Connect/user/{userId}
func (UnimplementedHandler) ConnectUserUserIdGet(ctx context.Context, params ConnectUserUserIdGetParams) (r *User, _ error) {
	return r, ht.ErrNotImplemented
}

// ConnectUserUserIdPost implements POST /Connect/user/{userId} operation.
//
// POST /Connect/user/{userId}
func (UnimplementedHandler) ConnectUserUserIdPost(ctx context.Context, params ConnectUserUserIdPostParams) (r *ConnectionRequest, _ error) {
	return r, ht.ErrNotImplemented
}

// ConnectUserUserIdVerifyPost implements POST /Connect/user/{userId}/verify operation.
//
// POST /Connect/user/{userId}/verify
func (UnimplementedHandler) ConnectUserUserIdVerifyPost(ctx context.Context, params ConnectUserUserIdVerifyPostParams) error {
	return ht.ErrNotImplemented
}

// ConnectUsersConnectedGet implements GET /Connect/users/connected operation.
//
// Get all users stored which may or may not have a connected account.
//
// GET /Connect/users/connected
func (UnimplementedHandler) ConnectUsersConnectedGet(ctx context.Context, params ConnectUsersConnectedGetParams) (r []User, _ error) {
	return r, ht.ErrNotImplemented
}

// ConnectUsersGet implements GET /Connect/users operation.
//
// Get all users stored which may or may not have a connected account.
//
// GET /Connect/users
func (UnimplementedHandler) ConnectUsersGet(ctx context.Context, params ConnectUsersGetParams) (r []User, _ error) {
	return r, ht.ErrNotImplemented
}

// ConnectUsersIdsGet implements GET /Connect/users/ids operation.
//
// Get all users which ids is in the list.
//
// GET /Connect/users/ids
func (UnimplementedHandler) ConnectUsersIdsGet(ctx context.Context, params ConnectUsersIdsGetParams) (r []User, _ error) {
	return r, ht.ErrNotImplemented
}