package form

type UtilEncreptForm struct {
	Data string `json:"data" form:"data" binding:"required"`
}

type UtilDecryptForm struct {
	Data string `json:"data" form:"data" binding:"required"`
}

type UtilUseDBSampleForm struct {
	Id int64 `json:"id" form:"id" binding:"required"`
}
