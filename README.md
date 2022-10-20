# workspaces

A small command line utility to quickly switch between repositories using a fuzzy finder.

<p align="center">
  <img src="https://user-images.githubusercontent.com/7422050/103565201-32ecd000-4ec0-11eb-9eb8-7e2c180519c6.png" alt="Workspaces Preview"/>
</p>

## Usage

Install by running `go install github.com/fawind/workspaces@latest`.

**Commands:**
* `workspaces dir`: Print the local directory of the selected repository. Clones the repo first if it does not exist locally yet.
* `workspaces github`: Print the Github url of the selected repository.
* `workspaces refresh`: Refresh the cached list of repositories.

This tool is meant to be used in combination with other cli utilities. For example using shell aliases:
```bash
# Navigate to local dir.
alias wscd="cd \$(workspaces dir | tail -1)"

# Open the repo url in the browser.
alias wsgh="workspaces github | xargs open"

# Navigate to local dir and open gradle project in Intellij.
alias wsidea="wscd && ./gradlew openIdea"
```

## Setup

1. Create a config file with the Github organizations you want to include:
    ```yaml
    # ~/.config/workspaces/config.yml
    workspaces:
      - directory: "~/repos/google"
        organization: "https://github.com/google"
      - directory: "~/repos/apache"
        organization: "https://github.com/apache"
    ```
2. (Optional) Create a secrets config file with your [Github access token](https://github.com/settings/tokens) (using the `repo` scope). This is only required to access private repos, Github Enterprise endpoints, or to increase your API limit.
    ```yaml
    # ~/.config/workspaces/.secrets.yml
    - endpoint: "https://github.com"
      token: "my-token"
    ```
3. Refresh the list of repositories. The list is cached in `~/.config/workspaces/.repos.csv`. You need to re-run this command to pick up new repositories.
    ```bash
    $ workspaces refresh
    ```
4. (Optional) Setup your shell aliases.
    ```bash
    # Navigate to local dir of slected repo. Clone repo if it does not exist yet.
    alias wscd="cd \$(workspaces dir | tail -1)"

    # Open the repo url in the browser.
    alias wsgh="workspaces github | xargs open"
    ```