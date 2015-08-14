package main

import (
  "fmt"
  "os"
  "github.com/google/go-github/github"
  "golang.org/x/oauth2"
)

func main() {
  owner       := os.Getenv("GITHUB_OWNER")
  repo        := os.Getenv("GITHUB_REPO")
  token       := os.Getenv("GITHUB_AUTH_TOKEN")
  filename    := os.Getenv("GITHUB_RELEASE_ASSET")
  releasename := os.Getenv("GITHUB_RELEASE_NAME")

  // 1. Open the release asset file
  file, err := os.Open(filename)
  if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }

  // 2. Configure authentication
  ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
  tc := oauth2.NewClient(oauth2.NoContext, ts)
  client := github.NewClient(tc)

  // 3. Create the release
  name := github.String(releasename)
  request := &github.RepositoryRelease{
    Name:    name,
    TagName: name,
  }
  release, _, err := client.Repositories.CreateRelease(owner, repo, request)
  if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }

  // 4. Upload the release asset file to the release
  opt := &github.UploadOptions{Name: repo}
  releaseasset, _, err := client.Repositories.UploadReleaseAsset(owner, repo, *release.ID, opt, file)
  if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
  fmt.Printf("%s: Asset %s. Version: %s\n", *releaseasset.Name, *releaseasset.State, *name)
}
