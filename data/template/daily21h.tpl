{{ .Greeting }}
**{{ .Time }} giờ 30** hôm nay có:
{{- if .Content.GuildFeast }}
- **Tiệc hội.**
{{- end }}
