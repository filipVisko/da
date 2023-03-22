To create a new worktree, use the `{{aka}}` subcommand followed by the name of the worktree.

The worktree folder will always be created under the root of the bare repo.

You can now navigate into the folder using the `cd` command to access the branch's files and make changes.

When creating git worktrees, a new branch with the same name is also created.
Creating a new worktree that has the same name as a branch in the remote will track that branch.

Adding branch-backed worktrees in bare repos requires you to clone them using `fidi clone`.
