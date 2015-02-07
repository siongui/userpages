/* set the variable needed by http post */
private String uploadFile="/sdcard/test.png";
private String actionUrl = "http://your_domain/your_post_url";
final String end = "\r\n";
final String twoHyphens = "--";
final String boundary = "*****++++++************++++++++++++";

URL url = new URL(actionUrl);
HttpURLConnection conn = (HttpURLConnection)url.openConnection();

conn.setDoInput(true);
conn.setDoOutput(true);
conn.setUseCaches(false);
conn.setRequestMethod("POST");

/* setRequestProperty */
conn.setRequestProperty("Connection", "Keep-Alive");
conn.setRequestProperty("Charset", "UTF-8");
conn.setRequestProperty("Content-Type", "multipart/form-data;boundary="+ boundary);

DataOutputStream ds = new DataOutputStream(conn.getOutputStream());
ds.writeBytes(twoHyphens + boundary + end);
ds.writeBytes("Content-Disposition: form-data; name=\"from\""+end+end+"auto"+end);
ds.writeBytes(twoHyphens + boundary + end);
ds.writeBytes("Content-Disposition: form-data; name=\"to\""+end+end+"ja"+end);
ds.writeBytes(twoHyphens + boundary + end);
ds.writeBytes("Content-Disposition: form-data; name=\"uploadedFile\";filename=\"" + uploadFile +"\"" + end);
ds.writeBytes(end);

FileInputStream fStream = new FileInputStream(uploadFile);
int bufferSize = 1024;
byte[] buffer = new byte[bufferSize];
int length = -1;

while((length = fStream.read(buffer)) != -1) {
  ds.write(buffer, 0, length);
}
ds.writeBytes(end);
ds.writeBytes(twoHyphens + boundary + twoHyphens + end);
/* close streams */
fStream.close();
ds.flush();
ds.close();

if(conn.getResponseCode() != HttpURLConnection.HTTP_OK){
  Toast.makeText(getBaseContext(), conn.getResponseMessage(), Toast.LENGTH_LONG);
}

StringBuffer b = new StringBuffer();
InputStream is = conn.getInputStream();
byte[] data = new byte[bufferSize];
int leng = -1;
while((leng = is.read(data)) != -1) {
  b.append(new String(data, 0, leng));
}
String result = b.toString();
