package mongo

// InitDB initializes the MongoStore
import (
	"strings"
	"time"

	"github.com/globalsign/mgo"
)

// New returns a new MongoStore pointer
func New(c Config) (*MongoStore, error) {
	// connStr := m.String()
	// session, err := mgo.DialWithTimeout(connStr, m.Timeout)
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
	return &MongoStore{DBName: c.Name, session: session}, nil
}

// Init initializes the DB handler
func Init(c Config) error {
	handler, err := New(c)
	if err != nil {
		return err
	}

	Mongo = *handler
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

// Validate MongoDB settings
func (c *Config) Validate() bool {
	if len(c.Name) < 1 {
		return false
	}

	if len(c.Port) < 1 {
		c.Port = "27017"
	}

	c.Timeout = time.Second * c.Timeout

	return true
}

// String returns the MongoDB Dial string
func (c *Config) String() string {
	dialString := "mongodb://"

	if len(c.ConnectionString) > 0 {
		dialString = c.ConnectionString
	} else {
		if len(c.Username) > 0 && len(c.Password) > 0 {
			dialString += (c.Username + ":" + c.Password + "@")
		}

		if len(c.Host) > 0 {
			// dialString += (strings.Join(m.Host, ","))
			for idx, h := range c.Host {
				if len(c.Port) > 0 {
					// Add port if port number not already mentioned in host
					if strings.Index(h, ":") < 0 {
						h += ":" + c.Port
					}
				}
				dialString += h
				if idx < (len(c.Host) - 1) {
					dialString += ","
				}
			}
		}

		if len(c.Name) > 0 {
			dialString += ("/" + c.Name)
		}

		if len(c.AuthSource) > 0 {
			dialString += "?authSource=" + c.AuthSource
		}
	}

	return dialString
}
