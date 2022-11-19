package client_simulation

type ClientSimulation interface {
	// Hashing pws with encryption
	RegisterCredentials(string, string) string
	// Where Digital Signature on hashed provided Token is happening
	LogInCredentials(string, string, string) string
}

func RegisterInterface(ci ClientSimulation, usr string, pw string) string {
	return ci.RegisterCredentials(usr, pw)
}
func LogInInterface(ci ClientSimulation, usr string, pw string, tk string) string {
	return ci.LogInCredentials(usr, pw, tk)
}
