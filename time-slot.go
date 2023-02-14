package timeslot

import (
	"sort"
)

type TimeSlot struct {
	Start int64
	End   int64
	Next  *TimeSlot
}

// MergeTimeSlot 合并重叠时间段
func MergeTimeSlot(timeSlots [][2]int64) [][2]int64 {
	var slot2 [][2]int64
	sort.Slice(timeSlots, func(i, j int) bool {
		return timeSlots[i][0] < timeSlots[j][1]
	})
	var (
		head *TimeSlot
		tmp  *TimeSlot
	)
	for _, t := range timeSlots {
		start := t[0]
		end := t[1]
		if start > end {
			continue
		}
		if tmp == nil {
			tmp = &TimeSlot{
				Start: start,
				End:   end,
				Next:  nil,
			}
			head = tmp
			continue
		}
		if start > tmp.End {
			tmp.Next = &TimeSlot{
				Start: start,
				End:   end,
				Next:  nil,
			}
			tmp = tmp.Next
			continue
		}
		if end > tmp.End {
			tmp.End = end
		}
	}
	for head != nil {
		slot2 = append(slot2, [2]int64{head.Start, head.End})
		head = head.Next
	}
	return slot2
}

// TimeTotal 多个时间段合并查询总时间长度
func TimeTotal(timeSlots [][2]int64) int64 {
	var total int64
	slot2 := MergeTimeSlot(timeSlots)
	for _, slot := range slot2 {
		total += slot[1] - slot[0]
	}
	return total
}
