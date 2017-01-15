"use strict";

function SaveSetting(s) {
  localStorage.mySetting = JSON.stringify(s);
}

function LoadSetting() {
  return JSON.parse(localStorage.mySetting);
}

function SetupSetting() {
  var enElm = document.getElementById("en");
  var jaElm = document.getElementById("ja");
  var orderElm = document.getElementById("order");

  var setting = {
    en: true,
    ja: true,
    order: "eng",
  };

  // check if there is saved setting in user browser
  if (typeof localStorage.mySetting === 'undefined') {
    setting.en = enElm.checked;
    setting.ja = jaElm.checked;
    setting.order = orderElm.options[orderElm.selectedIndex].value;
  } else {
    setting = LoadSetting();
    enElm.checked = setting.en;
    jaElm.checked = setting.ja;
    orderElm.value = setting.order;
  }

  enElm.addEventListener("click", function(e) {
    setting.en = enElm.checked;
    SaveSetting(setting);
  });

  jaElm.addEventListener("click", function(e) {
    setting.ja = jaElm.checked;
    SaveSetting(setting);
  });

  orderElm.addEventListener("change", function(e) {
    setting.order = orderElm.options[orderElm.selectedIndex].value;
    SaveSetting(setting);
  });
}

SetupSetting();
