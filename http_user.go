package pubuser

import (
  "testing"
)

func TestHttpJson(t *testing.T) {
  user, err := GetHttpJson("https://gist.github.com/evanphx/8106877/raw")

  if err != nil {
    t.Fatalf("Could not get url with JSON")
  }

  if user.UserName() != "evanphx" {
    t.Errorf("Unable to get username")
  }
}
