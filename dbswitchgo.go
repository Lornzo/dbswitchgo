package dbswitchgo

type DbSwitch struct {
	// 資料庫的位置
	DbHost string
	// 資料庫的Port
	DbPort string
	// 選用資料庫的名稱
	DbName string
	// 登入的帳號
	DbUser string
	// 登入的密碼
	DbPass string
	// 選用的table
	TableName string
	//
	SelectString string
	//
	Conditions []string

	OrderByString string
}

/** 設定連線資訊
 * @params {String} host 主機位置，可以是ip，也可以是domain
 * @params {String} port 目標DB所使用的port
 * @params {String} name 所想要使用的 DB 名稱
 * @params {String} user 登入DB所要用的帳號
 * @params {String} psw 登入DB所需要用的密碼
 */
func (d *DbSwitch) SetConnectionInfo(host, port, name, user, psw string) {
	d.DbHost = host
	d.DbPort = port
	d.DbUser = user
	d.DbPass = psw
}

func (d *DbSwitch) SetDatabase(database string) {
	d.DbName = database
}

/** 設定要取用的Table/Collection
 * @params {String} table 選定要操作的Table/Collection
 */
func (d *DbSwitch) SetTable(table string) {
	d.TableName = table
}

func (d *DbSwitch) SetCondition() {}
