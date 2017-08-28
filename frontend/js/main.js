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

var selectedLecture = new Set();

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
function addModal() {
  var lectureCells = $('td.selectable');
  for (var i = 0, l = lectureCells.length; i < l; i++) {
    var cell = lectureCells[i];

    var lectureName = cell.getAttribute('lecture-name');
    if (lectureName != null) {

      var jqueryCell = $(cell).find("div.modal[lecture-name='" + lectureName + "']")
        .modal('attach events', "td.selectable[lecture-name='" + lectureName + "']", 'show');
        // .modal('attach events', "div.each-subject.modal-subject-name[lecture-name='" + lectureName + "']", 'hide')

      jqueryCell.find("div.each-subject.modal-subject-name[lecture-name='" + lectureName + "']").click(function(){

        var selectedLectureNumber = $(this).attr("lecture-number");

        if (selectedLecture.has(selectedLectureNumber) == false) {
          // add
          function addLecture(lecture) {
            selectedLecture.add(selectedLectureNumber);
            $(lecture).addClass("selected");

            var lectureName = $(lecture).attr("lecture-name");
            $("td.selectable[lecture-name='" + lectureName + "']").addClass("selected");

            console.log("lecture " +selectedLectureNumber+ " is added to var selectedLecture");
          };
          
          sameLectures = $(this).parent().find("[lecture-number='"+ selectedLectureNumber +"']");
          for (var i = 0, l = sameLectures.length; i < l; i++) {
            addLecture(sameLectures[i]);
          }

          // addLecture(this);

        } else {
          // delete
          function deleteLecture(lecture) {
            selectedLecture.delete(selectedLectureNumber);
            $(lecture).removeClass("selected");

            // when GwamokNm is same
            sameNameLectures = $(lecture).parent().find(".selected");
            if (sameNameLectures.length == 0){
              // when there is no selected subject
              var lectureName = $(lecture).attr("lecture-name");
              $("td.selectable[lecture-name='" + lectureName + "']").removeClass("selected");
            }

            console.log("lecture " +selectedLectureNumber+ " is removed from var selectedLecture");
          };

          // when SuupNo2 is same
          sameLectures = $(this).parent().find("[lecture-number='"+ selectedLectureNumber +"']");
          for (var i = 0, l = sameLectures.length; i < l; i++) {
            deleteLecture(sameLectures[i]);
          }
        }
      });
    } else {
      // No lecture-name. do nothing
    }
  }
}
// addModal();


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

// This function gives 'selected' class to  already selected lectures 
function alreadySelectedLecture() {
  selectedLecture.forEach(function(lectureNumber){
    console.log(lectureNumber);

    var eachLecture = $("div.each-subject.modal-subject-name[lecture-number='"+lectureNumber+"']");
    var lectureName = eachLecture.attr("lecture-name");

    eachLecture.addClass("selected");
    $("td.selectable[lecture-name='" + lectureName + "']").addClass("selected");
  });
}


function showSearchResultLoader() {
  var loaderHTML = `
    <div class="ui active inverted dimmer">
      <div class="ui large loader"></div>
    </div>`;
  $("#searchResultLoader")[0].innerHTML = loaderHTML;
};

function hideSearchResultLoader() {
  $("#searchResultLoader")[0].innerHTML = "";
}

$("input[name='receipt_major']").change(function(event){
  displayMajorTable();
  showSearchResultLoader();

  // show loader
  // showSearchResultLoader();

  $('#echoInputValue').text(event.target.value);

  // draw table
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/db/getSubjectTable");
  xhr.setRequestHeader("Content-Type", "application/json");

  xhr.onreadystatechange = function() {
    if(xhr.readyState === 4 && xhr.status === 200) {
      // console.log(xhr.response);
      //fill the inner html
      document.getElementById('searchResult').innerHTML = xhr.response;
      addModal();
      alreadySelectedLecture();
      hideSearchResultLoader();
    }
  };

  xhr.send(JSON.stringify({
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
      // console.log(x.response);
      alert(x.response)
    }
  }

  x.send();
});
