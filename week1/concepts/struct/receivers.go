package concepts

type Rectangle struct {
	Width, Height int
}

// 方法：值接收器
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// 方法：指標接收器
func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}

// 指標接收器 vs 值接收器
// 值接收器：不修改原始數據，適合不可變操作。
// 指標接收器：可修改原始數據，適合需要改變結構體狀態或避免拷貝大的結構體。
