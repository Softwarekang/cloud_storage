package extension

import (
	"user-service/app/com/wpbs/SQL"
)

var (
	sqlClients = make(map[string]func() SQL.SQLClient)
)

func SetSQlClient(name string, fnc func() SQL.SQLClient) {
	sqlClients[name] = fnc
}

func GetSQLClient(name string) SQL.SQLClient {
	if name == ""{
		name = "xorm"
	}
	if sqlClients[name] == nil {
		panic("sqlclient for" + name + "not existing, please import the pkg")
	}

	return sqlClients[name]()
}
