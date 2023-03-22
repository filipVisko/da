To pull branch changes into its specific worktree, use `{{aka}}` followed by the name of the branch.

This subcommand can be used from any directory within the bare repo.
For example, this can be useful to pull changes into your `main` worktree while in the wd of a `dev` worktree.

Will call either `rebase` or `merge` to reconcile diverging branches depending on your configuration (does not accept flags).
