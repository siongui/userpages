package taobaoitem2rst

import (
	"bytes"
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

var imgRst = `.. image:: {{ .ImgURL }}
   :alt: {{ .Title }}
   :target: {{ .URL }}
   :align: center`

func getTaobaoItemImgRst(url string, r *http.Request) string {
	nUrl := NormalizeURL(url)
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get(nUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	ii := getTaobaoItemInfo(resp.Body)
	ii.URL = nUrl

	tmpl := template.Must(template.New("imgRst").Parse(imgRst))
	var rst bytes.Buffer
	err = tmpl.Execute(&rst, &ii)
	if err != nil {
		panic(err)
	}

	return rst.String()
}
