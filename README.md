markdown resume
===============

Write resume in Markdown. Convert to HTML with Pandoc. Convert to PDF with Chrome headless.

Written by ChatGPT.

```bash
docker run --rm -v "$(pwd):/data" pandoc/core -f markdown -t html -o /data/resume.html /data/resume.md
go build -o html2pdf main.go
./html2pdf -input resume.html -output resume.pdf
```
