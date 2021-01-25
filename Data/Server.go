package Data

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/login/Control"
	"github.com/wuxia-server/login/DataTable"
)

func GetServerList() (serverList []*DataTable.Server) {
	serverList = make([]*DataTable.Server, 0)
	sqlstr := dal.MarshalGetSql(DataTable.NewServer())
	rows, err := Control.GateDB.Query(sqlstr)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			server := DataTable.NewServer()
			rows.Scan(
				&server.Id,
				&server.Title,
				&server.Status,
				&server.Host,
				&server.Port,
				&server.ServiceTime,
			)
			serverList = append(serverList, server)
		}
	}
	return
}

func GetServerById(serverId int64) (server *DataTable.Server) {
	server = DataTable.NewServer()
	sqlstr := dal.MarshalGetSql(server, "id")
	row := Control.GateDB.QueryRow(sqlstr, serverId)
	if row.Scan(
		&server.Id,
		&server.Title,
		&server.Status,
		&server.Host,
		&server.Port,
		&server.ServiceTime,
	) != nil {
		server = nil
	}
	return
}
