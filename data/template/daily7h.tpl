{{ .Greeting }}
**{{ .Time }} giờ** hôm nay có:
{{- if .Content.NetherworldGate }}
- **Cổng ma**
{{- end }}
{{- if ne .Content.DemonEncounter "" }}
- **Boss TKDD:** {{ .Content.DemonEncounter }}
{{- end }}
