To release new version execute command
```cmd
./gradlew clean build publish final -Pplatform={platform} -Prelease.version={version} -PreleaseBranchPatterns=main
```
Command will build library for specific platform and create git tag.
After releasing first distribution add `-Prelease.useLastTag=true` and omit version for other 2 distributions.
```cmd
./gradlew clean build publish final -Pplatform={platform} -Prelease.useLastTag=true -PreleaseBranchPatterns=main
```