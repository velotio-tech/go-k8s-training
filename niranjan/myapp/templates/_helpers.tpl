{{/*
Users labels
*/}}
{{- define "users.labels" -}}
app: {{- .Values.usersdeployment.labels.app | indent 1 }}
{{- end }}

{{/*
Orders labels
*/}}
{{- define "orders.labels" -}}
app: {{- .Values.usersdeployment.labels.app | indent 1 }}
{{- end }}

{{/*
MySQL labels
*/}}
{{- define "mysql.labels" -}}
app: {{- .Values.mysqldeployment.labels.app | indent 1 }}
tier: {{- .Values.mysqldeployment.labels.tier | indent 1 }}
{{- end }}
