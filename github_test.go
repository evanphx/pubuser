package pubuser

import (
  "testing"
)

func TestGithubUser(t *testing.T) {
  user, err := GetGithubUser("evanphx")

  if err != nil {
    t.Fatalf("Error getting user: %s", err)
  }

  if user.Name != "Evan Phoenix" {
    t.Errorf("Couldn't get full name")
  }

  if len(user.GHSSHKeys) == 0 {
    t.Errorf("Unable to pull SSH keys")
  }

  if len(user.Organizations) == 0 {
    t.Errorf("Unable to pull organizations")
  }
}

func TestGithubUserMissing(t *testing.T) {
  _, err := GetGithubUser("thisisnotanaccount")

  if err == nil {
    t.Fatalf("Error not reported")
  }
}
