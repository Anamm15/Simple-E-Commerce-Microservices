package helper

func CalculateOffset(page *int32, limit *int32) int32 {
	if page != nil && *page > 0 && limit != nil && *limit > 0 {
		return (*page - 1) * *limit
	}

	return 0
}
