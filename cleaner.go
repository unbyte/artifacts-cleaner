package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

type Cleaner struct {
	client *github.Client
}

func NewCleaner(token string) *Cleaner {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)

	return &Cleaner{
		client: client,
	}
}

func (i *Cleaner) Clean(owner, repo string) error {
	fmt.Printf("Starting clean %s/%s\n", owner, repo)
	artifacts, _, err := i.client.Actions.ListArtifacts(context.Background(), owner, repo, nil)
	if err != nil {
		return err
	}
	failed := make([]int64, 0, 10)
	successNum := 0
	failNum := 0
	total := len(artifacts.Artifacts)
	done := 0
	fmt.Printf("Total artifacts: %d\n", total)
	for _, artifact := range artifacts.Artifacts {
		if id := artifact.GetID(); id != 0 {
			_, err = i.client.Actions.DeleteArtifact(context.Background(), owner, repo, id)
			if err != nil {
				failNum++
				failed = append(failed, id)
			} else {
				successNum++
			}
		}
		done++
		fmt.Printf("\rprocessing %d/%d", done, total)
	}
	fmt.Println()
	fmt.Printf("Result: success %d, fail %d\n", successNum, failNum)
	fmt.Println("Failed artifacts:")
	for _, id := range failed {
		fmt.Printf("#%d\n", id)
	}
	return nil
}
