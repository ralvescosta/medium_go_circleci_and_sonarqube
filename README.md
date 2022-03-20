# How to configure CircleCI for GoLang Application

[**In progress**]

## Table of contents

  - [Introduction](#introduction)
  - [Initial Configuration](#initial-configuration)
  - [Build Job](#build-job)
  - [Lint Job](#lint-job)
  - [Test and coverage Job](#test-and-coverage-job)

## Introduction

One of the impotent thing in our projects is the CI process. Continuous Integration (CI) is the practice of automating the integration of code changes and guarantee the quality of the software. If CI is so important, why we don't configure this process in our personal projects? Maybe because we thing is so harder to configure or even it's take so much time. In this post I'm going to show you a simple way to configure a strong CI process using some of the best tools for that, [CircleCI](https://circleci.com/) and [SonarQuebe](https://sonarcloud.io/).

For this post we are configure a CI for a simple GoLang application creating a multistage CI, each stage we called Job, in the end of this post we're going to have four jobs: **Lint**, **Test and Coverage**, **Quality Analises with SonarQuebe** and **Build**. The proposal where is to explain the CI not build a GoLang application so we assume you already know the GoLang basics and some tools [Test Pkg](https://pkg.go.dev/cmd/go/internal/test), [GolangCI Lint](https://golangci-lint.run/) also the basics about Github and Github Actions.

The project that was built can be found in [this repository](https://github.com/ralvescosta/medium_go_and_circleci).

## Initial Configurations

First we need to create a yaml file to configure our CI processes, for CircleCI this file need to be create in a specific directory:

```bash

make .circleci

touch .circleci/config.yml

```

We start our config.yml like this:

```yml

version: 2.1

jobs:
  - job_name

workflows:
  - workflow_name
      jobs:
        - job_name

```

We can see two main tags: 'jobs' and 'workflows'. Basically the tag 'jobs' we define the job execution flow and in the 'workflows' how to execute the 'jobs'. Let's start with the build job:

## Build Job

```yml
version: 2.1

jobs:
  build:
    working_directory: ~/repo
    docker:
      - image: circleci/golang:1.17
    steps:
      - checkout
      - restore_cache:
          keys:
            - go-mod-v4-{{ checksum "go.sum" }}

      - run:
          name: Install Dependencies
          command: go mod download
      - save_cache:
          key: go-mod-v4-{{ checksum "go.sum" }}
          paths:
            - "/go/pkg/mod"

      - run:
          name: Run build
          command: |
            mkdir -p /tmp/artifacts/build
            go build -ldflags "-s -w" -o exec main.go
            mv exec /tmp/artifacts/build
      - store_artifacts:
          path: /tmp/artifacts/build

workflows:
  ci:
    jobs:
      - build
```

This job is self explanatory basically we downloaded the project packages and build our project. With this job we can run our first pipeline, but first we need to publish the config.yml in the repository and then we need to configure our project in CircleCi Projects.

<img src="./assets/1.png" />

At this time the CircleCi will run for the first time our pipeline 🚀🚀.


## Lint Job

Now let's create our lint job

```yml
jobs:
  lint:
    working_directory: ~/repo
    docker:
      - image: golangci/golangci-lint:v1.45
    steps:
      - checkout
      - run: golangci-lint run ./... --out-format=checkstyle --print-issued-lines=false --print-linter-name=false --issues-exit-code=0 --enable=revive > golanci-report.xml
      - persist_to_workspace:
          root: ~/repo
          paths: 
            - golanci-report.xml
```

For lint we used the [GolangCI Lint](https://golangci-lint.run/) but we have some configurations to do, how you can see the golangci-lint execution has a bunch of arguments, let's talk about that. The first argument '--out-format' needed to be 'checkstyle' for integrate better with sonar. Other flag importante to talk about is --issues-exit-code=0, we need to configure this flag with 0 because whe golangci execute and found something wrong in our project the executor will finish with exit code igual 1 and when our pipeline receives this result it failure, to avoid the pipeline failure and get the report to send for the sonar we needed to change the flag --issues-exit-code. Before the runner we have a a step to persiste the report file in a path to allowed others job to access it.

For this point we can now configure our workflows with theses jobs:

```yml
workflows:
  ci:
    jobs:
      - lint
      - build
        - requires:
            - lint

```

In our workflows we can see something diferente, it's because we wanted the build job execute only after the lint job finished their execution.

## Test and coverage Job

```yml
jobs:
  test_and_coverage:
    working_directory: ~/repo
    docker:
      - image: circleci/golang:1.17
    steps:
      - checkout
      - persist_to_workspace:
          root: ~/repo
          paths:
            - pkg
            - sonar-project.properties
      - restore_cache:
          keys:
            - go-mod-v4-{{ checksum "go.sum" }}

      - run:
          name: Install Dependencies
          command: go mod download

      - save_cache:
          key: go-mod-v4-{{ checksum "go.sum" }}
          paths:
            - "/go/pkg/mod"

      - run:
          name: Run unit tests
          command: |
            mkdir -p /tmp/test-reports
            gotestsum --junitfile /tmp/test-reports/unit-tests.xml
      - store_test_results:
          path: /tmp/test-reports

      - run:
          name: Run coverage
          command: |
            go test ./... -race -coverprofile=coverage.out -json > report.json
      - persist_to_workspace:
          root: ~/repo
          paths: 
            - coverage.out
            - report.json
```