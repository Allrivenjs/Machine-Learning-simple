package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Exercise struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Stage struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Duration  string     `json:"duration"`
	Folder    string     `json:"folder"`
	Color     string     `json:"color"`
	Exercises []Exercise `json:"exercises"`
}

type Progress struct {
	Completed []string `json:"completed"`
}

func projectRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..", "..")
}

func loadStages() ([]Stage, error) {
	data, err := os.ReadFile(filepath.Join(projectRoot(), "stages.json"))
	if err != nil {
		return nil, err
	}
	var stages []Stage
	return stages, json.Unmarshal(data, &stages)
}

func loadProgress() (Progress, error) {
	data, err := os.ReadFile(filepath.Join(projectRoot(), "progress.json"))
	if err != nil {
		return Progress{Completed: []string{}}, nil
	}
	var p Progress
	return p, json.Unmarshal(data, &p)
}

func saveProgress(p Progress) error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(projectRoot(), "progress.json"), data, 0644)
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func remove(slice []string, s string) []string {
	result := make([]string, 0, len(slice))
	for _, v := range slice {
		if v != s {
			result = append(result, v)
		}
	}
	return result
}

func handleToggle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	p, _ := loadProgress()
	if contains(p.Completed, req.ID) {
		p.Completed = remove(p.Completed, req.ID)
	} else {
		p.Completed = append(p.Completed, req.ID)
	}
	if err := saveProgress(p); err != nil {
		http.Error(w, "save error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"completed": contains(p.Completed, req.ID)})
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	stages, err := loadStages()
	if err != nil {
		http.Error(w, "cannot load stages.json: "+err.Error(), 500)
		return
	}
	progress, _ := loadProgress()

	totalEx := 0
	totalDone := 0
	for _, s := range stages {
		totalEx += len(s.Exercises)
		for _, ex := range s.Exercises {
			if contains(progress.Completed, ex.ID) {
				totalDone++
			}
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlHead(totalDone, totalEx))

	for _, stage := range stages {
		done := 0
		for _, ex := range stage.Exercises {
			if contains(progress.Completed, ex.ID) {
				done++
			}
		}
		fmt.Fprint(w, renderStage(stage, progress.Completed, done))
	}

	fmt.Fprint(w, htmlFoot())
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	stages, _ := loadStages()
	progress, _ := loadProgress()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"stages":    stages,
		"completed": progress.Completed,
	})
}

func handleView(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Query().Get("f")
	if f == "" {
		http.Error(w, "missing f param", http.StatusBadRequest)
		return
	}
	// prevent path traversal
	clean := filepath.Clean(f)
	if strings.HasPrefix(clean, "..") || filepath.IsAbs(clean) {
		http.Error(w, "invalid path", http.StatusForbidden)
		return
	}
	full := filepath.Join(projectRoot(), clean)
	data, err := os.ReadFile(full)
	if err != nil {
		http.Error(w, "file not found: "+clean, http.StatusNotFound)
		return
	}
	title := filepath.Base(clean)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, viewHTML(title, string(data)))
}

func viewHTML(title, content string) string {
	lines := strings.Split(content, "\n")
	var sb strings.Builder
	inCode := false
	sb.WriteString(fmt.Sprintf(`<!DOCTYPE html>
<html lang="es">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<title>%s</title>
<style>
  *{box-sizing:border-box;margin:0;padding:0}
  body{background:#0f0f13;color:#e2e8f0;font-family:'Segoe UI',system-ui,sans-serif;
       max-width:860px;margin:0 auto;padding:2rem 1.5rem;line-height:1.7}
  a.back{display:inline-block;margin-bottom:1.5rem;color:#6366f1;text-decoration:none;
         font-size:.85rem;border:1px solid #6366f133;padding:4px 12px;border-radius:6px}
  a.back:hover{background:#6366f120}
  h1{font-size:1.7rem;color:#f1f5f9;margin:1.2rem 0 .6rem;border-bottom:1px solid #2a2a3e;padding-bottom:.5rem}
  h2{font-size:1.25rem;color:#c4b5fd;margin:1.4rem 0 .5rem}
  h3{font-size:1.05rem;color:#93c5fd;margin:1rem 0 .4rem}
  p{margin:.4rem 0;color:#cbd5e1}
  ul,ol{margin:.4rem 0 .4rem 1.4rem;color:#cbd5e1}
  li{margin:.2rem 0}
  pre{background:#1e1e2e;border:1px solid #2a2a3e;border-radius:8px;
      padding:1rem 1.2rem;overflow-x:auto;margin:.8rem 0}
  code{font-family:'JetBrains Mono','Fira Code',monospace;font-size:.82rem;color:#a5f3fc}
  pre code{color:#e2e8f0}
  .checkbox{color:#10b981}
  .checkbox-empty{color:#475569}
  hr{border:none;border-top:1px solid #2a2a3e;margin:1.2rem 0}
  strong{color:#f1f5f9}
  em{color:#fbbf24}
</style>
</head>
<body>
<a class="back" href="/">← Volver al dashboard</a>
`, html.EscapeString(title)))

	for _, line := range lines {
		if strings.HasPrefix(line, "```") {
			if inCode {
				sb.WriteString("</pre>\n")
				inCode = false
			} else {
				sb.WriteString("<pre><code>")
				inCode = true
			}
			continue
		}
		if inCode {
			sb.WriteString(html.EscapeString(line) + "\n")
			continue
		}
		switch {
		case strings.HasPrefix(line, "# "):
			sb.WriteString("<h1>" + html.EscapeString(line[2:]) + "</h1>\n")
		case strings.HasPrefix(line, "## "):
			sb.WriteString("<h2>" + html.EscapeString(line[3:]) + "</h2>\n")
		case strings.HasPrefix(line, "### "):
			sb.WriteString("<h3>" + html.EscapeString(line[4:]) + "</h3>\n")
		case strings.HasPrefix(line, "---"):
			sb.WriteString("<hr>\n")
		case strings.HasPrefix(line, "- [x] "), strings.HasPrefix(line, "- [X] "):
			sb.WriteString(`<p><span class="checkbox">☑</span> ` + html.EscapeString(line[6:]) + "</p>\n")
		case strings.HasPrefix(line, "- [ ] "):
			sb.WriteString(`<p><span class="checkbox-empty">☐</span> ` + html.EscapeString(line[6:]) + "</p>\n")
		case strings.HasPrefix(line, "- "), strings.HasPrefix(line, "* "):
			sb.WriteString("<li>" + html.EscapeString(line[2:]) + "</li>\n")
		case line == "":
			sb.WriteString("<p></p>\n")
		default:
			escaped := html.EscapeString(line)
			escaped = strings.ReplaceAll(escaped, "**", "<strong>")
			if strings.Count(escaped, "<strong>")%2 != 0 {
				escaped = strings.Replace(escaped, "<strong>", "</strong>", strings.Count(escaped, "<strong>")-1)
			}
			sb.WriteString("<p>" + escaped + "</p>\n")
		}
	}
	if inCode {
		sb.WriteString("</pre>\n")
	}
	sb.WriteString("</body></html>")
	return sb.String()
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/toggle", handleToggle)
	http.HandleFunc("/api", handleAPI)
	http.HandleFunc("/view", handleView)

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	fmt.Printf("Dashboard corriendo en http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func htmlHead(done, total int) string {
	pct := 0
	if total > 0 {
		pct = done * 100 / total
	}
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="es">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>ML Roadmap — Dashboard</title>
<style>
  * { box-sizing: border-box; margin: 0; padding: 0; }
  body {
    background: #0f0f13;
    color: #e2e8f0;
    font-family: 'Segoe UI', system-ui, sans-serif;
    min-height: 100vh;
    padding: 2rem 1rem;
  }
  .header {
    text-align: center;
    margin-bottom: 2.5rem;
  }
  .header h1 {
    font-size: 2rem;
    font-weight: 700;
    background: linear-gradient(135deg, #6366f1, #ec4899);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    margin-bottom: 0.5rem;
  }
  .header p { color: #94a3b8; font-size: 0.95rem; }
  .overall-bar {
    max-width: 600px;
    margin: 1.5rem auto;
  }
  .bar-label {
    display: flex;
    justify-content: space-between;
    font-size: 0.85rem;
    color: #94a3b8;
    margin-bottom: 0.4rem;
  }
  .bar-track {
    background: #1e1e2e;
    border-radius: 99px;
    height: 12px;
    overflow: hidden;
  }
  .bar-fill {
    height: 100%%;
    border-radius: 99px;
    background: linear-gradient(90deg, #6366f1, #ec4899);
    transition: width 0.4s ease;
  }
  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(360px, 1fr));
    gap: 1.5rem;
    max-width: 1400px;
    margin: 0 auto;
  }
  .card {
    background: #16161f;
    border: 1px solid #2a2a3e;
    border-radius: 12px;
    overflow: hidden;
    transition: border-color 0.2s;
  }
  .card:hover { border-color: #3a3a5e; }
  .card-header {
    padding: 1rem 1.25rem 0.75rem;
    border-bottom: 1px solid #2a2a3e;
  }
  .stage-tag {
    font-size: 0.7rem;
    font-weight: 600;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    padding: 2px 8px;
    border-radius: 4px;
    display: inline-block;
    margin-bottom: 0.5rem;
  }
  .card-title {
    font-size: 1.05rem;
    font-weight: 600;
    color: #f1f5f9;
  }
  .card-meta {
    font-size: 0.78rem;
    color: #64748b;
    margin-top: 0.25rem;
  }
  .card-progress {
    padding: 0.6rem 1.25rem;
    display: flex;
    align-items: center;
    gap: 0.75rem;
    border-bottom: 1px solid #2a2a3e;
  }
  .mini-bar { flex: 1; background: #1e1e2e; border-radius: 99px; height: 6px; }
  .mini-fill { height: 100%%; border-radius: 99px; transition: width 0.4s; }
  .pct-label { font-size: 0.75rem; color: #64748b; width: 36px; text-align: right; }
  .exercises { padding: 0.5rem 0; }
  .ex-item {
    display: flex;
    align-items: center;
    gap: 0.6rem;
    padding: 0.35rem 1.25rem;
    cursor: pointer;
    border-radius: 6px;
    transition: background 0.15s;
    user-select: none;
  }
  .ex-item:hover { background: #1e1e2e; }
  .check {
    width: 16px; height: 16px;
    border: 2px solid #3a3a5e;
    border-radius: 4px;
    flex-shrink: 0;
    display: flex; align-items: center; justify-content: center;
    transition: all 0.15s;
  }
  .check.done {
    border-color: currentColor;
    background: currentColor;
  }
  .check.done::after {
    content: '';
    width: 9px; height: 5px;
    border-left: 2px solid #0f0f13;
    border-bottom: 2px solid #0f0f13;
    transform: rotate(-45deg) translateY(-1px);
  }
  .ex-id { font-size: 0.68rem; color: #475569; font-family: monospace; flex-shrink: 0; }
  .ex-title { font-size: 0.82rem; color: #94a3b8; }
  .ex-item.done .ex-title { color: #475569; text-decoration: line-through; }
  .links {
    padding: 0.75rem 1.25rem;
    border-top: 1px solid #2a2a3e;
    display: flex;
    gap: 0.75rem;
  }
  .link {
    font-size: 0.75rem;
    color: #6366f1;
    text-decoration: none;
    padding: 3px 8px;
    border: 1px solid #6366f133;
    border-radius: 4px;
    transition: background 0.15s;
  }
  .link:hover { background: #6366f120; }
  .completed-badge {
    font-size: 0.75rem;
    background: #10b98120;
    color: #10b981;
    padding: 2px 8px;
    border-radius: 4px;
    margin-left: auto;
  }
</style>
</head>
<body>
<div class="header">
  <h1>ML Roadmap</h1>
  <p>Desde cero hasta Machine Learning Engineer</p>
  <div class="overall-bar">
    <div class="bar-label">
      <span>Progreso total</span>
      <span>%d / %d ejercicios (%d%%)</span>
    </div>
    <div class="bar-track">
      <div class="bar-fill" style="width: %d%%"></div>
    </div>
  </div>
</div>
<div class="grid">
`, done, total, pct, pct)
}

func renderStage(s Stage, completed []string, done int) string {
	total := len(s.Exercises)
	pct := 0
	if total > 0 {
		pct = done * 100 / total
	}

	allDone := done == total && total > 0

	result := fmt.Sprintf(`
<div class="card">
  <div class="card-header">
    <div class="stage-tag" style="background: %s22; color: %s">Etapa %d</div>
    <div style="display:flex; align-items:center; gap:0.5rem">
      <div class="card-title">%s</div>
      %s
    </div>
    <div class="card-meta">%s &nbsp;·&nbsp; %d ejercicios</div>
  </div>
  <div class="card-progress">
    <div class="mini-bar">
      <div class="mini-fill" style="width:%d%%; background:%s"></div>
    </div>
    <div class="pct-label">%d%%</div>
  </div>
  <div class="exercises">
`, s.Color, s.Color, s.ID, s.Name,
		func() string {
			if allDone {
				return `<span class="completed-badge">✓ Completada</span>`
			}
			return ""
		}(),
		s.Duration, total, pct, s.Color, pct)

	for _, ex := range s.Exercises {
		isDone := contains(completed, ex.ID)
		checkClass := "check"
		itemClass := "ex-item"
		if isDone {
			checkClass += " done"
			itemClass += " done"
		}
		result += fmt.Sprintf(`
    <div class="%s" onclick="toggle('%s', this)" data-id="%s" style="--c:%s">
      <div class="%s" style="%s"></div>
      <span class="ex-id">%s</span>
      <span class="ex-title">%s</span>
    </div>`,
			itemClass, ex.ID, ex.ID, s.Color,
			checkClass,
			func() string {
				if isDone {
					return fmt.Sprintf("color:%s", s.Color)
				}
				return ""
			}(),
			ex.ID, ex.Title)
	}

	result += fmt.Sprintf(`
  </div>
  <div class="links">
    <a class="link" href="/view?f=%s/guia-lectura.md">📖 Guía de lectura</a>
    <a class="link" href="/view?f=%s/ejercicios.md">✏️ Ejercicios</a>
  </div>
</div>
`, s.Folder, s.Folder)

	return result
}

func htmlFoot() string {
	return `
</div>
<script>
async function toggle(id, el) {
  const res = await fetch('/toggle', {
    method: 'POST',
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify({id})
  });
  const data = await res.json();
  const check = el.querySelector('.check');
  const color = el.style.getPropertyValue('--c');
  if (data.completed) {
    check.classList.add('done');
    check.style.color = color;
    el.classList.add('done');
  } else {
    check.classList.remove('done');
    check.style.color = '';
    el.classList.remove('done');
  }
  updateProgress();
}

async function updateProgress() {
  const res = await fetch('/api');
  const data = await res.json();
  let total = 0, done = 0;
  for (const stage of data.stages) {
    total += stage.exercises.length;
    for (const ex of stage.exercises) {
      if (data.completed.includes(ex.id)) done++;
    }
  }
  const pct = total > 0 ? Math.round(done * 100 / total) : 0;
  document.querySelector('.bar-fill').style.width = pct + '%';
  document.querySelector('.bar-label span:last-child').textContent =
    done + ' / ' + total + ' ejercicios (' + pct + '%)';
}
</script>
</body>
</html>
`
}
