package resolver

import (
	"context"

	"github.com/romshark/dgraph_graphql_go/store"
	strerr "github.com/romshark/dgraph_graphql_go/store/errors"
)

// EditPost resolves Mutation.editPost
func (rsv *Resolver) EditPost(
	ctx context.Context,
	params struct {
		Post        string
		Editor      string
		NewTitle    *string
		NewContents *string
	},
) (*Post, error) {
	if params.NewTitle == nil && params.NewContents == nil {
		err := strerr.New(strerr.ErrInvalidInput, "no changes")
		rsv.error(ctx, err)
		return nil, err
	}

	transactRes, err := rsv.str.EditPost(
		ctx,
		store.ID(params.Post),
		store.ID(params.Editor),
		params.NewTitle,
		params.NewContents,
	)
	if err != nil {
		rsv.error(ctx, err)
		return nil, err
	}

	return &Post{
		root:      rsv,
		uid:       transactRes.UID,
		id:        store.ID(params.Post),
		creation:  transactRes.CreationTime,
		title:     transactRes.Title,
		contents:  transactRes.Contents,
		authorUID: transactRes.AuthorUID,
	}, nil
}