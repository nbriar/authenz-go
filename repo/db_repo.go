package repo

import (
	"github.com/jinzhu/gorm"
	c "github.com/nbriar/authenz-go/configuration"

	// Adapter for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBRepo is the primary datastore
var DBRepo *gorm.DB

func init() {
	var err error
	DBRepo, err = gorm.Open(
		c.DB.Adapter,
		"host="+c.DB.Host+
			" user="+c.DB.User+
			" dbname="+c.DB.DBName+
			" password="+c.DB.DBPassword)

	if err != nil {
		panic("failed to connect database")
	}
}
