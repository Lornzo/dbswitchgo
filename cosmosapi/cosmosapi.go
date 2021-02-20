package cosmosapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"dbswitchgo"
	"encoding/base64"
	"fmt"
	"strings"
)

// CosmosAPIResponse ...
type CosmosAPIResponse struct {
	ResponseStatus int
	ResponseID     string        `json:"_rid"`
	Datas          []interface{} `json:"Documents"`
	Counter        int64         `json:"_count"`
}

// CosmosAPI ...
type CosmosAPI struct {
	// cosmos 的 key
	authorizationKey string
	// container的名稱(azure的資源名稱)
	containerName string
	// cosmos rest api 版本
	apiVersion string
	// 繼承自DbSwitch
	dbswitchgo.DbSwitch
}

// NewCosmosAPI ...
//  Cosmos Rest Api 物件
//  透過Cosmos Rest Api來連接Azure Cosmos DB
//  @return {struct} *CosmosApi
func NewCosmosAPI() *CosmosAPI {
	return new(CosmosAPI).SetAPIVersion("2018-12-31")
}

// SetContainer ...
//  設定想要操作的container
//  @params {String} container 想要操作的container
//  @return {*CosmosApi} this
func (thisObj *CosmosAPI) SetContainer(container string) *CosmosAPI {
	thisObj.containerName = container
	return thisObj
}

// SetAuthorizationKey ...
//  設定Cosmos Rest Api需要用到的AuthorizationKey
//  @params {String} key
//  @return {*CosmosApi} this
func (thisObj *CosmosAPI) SetAuthorizationKey(key string) *CosmosAPI {
	thisObj.authorizationKey = key
	return thisObj
}

// SetAPIVersion ...
//  設用Cosmos Rest Api 的版本
//  @params {String} v 版號，一般來說是日期
//  @return {*CosmosApi} this
func (thisObj *CosmosAPI) SetAPIVersion(v string) *CosmosAPI {
	thisObj.apiVersion = v
	return thisObj
}

// SetDatabase ...
//  設定Cosmos 要操作的Database
//  @params {String} database 資料庫名稱
//  @return {*CosmosAPI} this
func (thisObj *CosmosAPI) SetDatabase(database string) *CosmosAPI {
	thisObj.DbName = database
	return thisObj
}

// SetTable ... {override}
//  設定要取用的 Collection
//  @params {String} table 選定要操作的Table/Collection
//  @return this
func (thisObj *CosmosAPI) SetTable(table string) *CosmosAPI {
	thisObj.TableName = table
	return thisObj
}

// SetSelect ... {override}
//  設定SQL語法的選定
//  @params {String} selectString => COUNT(aa) AS BB
//  @return {*CosmosAPI} this
func (thisObj *CosmosAPI) SetSelect(selectString string) *CosmosAPI {
	thisObj.SelectString = selectString
	return thisObj
}

// SetOrderByString ... {orverride}
//  設定查詢的order by 尾巴
//  * 目前不可用
//  @params {String} orderBy 查詢尾巴字串
//  @return {*CosmosApi} this
func (thisObj *CosmosAPI) SetOrderByString(orderBy string) *CosmosAPI {
	thisObj.OrderByString = orderBy
	return thisObj
}

// getAPIToken ...
//  取得Cosmos Api用的token
//  @params {string} verb
func (thisObj *CosmosAPI) getAPIToken(verb string, resouceType string, resourceLink string, key string, dateString string) string {
	result := ""
	secret, err := base64.StdEncoding.DecodeString(key)
	if err == nil {
		tokenString := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", strings.ToLower(verb), strings.ToLower(resouceType), resourceLink, strings.ToLower(dateString), "")
		h := hmac.New(sha256.New, secret)
		h.Write([]byte(tokenString))
		result = base64.StdEncoding.EncodeToString(h.Sum(nil))
	}
	return result
}
