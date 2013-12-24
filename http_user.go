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

func TestHttpJsonParse(t *testing.T) {
  data := `{
  "username": "evanphx",
  "name": "Evan Phoenix",
  "ssh_keys": [
    "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDNKQcKibi+gHYqphMje9lOzwDbM9q4NHPbeuD0022X7g22QNuQDHwsAlJ/JpGKU240302sahy7cMfBNU5+Rp2kzCI/XE9NdCIC94TlDtV+l9E7n8uI68/Of1i3xupmLdv8NB6Y8F/RQ2qOfZxsbdHHdLsSTxBOc7jWV98eaPyLLuV7N+q76ADW9d+0e9PmGwOfkchMXFq7N5IlSmzL3NCkcd6LzCCV0c1eh3/MbdzY/4vavmrDfRct8EI6Jy14lDexymQ8eG/w2zrkv0Zt8ZMkCdJPk+pmm7PQ5TlJs1GTVLAMVrGjtuvLRs1CNkRGEJ3ye/1IE53URaYaCGVjpn/v evan@aero.local"
  ],
  "groups": ["test"]
}`

  user, err := ParseJSON([]byte(data))

  if err != nil {
    t.Fatalf("Could not get url with JSON")
  }

  if user.UserName() != "evanphx" {
    t.Errorf("Unable to get username")
  }
}
