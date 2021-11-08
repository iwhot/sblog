package page

// GetOffset 获取偏移量
func GetOffset(page, pageSize int) int {
	if page < 1 {
		return 0
	}

	return (page - 1) * pageSize
}