# Set java version

[![Step changelog](https://shields.io/github/v/release/BirmacherAkos/bitrise-step-set-java-version?include_prereleases&label=changelog&color=blueviolet)](https://github.com/BirmacherAkos/bitrise-step-set-java-version/releases)

Set the java version used during the build

<details>
<summary>Description</summary>

You can select which installed java version to be used during the build run
</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `set_java_version` | Select the instlled java version you want to use during the build run.  You can check [here](https://github.com/bitrise-io/bitrise.io/tree/master/system_reports) which java versions are installed on each Bitrise stack.  | required | `11` |
</details>

<details>
<summary>Outputs</summary>

| Environment Variable | Description |
| --- | --- |
| `JAVA_HOME` | JAVA_HOME is an environment variable points to the file system location where the JDK or JRE was installed. |
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/BirmacherAkos/bitrise-step-set-java-version/pulls) and [issues](https://github.com/BirmacherAkos/bitrise-step-set-java-version/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
