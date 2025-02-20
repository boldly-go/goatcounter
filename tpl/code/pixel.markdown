The endpoint returns a small 1×1 GIF image. You don't need to use the JavaScript
integration and can load this directly:

    <img src="{{.SiteURL}}/count?p=/test">

Or you can build your own JavaScript integration if you want. Use the
[API](/code/backend) if you want to send data from a script; `/count` is only
intended to be loaded by the visitor's browser.

The tracking pixel won’t allow recording the referrer or screen size, and may
also increase the number of bot requests (we do our best to filter this out, but
it’s hard to get all of them, since many spam scrapers and such disguise
themselves as regular browsers).

Wrap in a `<noscript>` tag to use this only for people without JavaScript:

    {{template "code" .}}
    <noscript>
        <img src="{{.SiteURL}}/count?p=/test-noscript">
    </noscript>

---

This accepts the following query parameters:

| Query | count.js   | Description                                                 |
| :---- | :--------  | :----------                                                 |
| `p`   | `path`     | Page path or event name.                                    |
| `t`   | `title`    | Page title.                                                 |
| `r`   | `referrer` | Referrer value; usually the Referer header.                 |
| `e`   | `event`    | event; as boolean (`true`, `false`, `1`, `0`, `on`, `off`). |
| `q`   | -          | Query parameters, for getting campaigns.                    |
| `s`   | -          | screen size, as `width,height,scale`.                       |
| `b`   | -          | Flag this as a "bot request"; number.                       |
| `rnd` | -          | Ignored; intended as a "cache buster".                      |

These parameters are guaranteed to be stable; any future incompatible changes
will use a new endpoint. Building your own JavaScript integration should be
safe, although you may need to modify it if new features get added.

`rnd` is useful as sometimes browsers and procies have their own opinion about
what can or can't be cached in spite of what the cache headers say.

The `b` accepts an integer constant from the [zgo.at/isbot][isbot] library and
should be >=150. See the [count.js source][cjs] how to detect this. Current
values:

- `150` – Phantom headless browser.
- `151` – Nightmare headless browser.
- `152` – Selenium headless browser.
- `153` – Generic WebDriver-based headless browser.

[isbot]: https://github.com/zgoat/isbot/blob/master/isbot.go#L46
[cjs]: https://github.com/zgoat/goatcounter/blob/master/public/count.js#L54
