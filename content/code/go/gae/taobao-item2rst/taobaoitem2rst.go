package taobaoitem2rst

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
  <title>Taobao Item to Rst</title>
</head>
<body>
  <form action="/" method="post">
    URL: <input name="url" size="80">
    <button>Send</button>
  </form><br>
  <textarea id="ta" rows="5" cols="80">{{.Textarea}}</textarea><br>
  <button type="button" id="copy">Copy textarea to clipboard</button>

  <br><br>
  <a target="_blank" href="http://html2rst.golden-operator-130720.appspot.com/">HTML to reStructuredText</a>
  <br><br>
  <a target="_blank" href="http://v1.golden-operator-130720.appspot.com/">URL to reStructuredText</a>

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

var tmpl = template.Must(template.New("taobaoitem2rst").Parse(index))

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	val := TemplateValue{}
	if r.Method == "POST" {
		val.Textarea = getTaobaoItemImgRst(r.PostFormValue("url"), r)
	}

	if err := tmpl.Execute(w, &val); err != nil {
		panic(err)
	}
}
