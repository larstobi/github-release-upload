package main // import "github.com/larstobi/github-release-upload"

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

  // Open the release asset file
  file, err := os.Open(filename)
  if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }

  // Configure authentication
  ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
  tc := oauth2.NewClient(oauth2.NoContext, ts)
  client := github.NewClient(tc)

  // Make a commit to Releases.md
  message := releasename
  content := []byte(releasename)
  repositoryContentsOptions := &github.RepositoryContentFileOptions{
    Message:   &message,
    Content:   content,
    Committer: &github.CommitAuthor{Name: github.String("bonnyrigg"),
                Email: github.String("support@digipost.no")},
  }
  path := "releases/"
  path += releasename
  _, _, err = client.Repositories.CreateFile(owner, repo,
    path, repositoryContentsOptions)
  if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }

  // Create the release
  name := github.String(releasename)
  request := &github.RepositoryRelease{
    Name:    name,
    TagName: name,
  }
  release, _, err := client.Repositories.CreateRelease(owner, repo, request)
  if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }

  // Upload the release asset file to the release
  opt := &github.UploadOptions{Name: filename}
  releaseasset, _, err := client.Repositories.UploadReleaseAsset(owner, repo, *release.ID, opt, file)
  if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
  fmt.Printf("%s: Asset %s. Version: %s\n", *releaseasset.Name, *releaseasset.State, *name)
}
