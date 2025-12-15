## To Do(lol):

### Where we are:
    - users can login, they get refresh tokens and access tokens. Should review it again might have some janky logic for access and refresh token.

### What should be done:
    - refresh tokens should be correctly stored in redis. They are not now.
    - Add get by name logic to redis. this will be used to check if the user that just logged in, already has a refresh token, this way we wont store too many refresh tokens for a single user just becuase they logged in too many times. (security risk in some ways)
    - change the token login logic. It should just cache hit and give back the stuff. shouldn't parse the token.
    - write unit tests and db tests for the ath logic.
    - **Big Deal:** Add the whole todo logic. --> in progress
