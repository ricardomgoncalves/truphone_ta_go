package postgres

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"time"
)

func CreateTables(db *gorm.DB) error {
	return db.AutoMigrate(&familyRow{}, &memberRow{}).Error
}

func Populate(db *gorm.DB) error {
	familyUuid := uuid.New()
	familyRow := newFamilyRow(&family.Family{
		Id:          familyUuid,
		Name:        "Family 1",
		CountryCode: "PT",
	})
	if err := db.Create(familyRow).Error; err != nil {
		return err
	}

	fatherUuid := uuid.New()
	motherUuid := uuid.New()
	fatherRow := newMemberRow(&family.Member{
		Id:         fatherUuid,
		FamilyId:   familyUuid,
		FirstName:  "Father",
		MiddleName: "0",
		LastName:   "1",
		SpouseId:   &motherUuid,
		Birthday:   time.Now(),
	})
	motherRow := newMemberRow(&family.Member{
		Id:         motherUuid,
		FamilyId:   familyUuid,
		FirstName:  "Mother",
		MiddleName: "0",
		LastName:   "1",
		SpouseId:   &fatherUuid,
		Birthday:   time.Now(),
	})
	childRow := newMemberRow(&family.Member{
		Id:         uuid.New(),
		FamilyId:   familyUuid,
		FirstName:  "Child",
		MiddleName: "0",
		LastName:   "1",
		FatherId:   &fatherUuid,
		MotherId:   &motherUuid,
		Birthday:   time.Now(),
	})

	if err := db.Create(fatherRow).Error; err != nil {
		return err
	}
	if err := db.Create(motherRow).Error; err != nil {
		return err
	}
	if err := db.Create(childRow).Error; err != nil {
		return err
	}

	return nil
}
