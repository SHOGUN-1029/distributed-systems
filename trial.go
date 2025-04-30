package main

/*
var tmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Ray Docker Controller</title>
</head>
<body>
    <h1>Ray Container Management</h1>
    <button onclick="send('/start')">Start Container</button>
    <button onclick="send('/ray?cmd=ray start --head --port=6379 --dashboard-host=0.0.0.0')">Start Ray Head</button>
    <button onclick="send('/ray?cmd=ray status')">Ray Status</button>
    <button onclick="send('/ray?cmd=ray stop')">Stop Ray</button>
    <button onclick="send('/stop')">Stop Container</button>
    <pre id="output"></pre>

    <script>
    function send(url) {
        fetch(url)
            .then(response => response.text())
            .then(data => {
                document.getElementById('output').textContent = data;
            });
    }
    </script>
</body>
</html>
`

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/start", startContainer)
	http.HandleFunc("/ray", runRayCommand)
	http.HandleFunc("/stop", stopContainer)
	fmt.Println("[+] Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("webpage").Parse(tmpl))
	t.Execute(w, nil)
}

func startContainer(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("docker", "run", "-d", "--name", "ray-head", "-p", "6379:6379", "-p", "8265:8265", "rayproject/ray-ml:latest-py39", "sleep", "infinity")
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, string(output)+"\n"+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Started Container:\n%s", output)
}

func runRayCommand(w http.ResponseWriter, r *http.Request) {
	cmdStr := r.URL.Query().Get("cmd")
	cmd := exec.Command("docker", "exec", "ray-head", "bash", "-c", cmdStr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, string(output)+"\n"+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Command Output:\n%s", output)
}

func stopContainer(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("docker", "rm", "-f", "ray-head")
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, string(output)+"\n"+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Stopped and Removed Container:\n%s", output)
}
*/
