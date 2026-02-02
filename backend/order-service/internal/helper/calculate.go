package helper

func CalculateOffset(page int32, limit int32) int32 {
	return (page - 1) * limit
}
