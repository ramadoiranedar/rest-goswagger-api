package rest_goswagger_api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewRuntime(conf *viper.Viper, log logrus.FieldLogger) (*Runtime, error) {
	// Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GetString("db.user"),
		conf.GetString("db.pass"),
		conf.GetString("db.host"),
		conf.GetString("db.port"),
		conf.GetString("db.name"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.WithFields(logrus.Fields{
			"module": "runtime",
		}).Errorf("DB connection : %v", err)
		os.Exit(1)
	}

	rt := &Runtime{
		Db:   db,
		Conf: conf,
		Log:  log,
	}

	// run auto migrate the database
	rt.RunMigration()

	return rt, nil
}

type Runtime struct {
	Db   *gorm.DB
	Conf *viper.Viper
	Log  logrus.FieldLogger
}

func (r *Runtime) DB() *gorm.DB {
	return r.Db
}

func (r *Runtime) Logger() logrus.FieldLogger {
	return r.Log
}

func (r *Runtime) Config() *viper.Viper {
	return r.Conf
}

// TODO: close everything needed
func (r *Runtime) Close() {
	return
}

func (r *Runtime) SetError(code int, msg string, args ...interface{}) error {
	return errors.New(int32(code), msg)
}

func (r *Runtime) GetError(err error) errors.Error {
	if v, ok := err.(errors.Error); ok {
		return v
	}

	return errors.New(http.StatusInternalServerError, err.Error())
}

func (r *Runtime) RunMigration() {
	r.Log.Infof("Migrating Database Tables")
	r.Db.AutoMigrate()
}
