package pubuser

import (
  "testing"
)

func TestGithubProtocol(t *testing.T) {
  user, err := Find("github:evanphx")

  if err != nil {
    t.Fatalf("Unable to find user: %s", err)
  }

  if user.UserName() != "evanphx" {
    t.Errorf("Error getting user name")
  }
}
