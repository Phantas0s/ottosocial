![ottosocial logo](./logo.png)
![travis CLI](https://api.travis-ci.org/Phantas0s/ottosocial.svg?branch=master&style=for-the-badge) ![Hits-of-Code](https://hitsofcode.com/github/phantas0s/devdash) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) ![Go Report Card](https://goreportcard.com/badge/github.com/Phantas0s/ottosocial) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/8d788aa393a44def853e7cd90b403164)](https://www.codacy.com/manual/Phantas0s/ottosocial?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=Phantas0s/ottosocial&amp;utm_campaign=Badge_Grade)

ottosocial is a CLI which allows you to schedule and send automatically messages on social media, from a CSV file.

You can send threads by separating your tweets with a pipe `|`. As a result, you can't use the pipe as a character in a tweet.

*For now, only Twitter is supported.*

*ottosocial is in beta. Breaking changes might come.*

## Menu

* [Installation](#installation)
* [Getting Started](#getting-started)
* [Commands](#commands)
* [Configuration](#configuration)
* [CSV Format](#csv-format)
* [Contribute](#contribute)
* [Licence](#licence)

## Installation

### General

You can simply grab the [latest released binary file](https://github.com/Phantas0s/ottosocial/releases/latest) and download the version you need, depending on your OS.

### Linux script

If you use a Linux based OS, here's a simple way to download ottosocial and move it in `/usr/local/bin` to be able to use it globally.

```shell
curl -LO https://raw.githubusercontent.com/Phantas0s/ottosocial/master/install/linux.sh && \
./linux.sh && \
rm linux.sh
```

### Manual installation

You need to clone this repository and build the binary in the root directory.

## Getting Started

First, you need to authorize ottosocial to send tweets from your twitter account. To do so:

1. Go to [Twitter app management](https://developer.twitter.com/en/apps). Log in if necessary.
2. Click on "Create an app".
3. Fill the required fields.
4. You'll need to copy the following to authorize ottosocial to send tweets:
    * API key, 
    * API secret key, 
    * Access Token 
    * Access Token Secret provided, 

See below the configuration to include everything to the configuration file.

## Commands

For now, only the command `csv` is available.

I advise you to use a [terminal multiplexer](https://thevaluable.dev/tmux-boost-productivity-terminal/) or `&` to run ottosocial in the background.

## Configuration

Some configuration is required in order to run ottosocial. Here's the detail:

| Name             | Description                     | Flag                       | Type         | Required   | default   |
| ---------------- | ------------------------------- | -------------------------- | ------------ | ---------- | --------- |
| key              | Twitter API key                 | `--key`, `-k`              | string       | yes        | -         |
| secret           | Twitter API secret key          | `--secret`, `-s`           | string       | yes        | -         |
| token            | Twitter Access Token            | `--token`, `-t`            | string       | yes        | -         |
| token-secret     | Twitter Access Token Secret     | `--token-secret`, `-j`     | string       | yes        | -         |
| logpath          | File path of the log file       | `--logpath`, `-l`          | string       | no         | -         |
| file             | File path of the CSV            | `--filepath`, `-f`         | string       | yes        | -         |
| verify           | Verify each tweet               | `--verify`, `-v`           | boolean      | no         | false     |

**NOTES**: 

* If `logpath` is not filled, the default output will be the shell. I advise you to use a logpath and tailing it at all time to see every error.
* If `verify` is `true`, ottosocial will stop if one (or more) tweets are invalid.

You can either way use flags, a configuration file, or both at the same time.

### Using flags

If you use only flags, you'll need to precise every required configuration field directly when you launch ottosocial.

For example:

```
ottosocial csv --key=1234 --secret=5678 --token=910-11 --token-secret=12KK --logpath=/tmp/ottosocial-logs -v -f ./test.csv
```

### Configuration File

This is the simplest way to use ottosocial. You'll need to create the file `~/ottosocial.yml` (directly in your home directory) and add your configuration there. For example:

```
---
key: 1234 # API key
secret: 5678 # API secret key
token: 910-11 # Access token
token-secret: 12KK # Access token secret
logpath: /tmp/ottosocial-logs
file: ~/my-tweets.csv
```

Then, you simply need to launch ottosocial.

## CSV Format

ottosocial can send tweets from a CSV. It needs to have two columns: `date` (when the tweet will be sent) and `message`.

For example:

```csv
2020-01-23 07:57,This is a test
2020-01-23 08:00,"This is a test, with quotes"
```

You need to respect the date format. If the date is in the past, the tweet will never be sent.

**NOTES**:

* If you modify the CSV, you need to relaunch ottosocial. I might implement hot reloading later, if somebody is interested in that.

## Contribute

Thanks a lot if you want to contribute to ottosocial!

I think the ["talk, then code"](https://dave.cheney.net/tag/contributing) practice is pretty good to avoid misunderstandings and hours of work for nothing.

Therefore:

"Every new feature or bug fix should be discussed with the maintainer(s) of the project before work commences. It’s fine to experiment privately, but do not send a change without discussing it first."

## Acknowledgement

The logo uses [Maria Letta mega doodle pack](https://github.com/MariaLetta/mega-doodles-pack). Thanks for these great illustrations.

## Licence

[Apache Licence 2.0](https://choosealicense.com/licenses/apache-2.0/)
