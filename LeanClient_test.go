package lean

import "testing"

func TestDetectDedicatedApiServer(t *testing.T) {
	appId := "TkfTW7NsI3fl8G5BJTNelEtn-gzGzoHsz"
	apiServer, err := DetectDedicatedApiServer(appId)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("api server: %s", apiServer)
}
