package database

import (

	"fmt"
	"log"
	"net/http"
//	"context"
//	"encoding/json"
	"time"
	"os"
	config "brypt-server/config"

//	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/mongo"
	// "github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/ftdc/bsonx"
	"github.com/mongodb/mongo-go-driver/x/mongo/driver/uuid"
)

var configuration = config.Configuration{}

type key string

/* **************************************************************************
** Managers
** *************************************************************************/
type Manager struct {
	ID                uuid.UUID					`bson:"_id,omitempty" json:"_id,omitempty"`
	Manager_name      string            `bson:"manager_name" json:"manager_name"`
}

/* **************************************************************************
** Clusters
** *************************************************************************/
type Cluster struct {
	ID                uuid.UUID 		`bson:"_id,omitempty" json:"_id,omitempty"`
	Connection_token  string            `bson:"connection_token" json:"connection_token"`
	Coord_ip          string            `bson:"coord_ip" json:"coord_ip"`
	Coord_port        string            `bson:"coord_port" json:"coord_port"`
	Comm_tech         string            `bson:"comm_tech" json:"comm_tech"`
}

/* **************************************************************************
** Networks
** *************************************************************************/
type Network struct {
	ID                uuid.UUID 		`bson:"_id,omitempty" json:"_id,omitempty"`
	Network_name      string            `bson:"network_name" json:"network_name"`
	Owner_name        string            `bson:"owner_name" json:"owner_name"`
	Managers          []Manager         `bson:"managers" json:"managers"`
	Direct_peers      int               `bson:"direct_peers" json:"direct_peers"`
	Total_peers       int               `bson:"total_peers" json:"total_peers"`
	Ip_address        string            `bson:"ip_address" json:"ip_address"`
	Port              int               `bson:"port" json:"port"`
	Connection_token  string            `bson:"connection_token" json:"connection_token"`
	Clusters          []Cluster         `bson:"clusters" json:"clusters"`
	Created_on        time.Time         `bson:"created_on" json:"created_on"`
	Last_accessed     time.Time         `bson:"last_accessed" json:"last_accessed"`
}

/* **************************************************************************
** Users
** *************************************************************************/
type User struct {
	ID                uuid.UUID 		`bson:"_id,omitempty" json:"_id,omitempty"`
	Username          string            `bson:"username" json:"username"`
	First_name        string            `bson:"first_name" json:"first_name"`
	Last_name         string            `bson:"last_name" json:"last_name"`
	Email             string            `bson:"email" json:"email"`
	Organization      string            `bson:"organization" json:"organization"`
	Networks          []Network         `bson:"networks" json:"networks"`
	Age               time.Time         `bson:"age" json:"age"`
	Join_date         time.Time         `bson:"join_date" json:"join_date"`
	Last_login        time.Time         `bson:"last_login" json:"last_login"`
	Login_attempts    int               `bson:"login_attempts" json:"login_attempts"`
	Login_token       string            `bson:"login_token" json:"login_token"`
	Region            string            `bson:"region" json:"region"`
}

/* **************************************************************************
** Nodes
** *************************************************************************/
type Node struct {
	ID                uuid.UUID			`bson:"_id,omitempty" json:"_id,omitempty"`
	Serial_number     string            `bson:"serial_number" json:"serial_number"`
	Type              string            `bson:"type" json:"type"`
	Created_on        time.Time         `bson:"created_on" json:"created_on"`
	Registered_on     time.Time         `bson:"registered_on" json:"registered_on"`
	Registered_to     string            `bson:"registered_to" json:"registered_to"`
	Connected_network string            `bson:"connected_network" json:"connected_network"`
}


const (
	hostKey = "123"
)

var (
	Client *mongo.Client
)

/* **************************************************************************
** Function: Setup
** URI:
** Description:
** *************************************************************************/
func Setup() {

	var err error

	configuration = config.GetConfig()

	connectionURL := configuration.Database.MongoURI
	if connectionURL == "" {
		panic( "Connection variable is not set!" )
	}

	Client, err = mongo.NewClient( connectionURL )
	if err != nil {
		panic( err )
	}

	fmt.Print( Client )

}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	
	print("In users handler!\n")
	WriteManager(w)
	/*switch r.Method {
	case "GET":
		users_collection := Client.Database("brypt_server").Collection("brypt_users")

		var sort *options.FindOptions
		var err error

		//sort, err := users_collection.find({}, {"username":1, _id:0}).sort({"username":1})
		// TODO: implement SetSort
		// sort, err := mongo.Opt.Sort(bsonx.NewDocument(bsonx.EC.Int32("username", 1)))   // Sort by username?

		if err != nil { // Error handler for sort
			log.Fatal("Error in usersHandler() sorting: ", err)
		}

		cursor, err := users_collection.Find(nil, nil, sort)  // Get a cursor to the start of the collection?
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal("Internal server error")
			return
		}

		defer cursor.Close(context.Background())

		var users []User  // Create an array to store data from users table

		for cursor.Next(nil) {  // Iterate through users_collection
			user := User{}
			err := cursor.Decode(&user) // Catch any errors while decoding the user object
			if err != nil { // Log the error caught
				log.Fatal("Users collection decode error: ", err)
			}
			users = append(users, user) // Append the stored object to our users array
		}

		if err := cursor.Err(); err != nil {  // Check for a cursor error
			log.Fatal("Cursor error: ", err)
		}

		jsonstr, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal("Internal server error")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonstr)  // Returns the entire json users collection
		return

	case "PUT":
		r.ParseForm()
		//	users_collection := Client.Database("brypt_server").Collection("brypt_users")

		// TODO: Perform error checking
		newUser := bsonx.NewDocument(bsonx.EC.String("username", r.Form.Get("username")),
		bsonx.EC.String("first_name", r.Form.Get("first_name")),
		bsonx.EC.String("last_name", r.Form.Get("last_name")))
		WriteUsers(newUser, w)
		return
	}*/
	return
}


func WriteUsers(newUser *bsonx.Document, w http.ResponseWriter){
	users_collection := Client.Database("brypt_server").Collection("brypt_users")

	_, err := users_collection.InsertOne(nil, newUser)
	if err != nil {
		log.Println("Error inserting new user: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	return
}

func WriteNetworks(n Network) {
	// TODO
}

func WriteNodes(n Node) {
	// TODO
}

func WriteManager(w http.ResponseWriter){
	newManager := bsonx.NewDocument(bsonx.EC.String("Manager_name", "testname"))

	m_collection := Client.Database("brypt_server").Collection("brypt_managers")

	_, err := m_collection.InsertOne(nil, newManager)
	if err != nil {
		log.Println("Error inserting new manager: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	return
}

func CreateClient() {

	var err error
	print("\n\nSetting up client...\n")

	configuration = config.GetConfig()

	connectionURL := configuration.Database.MongoURI
	if connectionURL == "" {
		panic( "Connection variable is not set!" )
	}

	cert_path, cert_exists := os.LookupEnv("MONGODB_CERT_PATH")

	if cert_exists {  // If user has certification, create a new client with cert info
		print(cert_path)
		Client, err = mongo.NewClient(connectionURL)
		//Client, err = mongo.NewClientWithOptions(connectionURL, mongo.ClientOpt.SSLCaFile(cert_path))
	} else {  // Else create a new client without cert info
		Client, err = mongo.NewClient(connectionURL)
	}

	if err != nil {
		panic( err )
	}

	print("\nFinished setting up client...\n\n")
	fmt.Print( Client )
	return
}

// /* **************************************************************************
// ** Function: CreateClient
// ** URI:
// ** Description: Creates a database client
// ** *************************************************************************/
//
// func CreateClient() {
//
// 				var err error
//
// 				connection_url, url_exists := os.LookupEnv("COMPOSE_MONGODB_URL")
// 				if !url_exists) {
// 								log.Fatal("COMPOSE_MONGODB_URL environmental variable is not set. This needs to be set to ...")
// 				}
//
// 				cert_path, cert_exists := os.LookupEnv("MONGODB_CERT_PATH")
//
// 				if cert_exists {  // If user has certification, create a new client with cert info
// 								Client, err = mongo.NewClientWithOptions(connection_url, mongo.ClientOpt.SSLCaFile(cert_path))
// 				} else {  // Else create a new client without cert info
// 								Client, err = mongo.NewClient(connection_url)
// 				}
//
// 				if err != nil {
// 								log.Fatal(err)  // Log any errors which come up during client connection
// 				}
//
// }

/* **************************************************************************
** Function: Connect
** URI:
** Description: Creates client connection
** *************************************************************************/
func Connect() {
	var err error

	err = Client.Connect(nil)

	if err != nil {
		log.Fatal(err)  // Log any errors thrown during connection
	}

 }


// /* **************************************************************************
// ** Function: Disconnect
// ** URI:
// ** Description: Disconnects client
// ** *************************************************************************/
// func Disconnect() {
// 	defer Client.Disconnect(nil)	// Disconnection client
// }
