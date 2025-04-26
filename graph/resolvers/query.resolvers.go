package resolvers

import (
	"context"
	"errors"
	"fmt"
	"graphql-go/graph"
	"graphql-go/graph/model"
)

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

func (r *queryResolver) Quiz(ctx context.Context, id string) (*model.Quiz, error) {
	for _, q := range r.quizzes {
		if q.ID == id {
			return q, nil
		}
	}
	return nil, errors.New("quiz not found")
}

func (r *queryResolver) Quizzes(ctx context.Context, pageSize *int32, pageToken *string) (*model.QuizzesConnection, error) {
	var size int32 = 10
	if pageSize != nil {
		size = *pageSize
	}

	startIndex := 0
	if pageToken != nil && *pageToken != "" {
		_, err := fmt.Sscanf(*pageToken, "%d", &startIndex)
		if err != nil {
			return nil, errors.New("invalid page token")
		}
	}

	endIndex := startIndex + int(size)
	if endIndex > len(r.quizzes) {
		endIndex = len(r.quizzes)
	}

	items := r.quizzes[startIndex:endIndex]

	var nextPageToken *string
	if endIndex < len(r.quizzes) {
		token := fmt.Sprintf("%d", endIndex)
		nextPageToken = &token
	}

	return &model.QuizzesConnection{
		Items:         items,
		NextPageToken: nextPageToken,
	}, nil
}

func (r *queryResolver) QuestionsByQuiz(ctx context.Context, quizID string) ([]*model.Question, error) {
	var result []*model.Question
	for _, q := range r.questions {
		if q.QuizID == quizID {
			result = append(result, q)
		}
	}
	return result, nil
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*model.User, error) {
	userID, ok := ctx.Value("userID").(string)
	if !ok || userID == "" {
		return nil, errors.New("not authenticated")
	}

	for _, u := range r.users {
		if u.ID == userID {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *queryResolver) Users(ctx context.Context, pageSize *int32, pageToken *string) (*model.UsersConnection, error) {
	var size int32 = 10
	if pageSize != nil {
		size = *pageSize
	}

	startIndex := 0
	if pageToken != nil && *pageToken != "" {
		_, err := fmt.Sscanf(*pageToken, "%d", &startIndex)
		if err != nil {
			return nil, errors.New("invalid page token")
		}
	}

	endIndex := startIndex + int(size)
	if endIndex > len(r.users) {
		endIndex = len(r.users)
	}

	items := r.users[startIndex:endIndex]

	var nextPageToken *string
	if endIndex < len(r.users) {
		token := fmt.Sprintf("%d", endIndex)
		nextPageToken = &token
	}

	return &model.UsersConnection{
		Items:         items,
		NextPageToken: nextPageToken,
	}, nil
}
