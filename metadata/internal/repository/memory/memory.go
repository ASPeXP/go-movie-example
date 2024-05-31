package memory

import (
	"context"
	"sync"

	"github.com/aspexp/movieexample/metadata/internal/repository"
	"github.com/aspexp/movieexample/metadata/pkg/model"
)

// repository define a memory movie metadata repository.
type Repository struct {
	sync.RWMutex 
	data map[string]*model.Metadata 

}
//New creates a new memory repositiory.
func New() *Repository {
	return &Repository{ data:map[string]*model.Metadata{}}
}
//Get retrieve movie metadata for by movie id 
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()
	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil 
}
//put adds movie metadata for a given movie id.
func (r *Repository) Put (_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata 
	return nil  

}