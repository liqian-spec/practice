package migrate

import (
	"github.com/practice/pkg/console"
	"github.com/practice/pkg/database"
	"github.com/practice/pkg/file"
	"gorm.io/gorm"
	"os"
)

type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

func NewMigrator() *Migrator {

	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}

	migrator.createMigrationsTable()

	return migrator
}

func (migrator *Migrator) createMigrationsTable() {
	migration := Migration{}

	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

func (migrator *Migrator) Up() {

	migrateFiles := migrator.readAllMigrationFiles()

	batch := migrator.getBatch()

	migrations := []Migration{}
	migrator.DB.Find(&migrations)

	runed := false

	for _, mfile := range migrateFiles {

		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to date.")
	}
}

func (migrator *Migrator) Rollback() {

	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	migrations := []Migration{}
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}

}

func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {

	runed := false

	for _, _migration := range migrations {

		console.Warning("rollback " + _migration.Migration)

		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}

		runed = true

		migrator.DB.Delete(&_migration)

		console.Success("finish " + mfile.FileName)
	}
	return runed
}

func (migrator *Migrator) getBatch() int {

	batch := 1

	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {

	files, err := os.ReadDir(migrator.Folder)
	console.ExitIf(err)

	var migrateFiles []MigrationFile
	for _, f := range files {

		fileName := file.FileNameWithoutExtension(f.Name())

		mfile := getMigrationFile(fileName)

		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrateFiles, mfile)
		}
	}
	return migrateFiles
}

func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {

	if mfile.Up != nil {
		console.Warning("migrating" + mfile.FileName)
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		console.Success("migrated " + mfile.FileName)
	}

	err := migrator.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}
