学习笔记

```
➜  Go-001 git:(main) ✗ curl 127.0.0.1:8080            
hello golang, service is running.
➜  Go-001 git:(main) ✗ curl 127.0.0.1:8081
hello golang, manage is running.
➜  Go-001 git:(main) ✗ curl 127.0.0.1:8081/close
closing%                                                                                                                                                  
➜  Go-001 git:(main) ✗



➜  Week03 git:(main) ✗ go run main.go
receive service close! api shutdown
service closed http: Server closed 
service close <nil> 
manage closed http: Server closed 
manage close <nil> 
group err <nil> 
➜  Week03 git:(main) ✗ 
```