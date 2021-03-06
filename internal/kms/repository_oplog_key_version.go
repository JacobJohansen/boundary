package kms

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/boundary/internal/errors"
	wrapping "github.com/hashicorp/go-kms-wrapping"
)

// CreateOplogKeyVersion inserts into the repository and returns the new key
// version with its PrivateId.  There are no valid options at this time.
func (r *Repository) CreateOplogKeyVersion(ctx context.Context, rkvWrapper wrapping.Wrapper, oplogKeyId string, key []byte, opt ...Option) (*OplogKeyVersion, error) {
	if rkvWrapper == nil {
		return nil, fmt.Errorf("create oplog key version: missing root key version wrapper: %w", errors.ErrInvalidParameter)
	}
	rootKeyVersionId := rkvWrapper.KeyID()
	switch {
	case !strings.HasPrefix(rootKeyVersionId, RootKeyVersionPrefix):
		return nil, fmt.Errorf("create oplog key version: root key version id %s doesn't start with prefix %s: %w", rootKeyVersionId, RootKeyVersionPrefix, errors.ErrInvalidParameter)
	case rootKeyVersionId == "":
		return nil, fmt.Errorf("create oplog key version: missing root key version id: %w", errors.ErrInvalidParameter)
	}
	if oplogKeyId == "" {
		return nil, fmt.Errorf("create oplog key version: missing oplog key id: %w", errors.ErrInvalidParameter)
	}
	if len(key) == 0 {
		return nil, fmt.Errorf("create oplog key version: missing key: %w", errors.ErrInvalidParameter)
	}
	kv := AllocOplogKeyVersion()
	id, err := newOplogKeyVersionId()
	if err != nil {
		return nil, fmt.Errorf("create oplog key version: %w", err)
	}
	kv.PrivateId = id
	kv.RootKeyVersionId = rootKeyVersionId
	kv.Key = key
	kv.OplogKeyId = oplogKeyId
	if err := kv.Encrypt(ctx, rkvWrapper); err != nil {
		return nil, fmt.Errorf("create oplog key version: encrypt: %w", err)
	}

	var returnedKey interface{}
	_, err = r.writer.DoTx(
		ctx,
		db.StdRetryCnt,
		db.ExpBackoff{},
		func(_ db.Reader, w db.Writer) error {
			returnedKey = kv.Clone()
			// no oplog entries for root key version
			if err := w.Create(ctx, returnedKey); err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("create oplog key version: %w for %s oplog key id", err, kv.OplogKeyId)
	}
	return returnedKey.(*OplogKeyVersion), err
}

// LookupOplogKeyVersion will look up a key version in the repository.  If
// the key version is not found, it will return nil, nil.
func (r *Repository) LookupOplogKeyVersion(ctx context.Context, keyWrapper wrapping.Wrapper, privateId string, opt ...Option) (*OplogKeyVersion, error) {
	if privateId == "" {
		return nil, fmt.Errorf("lookup oplog key version: missing private id: %w", errors.ErrInvalidParameter)
	}
	if keyWrapper == nil {
		return nil, fmt.Errorf("lookup oplog key version: missing key wrapper: %w", errors.ErrInvalidParameter)
	}
	k := AllocOplogKeyVersion()
	k.PrivateId = privateId
	if err := r.reader.LookupById(ctx, &k); err != nil {
		return nil, fmt.Errorf("lookup oplog key version: failed %w for %s", err, privateId)
	}
	if err := k.Decrypt(ctx, keyWrapper); err != nil {
		return nil, fmt.Errorf("lookup oplog key version: decrypt: %w", err)
	}
	return &k, nil
}

// DeleteOplogKeyVersion deletes the key version for the provided id from the
// repository returning a count of the number of records deleted.  All options
// are ignored.
func (r *Repository) DeleteOplogKeyVersion(ctx context.Context, privateId string, opt ...Option) (int, error) {
	if privateId == "" {
		return db.NoRowsAffected, fmt.Errorf("delete oplog key version: missing private id: %w", errors.ErrInvalidParameter)
	}
	k := AllocOplogKeyVersion()
	k.PrivateId = privateId
	if err := r.reader.LookupById(ctx, &k); err != nil {
		return db.NoRowsAffected, fmt.Errorf("delete oplog key version: failed %w for %s", err, privateId)
	}

	var rowsDeleted int
	_, err := r.writer.DoTx(
		ctx,
		db.StdRetryCnt,
		db.ExpBackoff{},
		func(_ db.Reader, w db.Writer) (err error) {
			dk := k.Clone()
			// no oplog entries for the key version
			rowsDeleted, err = w.Delete(ctx, dk)
			if err == nil && rowsDeleted > 1 {
				return errors.ErrMultipleRecords
			}
			return err
		},
	)
	if err != nil {
		return db.NoRowsAffected, fmt.Errorf("delete oplog key version: %s: %w", privateId, err)
	}
	return rowsDeleted, nil
}

// LatestOplogKeyVersion searches for the key version with the highest
// version number.  When no results are found, it returns nil,
// errors.ErrRecordNotFound.
func (r *Repository) LatestOplogKeyVersion(ctx context.Context, rkvWrapper wrapping.Wrapper, oplogKeyId string, opt ...Option) (*OplogKeyVersion, error) {
	if oplogKeyId == "" {
		return nil, fmt.Errorf("latest oplog key version: missing oplog key id: %w", errors.ErrInvalidParameter)
	}
	if rkvWrapper == nil {
		return nil, fmt.Errorf("latest oplog key version: missing root key version wrapper: %w", errors.ErrInvalidParameter)
	}
	var foundKeys []*OplogKeyVersion
	if err := r.reader.SearchWhere(ctx, &foundKeys, "oplog_key_id = ?", []interface{}{oplogKeyId}, db.WithLimit(1), db.WithOrder("version desc")); err != nil {
		return nil, fmt.Errorf("latest oplog key version: failed %w for %s", err, oplogKeyId)
	}
	if len(foundKeys) == 0 {
		return nil, errors.ErrRecordNotFound
	}
	if err := foundKeys[0].Decrypt(ctx, rkvWrapper); err != nil {
		return nil, fmt.Errorf("latest oplog key version: %w", err)
	}
	return foundKeys[0], nil
}

// ListOplogKeyVersions will lists versions of a key.  Supports the WithLimit option.
func (r *Repository) ListOplogKeyVersions(ctx context.Context, rkvWrapper wrapping.Wrapper, oplogKeyId string, opt ...Option) ([]DekVersion, error) {
	if oplogKeyId == "" {
		return nil, fmt.Errorf("list oplog key versions: missing oplog key id %w", errors.ErrInvalidParameter)
	}
	if rkvWrapper == nil {
		return nil, fmt.Errorf("list oplog key versions: missing root key version wrapper: %w", errors.ErrInvalidParameter)
	}
	var versions []*OplogKeyVersion
	err := r.list(ctx, &versions, "oplog_key_id = ?", []interface{}{oplogKeyId}, opt...)
	if err != nil {
		return nil, fmt.Errorf("list oplog key versions: %w", err)
	}
	for i, k := range versions {
		if err := k.Decrypt(ctx, rkvWrapper); err != nil {
			return nil, fmt.Errorf("list oplog key versions: error decrypting key num %d: %w", i, err)
		}
	}
	dekVersions := make([]DekVersion, 0, len(versions))
	for _, version := range versions {
		dekVersions = append(dekVersions, version)
	}
	return dekVersions, nil
}
