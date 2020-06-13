package duplicates

import (
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/graphs"
)

func FindPossibleDuplicates(inputMembers []family.Member) []family.Member {
	possibleDuplicatesIds := make(map[string]struct{})
	possibleDuplicates := make([]family.Member, 0)

	members := make([]family.Member, len(inputMembers))
	copy(members, inputMembers)

	graph := graphs.NewFamilyGraph()
	roots := make([]family.Member, 0)
	for _, member := range inputMembers {
		graph.AddNode(member)

		if member.FatherId != nil {
			graph.AddEdge(*member.FatherId, member)
		}

		if member.MotherId != nil {
			graph.AddEdge(*member.MotherId, member)
		}

		if member.FatherId == nil || member.MotherId == nil {
			roots = append(roots, member)
		}
	}

	possibleDuplicates = append(possibleDuplicates, getPossibleDuplicates(possibleDuplicatesIds, roots)...)
	graph.Traverse(func(children []family.Member) {
		possibleDuplicates = append(possibleDuplicates, getPossibleDuplicates(possibleDuplicatesIds, children)...)
	})

	return possibleDuplicates
}

func getPossibleDuplicates(duplicatesIds map[string]struct{}, members []family.Member) []family.Member {
	duplicateMembers := make([]family.Member, 0)

	for i := 0; i < len(members); i++ {
		for j := i + 1; j < len(members); j++ {
			if !isPossibleDuplicates(duplicatesIds, members[i], members[j]) {
				continue
			}

			if _, ok := duplicatesIds[members[i].Id]; !ok {
				duplicatesIds[members[i].Id] = struct{}{}
				duplicateMembers = append(duplicateMembers, members[i])
			}

			if _, ok := duplicatesIds[members[j].Id]; !ok {
				duplicatesIds[members[j].Id] = struct{}{}
				duplicateMembers = append(duplicateMembers, members[j])
			}
		}
	}

	return duplicateMembers
}

func isPossibleDuplicates(duplicatesIds map[string]struct{}, m1, m2 family.Member) bool {
	if !m1.HasCommonName(m2) {
		return false
	}

	if !m1.HasSimilarBirthday(m2) {
		return false
	}

	if (m1.IsMissingMother() || m2.IsMissingMother()) && (m1.IsMissingFather() || m2.IsMissingFather()) {
		return true
	}

	return checkMembersMothers(duplicatesIds, m1, m2) && checkMembersFathers(duplicatesIds, m1, m2)
}

func checkMembersMothers(duplicatesIds map[string]struct{}, m1, m2 family.Member) bool {
	if m1.HasSameMother(m2) {
		return true
	}

	if m1.IsMissingMother() || m2.IsMissingMother() {
		return true
	}

	m1MotherPossibleDuplicate := false
	m2MotherPossibleDuplicate := false

	if !m1.IsMissingMother() {
		_, ok := duplicatesIds[*m1.MotherId]
		m1MotherPossibleDuplicate = ok
	}

	if !m2.IsMissingMother() {
		_, ok := duplicatesIds[*m2.MotherId]
		m2MotherPossibleDuplicate = ok
	}

	return m1MotherPossibleDuplicate && m2MotherPossibleDuplicate
}

func checkMembersFathers(duplicatesIds map[string]struct{}, m1, m2 family.Member) bool {
	if m1.HasSameFather(m2) {
		return true
	}

	if m1.IsMissingFather() || m2.IsMissingFather() {
		return true
	}

	m1FatherPossibleDuplicate := false
	m2FatherPossibleDuplicate := false

	if !m1.IsMissingFather() {
		_, ok := duplicatesIds[*m1.FatherId]
		m1FatherPossibleDuplicate = ok
	}

	if !m2.IsMissingFather() {
		_, ok := duplicatesIds[*m2.FatherId]
		m2FatherPossibleDuplicate = ok
	}

	return m1FatherPossibleDuplicate && m2FatherPossibleDuplicate
}
