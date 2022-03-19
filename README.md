# How to configure CircleCI for GoLang Application

**In progress**

One of the impotent thing in our projects is the CI process. Continuous Integration (CI) is the practice of automating the integration of code changes and guarantee the quality of the software. If CI is so important, why we don't configure this process in our personal projects? Maybe because we thing is so harder to configure or even it's take so much time. In this post I'm going to show you a simple way to configure a strong CI process using some of the best tools for that, [CircleCI](https://circleci.com/) and [SonarQuebe](https://sonarcloud.io/).

For this post we are configure a CI for a simple GoLang application creating a multistage CI each stage we called Job, in the end of this post we're going to have four jobs: **Lint**, **Test and Coverage**, **Quality Analises with SonarQuebe** and **Build**. The proposal where is to explain the CI not build a GoLang application so we assume you already know the GoLang basics and some tools [Test Pkg](https://pkg.go.dev/cmd/go/internal/test), [GolangCI Lint](https://golangci-lint.run/) also the basics about Github and Github Actions.

The project that was built can be found in [this repository](https://github.com/ralvescosta/medium_go_and_circleci)