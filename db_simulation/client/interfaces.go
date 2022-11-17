package client_simulation

type ClientSimulation interface {
	// Hashing pws with encryption
	registerCredentials(string, string) bool
	// Where Digital Signature on hashed provided pw and decrypted hashed pw from database.
	logInCredentials(string, string) bool
	// Delete Credentials ( similar Digital Signature to the log/in )
	deleteCredentials(string, string) bool
}
