package database

func Drop() {
	for _, task := range migrationTasks {
		if task.DropOnFlush {
			DB.Migrator().DropTable(task.Model)
		}
	}
}
