package queue

import (
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
)

type MemberQueue struct {
	items []family.Member
}

func NewMemberQueue() *MemberQueue {
	q := MemberQueue{}
	q.items = []family.Member{}
	return &q
}

func (s *MemberQueue) Enqueue(t family.Member) {
	s.items = append(s.items, t)
}

func (s *MemberQueue) Dequeue() family.Member {
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return item
}

func (s *MemberQueue) IsEmpty() bool {
	return len(s.items) == 0
}
