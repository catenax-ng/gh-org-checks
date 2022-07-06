# GitHub organization checks

This repository contains a workflow, that periodically will check all repositories of the catenax-ng organization
for compliance with defaults defined our FOSS-, Security- and DevSecOps-experts.

## Run checks

The following checks are run:

- existence of a README.md file at the root of the repository
- existence of a NOTICE.md file at the root of the repository. Infos [here](https://github.com/catenax-ng/foss-example#files-from-directory-catenax-ng-repositories)
- existence of a SECURITY.md file at the root of the repository. Infos [here](https://github.com/catenax-ng/foss-example#files-from-directory-catenax-ng-repositories)
- is the trviy scan executed as workflow
- is the checkov scan executed as workflow

## Result overview

The results of the check are presented directly on the summary page of the workflow as job summary.
It will be shown as table, with a row for each repository analyzed and columns for each check.
