{{- /* Common chart helper templates */ -}}
{{- define "face-detector.name" -}}
{{- printf "%s" .Chart.Name -}}
{{- end -}}

{{- define "face-detector.fullname" -}}
{{- $name := printf "%s-%s" .Release.Name (include "face-detector.name" .) -}}
{{- if lt (len $name) 64 -}}
{{- $name -}}
{{- else -}}
{{- trunc 63 $name | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
