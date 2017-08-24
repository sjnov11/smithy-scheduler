new Vue({
  el: '#culturalList',
  data: {
    culturals: [
      '월요일',
      '화요일',
      '수요일',
      '목요일',
      '금요일',
      '토요일',
      '일요일',
      '가상대학영역',
      '경영클러스터영역',
      '고전읽기영역',
      '과학과기술영역',
      '과학기술과환경영역',
      '교양인영역',
      '글로벌언어와문화영역',
      '문학과예술영역',
      '미래산업과창업영역',
      '봉사인영역',
      '사회와세계영역',
      '산학협력영역',
      '세계인영역',
      '소프트웨어영역',
      '실용인영역',
      '언어와세계문화영역',
      '역사와철학영역',
      '영역없음',
      '인간과사회영역',
      '인문과예술영역',
      '일반영역',
      '학연산클러스터영역'
    ]
  }
});

var majorIsSelected = true;
var subjects;

$('.ui.dropdown')
  .dropdown()
;

$('#btn_login').click(function() {
  $('.login.page').dimmer('show');
});

$('#cCondition').click(function() {
  $('.config-c-condition.page').dimmer('show');
});


$('#btn_major').click(function() {
  $(this).removeClass("basic");
  $('#btn_cultural').addClass("basic");

  $('#majorSearchForm').css('display','');
  $('#culturalSearchForm').css('display', 'none');

  majorIsSelected = true;
  displayMajorTable();


  // Update input value
  var formText = $("input[name='receipt_major']").val();
  $('#echoInputValue').text(formText)
});

$('#btn_cultural').click(function() {
  $(this).removeClass("basic");
  $('#btn_major').addClass("basic");

  $('#culturalSearchForm').css('display','');
  $('#majorSearchForm').css('display', 'none');

  majorIsSelected = false;
  hideMajorTable();

  // Update input value
  var formText = $("input[name='receipt_cultural']").val();
  $('#echoInputValue').text(formText)
});


$('#btn_major_selected').click(function() {
  location.href="major-selected.html"
});

$('button.lecture-select').click(function() {
  var cell = $(this).parent().parent().parent().parent();

  if(cell.hasClass('positive')) {
    cell.removeClass('positive');
  } else {
    cell.addClass('positive');
  }
});

// add tooltips to the cells of the lecture tables.
function addTooltip() {
  var lectureCells = $('td.selectable');
  for (var i = 0, l = lectureCells.length; i < l; i++) {
    var cell = lectureCells[i];

    var lectureName = cell.getAttribute('lecture-name');
    if (lectureName != null) {
      $(cell).popup({
        // Find div which has 'lecture-name' attribute to show tooltip
        popup : $("div.lecture-info.popup[lecture-name='" + lectureName +"']"),
        on    : 'click'
      });
    } else {
      // No lecture-name. do nothing
    }
  }
}
addTooltip();


$('.combiation-start.button').popup({
  popup : $(".combination-complete.popup"),
  on    : 'click'
});


$('.combination-result.modal')
  .modal('attach events', '.time-table-thumbnail', 'show')
;

$('.combination-result.modal')
  .modal('attach events', '.time-table', 'hide')
;

$("input[name='receipt_major']").change(function(event){
    // alert(event.target.value + " has been selected.");
    displayMajorTable();

    $('#echoInputValue').text(event.target.value);

  //ajax call to get data from db
  var x = new XMLHttpRequest();
  x.open("POST", "/db/getDataByMajor");
  x.setRequestHeader("Content-Type", "application/json");

  x.onreadystatechange = function() {
    if(x.readyState === 4 && x.status === 200) {
      // console.log(x.response);
      subjects = JSON.parse(x.response);
      console.log("var subjects has got subject data");

      // fill the inner html
      // Mustache.js is needed for creating HTML code
      // document.getElementById('searchResult').innerHTML = generateSearchResultHTML(subjects);
    }
  };

  x.send(JSON.stringify({
    major: event.target.value
  }));

});

function displayMajorTable() {
  if (majorIsSelected && $("input[name='receipt_major']").val() !='') {
    $('.search-result-table').css('display','');

    // hide intro text
    $('.main-intro').css('display', 'none');
  } else {
    $('.search-result-table').css('display','none');
  }

};


function hideMajorTable() {
  $('.search-result-table').css('display','none');
}

$("input[name='receipt_cultural']").change(function(event){
  // alert(event.target.value + " has been selected.");
  $('#echoInputValue').text(event.target.value);
});


$('#btn_request_post').click(function(event) {
  event.preventDefault();

  var x = new XMLHttpRequest();

  x.open("POST", "/db/");
  x.onreadystatechange = function() {
    if(x.readyState === 4 && x.status === 200) {
      console.log(x.response);
      alert(x.response)
    }
  }
  x.send();
});
