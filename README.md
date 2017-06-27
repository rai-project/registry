
# Registry [![Build Status](https://travis-ci.org/rai-project/registry.svg?branch=master)](https://travis-ci.org/rai-project/registry)

## Config

~~~
registry:
  provider: etcd
  endpoints:
    - 127.0.0.1
  username: root
  password: foo
  timeout: 5s
  certificate: XXX
  bucket: name_of_bucket
  header_timeout_per_request: 1m
  auto_sync: true
~~~
