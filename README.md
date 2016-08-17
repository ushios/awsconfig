ushios/awsconfig
=================

Get aws config from IAMRole or credentials file.


[![Build Status](https://travis-ci.org/ushios/awsconfig.svg?branch=master)](https://travis-ci.org/ushios/awsconfig)
[![Coverage Status](https://coveralls.io/repos/github/ushios/awsconfig/badge.svg?branch=master)](https://coveralls.io/github/ushios/awsconfig?branch=master)


Installation
=============

```bash
$ go get github.com/ushios/awsconfig
```

Documentation
==============

[![GoDoc](https://godoc.org/github.com/ushios/awsconfig?status.svg)](https://godoc.org/github.com/ushios/awsconfig)


Samples
=======


### Get s3 client

```go
import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

s3client := s3.New(session.New(awsconfig.Config()))
```
