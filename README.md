# time-slot
多个时间段合并计算
# 使用实例
```golang
package main

import (
	"fmt"
	timeslot "github.com/mosesyu95/time-slot"
)

func main() {
	slot := [][2]int64{{1, 2}, {4, 6}, {6, 8}, {5, 3}, {11, 22}, {2, 7}, {6, 9}, {1, 6}, {33, 44}, {9, 10}}
	slots := timeslot.MergeTimeSlot(slot)
	fmt.Println(slots)
	total := timeslot.TimeTotal(slots)
	fmt.Println(total)
}
```
