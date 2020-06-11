package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"time"
)

func CreateTables(db *gorm.DB) error {
	return db.AutoMigrate(&familyRow{}, &memberRow{}).Error
}

func Populate(db *gorm.DB) error {
	familyUuid := "9fadb3cc-74ee-4ff7-8bd5-ffa1d34da038"
	fatherUuid := "11810f35-309a-4836-b7e9-1fee57bed924"
	motherUuid := "5d32fc95-ce3e-4b18-8680-fce1e6f8e3ea"
	childUuid := "8d9d3e48-4b34-40de-b986-1042b1a42f86"

	familyRow := newFamilyRow(&family.Family{
		Id:          familyUuid,
		Name:        "Family 1",
		CountryCode: "PT",
	})
	if err := db.Create(familyRow).Error; err != nil {
		return err
	}

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
		Id:         childUuid,
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
