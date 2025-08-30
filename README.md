# TODO-API
Golang으로 구현한 간단한 Todo API 서버입니다. 사용자는 사용자 등록 및 로그인을 통해 인증한 후, 자신의 Todo 목록을 관리할 수 있습니다.
- RESTful API
- JWT 인증
- PostgreSQL DB 연동
- swagger API 문서화

## API 목록
- 사용자 등록: POST /auth/register
- 로그인: POST /auth/login
- Todo 생성: POST /todos
- Todo 조회: GET /todos
- Todo 수정: PUT /todos/{id}
- Todo 삭제: DELETE /todos/{id}
- API와 관련된 상세 정보는 API 문서(http://localhost:8080/docs/index.html) 를 참고하십시오.

## 빌드
```bash
docker compose build
```

## 실행
```bash
docker compose up

# 종료하려면 Ctrl+C 로 빠져나온 후
# docker compose down
```

## 요청 및 응답 예시
### 사용자 등록
```bash
curl -X 'POST' 'http://localhost:8080/auth/register' \
-H 'accept: application/json' \
-H 'Content-Type: application/json' \
-d '{
"username": "test001",
"password": "test1234"
}'
```
```json
{"message":"user registered"}
```
---
### 로그인
```bash
curl -X 'POST' 'http://localhost:8080/auth/login' \
-H 'accept: application/json' \
-H 'Content-Type: application/json' \
-d '{
"username": "test001",
"password": "test1234"
}'
```
```json
{"message":"login success","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTY2Mjk3MDMsInVzZXJuYW1lIjoidGVzdDAwMSJ9.QAphkxj-rRwLUp6Pi2SDMgPuI86bv7Os4qV8SFCCLNk"}}
```
---
### TODO 생성
```bash
curl -X 'POST' 'http://localhost:8080/todos' \
-H 'accept: application/json' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTY2Mjk3MDMsInVzZXJuYW1lIjoidGVzdDAwMSJ9.QAphkxj-rRwLUp6Pi2SDMgPuI86bv7Os4qV8SFCCLNk' \
-H 'Content-Type: application/json' \
-d '{
"completed": false,
"title": "todo001"
}'
```
```json
{"message":"success"}
```
---
### TODO 조회
```bash
curl -X 'GET' 'http://localhost:8080/todos' \
-H 'accept: application/json' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTY2Mjk3MDMsInVzZXJuYW1lIjoidGVzdDAwMSJ9.QAphkxj-rRwLUp6Pi2SDMgPuI86bv7Os4qV8SFCCLNk' \
-H 'Content-Type: application/json'
```
```json
[{"id":"575c2fe7-aecd-4c8d-931e-a9ef8553c7cd","title":"todo001","completed":false,"username":"test001"}]
```
---
### TODO 업데이트
```bash
curl -X 'PUT' 'http://localhost:8080/todos/575c2fe7-aecd-4c8d-931e-a9ef8553c7cd' \
-H 'accept: application/json' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTY2Mjk3MDMsInVzZXJuYW1lIjoidGVzdDAwMSJ9.QAphkxj-rRwLUp6Pi2SDMgPuI86bv7Os4qV8SFCCLNk' \
-H 'Content-Type: application/json' \
-d '{
"completed": true,
"title": "todo002"
}'
```
```json
{"message":"success"}
```
---
### TODO 삭제
```bash
curl -X 'DELETE' 'http://localhost:8080/todos/575c2fe7-aecd-4c8d-931e-a9ef8553c7cd' \
-H 'accept: application/json' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTY2Mjk3MDMsInVzZXJuYW1lIjoidGVzdDAwMSJ9.QAphkxj-rRwLUp6Pi2SDMgPuI86bv7Os4qV8SFCCLNk' \
-H 'Content-Type: application/json'
```
```json
{"message":"success"}
```
