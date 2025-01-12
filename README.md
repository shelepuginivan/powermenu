<div align="center">

# Powermenu

A configurable power menu for Wayland compositors based on GTK3 and Gtk Layer Shell.

![Example powermenu window](https://github.com/user-attachments/assets/d88d3710-135f-4c62-bc4b-1e5d3e529b0b)

</div>


## What is powermenu?

**Powermenu** is a simple yet configurable power management menu. It is built with [`gotk3`](https://github.com/gotk3/gotk3) and [`gotk3-layershell`](https://github.com/dlasky/gotk3-layershell) &mdash; awesome Go bindings for GTK3 and gtk-layer-shell respectively.

Although it is designed mostly for users of standalone Wayland compositors such as [Sway](https://github.com/swaywm/sway), [Hyprland](https://github.com/hyprwm/Hyprland), and [Niri](https://github.com/YaLTeR/niri), powermenu can also be used with all compositors and desktop environments that implement the `zwlr_layer_shell_v1` protocol.


## Installation

```shell
git clone https://github.com/shelepuginivan/powermenu.git
cd powermenu
sudo make install-local
```

This will install powermenu in `/usr/local/bin`, which is the recommended location for user-installed binaries. If you prefer to install it in `/usr/bin`, run the following commands instead:

```shell
git clone https://github.com/shelepuginivan/powermenu.git
cd powermenu
sudo make install
```

## Usage

powermenu is fully keyboard controlled. The following table describes supported key bindings.

|                              Action                               |                                    Key bindings                                    |
| :---------------------------------------------------------------: | :--------------------------------------------------------------------------------: |
|                         Run active option                         |                                  <kbd>Enter</kbd>                                  |
|                    Quit without doing anything                    |                                   <kbd>Esc</kbd>                                   |
| Select next option (right or bottom depending on layout settings) |           <kbd>Right</kbd>, <kbd>l</kbd>, <kbd>Down</kbd>, <kbd>j</kbd>            |
| Select previous option (left or top depending on layout settings) |             <kbd>Left</kbd>, <kbd>h</kbd>, <kbd>Up</kbd>, <kbd>k</kbd>             |
| Select option by index (depends on the number of enabled options  | <kbd>1</kbd>, <kbd>2</kbd>, <kbd>3</kbd>, <kbd>4</kbd>, <kbd>5</kbd>, <kbd>6</kbd> |


## Configuration

The configuration file of powermenu is located at `$XDG_CONFIG_HOME/powermenu/config.yaml`, which defaults to `~/.config/powermenu/config.yaml` on most systems.

The [default configuration file](https://github.com/shelepuginivan/powermenu/blob/main/examples/config.yaml) can be found in `examples/`.


## Styling

The style file of powermenu is located at `$XDG_CONFIG_HOME/powermenu/style.css`, which defaults to `~/.config/powermenu/style.css` on most systems.

While some CSS features are supported, some aren't. See [Overview of CSS in GTK](https://docs.gtk.org/gtk3/css-overview.html) for more information.
