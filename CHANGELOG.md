# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/) and this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

### Added

### Changed

### Deprecated

### Removed

### Fixed

### Security

## [1.0.3] - 2019-08-01

### Added

* Added a forward-compatibility module in case of breaking changes for seamless upgrades.

### Changed

* Changed config directory to `~/.config/gipns`.

```bash
mv ~/.config/gi ~/.config/gipns
```

### Fixed

* Fixed `curl` test `DELETE` script.

## [1.0.2] - 2019-07-31

### Added

* Adds a few selected peers to its swarm on startup for faster discovery.
* Added a few unit tests.
* Cross-compile script with lots of obscure operating systems and architectures.

### Removed

* Removed mention of the old license on startup.

### Fixed

* Fixed pipeline on CircleCI.

## [1.0.1] - 2019-07-26

### Added

* Added the following API endpoints :

| Endpoint           | Method   | Description              |
| ------------------ | -------- | ------------------------ |
| `/api/repos`       | `GET`    | List of all repositories |
| `/api/repos`       | `POST`   | Create a repository      |
| `/api/repos/{url}` | `GET`    | Get this repository      |
| `/api/repos/{url}` | `DELETE` | Delete this repository   |

* Added a list of repositories
* Added **Delete** buttons on the list of repositories

### Changed

* Changed ownership from [@NatoBoram](https://github.com/NatoBoram) to [@Permaweb](https://github.com/Permaweb)
* Changed name from [Git to IPNS](https://gitlab.com/NatoBoram/git-to-ipfs) to [Permaweb Host](https://github.com/Permaweb/Host)
* Changed license from [GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/) to [MIT](https://choosealicense.com/licenses/mit/)
* Spinners are now blue when they're actually doing something and black when they're not

### Removed

* Removed the following API endpoint :
  * `/api/add`

### Fixed

* Fixed some potential issues when refreshing old repositories that contain errors

## [1.0.0] - 2019-07-25

Initial release. It has a web interface, can `git clone` repositories, add them to IPFS Cluster, create IPFS keys, then add the repo to IPNS.
There's no authentication.
The web interface doesn't check for errors, so there's a few bugs there.

## Types of changes

* `Added` for new features.
* `Changed` for changes in existing functionality.
* `Deprecated` for soon-to-be removed features.
* `Removed` for now removed features.
* `Fixed` for any bug fixes.
* `Security` in case of vulnerabilities.
