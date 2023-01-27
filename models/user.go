package models

// create a User Struct containing the data we want to store
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// declare variables
var (
	users  []*User
	nextID = 1
)

// create a function to get users
func GetUsers() []*User {
	return users
}

// create a function to add a user
func AddUser(u User) (User, error) {
	// take the user ID and assign it the nextID
	u.ID = nextID
	// increment nextID by one.
	nextID++
	users = append(users, &u)
	// return the manipulated user object (u), also return nil because there is no error
	return u, nil
}