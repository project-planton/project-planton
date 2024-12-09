package module

var vars = struct {
	TeamId          string
	PostgresVersion string
}{
	// TeamId required by zalando operator
	// a random word has been chosen here. This could be any valid string.
	TeamId:          "db",
	PostgresVersion: "14",
}
