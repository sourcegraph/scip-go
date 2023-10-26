# Development docs

## Cutting releases

1. Land a PR with the following changes:

   - A ChangeLog entry with `## vM.N.P`
   - Updated version in `internal/index/version.txt`

2. On the `main` branch, run the following:

    ```bash
    NEW_VERSION="M.N.P" ./dev/publish-release.sh
    ```
