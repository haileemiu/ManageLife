# ManageLife

My personal task management (aka "Adulting") application.

## Development Notes

- `go run .`
- http://localhost:4000/api/tasks/
- Rebuild = ctrl + shft + p; usually only when **change .devcontainer files**
- `go mod tidy`

- Use "C:\Users\HaileeMiu\source\repos\ManageLife1" C# project as a reference

```js
fetch('http://localhost:4000/api/tasks', {
  method: 'POST',
  body: JSON.stringify({
    title: "chrome title",
    notes: "chrome notes"
  }),
  headers: {
    'Content-type': 'application/json; charset=UTF-8'
  }
})
.then(res => res.json())
.then(console.log)
```
## Notes to save and reference later in the project 

- `dateOnly := time.Date(2023, time.October, 7, 0, 0, 0, 0, time.UTC)`

## Troubleshooting

- container terminal not as expected: ctrl + shft + p --> rebuild

## Packages Using

- [Chi Router](https://github.com/go-chi/chi)
- go ent
  - `go generate ./..`
  - `go run -mod=mod entgo.io/ent/cmd/ent new Task`
  - `go generate ./ent`

# Notes

```go
package main

func main() {
	var x int = 3

	z := &x // & = address of (creates a pointer to the location in memory where x is)
	zz := *z // * = dereference (undoes a pointer to a memory location returning the original memory at the location)

	var y *int // pointer to an int
}

```
- entClient:
  - Object Relational Mapper: Manages the relationships of your database
  - Maps and manages your schema to the db model so that you don't have to
- Ability to create complicated queries with little work on my side