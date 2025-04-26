package resolvers

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"graphql-go/graph"
	"graphql-go/graph/model"
	"time"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) CreateQuiz(ctx context.Context, input model.QuizInput) (*model.Quiz, error) {
	_, ok := ctx.Value("userID").(string)
	if !ok {
		return nil, errors.New("not authenticated")
	}

	quiz := &model.Quiz{
		ID:      uuid.New().String(),
		Title:   input.Title,
		Results: input.Results,
	}

	r.quizzes = append(r.quizzes, quiz)
	return quiz, nil
}

func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.QuestionInput) (*model.Question, error) {
	_, ok := ctx.Value("userID").(string)
	if !ok {
		return nil, errors.New("not authenticated")
	}

	quizExists := false
	for _, q := range r.quizzes {
		if q.ID == input.QuizID {
			quizExists = true
			break
		}
	}
	if !quizExists {
		return nil, errors.New("quiz not found")
	}

	options := []string{}
	optionsWeights := map[string][]float64{}

	for _, ow := range input.OptionsWeights {
		options = append(options, ow.Option)

		weights := make([]float64, len(ow.Weights))
		for i, w := range ow.Weights {
			weights[i] = float64(w)
		}

		optionsWeights[ow.Option] = weights
	}

	question := &model.Question{
		ID:      uuid.New().String(),
		QuizID:  input.QuizID,
		Body:    input.Body,
		Options: options,
	}

	r.questions = append(r.questions, question)
	return question, nil
}

func (r *mutationResolver) Register(ctx context.Context, username string, password string) (*model.User, error) {
	for _, u := range r.users {
		if u.Username == username {
			return nil, errors.New("username already exists")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now().Format(time.RFC3339)
	user := &model.User{
		ID:        uuid.New().String(),
		Username:  username,
		Password:  string(hashedPassword),
		CreatedAt: now,
		LastLogin: now,
	}

	r.users = append(r.users, user)
	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.LoginResponse, error) {
	var user *model.User
	var hashedPassword []byte

	for _, u := range r.users {
		if u.Username == username {
			user = u
			hashedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			break
		}
	}

	if user == nil {
		return nil, errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	token := uuid.New().String()
	r.userTokens[token] = user.ID

	user.LastLogin = time.Now().Format(time.RFC3339)

	return &model.LoginResponse{
		Token:  token,
		UserID: user.ID,
		User:   user,
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	currentUserID, ok := ctx.Value("userID").(string)
	if !ok {
		return false, errors.New("not authenticated")
	}

	if currentUserID != userID {
		return false, errors.New("not authorized")
	}

	for i, u := range r.users {
		if u.ID == userID {
			r.users = append(r.users[:i], r.users[i+1:]...)

			for token, id := range r.userTokens {
				if id == userID {
					delete(r.userTokens, token)
				}
			}

			return true, nil
		}
	}

	return false, errors.New("user not found")
}

func (r *mutationResolver) EvaluateAnswers(ctx context.Context, quizID string, answers []*model.AnswerInput) (*model.EvaluationResult, error) {
	var quiz *model.Quiz
	for _, q := range r.quizzes {
		if q.ID == quizID {
			quiz = q
			break
		}
	}
	if quiz == nil {
		return nil, errors.New("quiz not found")
	}

	if len(quiz.Results) == 0 {
		return nil, errors.New("quiz has no results defined")
	}

	resultIndex := len(answers) % len(quiz.Results)
	result := quiz.Results[resultIndex]

	return &model.EvaluationResult{
		Result: result,
	}, nil
}

//type quizResolver struct{ *Resolver }
//
//func (r *Resolver) Quiz() QuizResolver { return &quizResolver{r} }
//
//func (r *quizResolver) Questions(ctx context.Context, obj *model.Quiz) ([]*model.Question, error) {
//	var result []*model.Question
//	for _, q := range r.questions {
//		if q.QuizID == obj.ID {
//			result = append(result, q)
//		}
//	}
//	return result, nil
//}
//
//type questionResolver struct{ *Resolver }
//
//func (r *Resolver) Question() QuestionResolver { return &questionResolver{r} }
//
//func (r *questionResolver) Quiz(ctx context.Context, obj *model.Question) (*model.Quiz, error) {
//	for _, q := range r.quizzes {
//		if q.ID == obj.QuizID {
//			return q, nil
//		}
//	}
//	return nil, errors.New("quiz not found")
//}
