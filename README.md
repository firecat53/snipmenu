# Snipmenu

A Dmenu/Rofi/Bemenu snippet manager compatible with the
[Pet][1] TOML snippet file format.

I wanted something similar to [keepmenu](https://github.com/firecat53/keepmenu)
or [bitwarden-menu](https://github.com/firecat53/keepmenu) that can type
snippets equally into a terminal or graphical window with necessarily relying on
copy/paste.

#### Differences from Pet:

- Uses graphical launcher instead of terminal (unless you use bemenu in
  ncurses mode).
- There's no `exec` option. The command or snippet is typed into the focused
  window using xdotool or ydotool and can either be edited in place or
  executed.
- Copy/paste using xsel or xclip is available if desired.
- No Github/Gitlab sync.

## Installation

`go get github.com/firecat53/snipmenu.go` OR [download binary][4]

Ensure `$GOPATH/bin` is in your `$PATH`.

## Requirements

- [Dmenu][2],
  [Rofi][3] or
  [Bemenu][4]
- xdotool
- xsel or xclip
- (optional) ydotool if using a wlroots window manager like Sway

## Features

- Add/edit/delete information snippets:
    - Command line invocations
    - Addresses, email addresses
- Compatible with [Pet][1] snippet format
- Snippets can contain:
    - Description
    - The snippet
    - Tags
    - Example command output
- Variable substitution when executing commands using `<variable>`

## License

MIT

## Usage

`snipmenu`

Menu options should be self explanatory. To use variables:

    [[snippets]]
      command = "ssh <host>:<port=22>"
      description = "SSH to hostname"
      
When selected, the user will be prompted for `host` and `port` prior to typing
the command. Note that defaults can be defined using `<param=value>`.

## Snippet file format

    [[snippets]]
      description = "ping"
      command = "ping <ip=8.8.8.8>"
      tag = ["network", "google"]
      output = ""

    [[snippets]]
      command = "echo | openssl s_client -connect example.com:443 2>/dev/null |openssl x509 -dates -noout"
      description = "Show expiration date of SSL certificate"
      tag = ["network", "ssl"]
      output = """
    notBefore=Nov  3 00:00:00 2015 GMT
    notAfter=Nov 28 12:00:00 2018 GMT"""

[1]: https://github.com/knqyf263/pet/ "Pet"
[2]: https://tools.suckless.org/dmenu/ "Dmenu"
[3]: https://davedavenport.github.io/rofi/ "Rofi"
[4]: https://github.com/Cloudef/bemenu "Bemenu"
[5]: https://github.com/firecat53/snipmenu/releases "Snipmenu Releases"
