package iam

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/watchtower/internal/db"
	"github.com/hashicorp/watchtower/internal/iam/store"
	"google.golang.org/protobuf/proto"
)

// AuthType defines the possible types for AuthMethod
type AuthType uint32

const (
	AuthUnknown  AuthType = 0
	AuthUserPass AuthType = 1
	AuthOIDC     AuthType = 2
)

func (a AuthType) String() string {
	return [...]string{
		"unknown",
		"userpass",
		"oidc",
	}[a]
}

// AuthMethod are the authentication methods available at the Organization Scope within WatchTower.
type AuthMethod struct {
	*store.AuthMethod
	tableName string `gorm:"-"`
}

// check that required interfaces are implemented: Resource, Clonable, db.VetForWriter
var _ Resource = (*AuthMethod)(nil)
var _ Clonable = (*AuthMethod)(nil)
var _ db.VetForWriter = (*AuthMethod)(nil)

// NewAuthMethod creates a new in memory AuthMethod for an Organization
// and authentication type.
func NewAuthMethod(organizationId string, authType AuthType, opt ...Option) (*AuthMethod, error) {
	opts := getOpts(opt...)
	withName := opts.withName
	if authType == AuthUnknown {
		return nil, errors.New("error unknown auth type")
	}
	if organizationId == "" {
		return nil, errors.New("error organization id is unset for new auth method")
	}
	publicId, err := db.NewPublicId("authm")
	if err != nil {
		return nil, fmt.Errorf("error generating public id %w for new auth method", err)
	}
	a := &AuthMethod{
		AuthMethod: &store.AuthMethod{
			PublicId: publicId,
			ScopeId:  organizationId,
			Type:     authType.String(),
			Name:     withName,
		},
	}
	return a, nil
}

// Clone creates a clone of the AuthMethod
func (a *AuthMethod) Clone() interface{} {
	cp := proto.Clone(a.AuthMethod)
	return &AuthMethod{
		AuthMethod: cp.(*store.AuthMethod),
	}
}

// VetForWrite implements db.VetForWrite() interface
func (a *AuthMethod) VetForWrite(ctx context.Context, r db.Reader, opType db.OpType, opt ...db.Option) error {
	if a.PublicId == "" {
		return errors.New("error public id is empty string for user write")
	}
	if a.ScopeId == "" {
		return errors.New("error scope id not set for user write")
	}
	// make sure the scope is valid for auth methods
	if err := a.scopeIsValid(ctx, r); err != nil {
		return err
	}
	return nil
}

// scopeIsValid makes sure the scope is an Organization
func (p *AuthMethod) scopeIsValid(ctx context.Context, r db.Reader) error {
	ps, err := LookupScope(ctx, r, p)
	if err != nil {
		return err
	}
	if ps.Type != OrganizationScope.String() {
		return errors.New("error scope is not an organization")
	}
	return nil
}

// GetScope returns the scope for the AuthMethod
func (p *AuthMethod) GetScope(ctx context.Context, r db.Reader) (*Scope, error) {
	return LookupScope(ctx, r, p)
}

// ResourceType returns the type of the AuthMethod
func (*AuthMethod) ResourceType() ResourceType { return ResourceTypeAuthMethod }

// Actions returns the  available actions for AuthMethod
func (*AuthMethod) Actions() map[string]Action {
	actions := CrudActions()
	return actions
}

// TableName returns the tablename to override the default gorm table name
func (p *AuthMethod) TableName() string {
	if p.tableName != "" {
		return p.tableName
	}
	return "iam_auth_method"
}

// SetTableName sets the tablename and satisfies the ReplayableMessage interface
func (p *AuthMethod) SetTableName(n string) {
	if n != "" {
		p.tableName = n
	}
}
