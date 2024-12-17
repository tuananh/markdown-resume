markdown resume
===============

Write resume in Markdown. Convert to HTML with Pandoc. Convert to PDF with Chrome headless.

Written by ChatGPT.

```bash
docker run --rm -v "$(pwd):/workspace" pandoc/core /workspace/resume.md -f markdown -t html -c /workspace/style.css -s -o /workspace/resume.html
go build -o html2pdf main.go
./html2pdf -input resume.html -output resume.pdf
```
