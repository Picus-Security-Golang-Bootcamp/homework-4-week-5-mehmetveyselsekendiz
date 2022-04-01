package service


func Migrate(db *gorm.DB) {
	workplacePrototype := &Workplace{}
	workerPrototype := &Worker{}
	db.AutoMigrate(workplacePrototype, workerPrototype)
 }