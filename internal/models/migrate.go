package models

import (
	"git.code.oa.com/trpc-go/trpc-go/log"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"sort"
)

func autoMigrate() error {
	db := GetDBConn()
	m := gormigrate.New(db, gormigrate.DefaultOptions, migration())

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
		return err
	}
	log.Infof("Migration did run successfully")

	return nil
}

func migration() []*gormigrate.Migration {
	migrates := map[string]func(tx *gorm.DB) error{
		"2023122601": migrate2023122601,
	}

	keys := make([]string, 0)
	for key := range migrates {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	mig := make([]*gormigrate.Migration, 0)
	for _, k := range keys {
		mig = append(mig, &gormigrate.Migration{ID: k, Migrate: migrates[k],
			Rollback: func(tx *gorm.DB) error {
				return nil
			}})
	}
	return mig
}

func migrate2023122601(tx *gorm.DB) error {
	if err := tx.Migrator().CreateTable(&ApiApplies{}); err != nil {
		return err
	}
	return nil
}
