package repositories

import (
	"errors"
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
func (rs *Repositories) Update(dbObjectToUpdate interface{}, InObject interface{}) error {
	destValue := reflect.ValueOf(dbObjectToUpdate)
	srcValue := reflect.ValueOf(InObject)

	if destValue.Kind() != reflect.Ptr || destValue.IsNil() {
		return errors.New("dbObjectToUpdate must be a non-nil pointer")
	}
	destValue = destValue.Elem()

	if srcValue.Kind() != reflect.Struct {
		return errors.New("InObject must be a struct")
	}

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		srcType := srcValue.Type().Field(i)

		if !srcField.CanInterface() {
			continue
		}

		destField := destValue.FieldByName(srcType.Name)

		if destField.IsValid() && destField.CanSet() {
			if !reflect.DeepEqual(srcField.Interface(), reflect.Zero(srcField.Type()).Interface()) {
				destField.Set(srcField)
			}
		}
	}
	return rs.Save(dbObjectToUpdate) // this saves/updates the object to db
}
