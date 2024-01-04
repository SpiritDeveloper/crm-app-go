package db

import (
	"crm-app-go/config"
	. "crm-app-go/model"
	"fmt"
	"log"
	"net/url"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type gormDatabase struct {
	client *gorm.DB
	once   sync.Once
}

func NewGormDatabase() IDatabaseEngine {
	return &gormDatabase{}
}

func InitDatabase(g *gormDatabase, config *config.Database) {
	dsn := url.URL{
		User:     url.UserPassword(config.User, config.Password),
		Scheme:   config.Name,
		Host:     fmt.Sprintf("%s:%s", string(config.Server), string(config.Port)),
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	log.Println("Database", dsn)
	db, err := gorm.Open(config.Engine, dsn.String())
	if err != nil {
		log.Println("Database connection failed : ", err)
	} else {
		log.Println("Database connection established!")
	}
	log.Println("POSTGRES connection running on port 5432")
	g.client = db
}

// Making sure gormClient only initialise once as singleton
func (g *gormDatabase) GetDatabase(config config.Database) *gorm.DB {
	if g.client == nil {
		g.once.Do(func() {
			InitDatabase(g, &config)
		})
	}
	return g.client
}

func (g *gormDatabase) RunMigration() {
	if g.client == nil {
		panic("Initialise gorm db before running migrations")
	}
	// g.client.AutoMigrate(&User{})

	g.client.AutoMigrate(&Lead{})
	g.client.AutoMigrate(&Token{})
	g.client.AutoMigrate(&Log{})
	g.client.AutoMigrate(&Configuration{})
	//We need to add foreign keys manually.
}
