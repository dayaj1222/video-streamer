package streamer

import (
	"net/http"
	"os"
	"path/filepath"
)

type Streamer struct {
	VideoDir string
	Videos   []Video
}

type Video struct {
	Name string
}

func (s *Streamer) GetAllVideos() ([]string, error) {
	files := []string{}

	entries, err := os.ReadDir(s.VideoDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

func (s *Streamer) StreamVideo(w http.ResponseWriter, r *http.Request, filename string) error {

	path := filepath.Join(s.VideoDir, filename)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return err
	}

	http.ServeContent(w, r, filename, fi.ModTime(), file)

	return nil
}
