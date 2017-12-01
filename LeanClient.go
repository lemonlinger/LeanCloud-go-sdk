package lean

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	ApiVersion = "1.1"
)

var (
	apiServer = "api.leancloud.cn"
)

type LeanClient struct {
	appId, appKey, masterKey string
	useSign                  bool
	Installation, User, Role *Collection
	File                     *Collection
}

//create a new lean client.
//You only need one client in a go application
//application will not check your keys if you don't call API that needs it
func NewClient(appId, appKey, masterKey string) *LeanClient {

	ret := &LeanClient{
		appId:     appId,
		appKey:    appKey,
		masterKey: masterKey,
	}

	installation := ret.Collection("_Installation")
	ret.Installation = &installation
	ret.Installation.classSubfix = "/installaions"

	user := ret.Collection("_User")
	ret.User = &(user)
	ret.User.classSubfix = "/users"

	role := ret.Collection("_Role")
	ret.Role = &(role)
	ret.Role.classSubfix = "/roles"

	file := ret.Collection("_File")
	ret.File = &(file)
	ret.File.classSubfix = "/files"

	if s, err := DetectDedicatedApiServer(appId); err != nil {
		apiServer = s
	}
	return ret
}

func DetectDedicatedApiServer(appId string) (string, error) {
	routerUrl := "https://app-router.leancloud.cn/2/route?appId=" + appId
	resp, err := http.Get(routerUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type routerResult struct {
		ApiServer       string `json:"api_server"`
		EngineServer    string `json:"engine_server"`
		PushServer      string `json:"push_server"`
		RTMRouterServer string `json:"rtm_router_server"`
		StatsServer     string `json:"stats_server"`
		TTL             int    `json:"ttl"`
	}
	var result routerResult
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	return result.ApiServer, nil
}

func GetUrlBase() string {
	return "https://" + apiServer + "/" + ApiVersion
}
