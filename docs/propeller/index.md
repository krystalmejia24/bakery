---
title: Propeller
nav_order: 3
has_children: true
has_toc: false
---

# Propeller
Propeller is the ViacomCBS live-streaming platform. Propeller can orchestrate and manage all of your cloud based resources needed for a live streaming environment. If you have a stream running out of Propeller, you can stream that channel via Bakery! 

For help on managing playback of your Propeller channels via Bakery, check out the documentation below. 

## Playback

Bakery can be used to manage your Propeller playback URLs. 

### Channels

To request a Propeller channel via Bakery:

    https://bakery.dev.cbsi.video/propeller/<org-id>/<channel-id>.m3u8

Bakery will then set the Playback URL with the following priority, depending on your channel settings:

1. Startfruit
2. DAI
3. Carrier
4. Playback URL
5. Archive
{: .lh-tight }

As long as you're channel was set to archive, Bakery will automatically proxy the archive stream when your Propeller channel has ended. 

### Clips

To request a Propeller clip via Bakery:

    https://bakery.dev.cbsi.video/propeller/<org-id>/clip/<clip-id>.m3u8

**Note** Clips are not enabled for DASH

### Outputs

To request a Propeller output via Bakery:

    https://bakery.dev.cbsi.video/propeller/<org-id>/<channel-id>/<output-id>.m3u8

**Note** Outputs are not currently enabled for DASH but a feature implementation is in the backlog. To prioritize this feature, feel free to reach out on slack!

### Explicit Origins

To have Bakery use a specific origin url for the output, add another path item specifying the origin to be used:

    http://bakery.dev.cbsi.video/propeller/<org-id>/<channel-id>/<output-id>/<origin>.m3u8

Supported values for `origin` include

|**origin**|**description**|
|----------|---------------|
| dai      | ad-supported stream from Google DAI|
| captions | auto-captioned stream from Dictate |
| cdn      | raw cdn stream url |

#### Origin Fallback

If you're not sure what origins are available for your channel, or to gracefully handle channel configuration
changes that may remove or add origins while the channel is running, you can optionally specify multiple origins
in this last path element. The `-` delimiter separates multiple origins. If the path defines more than one origin,
Bakery will loop through the origins in order and choose the first valid origin found on the channel.

For example, in order to specify a Bakery url that prefers an auto-captioned stream, but falls back to the CDN origin
if auto-captions is disabled on the channel, you'd construct the url:

        http://bakery.dev.cbsi.video/propeller/<org-id>/<channel-id>/<output-id>/captions-cdn.m3u8

## Help

For more information on Propeller, check out the documentation <a href="https://cbsinteractive.github.io/propeller/">here</a> or reach out to the Propeller team on <a href="https://cbs.slack.com/app_redirect?channel=i-vidtech-propeller" target="_blank">Slack</a> to get all set up!