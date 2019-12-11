// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 
//
package v1

import (
	"context"
)

//go:generate mockgen -destination=mock/mock.go -package=mock optisam-backend/account-service/pkg/repository/v1 Account

//Account interface
type Account interface {
	UpdateAccount(ctx context.Context, userID string, req *UpdateAccount) error

	// CreateAccount creates a user in database
	CreateAccount(ctx context.Context, acc *AccountInfo) error

	AccountInfo(ctx context.Context, userID string) (*AccountInfo, error)
	// UserOwnedGroups returns total number of groups owned by user and a slice of
	// groups based on GroupQueryParams
	UserOwnedGroups(ctx context.Context, userID string, params *GroupQueryParams) (int, []*Group, error)

	// UserOwnedGroups returns total number of groups directly belongs to user and a slice of
	// groups based on GroupQueryParams
	UserOwnedGroupsDirect(ctx context.Context, userID string, params *GroupQueryParams) ([]*Group, error)

	// CreateGroup creates a group in database
	CreateGroup(ctx context.Context, userID string, group *Group) (*Group, error)

	// UpdateGroup updates the given group;
	UpdateGroup(ctx context.Context, groupID int64, update *GroupUpdate) error

	// DeleteGroup deletes a group with given id.
	DeleteGroup(ctx context.Context, groupID int64) error

	// ChildGroups fetches child groups of group with given id and query parameters.
	ChildGroupsDirect(ctx context.Context, groupID int64, params *GroupQueryParams) ([]*Group, error)

	// ChildGroups fetches All child groups of group with given id and query parameters.
	ChildGroupsAll(ctx context.Context, groupID int64, params *GroupQueryParams) ([]*Group, error)

	GroupInfo(ctx context.Context, groupID int64) (*Group, error)

	GroupExistsByFQN(ctx context.Context, fullyQN string) (bool, error)

	UserExistsByID(ctx context.Context, userID string) (bool, error)

	// UsersAll fetches all the users present
	UsersAll(ctx context.Context) ([]*AccountInfo, error)

	//GroupUsers fetches all the users present in the group with given group id
	GroupUsers(ctx context.Context, groupID int64) ([]*AccountInfo, error)

	//AddGroupUsers adds user to the group with given group id and user id
	AddGroupUsers(ctx context.Context, groupID int64, userIDs []string) error

	//DeleteGroupUsers adds  selected users from the group with given group id and user id
	DeleteGroupUsers(ctx context.Context, groupID int64, userIDs []string) error
	//GroupExistsByID(ctx context.Context, groupID int64) (bool, error)

	// UserOwnsGroupByID checks if the user owns the group either directly or as a subgroup
	UserOwnsGroupByID(ctx context.Context, userID string, groupID int64) (bool, error)

	// CheckPassword check for users password in database
	CheckPassword(ctx context.Context, userID, password string) (bool, error)

	// ChangePassword changes users password in daatbase
	ChangePassword(ctx context.Context, userID, password string) error

	// IsGroupRoot returns true if group is the root group
	IsGroupRoot(ctx context.Context, groupID int64) (bool, error)
}
