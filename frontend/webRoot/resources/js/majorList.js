var vueMajorList = new Vue({
  el: '#majorList',
  data: {
    majorNames: []
  }
});

function fillMajorList() {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/getMajorNameList");

  xhr.onreadystatechange = function() {
    if(xhr.readyState === 4 && xhr.status === 200) {
      vueMajorList.majorNames = JSON.parse(xhr.response);
    }
  };
  xhr.send();
};
fillMajorList();
