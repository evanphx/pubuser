package pubuser

import (
  "encoding/json"
  "net/url"
)

type User interface {
  UserName() string
  FullName() string
  GravatarID() string
  SSHKeys() []string
  Groups() []string
}

func Find(identifier string) (User, error) {
  uri, err := url.Parse(identifier)
  if err != nil {
    return nil, err
  }

  if uri.Scheme == "github" {
    return GetGithubUser(uri.Opaque)
  }

  if uri.Scheme == "http" || uri.Scheme == "https" {
    return GetHttpJson(identifier)
  }

  return nil, eNoSuchUser
}

func ToJson(user User) ([]byte, error) {
  d := &HttpUserData {
    UserName: user.UserName(),
    Name: user.FullName(),
    GravatarID: user.GravatarID(),
    SSHKeys: user.SSHKeys(),
    Groups: user.Groups(),
  }

  return json.Marshal(d)
}
