# Terraform UptimeRobot Provider

[![All Contributors](https://img.shields.io/badge/all_contributors-14-orange.svg?style=flat-square)](#contributors-)
[![Build Status](https://github.com/louy/terraform-provider-uptimerobot/workflows/test/badge.svg)](https://github.com/louy/terraform-provider-uptimerobot/actions?query=branch%3Amaster+workflow%3Atest)

## Migrating from the old provider

If you're using the provider this was forked from, you can run the following to use the
new provider:

```bash
terraform state replace-provider louy/uptimerobot vexxhost/uptimerobot
```

## Getting started

To install this provider, check out the installation instructions on [Terraform's registry page](https://registry.terraform.io/providers/vexxhost/uptimerobot/latest).

```tf
terraform {
  required_providers {
    uptimerobot = {
      source = "vexxhost/uptimerobot"
      version = "0.8.2"
    }
  }
}

provider "uptimerobot" {
  api_key = "[YOUR MAIN API KEY]" # or pass via environment variable UPTIMEROBOT_API_KEY
}

data "uptimerobot_account" "account" {}

data "uptimerobot_alert_contact" "default_alert_contact" {
  friendly_name = data.uptimerobot_account.account.email
}

resource "uptimerobot_alert_contact" "slack" {
  friendly_name = "Slack Alert"
  type          = "slack"
  value         = "https://hooks.slack.com/services/XXXXXXX"
}

resource "uptimerobot_monitor" "main" {
  friendly_name = "My Monitor"
  type          = "http"
  url           = "http://example.com"
  # pro allows 60 seconds
  interval      = 300

  alert_contact {
    id = uptimerobot_alert_contact.slack.id
    # threshold  = 0  # pro only
    # recurrence = 0  # pro only
  }

  alert_contact {
    id = data.uptimerobot_alert_contact.default_alert_contact.id
  }
}

resource "uptimerobot_monitor" "custom_port" {
  url           = "doe.john.me"
  type          = "port"
  sub_type      = "custom"
  port          = 5678
  friendly_name = "Custom port"
}

resource "uptimerobot_status_page" "main" {
  friendly_name  = "My Status Page"
  custom_domain  = "status.example.com"
  password       = "WeAreAwsome"
  sort           = "down-up-paused"
  monitors       = [uptimerobot_monitor.main.id]
}

resource "aws_route53_record" {
  zone_id = "[MY ZONE ID]"
  type    = "CNAME"
  records = [uptimerobot_status_page.main.dns_address]
}

```

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://louy.alakkad.me"><img src="https://avatars3.githubusercontent.com/u/349850?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Louay Alakkad</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=louy" title="Code">💻</a> <a href="#maintenance-louy" title="Maintenance">🚧</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=louy" title="Tests">⚠️</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=louy" title="Documentation">📖</a> <a href="#tool-louy" title="Tools">🔧</a></td>
    <td align="center"><a href="https://nhamlh.space"><img src="https://avatars3.githubusercontent.com/u/11173217?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Nham Le</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=nhamlh" title="Code">💻</a></td>
    <td align="center"><a href="http://blog.smartcube.co.za"><img src="https://avatars0.githubusercontent.com/u/237513?v=4?s=100" width="100px;" alt=""/><br /><sub><b>David Rubin</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=drubin" title="Code">💻</a> <a href="#maintenance-drubin" title="Maintenance">🚧</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=drubin" title="Tests">⚠️</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=drubin" title="Documentation">📖</a> <a href="#ideas-drubin" title="Ideas, Planning, & Feedback">🤔</a> <a href="#question-drubin" title="Answering Questions">💬</a></td>
    <td align="center"><a href="https://ijohan.nl"><img src="https://avatars2.githubusercontent.com/u/365827?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Johan Bloemberg</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=aequitas" title="Code">💻</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=aequitas" title="Tests">⚠️</a> <a href="#ideas-aequitas" title="Ideas, Planning, & Feedback">🤔</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=aequitas" title="Documentation">📖</a></td>
    <td align="center"><a href="https://twitch.tv/sebbity"><img src="https://avatars1.githubusercontent.com/u/564860?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Seb Patane</b></sub></a><br /><a href="#platform-Novex" title="Packaging/porting to new platform">📦</a></td>
    <td align="center"><a href="https://github.com/leeif"><img src="https://avatars1.githubusercontent.com/u/15794005?v=4?s=100" width="100px;" alt=""/><br /><sub><b>YIFAN LI</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=leeif" title="Code">💻</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=leeif" title="Tests">⚠️</a></td>
    <td align="center"><a href="https://nicolas.lamirault.xyz"><img src="https://avatars0.githubusercontent.com/u/29233?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Nicolas Lamirault</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=nlamirault" title="Documentation">📖</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/issues?q=author%3Anlamirault" title="Bug reports">🐛</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=nlamirault" title="Code">💻</a></td>
  </tr>
  <tr>
    <td align="center"><a href="http://ochrona.jawne.info.pl"><img src="https://avatars1.githubusercontent.com/u/3618479?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Adam Dobrawy</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=ad-m" title="Documentation">📖</a></td>
    <td align="center"><a href="http://fewbar.com/"><img src="https://avatars2.githubusercontent.com/u/470880?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Clint Byrum</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/issues?q=author%3ASpamapS" title="Bug reports">🐛</a></td>
    <td align="center"><a href="https://carrondo.net"><img src="https://avatars1.githubusercontent.com/u/2323546?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Tiago Carrondo</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/issues?q=author%3Atcarrondo" title="Bug reports">🐛</a></td>
    <td align="center"><a href="https://github.com/bpjbauch"><img src="https://avatars1.githubusercontent.com/u/13983135?v=4?s=100" width="100px;" alt=""/><br /><sub><b>JB</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/issues?q=author%3Abpjbauch" title="Bug reports">🐛</a></td>
    <td align="center"><a href="https://caarlos0.dev"><img src="https://avatars3.githubusercontent.com/u/245435?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Carlos Alexandro Becker</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=caarlos0" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/jjungnickel"><img src="https://avatars3.githubusercontent.com/u/160383?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Jan Jungnickel</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=jjungnickel" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/bd0zer"><img src="https://avatars2.githubusercontent.com/u/32301353?v=4?s=100" width="100px;" alt=""/><br /><sub><b>bd0zer</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/issues?q=author%3Abd0zer" title="Bug reports">🐛</a></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/randrusiak"><img src="https://avatars2.githubusercontent.com/u/29524175?v=4?s=100" width="100px;" alt=""/><br /><sub><b>randrusiak</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=randrusiak" title="Code">💻</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/issues?q=author%3Arandrusiak" title="Bug reports">🐛</a></td>
    <td align="center"><a href="https://keybase.io/jasonrogena"><img src="https://avatars3.githubusercontent.com/u/2384176?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Jason Rogena</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=jasonrogena" title="Code">💻</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=jasonrogena" title="Tests">⚠️</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
