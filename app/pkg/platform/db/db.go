package db

// InitDB initializes the MongoStore
import (
	"time"

	"github.com/aadi756/cardflow/app/pkg/platform/db/mongo"
	"github.com/globalsign/mgo"
)

// New returns a new MongoStore pointer
func New(c Config) (*mongo.MongoStore, error) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Database:       c.Name,
		ReplicaSetName: c.ReplicaSet,
		Addrs:          c.Host,
		Username:       c.Username,
		Password:       c.Password,
		Timeout:        c.Timeout,
		WriteTimeout:   time.Second * 10,
		ReadTimeout:    time.Second * 5,
		Source:         c.AuthSource,
		Direct:         c.Direct,
		AppName:        c.AppName,
		ReadPreference: &c.ReadPreference,
	})

	if err != nil {
		return nil, err
	}

	session.SetSafe(&mgo.Safe{WMode: "majority"})
	err = session.Ping()
	if err != nil {
		return nil, err
	}
	return &mongo.MongoStore{DBName: c.Name, Session:session}, nil
}

// Init initializes the DB handler
func Init(c Config) error {
	handler, err := New(c)
	if err != nil {
		return err
	}

	mongo.Mongo = *handler
	return nil
}

// Config holds the config required for MongoDB
type Config struct {
	AppName          string   `json:"appName,omitempty"`
	Name             string   `json:"name,omitempty"`
	Host             []string `json:"host,omitempty"`
	Port             string   `json:"port,omitempty"`
	ConnectionString string   `json:"connectionString,omitempty"`
	AuthSource       string   `json:"authSource,omitempty"`
	ReplicaSet       string   `json:"replicaSet,omitempty"`

	Username       string             `json:"username,omitempty"`
	Password       string             `json:"password,omitempty"`
	Direct         bool               `json:"direct,omitempty"`
	Timeout        time.Duration      `json:"timeout,omitempty"`
	ReadPreference mgo.ReadPreference `json:"readPreference,omitempty"`
}

