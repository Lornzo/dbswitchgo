package cosmosapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SetCondition ...
//  設定讀取資料表的條件
//  @params {[]string} conditions => {"aa='bb',cc='dd'"}
func (thisObj *CosmosAPI) SetCondition(conditions []string) *CosmosAPI {
	thisObj.Conditions = conditions
	return thisObj
}

// FetchData ...
//  從資料表有條件的讀取資料
//  @return
func (thisObj *CosmosAPI) FetchData() CosmosAPIResponse {
	dateString := time.Now().UTC().Format(http.TimeFormat)
	signToken := thisObj.getAPIToken("POST", "docs", fmt.Sprintf("dbs/%s/colls/%s", thisObj.DbName, thisObj.TableName), thisObj.authorizationKey, dateString)
	params := fmt.Sprintf("type=%s&ver=%s&sig=%s", "master", "1.0", signToken)
	authorization := url.QueryEscape(params)
	client := http.Client{}
	url := fmt.Sprintf("https://%s.documents.azure.com/dbs/%s/colls/%s/docs", thisObj.containerName, thisObj.DbName, thisObj.TableName)

	queryString := "SELECT * FROM table"
	if thisObj.SelectString != "" {
		queryString = fmt.Sprintf("SELECT %s FROM table", thisObj.SelectString)
	}

	if len(thisObj.Conditions) > 0 {
		queryString = fmt.Sprintf("%s WHERE table.%s", queryString, strings.Join(thisObj.Conditions, " AND table."))
	}

	if thisObj.OrderByString != "" {
		queryString = fmt.Sprintf("%s %s", queryString, thisObj.OrderByString)
	}

	jsonStr := []byte(fmt.Sprintf(`{"query":"%s","paramates":[]}`, queryString))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Add("Authorization", authorization)
	req.Header.Add("x-ms-date", dateString)
	req.Header.Add("x-ms-version", "2018-12-31")
	req.Header.Add("x-ms-documentdb-isquery", "true")
	req.Header.Add("x-ms-documentdb-query-enablecrosspartition", "true")
	req.Header.Add("x-ms-max-item-count", "-1")
	req.Header.Add("Content-Type", "application/query+json")

	if err != nil {
		fmt.Println(err.Error())
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	result := CosmosAPIResponse{ResponseStatus: -1}
	jsonString, err := ioutil.ReadAll(response.Body)

	if err == nil {
		err = json.Unmarshal([]byte(jsonString), &result)
		if err == nil {
			result.ResponseStatus = 1
			// for key, val := range result.Datas {
			// 	rec, _ := val.(map[string]interface{})
			// 	fmt.Println(fmt.Sprintf("%d => %s", key, rec["_ts"]))
			// }
		}
	}
	return result
}
