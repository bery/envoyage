package main

var rootPage = `
<html>
<meta charset="utf-8">

<head>
<title>Kubernetes app demo</title>
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
<style>
td {
	font-size: 70%;
}
a {
	font-weight: bold;
}
body {
	padding: 10px;
	background-color: {{ .Color }};
}
div.doc {
	font-size: 90%;
	background-color: #fafafa;
	padding: 10px;
	border-radius: 2px;
	font-weight: 400;
}

li.pod a, li.deploy a, li.rs a, li.svc a {
	display: block;
	padding: 2px;
	margin: 2px;
	color: white;
	text-align: center;
}

li a, li a:hover {
	color: white;
	cursor: crosshair;
	text-decoration: none;
}

li.pod {
	background-color: #050505;
}
li.deploy {
	background-color: #ffc674;
}
li.rs {
	background-color: #f0a9d5
}
li.svc {
	background-color: #dcea64;
}
table td { word-wrap:break-word; }
</style>
{{ if .PageRefresh }}
<meta http-equiv="refresh" content="2">
{{ end }}
</head>

<body>
<div class="container">
<h1>{{ .Namespace }}</h1>
<div class="row">

<div class="col-sm-6">

{{ if .Hits }}
<div class="alert alert-info">This worker returned page <strong>{{ .Hits }}</strong> times.</div>
{{ end }}


<div class="alert alert-info">Metrics exported at <a href="/metrics">/metrics</a></div>


{{ if .RedisHost }}
<div class="alert alert-info">Redis server is <code>{{ .RedisHost }}</code>, path <code>{{ .RedisPath }}</code></div>
{{ else }}
<div class="alert alert-warning">Redis server not used, set <code>REDIS_SERVER</code> to use it.</div>
{{ end }}

{{ if ne .RedisError "" }}
<div class="alert alert-danger">Redis connection failed: <code>{{ .RedisError }}</code></div>
{{ end }}

{{ if .ArgoRedisHost }}
<div class="alert alert-info">Argo Redis server is <code>{{ .ArgoRedisHost }}</code></div>
{{ end }}

{{ if ne .ArgoRedisError "" }}
<div class="alert alert-danger">Argo Redis connection failed: <code>{{ .ArgoRedisError }}</code></div>
{{ end }}

{{ if .Cmd }}
<div class="alert alert-info">Started with command <code>{{ .Cmd }}</code></div>
{{ end }}

{{ if .FailureProbability }}
<div class="alert alert-info">Request will be failing with probabilty <code>{{ .FailureProbability }}</code></div>
{{ end }}


{{ if .ConfFile }}
<div class="alert alert-info">Config file <code>/etc/kad/config.yml</code> content:<br><code><pre>{{ .ConfFile }}<pre></code></div>
{{ else }}
<div class="alert alert-warning">Config file <code>/etc/kad/config.yml</code> is empty.</code></div>
{{ end }}

{{ if not .Ready }}
<div class="alert alert-danger">This replica isn't ready.</div>
{{ end }}


{{ if .KubernetesError }}
<div class="alert alert-warning">Failed accesing Kubernetes: <code>{{ .KubernetesError }}</code></div>
{{ end }}

{{ if .PersistentFiles }}
<div class="alert alert-info">
Persistent files:<br>
<ul>
{{ with .PersistentFiles }}
{{ range . }}
	<li><code>{{ . }}</code></li>
{{ end }}
{{ end }}
</ul>
</div>
{{ end }}

{{ if ne .RemoteAddr "" }}
<div class="alert alert-info">Remote address: <code>{{ .RemoteAddr }}</code></div>
{{ end }}


<div class="doc">
<b>Aws instances</b>
{{ if ne .AwsError "" }}
<div class="alert alert-danger">Failed to get aws instances: <code>{{ .AwsError }}</code></div>
{{ else }}
<ul>
	{{ range $k, $v := .AwsInstances }}
		<li><code>{{ $k }}</code>: {{ $v }}</li>
	{{ end }}
</ul>
{{ end }}
</div>
<br />
<div class="doc">
<b>Endpoints (port {{ .Listen }}):</b>
<ul>
	<li><a>/heavy</a> - run many parallel goroutines printing /dev/null</li>
	<li><a>/slow</a> - wait 3 second before server reply</li>
	<li><a>/check/live</a> - liveness probe, always OK</li>
	<li><a>/check/ready</a> - readiness probo, ready if file <code>/tmp/notready</code> doesn't exist</li>
	<li><a>/metrics</a> - <a href="https://prometheus.io/">Prometheus</a> metrics</li>
	<li><a>/hostname</a> - prints hostname
</ul>

<b>Admin endpoints (port {{ .ListenAdmin }}):</b>
<ul>
	<li><a>/action/terminate</a> - Disable readiness probe, wait 15s and exit</li>
	<li><a>/check/live</a> - liveness probe, always OK</li>
	<li><a>/check/ready</a> - readiness probo, ready if file <code>/tmp/notready</code> doesn't exist</li>
	<li><a>/malware</a> - malware endpoint, exposes all cluster secrets and environment variables</li>
</ul>

<b>Command options:</b>
<ul>
	<li><a>--color</a> - Set background color</li>
	<li><a>--fail</a> - Terminate with non-zero exit code (immediatelly)</li>
	<li><a>--failure-probability</a> - Request to / will be failing with this probability</li>
</ul>


<p>
Server is expecting configuration file <code>{{ .ConfigFilePath }}</code>. It will run without configuration but error mesage will be printed.
</p>

</div>

</div>
<div class="col-sm-6">

{{ if not .KubernetesError }}
<div class="doc">

<p>
Kubernetes at <a href="{{ .KubernetesHost }}">{{ .KubernetesHost }}</a>
</p>

Pods
<ul>
{{ range $i := .Resources.Pods }}
<li class="pod"><a href="/kubernetes/delete/pod/{{ $i.ObjectMeta.Name }}">{{ $i.ObjectMeta.Name }}</a></li>
{{ end }}
</ul> 

Deployments
<ul>
{{ range $i := .Resources.Deployments }}
<li class="deploy"><a href="/kubernetes/delete/deploy/{{ $i.ObjectMeta.Name }}">{{ $i.ObjectMeta.Name }}</a></li>
{{ end }}
</ul> 

ReplicaSets
<ul>
{{ range $i := .Resources.ReplicaSets }}
<li class="rs"><a href="/kubernetes/delete/rs/{{ $i.ObjectMeta.Name }}">{{ $i.ObjectMeta.Name }}</a></li>
{{ end }}
</ul> 

Services
<ul>
{{ range $i := .Resources.Services }}
<li class="svc"><a href="/kubernetes/delete/svc/{{ $i.ObjectMeta.Name }}">{{ $i.ObjectMeta.Name }}</a></li>
{{ end }}
</ul> 
Secrets
<ul>
{{ range $i := .Resources.Secrets }}
<li class="svc"><a href="/kubernetes/delete/secret/{{ $i.ObjectMeta.Name }}">{{ $i.ObjectMeta.Name }}</a></li>
{{ range $k, $v := $i.Data }}
	<li><code>{{ $k }}</code>: {{ toString $v | substr 0 30 }}</li>
{{ end }}
{{ end }}
</ul> 

</div>
{{ end }}

{{ if .Keys }}
<div class="doc">

<p>
Argo redis keys
</p>

<ul>
{{ range $i := .Keys }}
{{ toString $i }}
{{ end }}
</ul> 
</div>
{{ end }}

<table class="table table-hover">
<thead>
<tr><th>Variable name</th><th>Value</th></tr>
</thead>
<tbody>
{{ range $v := .Vars }}
<tr class="{{ if $v.Dangerous }}table-danger{{ end }}"><td>{{ $v.Name }}</td><td>{{ $v.Value }}</td></tr>
{{ end }}
</tbody>
</table>
</div>

</div> <!-- row -->


<table class="table table-hover">
<thead>
<tr><th>Header name</th><th>Value</th></tr>
</thead>
<tbody>
{{ with .Headers }}
{{ range . }}
	<tr><td>{{ .Name }}</td><td>{{ .Value }}</td></tr>
{{ end }}
{{ end }}

</tbody>
</table>


</div> <!-- container -->

</body>
</html>
`
