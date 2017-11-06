# Execute shell command from http request with authentication

## Use

this process run in the host or vm host and provide http RESTful api receving requests,
and run the CMD field in http request in sync way, and return the result to http client.

## Example

1. Use config.yml in repo, command this:

```
$ ./http4shell -c config.yml -t -u cloud
User cloud token is:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNsb3VkIn0.FeD2yXXGQe290fFO7_FP_XoRubukCiboa2PYID3bAcc
```

2. Run the agent in host

```
$ ./http4shell -c config.yml
```

3. In remote client request in curl as follow:

```
$  curl -XPOST http://server:8080/run -i -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNsb3VkIn0.FeD2yXXGQe290fFO7_FP_XoRubukCiboa2PYID3bAcc" -d '{"type":"shell","cmd":"ls"}'
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 03 Nov 2017 10:44:14 GMT
Content-Length: 117

{"Errno":0,"Errmsg":"OK","Body":"api\nauth\ncmd\nconfig\nconfig.yml\nhttp\nhttp4shell\nmain.go\nREADME.md\nvendor\n"}
```
