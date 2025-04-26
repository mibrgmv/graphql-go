## Create a User

```graphql
mutation CreateUser {
    userCreate(user: { name: "John Doe", email: "john@example.com" }) {
        id
        name
        email
    }
}
```

## Create a Todo for that User

```graphql
mutation CreateTodo {
    todoCreate(todo: { text: "Learn GraphQL", userId: "U23" }) {
        id
        text
        done
        user {
            id
            name
        }
    }
}
```

## Query a Todo with User details

```graphql
query GetTodo {
    todo(id: "T42") {
        id
        text
        done
        user {
            id
            name
            email
        }
    }
}
```

## Complete a Todo

```graphql
mutation CompleteTodo {
    todoComplete(id: 42, updatedBy: 23) {
        id
        text
        done
    }
}
```

## List all Todos

```graphql
query ListTodos {
    todos(limit: 5, offset: 0) {
        id
        text
        done
        user {
            name
            email
        }
    }
}
```

## List all Users

```graphql
query ListUsers {
    users {
        id
        name
        email
        todos {
            id
            text
            done
        }
    }
}
```

## Delete a Todo (unimplemented)

```graphql
mutation DeleteTodo {
    todoDelete(id: 42, updatedBy: 23) {
        id
        text
    }
}
```

## Query a User with their Todos

```graphql
query GetUser {
    user(id: "U23") {
        id
        name
        email
        todos {
            id
            text
            done
        }
    }
}
```