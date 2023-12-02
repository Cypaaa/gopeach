package gopeach

import (
	"gorm.io/gorm"
)

// Database struct
type Database struct {
	conn      *gorm.DB
	dialector gorm.Dialector
}

// NewDatabase creates a new database connection
func NewDatabase(dialector gorm.Dialector) (*Database, error) {
	db := &Database{
		dialector: dialector,
	}
	err := db.Init()
	return db, err
}

// Conn returns the database connection
func (d *Database) Conn() *gorm.DB {
	return d.conn
}

// Init initializes the database connection
func (d *Database) Init() error {
	conn, err := gorm.Open(d.dialector, &gorm.Config{})
	if err != nil {
		return err
	}
	d.conn = conn
	return nil
}

// Close closes the database connection
func (d *Database) Close() error {
	db, err := d.conn.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

// Migrate migrates the database
func (d *Database) Migrate(i ...interface{}) error {
	return d.conn.AutoMigrate(i...)
}

// Save saves the model
func (d *Database) Save(i interface{}) error {
	return d.conn.Save(i).Error
}

// Delete deletes the model (soft delete: set deleted_at field)
func (d *Database) Delete(i interface{}) error {
	return d.conn.Delete(i).Error
}

// Destroy destroys the model (hard delete: delete entry from database)
func (d *Database) Destroy(i interface{}) error {
	return d.conn.Unscoped().Delete(i).Error
}

// First finds the first model by custom cudition from the same structure (findOne)
func (d *Database) First(i, c interface{}) error {
	return d.conn.First(i, c).Error
}

// FirstByQuery finds the first model by query (findOne)
func (d *Database) FirstByQuery(i interface{}, query interface{}, args ...interface{}) error {
	return d.conn.Where(query, args...).First(i).Error
}

// Find finds the models
func (d *Database) Find(i interface{}) error {
	return d.conn.Find(i).Error
}

// FindCond finds the models by custom condition from the same structure
func (d *Database) FindCond(i, c interface{}) error {
	return d.conn.Find(i, c).Error
}

// FindByQuery finds the model by query
func (d *Database) FindByQuery(i interface{}, query interface{}, args ...interface{}) error {
	return d.conn.Where(query, args...).Find(i).Error
}

// DeleteByQuery deletes the model by query
func (d *Database) DeleteByQuery(i interface{}, query interface{}, args ...interface{}) error {
	return d.conn.Where(query, args...).Delete(i).Error
}
