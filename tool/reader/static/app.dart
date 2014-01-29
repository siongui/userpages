import 'dart:html';
import 'dart:convert';

void showRssItems(String html) {
  DivElement view = query(".rss-site-items");
  view.children.clear();
  view.appendHtml(html);

  // toggle content (description) of rss item
  view.queryAll(".rss-item").forEach((Element elm) {
    elm.dataset["isShowContent"] = "false";
    elm.query(".rss-item-head").onClick.listen((Event e) {
      if (elm.dataset["isShowContent"] == "true") {
        elm.dataset["isShowContent"] = "false";
        elm.query(".rss-item-content").hidden = true;
      } else {
        elm.dataset["isShowContent"] = "true";
        elm.query(".rss-item-content").hidden = false;
      }
    });
    elm.query(".rss-item-content").hidden = true;
  });
}

void getContent(DivElement elm) {
  // @see https://www.dartlang.org/articles/json-web-service/

  HttpRequest request = new HttpRequest();

  // add an event handler that is called when the request finishes
  request.onReadyStateChange.listen((_) {
    if (request.readyState == HttpRequest.DONE &&
        (request.status == 200 || request.status == 0)) {
      // data saved OK.
      showRssItems(request.responseText);
    }
  });

  // POST the data to the server
  String url = "/site/";
  request.open("POST", url, async: false);

  Map<String, String> jsonData = {
    "Text": elm.dataset["text"],
    "Title": elm.dataset["title"],
    "Type": elm.dataset["type"],
    "XmlUrl": elm.dataset["xmlurl"],
    "HtmlUrl": elm.dataset["htmlurl"],
    "Favicon": elm.dataset["favicon"],
  };
  request.send(JSON.encode(jsonData));
}

void seedClicked(Event e) {
  DivElement elm = e.target as DivElement;

  // create new image element
  ImageElement img = new ImageElement();
  img.src = elm.dataset["favicon"];
  img.alt = elm.title;
  img.width = 16;
  img.height = 16;

  // create new anchor element
  AnchorElement a = new AnchorElement();
  a.attributes["href"] = elm.dataset["htmlurl"];
  a.attributes["target"] = "_blank";
  a.text = elm.text;

  AnchorElement a2 = new AnchorElement();
  a2.attributes["href"] = elm.dataset["xmlurl"];
  a2.attributes["target"] = "_blank";
  a2.text = "RSS";

  // append the anchor element to right view
  DivElement view = query(".rss-site-summary");
  view.children.clear();
  view.append(img);
  view.append(a);
  view.append(a2);

  getContent(elm);
}

void init() {
  queryAll(".seed-item").forEach((Element elm) {
    elm.onClick.listen(seedClicked);
  });
}

void main() {
  init();
}
