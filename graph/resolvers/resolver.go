package resolvers

import "graphql-go/graph/model"

type Resolver struct {
	quizzes    []*model.Quiz
	questions  []*model.Question
	users      []*model.User
	userTokens map[string]string
}

func NewResolver() *Resolver {
	return &Resolver{
		quizzes:    []*model.Quiz{},
		questions:  []*model.Question{},
		users:      []*model.User{},
		userTokens: make(map[string]string),
	}
}

func (r *Resolver) GetUserIDByToken(token string) (string, bool) {
	userID, exists := r.userTokens[token]
	return userID, exists
}
