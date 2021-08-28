# Set java version

[![Step changelog](https://shields.io/github/v/release/BirmacherAkos/bitrise-step-set-java-version?include_prereleases&label=changelog&color=blueviolet)](https://github.com/BirmacherAkos/bitrise-step-set-java-version/releases)

This step helps you in setting an already installed java version on the virtual machine.

<details>
<summary>Description</summary>

This step helps you in setting an already installed java version on the virtual machine. Mind you, that this step can't install any java version, it's only for changing between the already installed versions.

If you want to install other java version you can do it by [this guide](https://devcenter.bitrise.io/infrastructure/virtual-machines/#switching-to-a-java-version-not-installed-on-our-android-stacks).

### Troubleshooting

If the step fails to set the java version, you can use these [scripts](https://devcenter.bitrise.io/infrastructure/virtual-machines/#managing-java-versions) as a temporary workaround.

### Useful links

[Managing Java versions on Bitrise](https://devcenter.bitrise.io/infrastructure/virtual-machines/#managing-java-versions)

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
