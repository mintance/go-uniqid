# go-uniqid

Simple php uniqid() realisation on Golang

### How to use

##### Just include our package

```go get github.com/mintance/go-uniqid```

##### See samples

In PHP was:
```php
$id = uniqid("test", true);
```

In Go type:
```go
id := uniqid.New(uniqid.Params{"test", true})
```

