package signServer

import "time"

// 获取1970年1月1日 距今多少天
func getTodayNum() int64 {
	t := time.Now().Unix()
	diff := float64(t)
	d := int64(diff * 10 / (3600 * 24))

	if d%10 > 0 {
		d = d/10 + 1
	} else {
		d = d / 10
	}
	// 数字太多不利于存储, 因为会获取近7天的数据, 所以 尽量不要让减去的数据接近 1970-1-1 至今的天 靠近
	return d
}
