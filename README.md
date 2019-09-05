# psqlxtest

```go
func TestSomething(t *testing.T) {
	dbx, drop := psqlxtest.TmpDB(t)
	defer drop()

	// Run in an isolated postgres database instance.
	dbx.Exec(/* ... */)
}
```
