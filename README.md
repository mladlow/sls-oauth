# sls-oauth
I want to learn things about serverless oauth and about golang because my ambition knows no bounds.

## Building

```
make
```

## Deploying

```
serverless deploy --verbose --aws-profile <profile_name>
```

## Invoking

```
serverless invoke -f hello --aws-profile <profile_name>
```

## Removing
```
serverless remove --aws-profile <profile_name>
```

## Thoughts

I need a database. AWS Aurora Serverless seems like it does Postgres (what I want).

There's a Go library for Postgres (https://github.com/lib/pq).

I can (probably) use Flyway for migrations (https://flywaydb.org/), though this guy uses migrate (https://www.skitoy.com/p/aws-lambda-database-migrations/644/).

At this point, although it seems like overkill, probably should consider something like AWS CodeDeploy.

### 7/8/2019

Pivoted to DynamoDB. Will research how to actually structure data this way, but would prefer to hone Go and actual oauth skills, not AWS knowledge specifically.

Inserted an item into a DynamoDB table. Finally got permissions working and everything. Need to investigate variable storage and code refactoring.

Configured DynamoDB table using the AWS Console. Don't really want it linked to deploy & remove of my lambda functions anyway.

## Useful Links

* CloudFormation template basics (serverless.yml uses this) - https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/gettingstarted.templatebasics.html
* See also: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html

## Setup

1. I needed to install go. Thankfully there is a handy GSG. (https://golang.org/doc/install).
1. I have also on at least one computer needed to add `go` to my PATH.
1. After installation, I read about how to write go code. I wish more languages had this. I appreciate that it has testing instructions! (https://golang.org/doc/code.html).
1. I needed to install serverless (https://serverless.com). You need NPM for this, which thankfully I already have on this computer.
1. Serverless has three go templates for new projects. I went with the most vanilla-looking option (`aws-go`).
1. I needed to use `go get` to access the remote packages for AWS lambda.
1. I also needed to set up AWS lambda credentials on this computer (https://serverless.com/framework/docs/providers/aws/guide/credentials/).
