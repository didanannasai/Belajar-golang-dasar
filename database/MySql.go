package database

var connection string

// gunakan nama function "init" jika ingin langsung mengeksekusi fungsi tersebut saat packagenya dipanggil
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}