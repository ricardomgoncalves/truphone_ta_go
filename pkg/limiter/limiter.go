package limiter

import (
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
)

func LimitAgedFamilies(slice []family.AgeFamily, offset uint32, limit uint32) []family.AgeFamily {
	if limit == 0 {
		return slice
	}

	if offset > uint32(len(slice)) {
		return []family.AgeFamily{}
	}

	newSlice := make([]family.AgeFamily, limit)
	copy(newSlice, slice[offset:])
	return newSlice
}

func LimitMembers(slice []family.Member, offset uint32, limit uint32) []family.Member {
	if limit == 0 {
		return slice
	}

	if offset > uint32(len(slice)) {
		return []family.Member{}
	}

	newSlice := make([]family.Member, limit)
	copy(newSlice, slice[offset:])
	return newSlice
}
