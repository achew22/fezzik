package github

import (
	"context"

	"github.com/inigolabs/fezzik"
)

var GetInfoOperation string = `
query GetInfo($repo_owner : String!, $repo_name : String!) {
   viewer {
      login
      pullRequests(first: 100, states: [OPEN]) {
         nodes {
            id
            number
            title
            baseRefName
            headRefName
            mergeable
            reviewDecision
            repository {
               id
            }
            commits(first: 100) {
               nodes {
                  commit {
                     oid
                     messageHeadline
                     messageBody
                     statusCheckRollup {
                        state
                     }
                  }
               }
            }
         }
      }
   }
   repository(owner: $repo_owner, name: $repo_name) {
      id
   }
}
`

type GetInfoInputArgs struct {
	RepoOwner string `json:"repo_owner"`
	RepoName  string `json:"repo_name"`
}

type GetInfoResponse struct {
	Viewer struct {
		Login        string
		PullRequests struct {
			Nodes *[]*struct {
				Id             string
				Number         int
				Title          string
				BaseRefName    string
				HeadRefName    string
				Mergeable      MergeableState
				ReviewDecision *PullRequestReviewDecision
				Repository     struct {
					Id string
				}
				Commits struct {
					Nodes *[]*struct {
						Commit struct {
							Oid               string
							MessageHeadline   string
							MessageBody       string
							StatusCheckRollup *struct {
								State StatusState
							}
						}
					}
				}
			}
		}
	}
	Repository *struct {
		Id string
	}
}

func (c *client) GetInfo(ctx context.Context, input *GetInfoInputArgs) (
	*GetInfoResponse, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "GetInfo",
		Query:     GetInfoOperation,
		Variables: map[string]interface{}{
			"repo_owner": input.RepoOwner,
			"repo_name":  input.RepoName,
		},
	}

	var gqldata GetInfoResponse
	var gqlerrs fezzik.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

var CreatePullRequestOperation string = `
mutation CreatePullRequest($input : CreatePullRequestInput!) {
   createPullRequest(input: $input) {
      pullRequest {
         id
         number
      }
   }
}
`

type CreatePullRequestInputArgs struct {
	Input CreatePullRequestInput `json:"input"`
}

type CreatePullRequestResponse struct {
	CreatePullRequest *struct {
		PullRequest *struct {
			Id     string
			Number int
		}
	}
}

func (c *client) CreatePullRequest(ctx context.Context, input *CreatePullRequestInputArgs) (
	*CreatePullRequestResponse, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "CreatePullRequest",
		Query:     CreatePullRequestOperation,
		Variables: map[string]interface{}{
			"input": input.Input,
		},
	}

	var gqldata CreatePullRequestResponse
	var gqlerrs fezzik.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

var UpdatePullRequestOperation string = `
mutation UpdatePullRequest($input : UpdatePullRequestInput!) {
   updatePullRequest(input: $input) {
      pullRequest {
         number
      }
   }
}
`

type UpdatePullRequestInputArgs struct {
	Input UpdatePullRequestInput `json:"input"`
}

type UpdatePullRequestResponse struct {
	UpdatePullRequest *struct {
		PullRequest *struct {
			Number int
		}
	}
}

func (c *client) UpdatePullRequest(ctx context.Context, input *UpdatePullRequestInputArgs) (
	*UpdatePullRequestResponse, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "UpdatePullRequest",
		Query:     UpdatePullRequestOperation,
		Variables: map[string]interface{}{
			"input": input.Input,
		},
	}

	var gqldata UpdatePullRequestResponse
	var gqlerrs fezzik.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

var CommentPullRequestOperation string = `
mutation CommentPullRequest($input : AddCommentInput!) {
   addComment(input: $input) {
      clientMutationId
   }
}
`

type CommentPullRequestInputArgs struct {
	Input AddCommentInput `json:"input"`
}

type CommentPullRequestResponse struct {
	AddComment *struct {
		ClientMutationId *string
	}
}

func (c *client) CommentPullRequest(ctx context.Context, input *CommentPullRequestInputArgs) (
	*CommentPullRequestResponse, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "CommentPullRequest",
		Query:     CommentPullRequestOperation,
		Variables: map[string]interface{}{
			"input": input.Input,
		},
	}

	var gqldata CommentPullRequestResponse
	var gqlerrs fezzik.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

var MergePullRequestOperation string = `
mutation MergePullRequest($input : MergePullRequestInput!) {
   mergePullRequest(input: $input) {
      pullRequest {
         number
      }
   }
}
`

type MergePullRequestInputArgs struct {
	Input MergePullRequestInput `json:"input"`
}

type MergePullRequestResponse struct {
	MergePullRequest *struct {
		PullRequest *struct {
			Number int
		}
	}
}

func (c *client) MergePullRequest(ctx context.Context, input *MergePullRequestInputArgs) (
	*MergePullRequestResponse, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "MergePullRequest",
		Query:     MergePullRequestOperation,
		Variables: map[string]interface{}{
			"input": input.Input,
		},
	}

	var gqldata MergePullRequestResponse
	var gqlerrs fezzik.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

var ClosePullRequestOperation string = `
mutation ClosePullRequest($input : ClosePullRequestInput!) {
   closePullRequest(input: $input) {
      pullRequest {
         number
      }
   }
}
`

type ClosePullRequestInputArgs struct {
	Input ClosePullRequestInput `json:"input"`
}

type ClosePullRequestResponse struct {
	ClosePullRequest *struct {
		PullRequest *struct {
			Number int
		}
	}
}

func (c *client) ClosePullRequest(ctx context.Context, input *ClosePullRequestInputArgs) (
	*ClosePullRequestResponse, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "ClosePullRequest",
		Query:     ClosePullRequestOperation,
		Variables: map[string]interface{}{
			"input": input.Input,
		},
	}

	var gqldata ClosePullRequestResponse
	var gqlerrs fezzik.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}
