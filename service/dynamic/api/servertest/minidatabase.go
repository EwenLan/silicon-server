package servertest

import (
	"database/sql"

	"github.com/EwenLan/silicon-server/globaldefine"
	"github.com/EwenLan/silicon-server/service/dynamic/handler"
	"github.com/EwenLan/silicon-server/slog"

	_ "github.com/mattn/go-sqlite3"
)

type MiniDatabase struct {
	db *sql.DB
}

func (i *MiniDatabase) Init() {
	if i.db != nil {
		slog.Debugf("database has been loaded, not need to load again")
		return
	}
	db, err := sql.Open("sqlite3", "minidatabase.db")
	if (db == nil) || (err != nil) {
		slog.Errorf("cannot load sqlite3 database, err: %s", err)
		return
	}
	slog.Debugf("database has been loaded or created")
	i.db = db
	createTable := `
	create table if not exists minidatabase (
		id integer primary key autoincrement,
		name text not null,
		value text
	);
	`
	_, err = i.db.Exec(createTable)
	if err != nil {
		slog.Errorf("fail to create table, err: %s", err)
		return
	}
	slog.Debugf("create table successful")
}

var MiniDatabasePutImp = handler.JsonHandler{
	ServiceHandler: &MiniDatabasePut{},
}

type MiniDatabasePut struct {
	MiniDatabase
	request  globaldefine.MiniDatabaseRequest
	response globaldefine.MiniDatabaseResponse
}

func (i *MiniDatabasePut) HandleRequest() error {
	res, err := i.db.Exec("insert into minidatabase(name, value) values (?, ?)", i.request.Name, i.request.Value)
	if err != nil {
		slog.Errorf("fail to insert name: %s, value: %s, err: %s", i.request.Name, i.request.Value, err)
		return err
	}
	i.response.MiniDatabaseRequest = i.request
	i.response.ID, err = res.LastInsertId()
	if err != nil {
		slog.Errorf("fail to get last inserted id, err: %s", err)
		return err
	}
	return nil
}

func (i *MiniDatabasePut) Init() {
	i.request = globaldefine.MiniDatabaseRequest{}
	i.response = globaldefine.MiniDatabaseResponse{}
	i.MiniDatabase.Init()
}

func (i *MiniDatabasePut) GetRequestStruct() interface{} {
	return &i.request
}

func (i *MiniDatabasePut) GetResponseStruct() interface{} {
	return &i.response
}

type MiniDatabaseGet struct {
	MiniDatabase
	request  globaldefine.MiniDatabaseRequest
	response globaldefine.MiniDatabaseGetResponse
}

func (i *MiniDatabaseGet) HandleRequest() error {
	rows, err := i.db.Query("select * from minidatabase where name = ?", i.request.Name)
	if (err != nil) || (rows == nil) {
		slog.Errorf("fail to query name: %s, err: %s", i.request.Name, err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		newRow := globaldefine.MiniDatabaseResponse{}
		err = rows.Scan(&newRow.ID, &newRow.Name, &newRow.Value)
		if err != nil {
			slog.Errorf("fail to scan data, name: %s, err: %s", i.request.Name, err)
			return err
		}
		i.response.Ans = append(i.response.Ans, newRow)
	}
	return nil
}

func (i *MiniDatabaseGet) Init() {
	i.request = globaldefine.MiniDatabaseRequest{}
	i.response = globaldefine.MiniDatabaseGetResponse{}
	i.MiniDatabase.Init()
}

func (i *MiniDatabaseGet) GetRequestStruct() interface{} {
	return &i.request
}

func (i *MiniDatabaseGet) GetResponseStruct() interface{} {
	return &i.response
}

var MiniDatabaseGetImp = handler.JsonHandler{
	ServiceHandler: &MiniDatabaseGet{},
}
