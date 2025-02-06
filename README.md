# Meshery Extension: Kanvas Snapshot Helm Plugin

[Meshery Extensions](https://meshery.io/extension) are plugins or add-ons that enhance the functionality of the Meshery platform beyond its core capabilities. Meshery supports different types of extensions ([docs](https://docs.meshery.io/extensions/)):

- [Adapters](https://docs.meshery.io/concepts/architecture/adapters): Adapters allow Meshery to interface with the different cloud native infrastructure.
- [Load Generators](https://docs.meshery.io/extensibility/load-generators): for performance characterization and benchmarking
- [Integrations](https://docs.meshery.io/extensibility/integrations): model-based support for a broad variety of design and orchestration of cloud and cloud native platforms, tools, and technologies.
- [Providers](https://docs.meshery.io/extensibility/providers): for connecting to different cloud providers and infrastructure platforms
- [UI Plugins](https://docs.meshery.io/extensibility/ui): Meshery UI has a number of extension points that allow users to customize their experience with third-party plugins.
<!-- TODO: NEW PAGE NEEDED HERE: - CLI Plugins:  -->


**Table of Contents**

- [Meshery Extension: Kanvas Snapshot Helm Plugin](#meshery-extension-kanvas-snapshot-helm-plugin)
  - [Extension Overview](#extension-overview)
    - [Features](#features)
  - [Installation and Use](#installation-and-use)
    - [Prerequisites](#prerequisites)
    - [Usage](#usage)
  - [Contributing](#contributing)
  - [Join the Meshery community!](#join-the-meshery-community)
  - [Contributing](#contributing-1)
    - [Show Your Support](#show-your-support)
    - [License](#license)

## Extension Overview

The **Kanvas Snapshot Helm Plugin** allows users to generate a visual snapshot of their Helm charts directly from the command line. It simplifies the process of creating Meshery Snapshots, providing a visual representation of packaged Helm charts. This plugin integrates with Meshery Cloud and GitHub Actions to automate the workflow of snapshot creation, which is especially useful for Helm users who need to quickly visualize their chart configurations.

Helm charts can be complex, especially when custom configurations are applied via `values.yaml` files. This Meshery extension bridges the gap between Helm chart configurations and their visual representation by converting Helm charts into **Kanvas Snapshots**. These snapshots can be received either via email or as a URL displayed directly in the terminal.

### Features

1. **Snapshot Generation:** Create visual snapshots of Helm charts, complete with associated resources.
2. **Synchronous/Asynchronous Delivery:** Choose between receiving snapshots via email or directly in the terminal.
3. **Seamless Integration:** Leverages Meshery Cloud and GitHub Actions to handle snapshot rendering.
4. **Support for Packaged Charts:** Works with both packaged `.tar.gz` charts and unpackaged Helm charts.

## Installation and Use

To install the Meshery Snapshot Helm Plugin, use the following steps:

### Prerequisites

- `helm` must be [installed]( helm plugin install https://github.com/meshery/helm-kanvas-snapshot) on your system.
- (Optional) A free [Layer5 Cloud](https://cloud.layer5.io) user account.

**Plugin Installation**

1. Run the following command to install the Helm Kanvas Snapshot plugin:

   ```bash
   helm plugin install https://github.com/meshery/helm-kanvas-snapshot
   ```

3. Verify the installation by running:

   ```bash
   helm plugin list
   ```

   You should see the Kanvas Snapshot listed as `helm-kanvas-snapshot`.

Update to the latest version:

```bash
helm plugin update helm-kanvas-snapshot
```
### Usage

Once the plugin is installed, you can generate a snapshot using either a packaged or unpackaged Helm chart.

```bash
helm helm-kanvas-snapshot -f <chart-URI> [--name <snapshot-name>] [-e <email>]
```

- **`-f`**, **`--file`**: (required) path or URL to the Helm chart (required).
- **`--name`**: (optional) name for the snapshot. If not provided, a name will be auto-generated based on the chart name.
-  **`-e`**, **`--email`**: (optional) email address to notify when snapshot is ready. If not provided, a link to the snapshot will be displayed in the terminal.

**Example**

To generate a snapshot for a Helm chart located at `https://meshery.io/charts/meshery-v0.8.12.tgz`, you can use:

```bash
helm helm-kanvas-snapshot -f https://meshery.io/charts/meshery-v0.8.12.tgz --name meshery-chart
```

## Contributing

Please do! Thank you for your help in improving this Meshery extension! :balloon:

Start by forking the repository. After making your changes, submit a pull request. For more information, see the [Contributing Guide](CONTRIBUTING.md).


**1. Fork the Repository**

To get started, you'll first need to clone the Meshery Snapshot Helm Plugin repository from GitHub. Run the following command in your terminal:

```bash
git clone https://github.com/meshery/helm-kanvas-snapshot.git
```

**2. Navigate to the Plugin Directory**

Once the repository is cloned, navigate to the `helm-kanvas-snapshot` directory.

```bash
cd helm-kanvas-snapshot
```

**4. Build the binary locally**

```bash
make local.build
```

**5. Test the binary**

```bash
./helm-kanvas-snapshot -f https://meshery.io/charts/v0.8.0-meshery.tar.gz
```

This command will trigger the snapshot generation process. If everything is set up correctly, you should see a visual snapshot URL or receive the snapshot via email, depending on the options you specified.

<div>&nbsp;</div>

## Join the Meshery community!

<a name="contributing"></a><a name="community"></a>
Our projects are community-built and welcome collaboration. 👍 Be sure to see the <a href="https://docs.meshery.io/project/contributing#not-sure-where-to-start">Contributor Welcome Guide</a> and <a href="https://meshery.io/community#handbook">Community Handbook</a> for a tour of resources available to you and the <a href="https://layer5.io/community/handbook/repository-overview">Repository Overview</a> for a cursory description of repository by technology and programming language. Jump into community <a href="https://slack.meshery.io">Slack</a> or <a href="https://meshery.io/community#discussion-forums">discussion forum</a> to participate.

<p style="clear:both;">
<a href ="https://meshery.io/community"><img alt="MeshMates" src=".github/assets/images/readme/layer5-community-sign.png" style="margin-right:36px; margin-bottom:7px;" width="140px" align="left" /></a>
<h3>Find your MeshMate</h3>

<p>MeshMates are experienced community members, who will help you learn your way around, discover live projects, and expand your community network. Connect with a Meshmate today!</p>

Find out more on the <a href="https://meshery.io/community#meshmates">Meshery community</a>. <br />

</p>
<br /><br />
<div style="display: flex; justify-content: center; align-items:center;">
<div>
<a href="https://meshery.io/community"><img alt="Meshery Cloud Native Community" src="https://docs.meshery.io/assets/img/readme/community.png" width="140px" style="margin-right:36px; margin-bottom:7px;" width="140px" align="left"/></a>
</div>
<div style="width:60%; padding-left: 16px; padding-right: 16px">
<p>
✔️ <em><strong>Join</strong></em> any or all of the weekly meetings on <a href="https://meshery.io/calendar">community calendar</a>.<br />
✔️ <em><strong>Watch</strong></em> community <a href="https://www.youtube.com/@mesheryio?sub_confirmation=1">meeting recordings</a>.<br />
✔️ <em><strong>Fill-in</strong></em> a <a href="https://layer5.io/newcomers">community member form</a> to gain access to community resources.
<br />
✔️ <em><strong>Discuss</strong></em> in the <a href="https://meshery.io/community#discussion-forums">Community Forum</a>.<br />
✔️ <em><strong>Explore more</strong></em> in the <a href="https://meshery.io/community#handbook">Community Handbook</a>.<br />
</p>
</div><br /><br />
<div>
<a href="https://slack.meshery.io">
<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/readme/slack.svg"  width="110px" />
  <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/readme/slack.svg" width="110px" />
  <img alt="Shows an illustrated light mode meshery logo in light color mode and a dark mode meshery logo dark color mode." src="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/readme/slack.svg" width="110px" align="left" />
</picture>
</a>
</div>
</div>
<br /><br />
<p align="left">
&nbsp;&nbsp;&nbsp;&nbsp; <i>Not sure where to start?</i> Grab an open issue with the <a href="https://github.com/issues?q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Alayer5io+org%3Ameshery+org%3Aservice-mesh-performance+org%3Aservice-mesh-patterns+label%3A%22help+wanted%22+">help-wanted label</a>.
</p>
<br /><br />

<div>&nbsp;</div>

## Contributing

Please do! We're a warm and welcoming community of open source contributors. Please join. All types of contributions are welcome. Be sure to read the [Contributor Guides](https://docs.meshery.io/project/contributing) for a tour of resources available to you and how to get started.

<!-- <a href="https://youtu.be/MXQV-i-Hkf8"><img alt="Deploying Linkerd with Meshery" src="https://docs.meshery.io/assets/img/readme/deploying-linkerd-with-meshery.png" width="100%" align="center" /></a> -->

<div>&nbsp;</div>

### Show Your Support

<p align="center">
  <i>If you like Helm kanvas Snapshot, please <a href="../../stargazers">★</a> star this repository to show your support! 🤩</i>
</p>

### License

This repository and site are available as open-source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).
