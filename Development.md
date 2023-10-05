# Development docs

## Cutting releases

Push a tag to the `main` branch of the form `vM.N.P`.

```
git checkout main
NEW_VERSION="vM.N.P" bash -c "git tag $NEW_VERSION && git push $NEW_VERSION"
```
