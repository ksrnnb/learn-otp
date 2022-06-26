# otp
This project is a simple implement of 2 factor authentication using time-based one-time password.

## Getting started

### 1. Set up Google Authenticator
Install Google Authneticator app. ([iPhone](https://apps.apple.com/app/google-authenticator/id388497605) / [Android](https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2&hl=en&gl=US))

Scan QR code below. This QR code has example user's key for generating one-time password.

<img src="./otp_example.png">

If it is difficult to scan QR code, please enter account and key directly.
| | |
| ---- | ---- |
|account|arbitrary name|
|key|NBSWY3DP|

`NBSWY3DP` is a Base32 encoded value of `hello`.

When you finished set up, 6 digits code will be generated every 30 seconds.

### 2. Run server
```bash
make run
```

### 3. Try to login
Access to localhost:8080 and try to login.

You will be required to enter one-time password after id/password login.
Then, enter one-time password shown in Google Authenticator.

| | |
| ---- | ---- |
|id|hogehoge|
|password|hogehoge|

## NOT implement
- Resynchronization
- Preventing CSRF attacks

## References
- [RFC4226](https://datatracker.ietf.org/doc/html/rfc4226)
- [RFC6238](https://datatracker.ietf.org/doc/html/rfc6238)
- [Google Authenticator Key Uri Format](https://github.com/google/google-authenticator/wiki/Key-Uri-Format)
