package helper

func CalculateOffset(page, limit int32) int32 {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit
	return offset
}
