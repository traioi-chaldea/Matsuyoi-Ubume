{{ .Greeting }}
**{{ .Time }} giờ** hôm nay có:
{{- if ne .Content.GuildRaid "" }}
- **Kỳ Lân (hội):** {{ .Content.GuildRaid }}
{{- end }}
