# JWT Authentication Microservice

## To-Do

* [x] Read environment variables
* [x] Load variables from file
* [ ] Implement JWT with refresh token, add issued on to JWT header
* [ ] Store JWT in cookies with HTTPOnly flag, also return in response for those who wants to use HTML5 storage
* [ ] Implement JWT blacklist
* [ ] Implement group
* [ ] Implement signin
* [ ] Implement signup
* [ ] Implement update
* [ ] Implement remember me
* [ ] Implement account confirmation
* [ ] Implement account delete

* [ ] Embed user_id and salt to JWT header
* [ ] When token revoked, store user_id and salt to token revocation list
* [ ] When user logged out at all places, store user_id and time to token revocation list
* [ ] Verify identity for sensitive routes
* [ ] Set short access token time
* [ ] Set long refresh token time when remember me enabled
* [ ] Avoid storing token in localStorage on client side, try storing in cookie

* [ ] Implement logrus
* [ ] Send out emails
