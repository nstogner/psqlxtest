# psqlxtest

Create an isolated postgres database instance per test case.

## Example test

```go
func TestSomething(t *testing.T) {
	dbx, drop := psqlxtest.TmpDB(t)
	defer drop()

	// Run in an isolated postgres database instance.
	dbx.Exec(/* ... */)
}
```

## Example usage

```sh
# Run postgres
docker run --name my-test-postgres -p 5432:5432 -d postgres

# Run your tests
go test ./mypkg
```
