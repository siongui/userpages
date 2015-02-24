words_prefix_a = ["adwihjd", "aruwdn", "aeiuiwond", "aehuwgd", "aowqmgfv", "awqnmfwxsa", "aqgwqvsa", "aansqnsoq", "awqgwqenjio"];
words_prefix_b = ["bsaiuhw", "baww", "bahxnbzd", "bsqomzxb", "bbsjqbsoq", "bcsihwio", "bjaskdka", "baoshdsji", "bdaiohwi"];
words_prefix_c = ["csjioqjs", "ccwjiow", "csjqios", "cakopwd", "cdjiowhj", "ctjioend", "ckihsiow"];

function strMatch() {
  /* remove whitespace in the beginning and end of the string */
  var userInputStr = document.getElementById("Text1").value.replace(/(^\s+)|(\s+$)/g, "");

  document.getElementById("userInput").innerHTML = userInputStr;

  /* Here we give simple implementation for prefix matching */
  if (userInputStr.length == 0){
    document.getElementById("result").innerHTML = "";
    return;
  }

  var arrayName = "words_prefix_" + userInputStr[0];

  var matched_count = 0;
  var matched_array = new Array();
  /* Start to search the matched prefix,
 *      about the use of eval(), try Google Search keyword: javascript evaluate string as variable */
  for ( i=0; i < eval(arrayName).length; i++ ) {
    if (eval(arrayName)[i].indexOf(userInputStr) == 0) {// If the word starts with the string that users input
      matched_array.push(eval(arrayName)[i]);
      matched_count += 1;
    }
    if (matched_count == 25) {break;}//show no more than 25 words in prefix match
  }
  /* compile matched result into HTML code */
  var matched_result = new String();
  for (i=0; i<matched_array.length; i++) {
    matched_result = matched_result + matched_array[i] + "<br />";
  }
  document.getElementById("result").innerHTML = matched_result;// show matched result
}
