package link2rst

import (
	"html/template"
	"net/http"
)

type TemplateValue struct {
	Textarea string
}

var index = `<!doctype html>
<html>
<head>
  <title>HTML Link to Rst</title>
</head>
<body>
  <form action="/" method="post">
    URL: <input name="url" size="80">
    <button>Send</button>
  </form><br>
  <textarea id="ta" rows="5" cols="80">{{.Textarea}}</textarea><br>
  <button type="button" id="copy">Copy textarea to clipboard</button>
<script>
  var textareaElm = document.getElementById("ta");
  var copyElm = document.getElementById("copy");
  copyElm.onclick = function(event) {
    textareaElm.select();
    var isSuccessful = document.execCommand('copy');
    if (isSuccessful) {
      textareaElm.value = "Copy OK";
    } else {
      textareaElm.value = "Copy Fail";
    }
  }
</script>

</body>
</html>`

var tmpl = template.Must(template.New("link2rst").Parse(index))

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	val := TemplateValue{}
	if r.Method == "POST" {
		val.Textarea = getLinkRst(r.PostFormValue("url"), r)
	}

	if err := tmpl.Execute(w, &val); err != nil {
		panic(err)
	}
}
