# kit

_Toolkit for Learnpoint developers._

kit is a single binary with a few small tools for building and previewing static sites:

- **serve** — a static web server for local development, with live reload.
- **build** — a small static site generator: `src/` in, `docs/` out.
- **dev** — build, serve and rebuild on every change, in one command.
- **copy** — copy the files of a github repo into a folder, like a template.

Run `kit help` for all commands, and `kit help <command>` for the details of one.


## Install

windows (powershell):

    irm https://raw.githubusercontent.com/learnpoint/kit/main/scripts/install.ps1 | iex

macos and linux:

    curl -fsSL https://raw.githubusercontent.com/learnpoint/kit/main/scripts/install.sh | sh

Check that it works:

    kit version

Upgrade to the latest release:

    kit upgrade


## Serve a folder

From any folder with html files, run:

    kit serve

Open the address kit prints. When you edit a file in the folder, the browser reloads by itself. Stop the server with `Ctrl+C`.

- A request to a folder serves its `index.html`.
- If the root folder has a `404.html`, it's used for missing pages, like on github pages.
- serve only listens on your own machine. To open your site on a phone, see [No share command](#no-share-command).

Run `kit help serve` for the port and the other options.


## Building sites

### Quick start

Create a folder with a `src/` folder and one page in it:

    mysite
     └── src
          └── index.md

`src/index.md`:

    # Hello

    My first kit site.

From `mysite`, run:

    kit dev

Open the address kit prints. Edit `index.md`, save, and the browser reloads with the change. The generated site is in `docs/`.

### src/ and docs/

You edit files in `src/`. `kit build` generates the site into `docs/`, which is the folder github pages serves.

    mysite
     ├── docs                  generated, don't edit
     └── src
          ├── _layouts         optional
          │    └── default.html
          ├── _includes        optional
          │    └── footer.html
          ├── about.md
          └── index.html

- Pages are `.html`, `.md` and `.svg` files. Markdown pages become `.html`.
- Every other file (css, images, …) is copied as it is.
- `_layouts` and `_includes` are used during the build and never copied.
- `docs/` is fully owned by kit: a file in `docs/` without a source in `src/` is deleted. You can always delete `docs/` and rebuild.
- A build is all or nothing. If anything is wrong, kit prints the problems and leaves `docs/` untouched.

### Layouts

A layout wraps a page. It must contain `{{ content }}`, where the page is inserted:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <title>{{ title || My site }}</title>
</head>
<body>
    {{ content }}
</body>
</html>
```

`{{ title || My site }}` prints the page's `title`, or `My site` if the page has none.

### Frontmatter

A page picks its layout, and sets variables, in frontmatter at the top of the file:

```html
---
layout: default.html
title: Home
---
<h1>Hello</h1>
```

Frontmatter is flat `name: value` lines, not full yaml.

### Includes

An include is an html or svg file from `_includes`, inserted with a comment:

```html
<!-- footer.html -->
```

Pass props to an include in object syntax. Strings can use single or double quotes:

```html
<!-- button.html, { theme: 'primary', label: "Save" } -->
```

Inside the include, props and the page's frontmatter variables are available as `{{ name }}`:

```html
<button class="{{ theme || default }}">{{ label }}</button>
```

### Custom domain

Keep your `CNAME` file in `src/`. It's copied to `docs/` like any other file, so it survives a rebuild.


## Coming from piko

kit replaces piko. Your existing sites build to the same output. You can keep piko, but if you want to remove it:

    deno uninstall -g piko

Then install kit as above. What's different in kit compared to piko:

- The port is set with `-p` or `--port`, not as a positional argument, and the default port changed from 3333 to 4444.
- `build` has no `-f` flag. It always writes exactly the files that changed.
- There's no `share` command i kit, see below.
- serve only listens on 127.0.0.1, and sends no CORS header.
- Include props accept both single and double quotes.
- `docs/CNAME` is copied back to `src/CNAME` only when `src/CNAME` is missing. Keep `CNAME` in `src/`.
- `-v` and `--version` still work, but `kit version` is the documented form.

### No share command

To open your site on a phone, or share it with a colleague, use a cloudflare tunnel directly ([install cloudflared](https://developers.cloudflare.com/cloudflare-one/networks/connectors/cloudflare-tunnel/downloads/)). Replace `PORT` with the port kit prints:

    cloudflared tunnel --url localhost:PORT --http-host-header localhost:PORT

A shortcut for your powershell `$PROFILE`:

    function share($port) { cloudflared tunnel --url "localhost:$port" --http-host-header "localhost:$port" }

Then run `share PORT`.


## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).


## License

[MIT](LICENSE)
