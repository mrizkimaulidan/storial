package story

import "errors"

var (
	ErrStoryNotFound      = errors.New("story not found")
	ErrCoverImageNotFound = errors.New("cover image not found")
)
