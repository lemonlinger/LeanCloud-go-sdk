package lean

import (
	"github.com/parnurzeal/gorequest"
)

//be attention that EmailVerified and MobilePhoneVerified can be nil
type User struct {
	LeanClassesBase
	Salt                string                 `json:"salt,omitempty"`
	Email               string                 `json:"email, omitempty"`
	SessionToken        string                 `json:"sessionToken, omitempty"`
	Passowrd            string                 `json:"password,omitempty"`
	Username            string                 `json:"username,omitempty"`
	EmailVerified       *bool                  `json:"emailVerified, omitempty"`
	MobilePhoneNumber   string                 `json:"mobilePhoneNumber, omitempty"`
	AuthData            map[string]interface{} `json:"authData, omitempty"`
	MobilePhoneVerified *bool                  `json:"mobilePhoneVerified, omitempty"`
}

//will return nil if there are any error
func (c *LeanClient) Login(userName, pwd string) (*User, error) {
	requestBody := map[string]string{
		"username": userName,
		"password": pwd,
	}
	url := UrlBase + "/login"
	request := gorequest.New()
	superAgent := request.Post(url).
		Send(requestBody)
	agent := &Agent{
		superAgent: superAgent,
		client:     c,
	}
	if err := agent.Do(); nil != err {
		return nil, err
	}
	ret := &User{}
	if err := agent.ScanResponse(ret); nil != err {
		return nil, err
	}
	return ret, nil
}

//will return nil if there are any error
func (c *LeanClient) UserMe(token string) (*User, error) {
	url := UrlBase + "/users/me"
	request := gorequest.New()
	superAgent := request.Get(url)
	agent := &Agent{
		superAgent: superAgent,
		client:     c,
	}
	agent.UseSessionToken(token)

	if err := agent.Do(); nil != err {
		return nil, err
	}
	ret := &User{}
	if err := agent.ScanResponse(ret); nil != err {
		return nil, err
	}
	return ret, nil
}

func (c *LeanClient) UsersByMobilePhone(mobilePhone, smsCode string) (*User, error) {
	url := UrlBase + "/usersByMobilePhone"
	requestBody := map[string]string{
		"mobilePhoneNumber": mobilePhone,
		"smsCode":           smsCode,
	}
	request := gorequest.New()
	superAgent := request.Post(url).Send(requestBody)
	agent := &Agent{
		superAgent: superAgent,
		client:     c,
	}
	if err := agent.Do(); err != nil {
		return nil, err
	}
	user := &User{}
	if err := agent.ScanResponse(user); nil != err {
		return nil, err
	}
	return user, nil
}

func (c *LeanClient) UpdateUser(u User) error {
	url := UrlBase + "/users/" + u.ObjectId
	requestBody := make(map[string]string)
	if u.Passowrd != "" {
		requestBody["password"] = u.Passowrd
	}
	if u.MobilePhoneNumber != "" {
		requestBody["phone"] = u.MobilePhoneNumber
	}
	if u.Username != "" {
		requestBody["username"] = u.Username
	}
	if u.Email != "" {
		requestBody["email"] = u.Email
	}
	if len(requestBody) == 0 {
		return nil
	}
	request := gorequest.New()
	superAgent := request.Put(url).Send(requestBody)
	agent := &Agent{
		superAgent: superAgent,
		client:     c,
	}
	agent.UseSessionToken(u.SessionToken)
	return agent.Do()
}
