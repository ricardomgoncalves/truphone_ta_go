package age

import (
	"log"
	"math"
	"sort"
	"time"

	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
)

func FindFastestGrowingAge(families []family.Family, membersFamilies map[string][]family.Member) []family.AgeFamily {
	agedFamilies := make([]family.AgeFamily, 0)

	for _, fam := range families {
		members, ok := membersFamilies[fam.Id]
		if !ok {
			log.Printf("Family %s does not have any member.\n", fam.Id)
			continue
		}
		agedFamilies = append(agedFamilies, calculateFastestGrowingAgeFamily(fam, members))
	}

	sort.SliceStable(agedFamilies, func(i, j int) bool {
		if agedFamilies[i].Age == agedFamilies[j].Age {
			return len(membersFamilies[agedFamilies[i].Id]) > len(membersFamilies[agedFamilies[j].Id])
		}

		return agedFamilies[i].Age > agedFamilies[j].Age
	})

	return agedFamilies
}

func calculateFastestGrowingAgeFamily(fam family.Family, members []family.Member) family.AgeFamily {
	if len(members) == 0 {
		return family.AgeFamily{
			Family: fam,
		}
	}

	firstBorn := time.Time{}
	lastBorn := time.Time{}

	for _, member := range members {
		if firstBorn.IsZero() {
			firstBorn = member.Birthday
		}

		if lastBorn.IsZero() {
			lastBorn = member.Birthday
		}

		if firstBorn.After(member.Birthday) {
			firstBorn = member.Birthday
		}

		if lastBorn.Before(member.Birthday) {
			lastBorn = member.Birthday
		}
	}

	duration := lastBorn.Sub(firstBorn)
	return family.AgeFamily{
		Family: fam,
		Age:    int64(math.Ceil(duration.Seconds() / float64(len(members)))),
	}
}
