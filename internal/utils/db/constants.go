package dbutils

const (
	// TODO: Fall back to these values if the environment variables are not set
	DB_DEFAULT_TIMEZONE     = "Asia/Singapore"
	DB_DEFAULT_HOSTNAME     = "localhost"
	DB_DEFAULT_USER         = "postgres"
	DB_DEFAULT_PASSWORD     = ""
	DB_DEFAULT_PORT         = 5432
	DB_DEFAULT_NAME         = "sa_stories"
	DB_DEFAULT_NAME_TESTING = "sa_stories_testing"
)

// Not used for now
// TODO: Remove if not used
func GetOrDefault[T comparable](val T, deflt T) T {
	var empty T
	if val == empty {
		return deflt
	}
	return val
}
