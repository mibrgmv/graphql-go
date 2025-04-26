# 123

```
{
  "Authorization": "c0072717-a887-45d9-9344-877cbeb3502b"
}
```

# User

## Register a User

```graphql
mutation RegisterNewUser {
    register(username: "newuser", password: "securepassword") {
        id
        username
        createdAt
    }
}
```

## Login

```graphql
mutation LoginExistingUser {
    login(username: "newuser", password: "securepassword") {
        token
        user {
            id
            username
        }
    }
}
```

## Get Current User

```graphql
query GetCurrentUser {
    currentUser {
        id
        username
        createdAt
        lastLogin
    }
}
```

## Get Users

```graphql
query GetAllUsersWithPagination {
    users(pageSize: 5) {
        items {
            id
            username
            createdAt
            lastLogin
        }
        nextPageToken
    }
}
```

# Quiz

## Create Quiz
```graphql
mutation CreateQuiz {
  createQuiz(input: {
    title: "Personality Test",
    results: ["Introvert", "Extrovert", "Ambivert"]
  }) {
    id
    title
    results
  }
}
```

## Get Quiz
```graphql
query GetQuiz {
  quiz(id: "QUIZ_ID_FROM_CREATE_QUIZ") {
    id
    title
    results
    questions {
      id
      body
      options
    }
  }
}
```

## Get Quizzes
```graphql
query GetQuizzes {
  quizzes(pageSize: 10) {
    items {
      id
      title
      results
    }
    nextPageToken
  }
}
```

# Question

## Create Question
```graphql
mutation CreateQuestion {
  createQuestion(input: {
    quizId: "QUIZ_ID_FROM_CREATE_QUIZ",
    body: "Do you enjoy social gatherings?",
    optionsWeights: [
      {
        option: "Yes, always",
        weights: [0.2, 0.8, 0.5]
      },
      {
        option: "Sometimes",
        weights: [0.5, 0.5, 0.8]
      },
      {
        option: "No, rarely",
        weights: [0.8, 0.2, 0.3]
      }
    ]
  }) {
    id
    body
    options
    quizId
  }
}
```

## Get Question by Quiz
```graphql
query GetQuestionsByQuiz {
  questionsByQuiz(quizId: "QUIZ_ID_FROM_CREATE_QUIZ") {
    id
    body
    options
  }
}
```