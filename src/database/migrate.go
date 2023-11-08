package database

import "folly/src/lib/common"

type OnFlush string // "DROP" | "TRUNCATE" | "NO ACTION"

const (
	DROP     OnFlush = "DROP"
	TRUNCATE OnFlush = "TRUNCATE"
	NOACTION OnFlush = "NO ACTION"
)

type MigrationTask struct {
	Model           common.Entity
	DropOnFlush     bool
	TruncateOnFlush bool
}

var migrationTasks []MigrationTask

func RegisterMigration(task *MigrationTask) {
	migrationTasks = append(migrationTasks, *task)
}

func Migrate() error {

	models := []interface{}{}
	for _, task := range migrationTasks {
		models = append(models, task.Model)
	}
	err := DB.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}
