---
page_title: "uptimerobot_alert_contact Resource - terraform-provider-uptimerobot"
description: |-
  Set up an alert contact
---

# Resource: uptimerobot_alert_contact

Use this resource to create an alert contact

## Example Usage

```hcl
resource `uptimerobot_alert_contact` `slack` {
  friendly_name = `Slack Alert`
  type          = `slack`
  value         = `https://hooks.slack.com/services/XXXXXXX`
}
```

## Arguments Reference

- `friendly_name` - friendly name of the alert contact (for making it easier to distinguish from others).
- `type` - the type of the alert contact notified

  Possible values are the following:

  - `sms`
  - `e-mail`
  - `twitter` (or `twitter-dm`)
  - `boxcar`
  - `web-hook` (or `webhook`)
  - `pushbullet`
  - `zapier`
  - `pro-sms`
  - `pushover`
  - `slack`
  - `voice-call`
  - `splunk`
  - `pagerduty`
  - `opsgenie`
  - `telegram`
  - `ms-teams`
  - `google-chat` (or `hangouts`)
  - `discord`

- `value` - alert contact's address/phone/url

## Attributes Reference

- `id` - the ID of the alert contact.
- `status` - the status of the alert contact (`not activated`, `paused` or `active`)
