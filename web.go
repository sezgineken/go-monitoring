package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func startWebServer() {
	http.HandleFunc("/", viewLogHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/clear", clearLogsHandler)

	println("🌐 Gelişmiş Web arayüzü başlatıldı: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func reverse(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func trimLogPaths(paths []string) []string {
	var result []string
	for _, p := range paths {
		result = append(result, filepath.Base(p))
	}
	return result
}

func viewLogHandler(w http.ResponseWriter, r *http.Request) {
	files, _ := filepath.Glob("logs/*.log")
	if len(files) == 0 {
		http.Error(w, "Hiç log bulunamadı", 404)
		return
	}
	sort.Strings(files)

	selectedFile := r.URL.Query().Get("file")
	if selectedFile == "" {
		selectedFile = filepath.Base(files[len(files)-1])
	}

	fullPath := ""
	for _, f := range files {
		if filepath.Base(f) == selectedFile {
			fullPath = f
			break
		}
	}
	if fullPath == "" {
		http.Error(w, "Seçilen log dosyası bulunamadı", 404)
		return
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		http.Error(w, "Log okunamadı", 500)
		return
	}

	lines := strings.Split(string(data), "\n")
	reverse(lines)

	renderTemplate(w, lines, trimLogPaths(files), selectedFile)
}

func clearLogsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	files, _ := filepath.Glob("logs/*.log")
	for _, file := range files {
		os.Remove(file)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("🟢 Sistem İzleyici Web Arayüzü çalışıyor.\n"))
}

func renderTemplate(w http.ResponseWriter, lines, files []string, selected string) {
	const tmpl = `
<!DOCTYPE html>
<html lang="tr">
<head>
        <meta charset="UTF-8">
        <title>Log Görüntüleyici</title>
        <style>
                body { font-family: Arial, sans-serif; background: #1e1e1e; color: #f1f1f1; margin: 0; }
                .sidebar {
                        width: 220px; background: #2c2c2c; position: fixed; height: 100vh;
                        padding: 20px; box-shadow: 2px 0 5px rgba(0,0,0,0.3);
                }
                .sidebar a, .sidebar form button {
                        display: block; margin: 10px 0; text-decoration: none;
                        color: #f1f1f1; background: #444; padding: 10px;
                        border-radius: 5px; text-align: center;
                }
                .sidebar select { width: 100%; padding: 6px; background: #333; color: #fff; border: none; border-radius: 4px; }
                .main { margin-left: 240px; padding: 20px; }
                .log-line { padding: 8px; margin: 5px 0; border-left: 4px solid #666; background: #2a2a2a; border-radius: 4px; }
                .warning { border-left-color: #f00; color: #f88; font-weight: bold; }
                .info { border-left-color: #0af; }
                .error { border-left-color: #f00; background: #400; }
                #search, #levelFilter {
                        padding: 8px; margin-bottom: 10px; width: 100%;
                        background: #333; color: #fff; border: none; border-radius: 4px;
                }
                #refresh-btn {
                        padding: 8px 16px; background: #005; color: #fff; border: none; border-radius: 5px; cursor: pointer;
                }
        </style>
</head>
<body>

        <div class="sidebar">
                <h3>🔧 Menü</h3>
                <a href="/">📄 Tüm Loglar</a>
                <a href="/status">📊 Durum</a>
                <form method="POST" action="/clear">
                        <button onclick="return confirm('Tüm logları silmek istediğine emin misin?')">🗑️ Logları Temizle</button>
                </form>
                <h4>📁 Log Dosyası</h4>
                <form method="GET" onchange="this.submit()">
                        <select name="file">
                                {{range .Files}}<option value="{{.}}" {{if eq . $.Selected}}selected{{end}}>{{.}}</option>{{end}}
                        </select>
                </form>
        </div>

        <div class="main">
                <h1>📋 Loglar</h1>
                <input type="text" id="search" placeholder="Anahtar kelime ile ara... (örn: CPU, ⚠️)">
                <select id="levelFilter">
                        <option value="">🔎 Seviye Filtrele</option>
                        <option value="info">Info</option>
                        <option value="warning">Warning</option>
                        <option value="error">Error</option>
                </select>
                <button id="refresh-btn" onclick="window.location.reload()">🔁 Yenile</button>

                <div id="log-entries">
                        {{range .Lines}}
                                {{if .}}
                                        {{ $class := "" }}
                                        {{if (hasError .)}}{{ $class = "error" }}
                                        {{else if (hasWarning .)}}{{ $class = "warning" }}
                                        {{else if (hasInfo .)}}{{ $class = "info" }}
                                        {{end}}
                                        <div class="log-line {{ $class }}">{{.}}</div>
                                {{end}}
                        {{end}}
                </div>
        </div>

        <script>
                const searchBox = document.getElementById("search");
                const levelFilter = document.getElementById("levelFilter");
                const logs = document.querySelectorAll(".log-line");

                searchBox.addEventListener("input", filterLogs);
                levelFilter.addEventListener("change", filterLogs);

                function filterLogs() {
                        const query = searchBox.value.toLowerCase();
                        const level = levelFilter.value;

                        logs.forEach(function(log) {
                                const text = log.innerText.toLowerCase();
                                const matchQuery = query === "" || text.includes(query);
                                const matchLevel = level === "" || log.classList.contains(level);
                                log.style.display = (matchQuery && matchLevel) ? "block" : "none";
                        });
                }
        </script>

</body>
</html>
`

	t := template.Must(template.New("log").Funcs(template.FuncMap{
		"hasWarning": func(s string) bool {
			return strings.Contains(s, "⚠️") || strings.Contains(strings.ToLower(s), "warning")
		},
		"hasError": func(s string) bool {
			return strings.Contains(strings.ToLower(s), "error") || strings.Contains(s, "❌")
		},
		"hasInfo": func(s string) bool {
			return strings.Contains(strings.ToLower(s), "info") || strings.Contains(s, "ℹ️")
		},
	}).Parse(tmpl))

	t.Execute(w, struct {
		Lines    []string
		Files    []string
		Selected string
	}{
		Lines:    lines,
		Files:    files,
		Selected: selected,
	})
}
