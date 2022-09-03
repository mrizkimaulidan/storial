package chapter

import "errors"

var (
	ErrChapterNotFound          = errors.New("chapter not found")
	ErrCannotLikeYourOwnChapter = errors.New("cannot like your own chapter")
)
