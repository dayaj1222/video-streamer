# video-streamer

A lightweight HTTP video streaming server written in Go. Supports byte-range requests for seeking/scrubbing in any media player or browser.

## Features

- List available videos via REST API
- Stream videos with full HTTP range request support
- Zero dependencies beyond the Go standard library

## Getting Started

**Prerequisites:** Go 1.21+

```bash
git clone https://github.com/dayaj1222/video-streamer.git
cd video-streamer
```

Place `.mp4` (or any video) files in the `videos/` directory, then run:

```bash
go run main.go
```

Server starts on **http://localhost:8080**

## API

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/videos` | Returns a JSON array of available video filenames |
| `GET` | `/video?file=<filename>` | Streams the specified video (supports range requests) |

**Examples:**

```bash
# List videos
curl http://localhost:8080/videos

# Stream a video (download or pipe to player)
curl http://localhost:8080/video?file=example.mp4 -o out.mp4

# Open directly in browser
open "http://localhost:8080/video?file=example.mp4"
```

## Development (hot reload)

Requires [Node.js](https://nodejs.org) and `nodemon`:

```bash
npm install -g nodemon
node nodemon.js
```

## License

MIT — see [LICENSE](LICENSE)
