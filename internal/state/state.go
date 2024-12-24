package state

import (
	"slices"

	"ericarthurc.com/internal/database"
	"ericarthurc.com/internal/model"
	"github.com/puzpuzpuz/xsync/v3"
)

type State struct {
	DbPool  *database.DbPool
	PostMap *xsync.MapOf[string, model.Post]
	PostMeta
}

type PostMeta struct {
	FeaturedPostsMetaSorted    []model.Post
	NonFeaturedPostsMetaSorted []model.Post
}

func NewState(dbPool *database.DbPool) (*State, error) {
	postMap := xsync.NewMapOf[string, model.Post]()
	// postMap.Store("hello-world", model.Post{})

	// call database and get all the posts
	posts, err := model.GetAllPosts(dbPool)
	if err != nil {
		return nil, err
	}

	// Sort the posts by date
	slices.SortFunc(posts, func(a, b model.Post) int {
		return a.Date.Compare(b.Date)
	})

	var featured []model.Post
	var nonFeatured []model.Post

	for _, p := range posts {
		err := p.MarkdownToHTML()
		if err != nil {
			return nil, err
		}

		// Store the post in the postMap
		postMap.Store(p.Slug, p)

		if p.Featured {
			featured = append(featured, p)
		} else {
			nonFeatured = append(nonFeatured, p)
		}
	}

	return &State{
		DbPool:  dbPool,
		PostMap: postMap,
		PostMeta: PostMeta{
			FeaturedPostsMetaSorted:    featured,
			NonFeaturedPostsMetaSorted: nonFeatured,
		},
	}, err
}
