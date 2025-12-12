{{- /* Common chart helper templates */ -}}
{{- define "face-detector.name" -}}
{{- printf "%s" .Chart.Name -}}
{{- end -}}

{{- define "face-detector.fullname" -}}
{{- .Release.Name -}}
{{- end -}}
