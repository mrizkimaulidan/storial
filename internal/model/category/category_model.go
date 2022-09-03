package category

type CategoryResponse struct {
	Id          uint64 `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	StoryCounts uint64 `json:"storyCounts"`
}

type CategoryResponseByStorySlug struct {
	Id   uint64 `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}
