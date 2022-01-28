# Go Backend Coding Practice

This is my practice repo to learn about creating backend stuff & openid connect authentication stuff in golang.

## How to Start

- install golang
- `$git clone https://github.com/coding-CEO/go-backend-test.git`
- create _.env_ file in root directory and set following values,

        GOOGLE_OAUTH2_CLIENT_ID={YOUR CLIENT ID}
        GOOGLE_OAUTH2_CLIENT_SECRET={YOUR CLIENT SECRET}

- `$go run main.go`
- server should start now

## How to Use

1. start server
2. call _GET /googleGenerateUserOAuthCode_ API by setting request as mentioned in below API table from front-end, you will get google oauth url
3. redirect front-end to that url
4. perform google login
5. google will redirct to specific url that you set in step 2
6. call _GET /googleVerifyUserOAuthCode_ by setting request as mentioned in below API table from front-end
7. you will get JWT of your google account

## Basic Design

This is the basic design of what I want to build. \
Authenticate users of clients using single authentication-service-provider

    Clients:            ┌────>  a.com     b.in   c.org  <───────┐
                        │         ↓        ↓       ↓            │
    Front-End Repo:     │ ┌───>   login/signup interface <────┐ |
                        │ │                ↓                  │ │
    This Repo:          │ └─ authentication-service-provider ─┘ │
                        │     ↓            ↓          ↓         │
    OAuth Servers:      └─ google.com  apple.com  facebook.com ─┘

## API Routes

| Route                        | Type | body | output body            | req. cookie  | resp. cookie                                   | urlParams         | req. extras                 | details                                                                                                                                                                                                                                  |
| ---------------------------- | ---- | ---- | ---------------------- | ------------ | ---------------------------------------------- | ----------------- | --------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| /verifyUser                  | GET  | -    | -                      | token        | -                                              | -                 | -                           | This will verify the token from cookie and pass to next handler function. otherwise redirect to a central login page                                                                                                                     |
| /googleGenerateUserOAuthCode | GET  | -    | googleOAuthRedirectUrl | -            | Set-Cookie:state={xyz}, Set-Cookie:nonce={xyz} | targetRedirectUrl | set {withCredentials: true} | set verification cookies (state, nonce) and returns google oauth url(GOU) then front-end should redirect to GOU then GOU will automatically redirect to targetRedirectUrl where _code(OAuth code)_ param will automatically get attached |
| /googleVerifyUserOAuthCode   | GET  | -    | {JWT: 'xyz'}           | state, nonce | -                                              | state, code       | set {withCredentials: true} | verify the state from params & cookie, then verify the code(OAuth code) in the params, then verify nonce from idToken, then generate JWT and return as JSON. (maybe perform user registration in database here)                          |

## Who's Next

1. implement code_challenge stuff
2. maybe change the flow to reduce effort on front-end side
3. something else that will popup out of nowhere
