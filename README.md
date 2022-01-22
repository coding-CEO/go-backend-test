# Go Backend Coding Practice

This is my practice repo to learn about creating backend stuff, openid connect authentication stuff in golang.

## Basic Design

This is the basic design of what I want to build. \
Authenticate users of clients using single authentication-service-provider

    Clients:        ┌───>  a.com     b.in   c.org  <────────┐
                    │        ↓        ↓       ↓             │
    This Repo:      │   authentication-service-provider     │
                    │     ↓           ↓          ↓          │
    OAuth Servers:  └─ google.com  apple.com  facebook.com ─┘

## API Routes

| Route                        | Type | body | output body  | req. cookie  | resp. cookie                                   | urlParams         | expected HTTP Status Code | details                                                                                                                                                                                                         |
| ---------------------------- | ---- | ---- | ------------ | ------------ | ---------------------------------------------- | ----------------- | ------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| /verifyUser                  | GET  | -    | -            | token        | -                                              | -                 | 200 / 302                 | This will verify the token from cookie and pass to next handler function. otherwise redirect to a central login page                                                                                            |
| /googleGenerateUserOAuthCode | GET  | -    | -            | -            | Set-Cookie:state={xyz}, Set-Cookie:nonce={xyz} | targetRedirectUrl | 302                       | set verification cookies (state, nonce) and redirect to google oauth url then redirect to targetRedirectUrl where _code(OAuth code)_ param will automatically get attached                                      |
| /googleVerifyUserOAuthCode   | GET  | -    | {JWT: 'xyz'} | state, nonce | -                                              | state, code       | 200                       | verify the state from params & cookie, then verify the code(OAuth code) in the params, then verify nonce from idToken, then generate JWT and return as JSON. (maybe perform user registration in database here) |
