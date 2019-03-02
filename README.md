# getLatest

Get repository's latest version from GitHub tags.

### Usage

```console
$ glv --repo erlang/otp
{"name":"OTP-21.2.6","commit":{"sha":"f30b1052c7097a95faaba272feccc6190682a7f8","url":"https://api.github.com/repos/erlang/otp/commits/f30b1052c7097a95faaba272feccc6190682a7f8"},"zipball_url":"https://api.github.com/repos/erlang/otp/zipball/OTP-21.2.6","tarball_url":"https://api.github.com/repos/erlang/otp/tarball/OTP-21.2.6"}
```

### Build

```console
$ export GO111MODULE=on
$ make build
```
