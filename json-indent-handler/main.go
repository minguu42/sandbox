package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
)

func main() {
	l := slog.New(NewJSONIndentHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				a.Key = "message"
			}
			return a
		},
	}))

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			l.Info(fmt.Sprintf("number %d", i))
		}()
	}
	wg.Wait()
	// l.With("a", "b").WithGroup("G").With("c", "d").WithGroup("H").Info("msg", "e", "f")
	// {
	//  "time": "2024-07-01T09:00:00.000000+09:00",
	//  "level": "INFO",
	//  "source": {
	//    "function": "main.main",
	//    "file": "/Users/example/main.go",
	//    "line": 19
	//  },
	//  "message": "msg",
	//  "a": "b",
	//  "G": {
	//    "c": "d",
	//    "H": {
	//      "e": "f"
	//    }
	//  }
	// }
}
