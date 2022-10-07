# Go Generic Data Access Layer

This is generic repository, which can be used in any project instead of writing separate repository for each model.

### Usage

```go
type User struct {
	gorm.Model
	FirstName string
	LastName  string
}

db, _ := gorm.Open(...)

repository := NewRepository[User](db)

repository.Create(context.TODO(), &User{Name: "Sumit"})
```
