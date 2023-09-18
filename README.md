# Set Java version

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/bitrise-step-set-java-version?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/bitrise-step-set-java-version/releases)

This Step helps you activate pre-installed Java (JDK) version on the build virtual machine.

<details>
<summary>Description</summary>

This Step is not supposed to install any Java version on the fly. You can only pick from the already installed versions.
If you want to install other Java versions, check out [Using a Java version not installed on our Android stacks](https://devcenter.bitrise.io/infrastructure/virtual-machines/#using-a-java-version-not-installed-on-our-android-stacks).

### Configuring the Step
1. In the **Java version to be set globally for the build** input, select an installed Java version you wish to use during the build run.

### Troubleshooting
If the Step fails to set the Java version, you can use these [scripts](https://devcenter.bitrise.io/infrastructure/virtual-machines/#managing-java-versions) as a temporary workaround.

### Useful links
- [Managing Java versions on Bitrise](https://devcenter.bitrise.io/infrastructure/virtual-machines/#managing-java-versions)

</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `set_java_version` | Select the installed Java version you want to use during the build run. You can check [in system reports](https://stacks.bitrise.io) which Java versions are installed on each Bitrise stack.  | required | `11` |
</details>

<details>
<summary>Outputs</summary>

| Environment Variable | Description |
| --- | --- |
| `JAVA_HOME` | JAVA_HOME is an environment variable that points to the path where the JDK (Java Development Kit) is installed. |
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/bitrise-step-set-java-version/pulls) and [issues](https://github.com/bitrise-steplib/bitrise-step-set-java-version/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
