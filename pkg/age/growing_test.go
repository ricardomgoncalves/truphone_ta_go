package age

import (
	"testing"
	"time"

	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
)

func TestFindFastestGrowingAge(t *testing.T) {
	t.Run("should return empty slice on empty members", func(t *testing.T) {
		families := []family.Family{
			{
				Id: "id",
			},
		}
		members := map[string][]family.Member{}
		agedFamilies := FindFastestGrowingAge(families, members)
		assert.Equal(t, []family.AgeFamily{}, agedFamilies)
	})
	t.Run("should return calculated age on 1 family", func(t *testing.T) {
		families := []family.Family{
			{
				Id: "id",
			},
		}
		birthday := time.Now()
		members := map[string][]family.Member{
			"id": {
				{Birthday: birthday},
				{Birthday: birthday.Add(time.Second * 6)},
			},
		}
		agedFamilies := FindFastestGrowingAge(families, members)
		assert.Equal(t, []family.AgeFamily{
			{
				Family: family.Family{
					Id: "id",
				},
				Age: 3,
			},
		}, agedFamilies)
	})
	t.Run("should return calculated on families", func(t *testing.T) {
		families := []family.Family{
			{
				Id: "id",
			},
			{
				Id: "id1",
			},
		}
		birthday := time.Now()
		members := map[string][]family.Member{
			"id": {
				{Birthday: birthday},
				{Birthday: birthday.Add(time.Second * 2)},
				{Birthday: birthday.Add(time.Second * 4)},
			},
			"id1": {
				{Birthday: birthday},
				{Birthday: birthday.Add(time.Second * 6)},
			},
		}
		agedFamilies := FindFastestGrowingAge(families, members)
		assert.Equal(t, []family.AgeFamily{
			{
				Family: family.Family{
					Id: "id1",
				},
				Age: 3,
			},
			{
				Family: family.Family{
					Id: "id",
				},
				Age: 2,
			},
		}, agedFamilies)
	})
	t.Run("should return high number of members families first", func(t *testing.T) {
		families := []family.Family{
			{
				Id: "id",
			},
			{
				Id: "id1",
			},
		}
		birthday := time.Now()
		members := map[string][]family.Member{
			"id": {
				{Birthday: birthday},
				{Birthday: birthday.Add(time.Second * 6)},
			},
			"id1": {
				{Birthday: birthday},
				{Birthday: birthday.Add(time.Second * 2)},
				{Birthday: birthday.Add(time.Second * 7)},
			},
		}
		agedFamilies := FindFastestGrowingAge(families, members)
		assert.Equal(t, []family.AgeFamily{
			{
				Family: family.Family{
					Id: "id1",
				},
				Age: 3,
			},
			{
				Family: family.Family{
					Id: "id",
				},
				Age: 3,
			},
		}, agedFamilies)
	})
}

func TestCalculateFastestGrowingAgeFamily(t *testing.T) {
	t.Run("should return the same family as input", func(t *testing.T) {
		fam := family.Family{
			Id:          "id",
			Name:        "Fam",
			CountryCode: "PT",
		}
		agedFamily := calculateFastestGrowingAgeFamily(fam, []family.Member{})
		assert.Equal(t, fam, agedFamily.Family)
		assert.Equal(t, int64(0), agedFamily.Age)
	})
	t.Run("should return 0 on 1 member family", func(t *testing.T) {
		fam := family.Family{
			Id:          "id",
			Name:        "Fam",
			CountryCode: "PT",
		}
		members := []family.Member{
			{
				Id:       "id_1",
				Birthday: time.Now(),
			},
		}
		agedFamily := calculateFastestGrowingAgeFamily(fam, members)
		assert.Equal(t, fam, agedFamily.Family)
		assert.Equal(t, int64(0), agedFamily.Age)
	})
	t.Run("should return difference", func(t *testing.T) {
		fam := family.Family{
			Id:          "id",
			Name:        "Fam",
			CountryCode: "PT",
		}
		birthday := time.Now()
		members := []family.Member{
			{
				Id:       "id_1",
				Birthday: birthday.Add(time.Second * 3),
			},
			{
				Id:       "id_2",
				Birthday: birthday,
			},
		}
		agedFamily := calculateFastestGrowingAgeFamily(fam, members)
		assert.Equal(t, fam, agedFamily.Family)
		assert.Equal(t, int64(2), agedFamily.Age)
	})
	t.Run("should return difference with more members", func(t *testing.T) {
		fam := family.Family{
			Id:          "id",
			Name:        "Fam",
			CountryCode: "PT",
		}
		birthday := time.Now()
		members := []family.Member{
			{
				Id:       "id_1",
				Birthday: birthday.Add(time.Second * 3),
			},
			{
				Id:       "id_2",
				Birthday: birthday,
			},
			{
				Id:       "id_3",
				Birthday: birthday.Add(time.Second * 10),
			},
		}
		agedFamily := calculateFastestGrowingAgeFamily(fam, members)
		assert.Equal(t, fam, agedFamily.Family)
		assert.Equal(t, int64(4), agedFamily.Age)
	})
}
