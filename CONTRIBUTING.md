# <a name="contributing">Contributing Overview</a>

Please do! Thank you for your help in improving Meshery! :balloon:

---

<details>

  <summary><h3>Find the complete set of contributor guides at https://docs.meshery.io/project/contributing</h3></summary>

All contributors are welcome. Not sure where to start? Please see the [newcomers welcome guide](https://layer5.io/community/newcomers) for how, where, and why to contribute. This project is community-built and welcomes collaboration. Contributors are expected to adhere to our [Code of Conduct](CODE_OF_CONDUCT.md).

All set to contribute? Grab an open issue with the [help-wanted label](../../labels/help%20wanted) and jump in. Join our [Slack channel](https://slack.meshery.io) and engage in conversation. Create a [new issue](/../../issues/new/choose) if needed. All [pull requests](/../../pulls) should ideally reference an open [issue](/../../issues). Include keywords in your pull request descriptions, as well as commit messages, to [automatically close related issues in GitHub](https://help.github.com/en/github/managing-your-work-on-github/closing-issues-using-keywords).

**Sections**

- <a name="contributing">General Contribution Flow</a>
  - <a href="#commit-signing">Developer Certificate of Origin</a>
- Meshery Contribution Flow
  - <a href="#contributing-docs">Meshery Documentation</a>
  - <a href="#contributing-meshery">Meshery Backend</a>
    - <a href="#adapter">Writing a Meshery Adapter</a>
  - <a href="#contributing-ui">Meshery UI</a>
    Relevant coding style guidelines are the Go Code Review Comments and the Formatting and style section of Peter Bourgon's Go: Best Practices for Production Environments.
  - <a href="#contributing-mesheryctl">Mesheryctl Documentation</a>
    - <a href="https://docs.google.com/spreadsheets/d/1q63sIGAuCnIeDs8PeM-0BAkNj8BBgPUXhLbe1Y-318o/edit#gid=0">Command Reference and Tracker</a>

# <a name="contributing">General Contribution Flow</a>

To contribute to Meshery, please follow the fork-and-pull request workflow described [here](docs/CONTRIBUTING-gitflow.md).

## Issues & Pull Requests

### Creating an Issue

Before **creating** an Issue i.e for `features`/`bugs`/`improvements` please follow these steps:


1. Search existing Issues before creating a new Issue (look to see if the Issue has already been created).
1. If it doesn't exist create a new Issue giving as much context as possible (please take note and select the correct Issue type, for example `bug`, `documentation` or `feature`.
1. If you wish to work on the Issue once it has been triaged, please include this in your Issue description.

### Working on an Issue

Before working on an existing Issue please follow these steps:

1. Comment asking for the Issue to be assigned to you.
1. To best position yourself for Issues assignment, we recommend that you:
    1. Confirm that you have read the CONTRIBUTING.md.
    1. Have a functional development environment (have built and are able to run the project).
    1. Convey your intended approach to solving the issue.
    1. Put each of these items in writing in one or more comments.
1. After the Issue is assigned to you, you can start working on it.
1. In general, **only** start working on this Issue (and open a Pull Request) when it has been assigned to you. Doing so will prevent confusion, duplicate work (some of which may go unaccepted given its duplicity), incidental stepping on toes, and the headache involved for maintainers and contributors alike as Issue assignments collide and heads bump together.
1. Reference the Issue in your Pull Request (for example `This PR fixes #123`). so that the corresponding Issue is automatically closed upon merge of your Pull Request.

> Notes:
>
> - Check the `Assignees` box at the top of the page to see if the Issue has been assigned to someone else before requesting this be assigned to you. If the issue has a current Assignee, but appears to be inactive, politely inquire with the current Assignee as to whether they are still working on a solution and/or if you might collaborate with them.
> - Only request to be assigned an Issue if you know how to work on it.
> - If an Issue is unclear, ask questions to get more clarity before asking to have the Issue assigned to you; avoid asking "what do I do next? how do I fix this?" (see the item above this line)
> - An Issue can be assigned to multiple people, if you all agree to collaborate on the Issue (the Pull Request can contain commits from different collaborators)
> - Any Issues that has no activity after 2 weeks will be unassigned and re-assigned to someone else.

## Reviewing Pull Requests

We welcome everyone to review Pull Requests. It is a great way to learn, network, and support each other.

### DOs

- Use inline comments to explain your suggestions
- Use inline suggestions to propose changes
- Exercise patience and empathy while offering critiques of the works of others.

### DON'Ts

- Do not repeat feedback, this creates more noise than value (check the existing conversation), use GitHub reactions if you agree/disagree with a comment
- Do not blindly approve Pull Requests to improve your GitHub contributors graph

## <a name="commit-signing">Signing-off on Commits (Developer Certificate of Origin)</a>

To contribute to this project, you must agree to the Developer Certificate of
Origin (DCO) for each commit you make. The DCO is a simple statement that you,
as a contributor, have the legal right to make the contribution.

See the [DCO](https://developercertificate.org) file for the full text of what you must agree to
and how it works [here](https://github.com/probot/dco#how-it-works).
To signify that you agree to the DCO for contributions, you simply add a line to each of your
git commit messages:

```
Signed-off-by: Jane Smith <jane.smith@example.com>
```

In most cases, you can add this signoff to your commit automatically with the
`-s` or `--signoff` flag to `git commit`. You must use your real name and a reachable email
address (sorry, no pseudonyms or anonymous contributions). An example of signing off on a commit:

```
$ commit -s -m “my commit message w/signoff”
```

To ensure all your commits are signed, you may choose to add this alias to your global `.gitconfig`:

_~/.gitconfig_

```
[alias]
  amend = commit -s --amend
  cm = commit -s -m
  commit = commit -s
```

Or you may configure your IDE, for example, Visual Studio Code to automatically sign-off commits for you:

<a href="https://user-images.githubusercontent.com/7570704/64490167-98906400-d25a-11e9-8b8a-5f465b854d49.png" ><img src="https://user-images.githubusercontent.com/7570704/64490167-98906400-d25a-11e9-8b8a-5f465b854d49.png" width="50%"><a>

## <a name="contributing-docs">Documentation Contribution Flow</a>

Please contribute! Meshery documentation uses GitHub Pages to host the docs site. Learn more about [Meshery's documentation framework](https://docs.google.com/document/d/17guuaxb0xsfutBCzyj2CT6OZiFnMu9w4PzoILXhRXSo/edit?usp=sharing). The process of contributing follows this flow:

1. Create a fork, if you have not already, by following the steps described [here](docs/CONTRIBUTING-gitflow.md)
1. In the local copy of your fork, navigate to the docs folder.
   `cd docs`
1. Create and checkout a new branch to make changes within
   `git checkout -b <my-changes>`
1. Edit/add documentation.
   `vi <specific page>.md`
1. Add redirect link on the old page (only when a new page is created that replaces the old page)
1. Run site locally to preview changes.
   `make docs`

- **Note:** _From the Makefile, this command is actually running `$ bundle exec jekyll serve --drafts --livereload --config _config_dev.yml`. If this command causes errors try running the server without Livereload with this command: `$ bundle exec jekyll serve --drafts --config _config_dev.yml`. Just keep in mind you will have to manually restart the server to reflect any changes made without Livereload. There are two Jekyll configuration, `jekyll serve` for developing locally and `jekyll build` when you need to generate the site artifacts for production._

1. Commit, [sign-off](#commit-signing), and push changes to your remote branch.
   `git push origin <my-changes>`
1. Open a pull request (in your web browser) against our main repo: https://github.com/meshery/meshery.

_Alternatively, LiveReload is available as an option during development: with jekyll serve --livereload no more manual page refresh. 

`bundle exec jekyll serve --drafts --livereload --incremental --config _config_dev.yml`


## <a name="contributing-meshery">Meshery Contribution Flow</a>

Meshery is written in `Go` (Golang) and leverages Go Modules. UI is built on React and Next.js. To make building and packaging easier a `Makefile` is included in the main repository folder.

Relevant coding style guidelines are the [Go Code Review Comments](https://code.google.com/p/go-wiki/wiki/CodeReviewComments) and the _Formatting and style_ section of Peter Bourgon's [Go: Best
Practices for Production Environments](https://peter.bourgon.org/go-in-production/#formatting-and-style).

**Please note**: All `make` commands should be run in a terminal from within the Meshery's main folder.

### Prerequisites for building Meshery in your development environment:

1. Go version 1.21.1 must be installed if you want to build and/or make changes to the existing code. The binary `go1.21.1` should be available in your path. If you don't want to disturb your existing version of Go, then follow these [instructions](https://go.dev/doc/manage-install#:~:text=and%20run%20them.-,Installing%20multiple%20Go%20versions,-You%20can%20install) to keep multiple versions of Go in your system.
2. `GOPATH` environment variable should be configured appropriately
3. `npm` and `node` should be installed on your machine, preferably the latest versions.
4. Fork this repository (`git clone https://github.com/meshery/meshery.git`), and clone your forked version of Meshery to your development environment, preferably outside `GOPATH`.
5. `golangci-lint` should be installed if you want to test Go code, for MacOS and linux users.

#### Build and Run Meshery Server

Before you can access the Meshery UI, you need to install the UI dependencies,

```sh
make ui-setup
```

and then build and export the UI

```sh
make ui-build
```

To build & run Meshery Server, run the following command:

```sh
make server
```

Any time changes are made to the Go code, you will have to stop the server and run the above command again.
Once the Meshery server is up and running, you should be able to access Meshery on your `localhost` on port `9081` at `http://localhost:9081`.

**Please note**: If you see "Meshery Development Incompatible" while trying to sign into Meshery Server, then follow these steps:

<img src="./docs/assets/img/meshery-development-incompatible-error.png" width="50%">

Potential Solution:

-  Go to your meshery folder in your local-system where you’ve cloned it.
Execute:

- `git remote add upstream https://github.com/meshery/meshery`
- `git fetch upstream`
- Restart the meshery server
- Additionally, before restarting the server, if you like to pull the latest changes, you can do: `git pull upstream master`
### UI Development Server

If you want to work on the UI, it will be a good idea to use the included UI development server. You can run the UI development server by running the following command:

```
make ui
```

Once you have the server configured, and running successfully on the default port `http://localhost:9081`, you may proceed to access the Meshery UI at `http://localhost:3000`.
Any UI changes made now will automatically be recompiled and served in the browser.

To access the [Meshery UI Development Server](#ui-development-server) on port `3000`, you will need to select your **Cloud Provider** by navigating to `localhost:9081` after running the Meshery server.

**Please note**: When running `make server` on the macOS platform, some may face errors with the crypto module in Go. This is caused due to invalid C headers in Clang installed with XCode platform tools. Replacing Clang with gcc by adding `export CC=gcc` to .bashrc / .zshrc should fix the issue. More information on the issue can be found [here](https://github.com/golang/go/issues/30072)

**Please Note** : Little minor things where you can face some issues in the windows platform -

1. Meshery requires gcc at the `make server` step, **x64 windows** architecture can face issues while finding the best **GCC compiler**, You can install [tdm64-GCC](https://jmeubank.github.io/tdm-gcc/) which worked smoothly but many compilers other than that can cause issues, you also have to set an environment variable for this step.

2. Installing `make` in windows requires you to install [choco](https://chocolatey.org/install) first, which makes it easier to install `make` then, It requires security access which can only be done in admin mode.

#### Tests

Users can now test their code changes on their local machine against the CI checks implemented through golang-ci lint.

To test code changes on your local machine, run the following command:

```
make golangci-run
```

#### Building Docker image

To build a Docker image of Meshery, please ensure you have `Docker` installed to be able to build the image. Now, run the following command to build the Docker image:

```sh
make docker
```

#### <a name="adapter">Writing a Meshery Adapter</a>

Meshery uses adapters to provision and interact with different service meshes. Follow these instructions to create a new adapter or modify an existing adapter.

1. Get the proto buf spec file from Meshery repo:
   `wget https://raw.githubusercontent.com/meshery/meshery/master/server/meshes/meshops.proto`
1. Generate code
   1. Using Go as an example, do the following:
      - install the protocol buffer compiler: https://grpc.io/docs/protoc-installation/
      - add GOPATH to PATH: `export PATH=$PATH:$(go env GOPATH)/bin`
      - install the protocol compiler plugins for go:
               `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
               `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
      - create a directory _meshes_
      - Generate Go code:
         	`protoc --proto_path=. --go_out=meshes --go_opt=paths=source_relative --go-grpc_out=meshes --go-grpc_opt=paths=source_relative meshops.proto`

   1. For other languages, please refer to gRPC.io for language-specific guides.
1. Implement the service methods and expose the gRPC server on a port of your choice (e.g. 10000).

_Tip:_ The [Meshery adapter for Istio](https://github.com/meshery/meshery-istio) is a good reference adapter to use as an example of a Meshery adapter written in Go.

#### <a name="meshery-istio">Running Meshery Adapter (Meshery-Istio)</a>

**Meshery-Istio** is a pre-written example of Meshery Adapter written in Go. Follow these instructions to run meshery-istio to avoid errors related to Meshery Adapters

1. Fork [Meshery-Istio](https://github.com/meshery/meshery-istio)
2. Clone your fork locally
3. Run this command from the root directory of **meshery-istio**
   ```sh
   make run
   ```
4. Try connecting to port 10000 as Meshery Adapter URL

## <a name="contributing-ui">UI Contribution Flow</a>

helm-kanvas-snapshot is written in `Go` (Golang) and leverages Go Modules To make building and packaging easier a `Makefile` is included in the main repository folder.

```
make local.build
```

All contributors are invited to review pull requests. See this short video on [how to review a pull request](https://www.youtube.com/watch?v=isLfo7jfE6g&feature=youtu.be).

# New to Git?

Resources: https://lab.github.com and https://try.github.com/

### License

This repository and site are available as open-source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).

</details>

---
