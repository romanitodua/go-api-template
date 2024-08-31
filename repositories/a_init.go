package repositories

import (
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"reflect"
)

type Repositories struct {
	db *gorm.DB
	// domain specific repos
}

func InitRepositories() *Repositories {
	cnf := postgres.Config{DriverName: "cloudsqlpostgres", DSN: os.Getenv("DB_CONNECTION")}
	db, err := gorm.Open(postgres.New(cnf))
	if err != nil {
		panic(err)
	}
	return &Repositories{
		db: db,
	}
}
func (rs *Repositories) MigrateDB() error {
	err := rs.db.AutoMigrate()
	return err
}

func (rs *Repositories) Save(object interface{}) error {
	rt := reflect.TypeOf(object)
	switch rt.Kind() {
	case reflect.Slice:
		vo := reflect.ValueOf(object)
		if vo.Len() == 0 {
			return nil
		}
	}

	return rs.db.Save(object).Error
}
func (rs *Repositories) Delete(object interface{}) error {
	rt := reflect.TypeOf(object)
	switch rt.Kind() {
	case reflect.Slice:
		vo := reflect.ValueOf(object)
		if vo.Len() == 0 {
			return nil
		}
	}
	return rs.db.Delete(object).Error
}
func ignoreNotFound[T any](object *T, err error) (*T, error) {
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return object, err
}
