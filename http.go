package pubuser

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

type HttpUserData struct {
  UserName string `json:"username"`
  Name string `json:"name"`
  GravatarID string `json:"gravatar_id"`
  SSHKeys []string `json:"ssh_keys"`
  Groups []string `json:"groups"`
}

type HttpUser struct {
  url string
  data *HttpUserData
}

func (h *HttpUser) UserName() string {
  return h.data.UserName
}

func (h *HttpUser) FullName() string {
  return h.data.Name
}

func (h *HttpUser) GravatarID() string {
  return h.data.GravatarID
}

func (h *HttpUser) Groups() []string {
  return h.data.Groups
}

func (h *HttpUser) SSHKeys() []string {
  return h.data.SSHKeys
}

var eBadJson = fmt.Errorf("Invalid JSON provided")

func GetHttpJson(url string) (*HttpUser, error) {
  res, err := http.Get(url)

  if res.StatusCode != 200 {
    return nil, eNoSuchUser
  }

  if res.StatusCode != 200 {
    return nil, eNoSuchUser
  }

  if err != nil {
    return nil, err
  }

  buf, err := ioutil.ReadAll(res.Body)

  if err != nil {
    return nil, err
  }

  user := &HttpUserData{}

  err = json.Unmarshal(buf, user)

  if err != nil {
    return nil, eBadJson
  }

  return &HttpUser{url, user}, nil
}
