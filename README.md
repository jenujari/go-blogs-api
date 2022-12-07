Go lang API demo

---

Steps to run the project and test.

1. Run following command in root dir for running mysql container.

```
docker-compose up -d
```

2. Run following command to build and run go lang server.

```
go build -o application.exe && ./application.exe
```

3. Test apis on localhost:3000 as specified in .env file.

4. For Unit test run following command.

```
go test -v ./test
```