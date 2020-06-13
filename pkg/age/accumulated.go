package age

import (
	"log"
	"sort"
	"time"

	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
)

func FindAccumulatedAge(families []family.Family, membersFamilies map[string][]family.Member) []family.AgeFamily {
	agedFamilies := make([]family.AgeFamily, 0)

	for _, fam := range families {
		members, ok := membersFamilies[fam.Id]
		if !ok {
			log.Printf("Family %s does not have any member.\n", fam.Id)
			continue
		}
		agedFamilies = append(agedFamilies, calculateAccumulatedAgeFamily(fam, members))
	}

	sort.SliceStable(agedFamilies, func(i, j int) bool {
		if agedFamilies[i].Age == agedFamilies[j].Age {
			return len(membersFamilies[agedFamilies[i].Id]) > len(membersFamilies[agedFamilies[j].Id])
		}

		return agedFamilies[i].Age > agedFamilies[j].Age
	})

	return agedFamilies
}

func calculateAccumulatedAgeFamily(fam family.Family, members []family.Member) family.AgeFamily {
	if len(members) == 0 {
		return family.AgeFamily{
			Family: fam,
		}
	}

	var age time.Duration

	for _, member := range members {
		age += time.Now().Sub(member.Birthday)
	}

	return family.AgeFamily{
		Family: fam,
		Age:    int64(age.Seconds()),
	}
}
