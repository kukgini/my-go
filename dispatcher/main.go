package main

import (
    "log"
    "time"
    "io/fs"
    "io/ioutil"
    "github.com/kukgini/my-go/dispatcher/events"    
)

type DirectoryWatcher struct {
    Dispatcher
    path string // path to directory on disk
}

func (w *DirectoryWatcher) Init(path string) {
    w.path = path
    w.Dispatcher.Init(w)
}

func (w *DirectoryWatcher) Unseen(file fs.FileInfo) bool {
    return true 
}

func (w *DirectoryWatcher) Watch() error {
    for {
        files, _ := ioutil.ReadDir(w.path)
        for _, file := range files {
            if w.Unseen(file) {
                w.Dispatch(NewFile, file)
            }
        }
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    watcher := DirectoryWatcher{}
    watcher.Init(".")    
    watcher.Watch()
}
