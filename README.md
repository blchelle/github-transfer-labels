## GitHub Label Transferring Script
This is a script written in Go that uses the GitHub API to transfer a set of labels from one GitHub repository to another. Note, you will need admin access to both the source and target repository for this to work.

> Note: This will delete all of the existing labels in the target repository

**Step 1**

Clone the repository 
```
git clone https://github.com/blchelle/github-transfer-labels.git
```

**Step 2**

First you will need a personal access token generated by GitHub which will be used in the script. See [this guide](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) to generate a personal access token.

Copy the generated token for step 3.

**Step 3**

Create a `.env` file at the root of the repository with a single line
```
GH_TOKEN=<TOKEN_FROM_STEP_2>
```

**Step 4**

Install the required go modules
```sh
go get
```

**Step 5**

Run the script. The script takes 4 command line parameters
- source_owner: the username of the source repo owner
- source_repository: the name of the source repository
- target_owner: the username of the target repo owner
- target_repository: the name of the target repository

For example: Copy all labels from `blchelle/repo-with-good-labels` to `blchelle/repo-with-no-labels`
```
go run main.go blchelle repo-with-good-labels blchelle repo-with-no-labels
```
