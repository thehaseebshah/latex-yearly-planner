\clearpage
{{ template "breadcrumb_00_header.tpl" dict "Cfg" .Cfg "Body" .Body }}
{{ if $.Cfg.Dotted -}}
\Repeat{32}{\myLineGrayVskipTop}
{{- else -}}
\Repeat{32}{\myLineGrayVskipTop}
{{- end}}
\pagebreak
