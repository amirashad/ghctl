# Example repo for automatize repo creation with CI/CD

In this example we'll use CircleCI as a CI/CD tool

## 1) Create CircleCI config file

CircleCI config (`.circleci/config.yml`) will help you to use ghctl docker image to run ghctl commands inside

## 2) Create repository config file

Describe your repository config in some YAML file, e.g. `repositories/ms-example.yml`

## 3) Create CircleCI context 

CircleCI context will help you to keep environment variables secure. 
Create two environment variables in that context:
  - `GITHUB_ORG` environment variable with value for your organization
  - `GITHUB_TOKEN` environment variable with value for your API token. More about [creating API token](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line) 

## 4) Apply repo-creation

On push to master, CircleCI will apply your changes and will create repository with following command:
```command
$ ghctl apply -f ms-example.yml
```