# sls-oauth
I want to learn things about serverless oauth and about golang because my ambition knows no bounds.

## Building

```
make
```

## Deploying

```
serverless deploy --aws-profile <profile_name>
```

## Invoking

```
serverless invoke -f hello --aws-profile <profile_name>
```

## Setup

1. I needed to install go. Thankfully there is a handy GSG. (https://golang.org/doc/install)
1. After installation, I read about how to write go code. I wish more languages had this. I appreciate that it has testing instructions! (https://golang.org/doc/code.html)
1. I needed to install serverless (https://serverless.com). You need NPM for this, which thankfully I already have on this computer.
1. Serverless has three go templates for new projects. I went with the most vanilla-looking option (`aws-go`).
1. I needed to use `go get` to access the remote packages for AWS lambda.
1. I also needed to set up AWS lambda credentials on this computer (https://serverless.com/framework/docs/providers/aws/guide/credentials/).
