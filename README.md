# yt-go-gui

This is a GUI app using the audio/video downloader command-line [yt-dlp](https://github.com/yt-dlp/yt-dlp) written in Go with [Fyne](https://fyne.io/) as the GUI framework

This project is split in two repository

It goes on pair with [@wharton](https://github.com/wharton-git)'s extension [`url-video-detector`](https://github.com/wharton-git/url-video-detector)

> [!warning] Warning
>
> This project is still in its early stage, breaking change may (if not **will**) occur

## Installation

As for now it is only possible to use it by building the source code :

```bash
git clone https://github.com/dzyanino/yt-go-gui

cd yt-go-gui

go mod tidy # or "go get ." to get the dependences
```

And then run it :

```bash
go run .
```

Or you can also build it :

```bash
go build .
```

## Usage

`yt-go-gui` starts a server on a given port and listens to it (defaults to `43214`).

Then waits for HTTP request from the browser extension, which contains the video URL
