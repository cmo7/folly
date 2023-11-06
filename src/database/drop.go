package database

func Drop() error {
	for _, task := range migrationTasks {
		if task.DropOnFlush {
			err := DB.Migrator().DropTable(task.Model)
			if err != nil {
				return err
			}
		}
	}
	DB.Migrator().DropTable("user_roles")
	return nil
}
