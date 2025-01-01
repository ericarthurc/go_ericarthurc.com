package state

import (
	"slices"
	"sync"

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
	Mu                         sync.RWMutex
	FeaturedPostsMetaSorted    []model.Post
	NonFeaturedPostsMetaSorted []model.Post
}

func NewState(dbPool *database.DbPool) (*State, error) {
	postMap := xsync.NewMapOf[string, model.Post]()

	// call database and get all the posts
	posts, err := model.GetAllPosts(dbPool)
	if err != nil {
		return nil, err
	}

	var featured []model.Post
	var nonFeatured []model.Post

	var wg sync.WaitGroup
	errCh := make(chan error, len(posts))

	for _, p := range posts {
		wg.Add(1)

		go func(post *model.Post) {
			defer wg.Done()

			if err := post.MarkdownToHTML(); err != nil {
				errCh <- err
				return
			}

			post.SkillsToSVGs()

			// slices.Sort(post.Skills)
			slices.Sort(post.Categories)

			// store the post in the postMap
			postMap.Store(post.Slug, *post)

			if p.Featured {
				featured = append(featured, *post)
			} else {
				nonFeatured = append(nonFeatured, *post)
			}
		}(&p)
	}

	wg.Wait()

	close(errCh)

	// Check for any errors in the channel
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	slices.SortFunc(nonFeatured, func(a, b model.Post) int {
		return b.Date.Compare(a.Date)
	})
	slices.SortFunc(featured, func(a, b model.Post) int {
		return b.Date.Compare(a.Date)
	})

	return &State{
		DbPool:  dbPool,
		PostMap: postMap,
		PostMeta: PostMeta{
			FeaturedPostsMetaSorted:    featured,
			NonFeaturedPostsMetaSorted: nonFeatured,
		},
	}, err
}

func (s *State) UpdateState() error {
	return nil
}
