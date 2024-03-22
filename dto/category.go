package dto

import "errors"

const (
	MESSAGE_CATEGORY_CREATE_SUCCESS = "Category created successfully"
	MESSAGE_CATEGORY_FETCH_SUCCESS  = "Category fetched successfully"
)

var (
	ErrCategoryCreateFailed = errors.New("failed to create category")
	ErrCategoryNotFound     = errors.New("category not found")
)

type (
	CategoryUploadRequest struct {
		Name string `json:"name"`
	}

	CategoryGeneralResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)
