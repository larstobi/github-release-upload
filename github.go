package main

import (
  "fmt"
  "os"
  "github.com/google/go-github/github"
  "golang.org/x/oauth2"
)

func main() {
  owner:= os.Getenv("GITHUB_OWNER")
  repo := os.Getenv("GITHUB_REPO")
  token := os.Getenv("GITHUB_AUTH_TOKEN")
  filename := os.Getenv("GITHUB_RELEASE_ASSET")
  id := 1661168

  // Authentication
  ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
  tc := oauth2.NewClient(oauth2.NoContext, ts)

  client := github.NewClient(tc)

  opt := &github.UploadOptions{Name: "KOBFD"}

  file, err := os.Open(filename)

  if err != nil {
    panic(err)
  }

  releaseasset, _, err := client.Repositories.UploadReleaseAsset(owner, repo, id, opt, file)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  } else {
    fmt.Printf("%v\n", releaseasset)
  }

}
