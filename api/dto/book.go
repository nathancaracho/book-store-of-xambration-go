package dto

type Book struct {
	Name             string   `json:"name" binding:"required"`
	Description      string   `json:"description" binding:"required"`
	BriefDescription string   `json:"briefDescription" binding:"required"`
	Authors          []Author `json:"authors" binding:"required"`
	Value            float32  `json:"value" binding:"required"`
}
