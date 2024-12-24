package state

import (
	"ericarthurc.com/internal/database"
	"ericarthurc.com/internal/model"
	"github.com/puzpuzpuz/xsync/v3"
)

type State struct {
	DbPool  *database.DbPool
	PostMap *xsync.MapOf[string, model.Post]
}

func NewState(dbPool *database.DbPool) *State {
	postMap := xsync.NewMapOf[string, model.Post]()
	// postMap.Store("hello-world", model.Post{})

	return &State{
		DbPool:  dbPool,
		PostMap: postMap,
	}
}
