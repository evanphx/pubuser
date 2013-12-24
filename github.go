package pubuser

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

type GithubUser struct {
  Login string `json:"login"`
  Name string `json:"name"`
  GHGravatarID string `json:"gravatar_id"`
  GHSSHKeys []string
  Organizations []string
}

func (g *GithubUser) UserName() string {
  return g.Login
}

func (g *GithubUser) FullName() string {
  return g.Name
}

func (g *GithubUser) GravatarID() string {
  return g.GHGravatarID
}

func (g *GithubUser) SSHKeys() []string {
  return g.GHSSHKeys
}

func (g *GithubUser) Groups() []string {
  return g.Organizations
}

type githubKey struct {
  Key string `json:"key"`
}

type githubOrg struct {
  Name string `json:"login"`
}

var eNoSuchUser = fmt.Errorf("No such user")

func GetGithubUser(name string) (*GithubUser, error) {
  res, err := http.Get("https://api.github.com/users/" + name)

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

  user := &GithubUser{}

  err = json.Unmarshal(buf, user)

  if err != nil {
    return nil, err
  }

  res, err = http.Get("https://api.github.com/users/" + name + "/keys")

  if err != nil {
    return nil, err
  }

  buf, err = ioutil.ReadAll(res.Body)

  var githubKeys []*githubKey

  err = json.Unmarshal(buf, &githubKeys)

  keys := make([]string, len(githubKeys))

  for idx, k := range githubKeys {
    keys[idx] = k.Key
  }

  user.GHSSHKeys = keys

  res, err = http.Get("https://api.github.com/users/" + name + "/orgs")

  if err != nil {
    return nil, err
  }

  buf, err = ioutil.ReadAll(res.Body)

  var githubOrgs []*githubOrg

  err = json.Unmarshal(buf, &githubOrgs)

  orgs := make([]string, len(githubOrgs))

  for idx, k := range githubOrgs {
    orgs[idx] = k.Name
  }

  user.Organizations = orgs

  return user, nil
}
