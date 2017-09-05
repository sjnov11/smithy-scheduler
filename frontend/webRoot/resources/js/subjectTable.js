// This varaible is a map of SuupNo2 and BanSosokNm
// e.g.) selectedLecture.set("11073","컴퓨터전공")
var selectedLecture = new Map();

// majorDataMap is a map. A 'key' is a major name a 'value' is subjects.
// e.g.) subjectData.set("컴퓨터전공", subjectsOfTheMajor)
var majorDataMap = new Map();

// when receipt_major value is changed
$("input[name='receipt_major']").change(function(event){
  displayMajorTable();
  showSearchResultLoader();

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


  // get major data
  var xhrMajorData = new XMLHttpRequest();
  xhrMajorData.open("POST", "/db/getDataByMajor");
  xhrMajorData.setRequestHeader("Content-Type", "application/json");

  xhrMajorData.onreadystatechange = function() {
    if(xhrMajorData.readyState === 4 && xhrMajorData.status === 200) {
      var majorData = JSON.parse(xhrMajorData.response);
      // TODO: 전공 여러개 선택할 수 있을 때, Map에 여러개의 전공 데이터 들어가야 함
      majorDataMap.clear();
      majorDataMap.set(event.target.value, majorData);
    }
  };

  xhrMajorData.send(JSON.stringify({
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
      addIsuJehanPopup(jqueryCell);

      var eachSubject = jqueryCell.find("div.each-subject.modal-subject-name[lecture-name='" + lectureName + "']");
      addEachSubjectPopup(eachSubject);

      eachSubject.click(function(){

        var selectedLectureNumber = $(this).attr("lecture-number");
        var selectedLectureMajor = $(this).attr("lecture-major");

        if (selectedLecture.has(selectedLectureNumber) == false) {
          // add SuupNo2 and BanSosokNm to selectedLecture
          function addLecture(lecture) {
            selectedLecture.set(selectedLectureNumber, selectedLectureMajor);
            $(lecture).addClass("selected");

            var lectureName = $(lecture).attr("lecture-name");
            $("td.selectable[lecture-name='" + lectureName + "']").addClass("selected");

            console.log("lecture " +selectedLectureNumber+ " is added to var selectedLecture");
          };
          
          sameLectures = $(this).parent().find(".each-subject.modal-subject-name[lecture-number='"+ selectedLectureNumber +"']");
          
          for (var i = 0, l = sameLectures.length; i < l; i++) {
            addLecture(sameLectures[i]);
          }

          // add a lecture data to selectedLectureBasketList
          selectedLectureBasketList.add(selectedLectureNumber, selectedLectureMajor);

          showBasket();

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
          sameLectures = $(this).parent().find(".each-subject.modal-subject-name[lecture-number='"+ selectedLectureNumber +"']");
          for (var i = 0, l = sameLectures.length; i < l; i++) {
            deleteLecture(sameLectures[i]);
          }

          selectedLectureBasketList.remove(selectedLectureNumber);
        }
      });
    } else {
      // No lecture-name. do nothing
    }
  }
}
// addModal();



// This function gives 'selected' class to  already selected lectures 
function alreadySelectedLecture() {
  selectedLecture.forEach(function(majorName, lectureNumber){

    var eachLecture = $("div.each-subject.modal-subject-name[lecture-number='"+lectureNumber+"']");
    var lectureName = eachLecture.attr("lecture-name");

    eachLecture.addClass("selected");
    $("td.selectable[lecture-name='" + lectureName + "']").addClass("selected");
  });
}


function addIsuJehanPopup(modal) {
  modal.find('.isu-jehan-yes')
    .popup({
      popup : $('.isu-jehan-yn.popup'),
  })
  ;
}


function addEachSubjectPopup(eachSubject) {
  eachSubject
    .popup({
      popup : eachSubject.parent().find(".each-subject.popup[lecture-number='"+ $(this).attr("lecture-number") +"']"),
      position : "left center",
    });
}


