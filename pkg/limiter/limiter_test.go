package limiter

import (
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLimitAgedFamilies(t *testing.T) {
	t.Run("should return slice on limit 0", func(t *testing.T) {
		agedFamilies := []family.AgeFamily{
			{
				Family: family.Family{
					Id: "id",
				},
				Age: 1,
			},
		}
		agedFamiliesOutput := LimitAgedFamilies(agedFamilies, 0, 0)
		assert.Equal(t, agedFamilies, agedFamiliesOutput)
	})
	t.Run("should return empty slice on offset greater than length", func(t *testing.T) {
		agedFamilies := []family.AgeFamily{
			{
				Family: family.Family{
					Id: "id",
				},
				Age: 1,
			},
		}
		agedFamiliesOutput := LimitAgedFamilies(agedFamilies, 3, 1)
		assert.Equal(t, []family.AgeFamily{}, agedFamiliesOutput)
	})
	t.Run("should return limited slice", func(t *testing.T) {
		agedFamilies := []family.AgeFamily{
			{
				Family: family.Family{
					Id: "id",
				},
				Age: 1,
			},
			{
				Family: family.Family{
					Id: "id1",
				},
				Age: 2,
			},
		}
		agedFamiliesOutput := LimitAgedFamilies(agedFamilies, 0, 1)
		assert.Equal(t, []family.AgeFamily{
			{
				Family: family.Family{
					Id: "id",
				},
				Age: 1,
			},
		}, agedFamiliesOutput)
	})
	t.Run("should return offset slice", func(t *testing.T) {
		agedFamilies := []family.AgeFamily{
			{
				Family: family.Family{
					Id: "id",
				},
				Age: 1,
			},
			{
				Family: family.Family{
					Id: "id1",
				},
				Age: 2,
			},
		}
		agedFamiliesOutput := LimitAgedFamilies(agedFamilies, 1, 1)
		assert.Equal(t, []family.AgeFamily{
			{
				Family: family.Family{
					Id: "id1",
				},
				Age: 2,
			},
		}, agedFamiliesOutput)
	})
	t.Run("should return offset and limited slice", func(t *testing.T) {
		agedFamilies := []family.AgeFamily{
			{
				Family: family.Family{
					Id: "id",
				},
				Age: 1,
			},
			{
				Family: family.Family{
					Id: "id1",
				},
				Age: 2,
			},
			{
				Family: family.Family{
					Id: "id3",
				},
				Age: 7,
			},
		}
		agedFamiliesOutput := LimitAgedFamilies(agedFamilies, 1, 1)
		assert.Equal(t, []family.AgeFamily{
			{
				Family: family.Family{
					Id: "id1",
				},
				Age: 2,
			},
		}, agedFamiliesOutput)
	})
}

func TestLimitMembers(t *testing.T) {
	t.Run("should return slice on limit 0", func(t *testing.T) {
		members := []family.Member{
			{
				Id: "1",
			},
		}
		membersOutput := LimitMembers(members, 0, 0)
		assert.Equal(t, members, membersOutput)
	})
	t.Run("should return empty slice on offset greater than length", func(t *testing.T) {
		members := []family.Member{
			{
				Id: "1",
			},
		}
		membersOutput := LimitMembers(members, 3, 1)
		assert.Equal(t, []family.Member{}, membersOutput)
	})
	t.Run("should return limited slice", func(t *testing.T) {
		members := []family.Member{
			{
				Id: "id",
			},
			{
				Id: "id2",
			},
		}
		membersOutput := LimitMembers(members, 0, 1)
		assert.Equal(t, []family.Member{
			{
				Id: "id",
			},
		}, membersOutput)
	})
	t.Run("should return offset slice", func(t *testing.T) {
		members := []family.Member{
			{
				Id: "id",
			},
			{
				Id: "id2",
			},
		}
		membersOutput := LimitMembers(members, 1, 1)
		assert.Equal(t, []family.Member{
			{
				Id: "id2",
			},
		}, membersOutput)
	})
	t.Run("should return offset and limited slice", func(t *testing.T) {
		members := []family.Member{
			{
				Id: "id",
			},
			{
				Id: "id1",
			},
			{
				Id: "id2",
			},
		}
		membersOutput := LimitMembers(members, 1, 1)
		assert.Equal(t, []family.Member{
			{
				Id: "id1",
			},
		}, membersOutput)
	})
}
