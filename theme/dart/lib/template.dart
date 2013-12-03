part of web;


class template {
  String templates_dir;

  // constructor
  template(String templates_dir) {
    this.templates_dir = templates_dir;
  }

  String render(String tempalte, Map<String, Object> data) {
    final File file = new File(
        path.join(this.templates_dir, tempalte));
    return file.readAsStringSync();
  }
}
