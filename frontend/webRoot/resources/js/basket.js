function showBasket() {
  $("#basket").css("display","");
  $("#searchForm").addClass("subject-selected");
};

var selectedLectureList = new Vue({
  el: '#selectedLectureList',
  data: {
    lectureData: []
  }
});


function updateSelectedLectureList(){
  console.log(selectedLecture); 
}
